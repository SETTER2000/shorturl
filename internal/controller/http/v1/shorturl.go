package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/SETTER2000/shorturl/internal/usecase/repo"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
	"github.com/SETTER2000/shorturl/scripts"
)

type shorturlRoutes struct {
	s   usecase.IShorturl
	l   logger.Interface
	cfg *config.Config
}

func newShorturlRoutes(handler chi.Router, s usecase.IShorturl, l logger.Interface, cfg *config.Config) {
	sr := &shorturlRoutes{s, l, cfg}
	handler.Route("/user", func(r chi.Router) {
		r.Get("/urls", sr.urls)
		r.Delete("/urls", sr.delUrls2)
	})
	handler.Route("/shorten", func(r chi.Router) {
		r.Post("/", sr.shorten) // POST /
		r.Post("/batch", sr.batch)
	})
	handler.Route("/internal", func(r chi.Router) {
		r.Get("/stats", sr.stats)
	})
}

// @Summary     Return short URL
// @Description Redirect to long URL
// @ID          ShortLink
// @Tags  	    shorturl
// @Accept      text
// @Produce     text
// @Success     307 {object} string
// @Failure     500 {object} response
// @Router      /{key} [get]

func (r *shorturlRoutes) shortLink(w http.ResponseWriter, req *http.Request) {
	shorturl := chi.URLParam(req, "key")
	data := entity.Shorturl{Config: r.cfg}
	data.Slug = shorturl
	sh, err := r.s.ShortLink(req.Context(), &data)
	if err != nil {
		r.l.Error(err, "http - v1 - shortLink")
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	// При запросе удалённого URL с помощью хендлера GET /{id} нужно вернуть статус 410 Gone
	if sh.Del {
		w.WriteHeader(http.StatusGone)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Add("Content-Encoding", "gzip")
	w.Header().Set("Location", sh.URL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// GET /ping, который при запросе проверяет соединение с базой данных
// при успешной проверке хендлер должен вернуть HTTP-статус 200 OK
// при неуспешной — 500 Internal Server Error
func (r *shorturlRoutes) connect(res http.ResponseWriter, req *http.Request) {
	dsn, ok := os.LookupEnv("DATABASE_DSN")
	if !ok || dsn == "" {
		dsn = r.cfg.ConnectDB
		if dsn == "" {
			r.l.Info("connect DSN string is empty: %v\n", dsn)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		db, err := pgx.Connect(req.Context(), os.Getenv("DATABASE_DSN"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		defer db.Close(context.Background())

		fmt.Printf("connect... \n")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("connect... "))
	}
}

// @Summary     Return short URL
// @Description Redirect to long URL
// @ID          longLink
// @Tags  	    shorturl
// @Accept      text
// @Produce     text
// @Success     201 {object} string
// @Failure     500 {object} response
// @Router      / [post]
func (r *shorturlRoutes) longLink(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	data := entity.Shorturl{Config: r.cfg}
	data.URL = string(body)
	data.Slug = scripts.UniqueString()
	data.UserID = ctx.Value("access_token").(string)
	shorturl, err := r.s.LongLink(ctx, &data)
	if err != nil {
		if errors.Is(err, repo.ErrAlreadyExists) {
			data2 := entity.Shorturl{
				Config: r.cfg,
				URL:    data.URL,
				UserID: data.UserID}

			sh, err := r.s.ShortLink(ctx, &data2)
			if err != nil {
				r.l.Error(err, "http - v1 - longLink")
				http.Error(res, fmt.Sprintf("%v", err), http.StatusBadRequest)
				return
			}
			shorturl = sh.Slug
			res.Header().Set("Content-Type", http.DetectContentType(body))
			res.WriteHeader(http.StatusConflict)
		} else {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}
	d := scripts.GetHost(r.cfg.HTTP, shorturl)
	res.Header().Set("Content-Type", http.DetectContentType(body))
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(d))
}

// GET
func (r *shorturlRoutes) urls(res http.ResponseWriter, req *http.Request) {
	u := entity.User{}
	userID := req.Context().Value("access_token")
	if userID == nil {
		res.Write([]byte(fmt.Sprintf("Not access_token and user_id: %s", userID)))
	}
	u.UserID = fmt.Sprintf("%s", userID)
	user, err := r.s.UserAllLink(req.Context(), &u)
	if err != nil {
		r.l.Error(err, "http - v1 - urls")
		http.Error(res, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	encoded, err := json.Marshal(user.Urls)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%v", len(encoded))
	res.Header().Set("Content-Type", "application/json")
	if string(encoded) == "null" {
		res.WriteHeader(http.StatusNoContent)
	} else {
		res.WriteHeader(http.StatusOK)
	}
	res.Write(encoded)
}

// GET /api/internal/stats
func (r *shorturlRoutes) stats(w http.ResponseWriter, req *http.Request) {
	ip, err := resolveIP(req, resolveIPOpts{
		UseHeader:     r.cfg.ResolveIPUsingHeader,
		TrustedSubnet: r.cfg.HTTP.TrustedSubnet,
	})
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	ok := resolveSubNet(ip, r.cfg)
	if !ok {
		http.Error(w, fmt.Sprintf("%v", ErrForbidden), http.StatusForbidden)
		return
	}

	Static := &entity.Static{}
	urls, err := r.s.AllLink()
	if err != nil {
		r.l.Error(err, "http - v1 - stats")
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	users, err := r.s.AllUsers()
	if err != nil {
		r.l.Error(err, "http - v1 - stats")
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	Static.CountURLs = urls
	Static.CountUsers = users
	encoded, err := json.Marshal(&Static)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if string(encoded) == "null" {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(encoded)
}

// @Summary     Return JSON short URL
// @Description Redirect to long URL
// @ID          shorten
// @Tags  	    shorturl
// @Accept      json
// @Produce     json
// @Success     307 {object} string
// @Failure     500 {object} response
// @Router      /shorten [post]
func (r *shorturlRoutes) shorten(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	data := entity.Shorturl{Config: r.cfg}
	resp := entity.ShorturlResponse{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	data.Slug = scripts.UniqueString()
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	data.UserID = req.Context().Value(r.cfg.AccessTokenName).(string)
	err = r.s.Post(ctx, &data)
	resp.URL = scripts.GetHost(r.cfg.HTTP, data.Slug)
	if err != nil {
		if errors.Is(err, repo.ErrAlreadyExists) {
			data2 := entity.Shorturl{
				Config: r.cfg,
				URL:    data.URL,
				UserID: data.UserID}

			sh, err := r.s.ShortLink(ctx, &data2)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
			}
			// return shorturl
			resp.URL = scripts.GetHost(r.cfg.HTTP, sh.Slug)

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusConflict)
		} else {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}

	encoded, err := json.Marshal(resp)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(encoded)
}

func (r *shorturlRoutes) batch(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	data := entity.Shorturl{Config: r.cfg}
	CorrelationOrigin := entity.CorrelationOrigin{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(body, &CorrelationOrigin); err != nil {
		panic(err)
	}

	var rs entity.Response
	var sr entity.ShortenResponse
	UserID := ctx.Value(r.cfg.AccessTokenName).(string)
	for _, bt := range CorrelationOrigin {
		data.URL = bt.URL
		data.Slug = bt.Slug
		data.UserID = UserID
		err = r.s.Post(ctx, &data)
		if err != nil {
			if errors.Is(err, repo.ErrAlreadyExists) {
				res.WriteHeader(http.StatusConflict)
				return
			}
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		sr.Slug = data.Slug
		sr.URL = scripts.GetHost(r.cfg.HTTP, data.Slug)
		rs = append(rs, sr)
	}

	encoded, err := json.Marshal(rs)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(encoded)
}

// Асинхронный (КАБУДА!) хендлер DELETE /api/user/urls,
// который принимает список идентификаторов сокращённых URL для удаления
// в формате: [ "a", "b", "c", "d", ...]
// В случае успешного приёма запроса хендлер должен возвращать HTTP-статус 202 Accepted.
// Фактический результат удаления может происходить позже — каким-либо
// образом оповещать пользователя об успешности или неуспешности не нужно.
func (r *shorturlRoutes) delUrls(res http.ResponseWriter, req *http.Request) {
	var slugs []string
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(body, &slugs); err != nil {
		panic(err)
	}

	u := entity.User{}
	userID := req.Context().Value("access_token")
	if userID == nil {
		res.Write([]byte(fmt.Sprintf("Not access_token and user_id: %s", userID)))
	}
	u.UserID = fmt.Sprintf("%s", userID)
	u.DelLink = slugs

	//-- fanOut fanIn - multithreading
	const workersCount = 16
	inputCh := make(chan entity.User)
	// входные значения кладём в inputCh
	go func(u entity.User) {
		inputCh <- u
		close(inputCh)
	}(u)
	// здесь fanOut
	fanOutChs := fanOut(inputCh, workersCount)
	workerChs := make([]chan entity.User, 0, workersCount)
	for _, fanOutCh := range fanOutChs {
		workerCh := make(chan entity.User)
		newWorker(r, req, fanOutCh, workerCh)
		workerChs = append(workerChs, workerCh)
	}

	// здесь fanIn
	for v := range fanIn(workerChs...) {
		r.l.Info("%s\n", v.UserID)
	}
	//-- end fanOut fanIn

	//-- not multithreading
	//for i := 0; i < len(slugs); i++ {
	//	fmt.Printf("SLUG#%d: %s\n", i, slugs[i])
	//	err = r.s.UserDelLink(req.Context(), &u)
	//}
	//if err != nil {
	//	r.l.Error(err, "http - v1 - delUrls")
	//	http.Error(res, fmt.Sprintf("%v", err), http.StatusBadRequest)
	//	return
	//}
	//-- end not multithreading

	res.WriteHeader(http.StatusAccepted)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("Ok!"))
}

func (r *shorturlRoutes) delUrls2(res http.ResponseWriter, req *http.Request) {
	var slugs []string
	const workersCount = 10
	inputCh := make(chan entity.User)

	go func() {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(body, &slugs); err != nil {
			panic(err)
		}
		u := entity.User{}
		userID := req.Context().Value("access_token")
		if userID == nil {
			res.Write([]byte(fmt.Sprintf("Not access_token and user_id: %s", userID)))
		}
		u.UserID = fmt.Sprintf("%s", userID)
		u.DelLink = slugs
		inputCh <- u
		close(inputCh)
	}()

	// здесь fanOut
	fanOutChs := fanOut(inputCh, workersCount)
	workerChs := make([]chan entity.User, 0, workersCount)
	for _, fanOutCh := range fanOutChs {
		workerCh := make(chan entity.User)
		newWorker(r, req, fanOutCh, workerCh)
		workerChs = append(workerChs, workerCh)
	}

	// здесь fanIn
	for v := range fanIn(workerChs...) {
		r.l.Info("%s\n", v.UserID)
	}

	res.WriteHeader(http.StatusAccepted)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("Ok!"))
}

func newWorker(r *shorturlRoutes, req *http.Request, input, out chan entity.User) {
	go func() {
		us := entity.User{}
		for u := range input {
			fmt.Printf("UserID: %s, DelLink: %s count: %v ", u.UserID, u.DelLink, len(u.DelLink))
			r.s.UserDelLink(req.Context(), &u)
			out <- us
		}
		close(out)
	}()
	time.Sleep(50 * time.Millisecond)
}
func fanIn(inputChs ...chan entity.User) chan entity.User {
	// один выходной канал, куда сливаются данные из всех каналов
	outCh := make(chan entity.User)

	go func() {
		wg := &sync.WaitGroup{}

		for _, inputCh := range inputChs {
			wg.Add(1)

			go func(inputCh chan entity.User) {
				defer wg.Done()
				for item := range inputCh {
					outCh <- item
				}
			}(inputCh)
		}

		wg.Wait()
		close(outCh)
	}()

	return outCh
}

func fanOut(inputCh chan entity.User, n int) []chan entity.User {
	chs := make([]chan entity.User, 0, n)
	for i := 0; i < n; i++ {
		ch := make(chan entity.User)
		chs = append(chs, ch)
	}

	go func() {
		defer func(chs []chan entity.User) {
			for _, ch := range chs {
				close(ch)
			}
		}(chs)

		for i := 0; ; i++ {
			if i == len(chs) {
				i = 0
			}

			num, ok := <-inputCh
			if !ok {
				return
			}

			ch := chs[i]
			ch <- num
		}
	}()

	return chs
}

type resolveIPOpts struct {
	UseHeader     bool
	TrustedSubnet string
}

func resolveIP(r *http.Request, opts resolveIPOpts) (net.IP, error) {
	network := "tcp"
	address := "poaleell.com:80"
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Printf("DIAL err: %s\n", err)
	}
	defer conn.Close()
	localAddr := conn.RemoteAddr().(*net.TCPAddr)
	fmt.Printf("Address of Dial function Remote IP Addr %s: %v\n", address, localAddr.IP)
	fmt.Printf("ENV TRUSTED_SUBNET: %s\n", os.Getenv("TRUSTED_SUBNET"))

	if opts.TrustedSubnet == "" && os.Getenv("TRUSTED_SUBNET") == "" {
		return nil, fmt.Errorf("err from resolveIP opts.TrustedSubnet is empty")
	}

	if !opts.UseHeader {
		addr := r.RemoteAddr
		// метод возвращает адрес в формате host:port
		// нужна только подстрока host
		ipStr, _, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}
		// парсим ip
		ip := net.ParseIP(ipStr)
		if ip == nil {
			panic("unexpected parse ip error")
		}
		return ip, nil
	} else {
		// смотрим заголовок запроса X-Real-IP
		ipStr := r.Header.Get("X-Real-IP")
		// парсим ip
		ip := net.ParseIP(ipStr)
		if ip == nil {
			// если заголовок X-Real-IP пуст, пробуем X-Forwarded-For
			// этот заголовок содержит адреса отправителя и промежуточных прокси
			// в виде 203.0.113.195, 70.41.3.18, 150.172.238.178
			ips := r.Header.Get("X-Forwarded-For")
			// разделяем цепочку адресов
			ipStrs := strings.Split(ips, ",")
			// интересует только первый
			ipStr = ipStrs[0]
			// парсим
			ip = net.ParseIP(ipStr)
		}
		if ip == nil {
			return nil, fmt.Errorf("failed parse ip from http header")
		}
		return ip, nil
	}
}

// resolveSubNet - проверить входит ли ip-адрес в доверенную подсеть
func resolveSubNet(ip net.IP, cfg *config.Config) bool {
	netMask, err := strconv.Atoi(strings.Split(cfg.HTTP.TrustedSubnet, "/")[1:][0])
	if err != nil {
		log.Printf("Ошибка %e", err)
	}

	ipv4Addr, ipv4Net, err := net.ParseCIDR(cfg.HTTP.TrustedSubnet)
	if err != nil {
		log.Fatal(err)
	}

	// This mask corresponds to a /24 subnet for IPv4.
	ipv4Mask := net.CIDRMask(netMask, 32)
	fmt.Printf("ipv4Addr подсети: %v\n ", ipv4Addr.Mask(ipv4Mask))
	fmt.Printf("IP %v входит в подсеть %v? : %v\n", ip, cfg.HTTP.TrustedSubnet, ipv4Net.Contains(ip))
	return ipv4Net.Contains(ip)
}

// countHosts кол-во возможных хостов в подсети
func countHosts(b int) (int, error) {
	if b > 24 || b < 2 {
		return 0, ErrBadRequest
	}
	size := 7

	m := []int{254, 126, 62, 30, 14, 6, 2}
	masks := make(map[int]int, size)
	for i := 0; i < size; i++ {
		u := 24 + i
		masks[u] = m[i]
	}

	return masks[b], nil
}

//func resolveSubNetOld(ip net.IP, cfg *config.Config) bool {
//	maska := [4]int{255, 255, 255}
//	k := make(map[int]byte, 7)
//	k[24] = 255 - 255
//	k[25] = 255 - 127
//	k[26] = 255 - 63
//	k[27] = 255 - 31
//	k[28] = 255 - 15
//	k[29] = 255 - 7
//	k[30] = 255 - 3
//	log.Printf("Bytes:: %v", k[24])
//
//	num := int64(192)
//	m := 255
//	bitwiseIP := strings.Split(ip.String(), ".")
//	addressSubNet := strings.Split(cfg.HTTP.TrustedSubnet, "/")[:1][0]
//	log.Printf("Адрес подсети: %s\n", addressSubNet)
//	netMask, err := strconv.Atoi(strings.Split(cfg.HTTP.TrustedSubnet, "/")[1:][0])
//	if err != nil {
//		log.Printf("Ошибка %e", err)
//	}
//
//	cnt, err := countHosts(netMask)
//	if err != nil {
//		log.Printf("%e", err)
//	}
//
//	log.Printf("Кол-во доступных хостов по этой маске: %d\n", cnt)
//
//	ipv4Addr := net.ParseIP(ip.String())
//	// This mask corresponds to a /24 subnet for IPv4.
//	ipv4Mask := net.CIDRMask(netMask, 32)
//	fmt.Printf("ipv4Addr подсети: %v\n ", ipv4Addr.Mask(ipv4Mask))
//
//	maska[3] = int(k[netMask])
//	mn := fmt.Sprintf("255.255.255.%v", k[netMask])
//	log.Printf("Маска подсети:: %v - %v\n", mn, maska)
//	log.Printf("Маска netMask:: %v\n", netMask)
//
//	xBin := strconv.FormatInt(num, 2)
//	log.Printf("r.cfg.ResolveIPUsingHeader: %v\n", cfg.ResolveIPUsingHeader)
//	log.Printf("IP в заголовке X-Real-IP: %v - %v\n", bitwiseIP, ip.String())
//	log.Printf("Двоичное числа %d: %T, %v\n", num, xBin, xBin)
//	log.Printf("cfg.HTTP.TrustedSubnet: %v", cfg.HTTP.TrustedSubnet)
//	var b int = 192 & m
//	log.Printf("&: поразрядная конъюнкция (операция И или поразрядное умножение): %d & %v = %v\n", num, m, b)
//	return false
//}

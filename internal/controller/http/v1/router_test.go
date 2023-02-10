package v1

import (
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
	"github.com/go-chi/chi/v5"
	"testing"
)

//func TestLongURL(t *testing.T) {
//	type want struct {
//		code        int
//		response    string
//		contentType string
//	}
//	tests := []struct {
//		name string
//		want want
//	}{
//		{
//			name: "positive test #1",
//			want: want{
//				code:        201,
//				response:    `^http://localhost:(\d+)/(\d{19})(\w{3})$`,
//				contentType: "text/plain",
//			},
//		},
//	}
//	for _, tt := range tests {
//		// запускаем каждый тест
//		t.Run(tt.name, func(t *testing.T) {
//			request := httptest.NewRequest(http.MethodPost, "/", nil)
//
//			// создаём новый Recorder
//			w := httptest.NewRecorder()
//			// определяем хендлер
//			h := http.HandlerFunc(LongURL)
//			// запускаем сервер
//			h.ServeHTTP(w, request)
//			res := w.Result()
//
//			// проверяем код ответа
//			if res.StatusCode != tt.want.code {
//				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
//			}
//
//			// получаем и проверяем тело запроса
//			defer res.Body.Close()
//			resBody, err := io.ReadAll(res.Body)
//			if err != nil {
//				t.Fatal(err)
//			}
//			// regexp
//			matched, _ := regexp.MatchString(tt.want.response, string(resBody))
//			if !matched {
//				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
//			}
//
//			// заголовок ответа
//			if res.Header.Get("Content-Type") != tt.want.contentType {
//				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
//			}
//		})
//	}
//}
//
//func TestShortURL(t *testing.T) {
//	type want struct {
//		code                int
//		response            string
//		contentType         string
//		contentTypeLocation string
//	}
//	tests := []struct {
//		name string
//		want want
//	}{
//		{
//			name: "bad test #1",
//			want: want{
//				code:                400,
//				response:            `^http://(\w)+(:(\d+))?/./(\d{19})(\w{3})$`,
//				contentType:         "text/plain; charset=utf-8",
//				contentTypeLocation: `^http|https://(\w)+(:(\d+))?./(\d{19})(\w{3})$`,
//			},
//		},
//	}
//	for _, tt := range tests {
//		// запускаем каждый тест
//		t.Run(tt.name, func(t *testing.T) {
//			//request := httptest.NewRequest(http.MethodPost, "/", nil)
//			request := httptest.NewRequest(http.MethodGet, "/", nil)
//			w := httptest.NewRecorder()
//			h := http.HandlerFunc(ShortURL)
//			h.ServeHTTP(w, request)
//			res := w.Result()
//			// запускаем сервер
//			// определяем хендлер
//			// создаём новый Recorder
//			defer res.Body.Close()
//
//			// TODO status code 307 не знаю как проверить
//			// проверяем код ответа
//			if res.StatusCode != tt.want.code {
//				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
//			}
//
//			// получаем и проверяем тело запроса
//
//			//resBody, err := io.ReadAll(res.Body)
//			//if err != nil {
//			//	t.Fatal(err)
//			//}
//			//// regexp
//			//matched, _ := regexp.MatchString(tt.want.response, string(resBody))
//			//if !matched {
//			//	t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
//			//}
//			// заголовок ответа
//			//if res.Header.Get("Location") != tt.want.contentTypeLocation {
//			//	t.Errorf("Expected Location %s, got %s", tt.want.contentTypeLocation, res.Header.Get("Location"))
//			//}
//
//			// заголовок ответа
//			if res.Header.Get("Content-Type") != tt.want.contentType {
//				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
//			}
//		})
//	}
//}
//
//
//func TestStatusHandler(t *testing.T) {
//	type want struct {
//		code        int
//		response    string
//		contentType string
//	}
//	tests := []struct {
//		name string
//		want want
//	}{
//		{
//			name: "positive test #1",
//			want: want{
//				code:        200,
//				response:    `{"status":"ok"}`,
//				contentType: "application/json",
//			},
//		},
//	}
//	for _, tt := range tests {
//		// запускаем каждый тест
//		t.Run(tt.name, func(t *testing.T) {
//			request := httptest.NewRequest(http.MethodGet, "/status", nil)
//
//			// создаём новый Recorder
//			w := httptest.NewRecorder()
//			// определяем хендлер
//			h := http.HandlerFunc(StatusHandler)
//			// запускаем сервер
//			h.ServeHTTP(w, request)
//			res := w.Result()
//
//			// проверяем код ответа
//			if res.StatusCode != tt.want.code {
//				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
//			}
//
//			// получаем и проверяем тело запроса
//			defer res.Body.Close()
//			resBody, err := io.ReadAll(res.Body)
//			if err != nil {
//				t.Fatal(err)
//			}
//			if string(resBody) != tt.want.response {
//				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
//			}
//
//			// заголовок ответа
//			if res.Header.Get("Content-Type") != tt.want.contentType {
//				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
//			}
//		})
//	}
//}
//
//func Test_dbNewLink(t *testing.T) {
//	type fields struct {
//		Slug string
//		URL  string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		{
//			name: "Sample test returns always",
//			fields: fields{
//				Slug: "",
//				URL:  "",
//			},
//			want: `^(\d{19})(\w{3})$`, // сгенерированная строка
//		},
//		{
//			name: "Test Slug for length and content",
//			fields: fields{
//				Slug: "",
//				URL:  "https://to.ru/go/fo/23423/34534534/dfghgfgh",
//			},
//			want: `^(\d{19})(\w{3})$`, // сгенерированная строка
//		},
//	}
//	for _, tt := range tests {
//		// запускаем каждый тест
//		t.Run(tt.name, func(t *testing.T) {
//			lnk := Link{
//				Slug: tt.fields.Slug,
//				URL:  tt.fields.URL,
//			}
//			v, _ := dbNewLink(&lnk)
//
//			// пустой «приемник» ошибки, ведь мы уверены, что пример отработает нормально
//			matched, _ := regexp.MatchString(tt.want, v)
//			if !matched {
//				t.Errorf("Expected body %s, got %s", tt.want, v)
//			}
//		})
//	}
//}
//
//func Test_urlFunc(t *testing.T) {
//	type want struct {
//		code        int
//		response    string
//		contentType string
//	}
//	tests := []struct {
//		name string
//		want want
//	}{
//		{
//			name: "positive test #1",
//			want: want{
//				code:        200,
//				response:    `{"status":"ok"}`,
//				contentType: "application/json",
//			},
//		},
//	}
//	for _, tt := range tests {
//		// запускаем каждый тест
//		t.Run(tt.name, func(t *testing.T) {
//			request := httptest.NewRequest(http.MethodGet, "/status", nil)
//
//			// создаём новый Recorder
//			w := httptest.NewRecorder()
//			// определяем хендлер
//			h := http.HandlerFunc(StatusHandler)
//			// запускаем сервер
//			h.ServeHTTP(w, request)
//			res := w.Result()
//
//			// проверяем код ответа
//			if res.StatusCode != tt.want.code {
//				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
//			}
//
//			// получаем и проверяем тело запроса
//			defer res.Body.Close()
//			resBody, err := io.ReadAll(res.Body)
//			if err != nil {
//				t.Fatal(err)
//			}
//			if string(resBody) != tt.want.response {
//				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
//			}
//
//			// заголовок ответа
//			if res.Header.Get("Content-Type") != tt.want.contentType {
//				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
//			}
//		})
//	}
//}

func TestNewRouter(t *testing.T) {
	type args struct {
		handler *chi.Mux
		l       logger.Interface
		s       usecase.Shorturl
		cfg     config.HTTP
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewRouter(tt.args.handler, tt.args.l, tt.args.s, tt.args.cfg)
		})
	}
}

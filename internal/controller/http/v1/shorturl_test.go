package v1

//
//var cfg = &config.Config{
//	HTTP: config.HTTP{
//		BaseURL:       "http://localhost:8080",
//		ServerAddress: "localhost:8080",
//	},
//	Storage: config.Storage{
//		FileStorage: "storage.txt",
//		ConnectDB:   "postgres://postgres:123456@localhost:5432/postgres",
//	},
//	Cookie: config.Cookie{
//		AccessTokenName: "access_token",
//		SecretKey:       "RtsynerpoGIYdab_s234r",
//	},
//}

//func TestStatusHandler(t *testing.T) {
//	type want struct {
//		code        int
//		response    string
//		contentType string
//	}
//	// создаём массив тестов: имя и желаемый результат
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

//func Test_shorturlRoutes_urls(t *testing.T) {
//	//storage := &config.Storage {
//	//	// FILE_STORAGE_PATH путь до файла с сокращёнными URL (директории не создаёт)
//	//	FileStorage string `env:"FILE_STORAGE_PATH"`
//	//	// Строка с адресом подключения к БД, например для PostgreSQL (драйвер pgx): postgres://username:password@localhost:5432/database_name
//	//	ConnectDB string `env:"DATABASE_DSN"`
//	//}
//
//	//cfg, err := config.NewConfig()
//	//if err != nil {
//	//	fmt.Errorf("error test: %e", err)
//	//}
//	type fields struct {
//		//s   usecase.Shorturl
//		//l   logger.Interface
//		cfg *config.Config
//	}
//	//type args struct {
//	//	res http.ResponseWriter
//	//	req *http.Request
//	//}
//	type want struct {
//		code        int
//		response    string
//		contentType string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		//args   args
//		want want
//	}{
//		{
//			name: "positive test #1",
//			fields: fields{
//				cfg: cfg,
//			},
//			//args: args{
//			//	res:    `{"status":"ok"}`,
//			//},
//			want: want{
//				code:        200,
//				response:    `{"status":"ok"}`,
//				contentType: "application/json",
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &shorturlRoutes{
//				//s:   tt.fields.s,
//				//l:   tt.fields.l,
//				cfg: tt.fields.cfg,
//			}
//			request := httptest.NewRequest(http.MethodGet, "/urls", nil)
//			//request := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
//			//request := httptest.NewRequest(http.MethodGet, "/user/urls", nil)
//
//			// создаём новый Recorder
//			w := httptest.NewRecorder()
//			// определяем хендлер
//			h := http.HandlerFunc(r.urls)
//			// запускаем сервер
//			h.ServeHTTP(w, request)
//			res := w.Result()
//
//			//проверяем код ответа
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

//func Test_shorturlRoutes_urls(t *testing.T) {
//	type fields struct {
//		s   usecase.Shorturl
//		l   logger.Interface
//		cfg *config.Config
//	}
//	type args struct {
//		res http.ResponseWriter
//		req *http.Request
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &shorturlRoutes{
//				s:   tt.fields.s,
//				l:   tt.fields.l,
//				cfg: tt.fields.cfg,
//			}
//			r.urls(tt.args.res, tt.args.req)
//		})
//	}
//}

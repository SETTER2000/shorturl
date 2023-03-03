package v1

//func Test_shorturlRoutes_shorten(t *testing.T) {
//	type mockBehavior func(s *mock_usecase.MockShorturl, sh entity.Shorturl)
//	type fields struct {
//		s   usecase.Shorturl
//		l   logger.Interface
//		cfg *config.Config
//	}
//	tests := []struct {
//		name                string
//		inputBody           string
//		fields              fields
//		inputSH             entity.Shorturl
//		mockBehavior        mockBehavior
//		expectedStatusCode  int
//		expectedRequestBody string
//	}{
//		{
//			name:      "positive test #1",
//			inputBody: `{"url":"https://lphp.ru"}`,
//			inputSH: entity.Shorturl{
//				Slug:   "",
//				URL:    "https://lphp.ru",
//				UserID: "",
//			},
//			mockBehavior: func(s *mock_usecase.MockShorturl, sh entity.Shorturl) {
//				s.EXPECT().Shorten(context.Background(), sh).Return("https://lphp.ru", nil)
//			},
//			expectedStatusCode:  201,
//			expectedRequestBody: `{"result":"http://localhost:8080"}`,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &shorturlRoutes{
//				s:   tt.fields.s,
//				l:   tt.fields.l,
//				cfg: tt.fields.cfg,
//			}
//			c := gomock.NewController(t)
//			defer c.Finish()
//
//			sh := mock_usecase.NewMockShorturl(c)
//			//ShRepo := mock_usecase.NewMockShorturlRepo(c)
//			tt.mockBehavior(sh, tt.inputSH)
//
//			//shUsCs := &usecase.ShorturlUseCase{
//			//	repo: ShRepo,
//			//}
//
//			nr := chi.NewRouter()
//			nr.Post("/", r.shorten)
//		})
//	}
//}
//

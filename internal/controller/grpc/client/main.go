package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	// ...
	pb "github.com/SETTER2000/shorturl-service-api/api"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	grpcTarget := fmt.Sprintf("%s:%s", "", "8088")
	// устанавливаем соединение с сервером
	conn, err := grpc.Dial(grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// получаем переменную интерфейсного типа UsersClient,
	// через которую будем отправлять сообщения
	c := pb.NewIShorturlClient(conn)

	startTime := time.Now()
	logrus.Info("Starting gRPC Client at: ", startTime.Format(time.RFC3339))

	// функция, в которой будем отправлять сообщения
	TestShorturls(c)

	logrus.Infof("[completed in %v mill.sec] SUCCESS 🐶", time.Since(startTime).Milliseconds())
}

// TestShorturls -.
func TestShorturls(c pb.IShorturlClient) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// набор тестовых данных
	shorturls := []*pb.Shorturl{
		{Slug: "1042", Url: "https://lphp.ru/artpage/44.html", UserId: "168636114238440868834_", Del: false},
		{Slug: "1340", Url: "https://lphp.ru/artpage/34.html", UserId: "168636114238440868834_", Del: false},
	}

	// Post
	resp, err := c.Post(ctx, &pb.PostRequest{
		Shorturl: &pb.Shorturl{
			Slug:   "10",
			Url:    "https://lphp.ru/artpage/10.html",
			UserId: "168636114238440868834_",
			Del:    false,
		},
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.NotFound {
				// выведет, что пользователь не найден
				fmt.Println(`NOT FOUND Post:`, e.Message())
			} else {
				// в остальных случаях выводим код ошибки в виде строки и сообщение
				fmt.Println(e.Code(), e.Message())
			}
		} else {
			fmt.Printf("Не получилось распарсить ошибку %v", err)
		}
	}
	fmt.Printf("Client.Post response: %v\n", resp)

	// LongLink
	for _, shorturl := range shorturls {
		// добавляем URL
		resp, err := c.LongLink(ctx, &pb.LongLinkRequest{
			Shorturl: shorturl,
		})
		if err != nil {
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.NotFound {
					// выведет, что пользователь не найден
					fmt.Println(`NOT FOUND`, e.Message())
					//return nil, status.Errorf(codes.AlreadyExists, `Короткий URL для %s не создан, такой адрес уже существует.`, in.Shorturl.Url)
				} else {
					// в остальных случаях выводим код ошибки в виде строки и сообщение
					fmt.Println(e.Code(), e.Message())
				}
			} else {
				fmt.Printf("Не получилось распарсить ошибку %v", err)
			}
		}

		fmt.Printf("Client.LongLink response: %v\n", resp)
	}

	// UserDelLink
	for _, shorturl := range shorturls {
		u := pb.User{
			UserId:  shorturl.UserId,
			DelLink: []string{"1042", "1340"},
		}
		// добавляем URL
		resp, err := c.UserDelLink(ctx, &pb.UserDelLinkRequest{
			User: &u,
		})
		if err != nil {
			log.Fatalf("could not sh: %v", err)
		}
		if resp.Error != "" {
			fmt.Println(resp.Error)
		}

		fmt.Printf("Client.UserDelLink response: %v\n", true)
	}

	// ShortLink
	//resp, err = c.ShortLink(ctx, &pb.ShortLinkRequest{
	//	Shorturl: &pb.Shorturl{
	//		Slug:   "10",
	//	},
	//})
	//if err != nil {
	//	if e, ok := status.FromError(err); ok {
	//		if e.Code() == codes.NotFound {
	//			// выведет, что пользователь не найден
	//			fmt.Println(`NOT FOUND ShortLink:`, e.Message())
	//		} else {
	//			// в остальных случаях выводим код ошибки в виде строки и сообщение
	//			fmt.Println(e.Code(), e.Message())
	//		}
	//	} else {
	//		fmt.Printf("Не получилось распарсить ошибку %v", err)
	//	}
	//}
	//fmt.Printf("Client.Post response: %v\n", resp)

	//// если запрос будет выполняться дольше 200 миллисекунд, то вернётся ошибка
	//// с кодом codes.DeadlineExceeded и сообщением context deadline exceeded
	//ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	//defer cancel()
	//
	//// получаем информацию о пользователях
	//// во втором случае должна вернуться ошибка:
	//// пользователь с email serge@example.com не найден
	//for _, userEmail := range []string{"sveta@example.com", "serge@example.com"} {
	//	resp, err := c.GetUser(ctx, &pb.GetUserRequest{
	//		Email: userEmail,
	//	})
	//	if err != nil {
	//		// обработка ошибок
	//		if e, ok := status.FromError(err); ok {
	//			switch e.Code() {
	//			//if e.Code() == codes.NotFound {
	//			// выведет, что пользователь не найден
	//			//	fmt.Println(`NOT FOUND`, e.Message())
	//			//} else {
	//			// в остальных случаях выводим код ошибки в виде строки и сообщение
	//			//fmt.Println(e.Code(), e.Message())
	//
	//			// выводить только текст сообщения
	//			case codes.NotFound, codes.DeadlineExceeded:
	//				fmt.Println(e.Message())
	//			default:
	//				// выводится текст и код ошибки
	//				fmt.Println(e.Code(), e.Message())
	//			}
	//		} else {
	//			fmt.Printf("Не получилось распарсить ошибку %v", err)
	//		}
	//		log.Fatal(err)
	//	}
	//	if resp.Error == "" {
	//		fmt.Println(resp.User)
	//	} else {
	//		fmt.Println(resp.Error)
	//	}
	//}
	//
	//// получаем список email пользователей
	//emails, err := c.ListUsers(context.Background(), &pb.ListUsersRequest{
	//	Offset: 0,
	//	Limit:  100,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(emails.Count, emails.Emails)
}
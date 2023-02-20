package encryp

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/scripts"
	"net/http"
)

const (
	secretSecret = "RtsynerpoGIYdab_s234r"
	cookieName   = "access_token"
)

/*
https://practicum.yandex.ru/learn/go-developer/courses/9908027e-ac38-4005-a7c9-30f61f5ed23f/sprints/89180/topics/40590a94-b05e-46a9-8b71-3f13c57bfe86/lessons/f4fa6991-c8d9-4e92-9328-a8e234ec5867/
Чтобы шифровать данные произвольной длины, нужен алгоритм, который делил бы данные на блоки,
преобразовывал и подавал их на вход AES. Стоит взять алгоритм GCM.
Для работы алгоритма GCM нужно дополнительно сгенерировать вектор инициализации из 12 байт.
Вектор должен быть уникальным для каждой процедуры шифрования. Если переиспользовать один и
тот же вектор, можно атаковать алгоритм, подавая на вход данные с разницей в один байт, и по
косвенным признакам вычислить ключ шифрования.
*/

func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//type User struct {
//	idUser string `json:"id_user"`
//}

var u *entity.User

// EncryptionKeyCookie - middleware, которая устанавливает симметрично подписанную и зашифрованную куку
// кука устанавливается любому запросу не имеющему соответствующую куку или не прошедшая идентификацию
// в куке зашифрован, сгенерированный идентификатор пользователя
func EncryptionKeyCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := scripts.UniqueString()
		idUser := ""
		at, err := r.Cookie(cookieName)
		if err == http.ErrNoCookie {
			fmt.Printf("IDD:: %s", id)
			// создать токен
			token, err := EncryptToken(id, secretSecret)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
			//sessionLifeNanos := 10000000
			http.SetCookie(w, &http.Cookie{
				Name:  cookieName,
				Path:  "/",
				Value: token,
				//Expires: time.Now().Add(time.Nanosecond * time.Duration(sessionLifeNanos)),
			})

			idUser, err = DecryptToken(token, secretSecret)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
			ctx = context.WithValue(ctx, cookieName, idUser)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		cookie := at.Value
		idUser, err = DecryptToken(cookie, secretSecret)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		ctx = context.WithValue(ctx, cookieName, idUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// EncryptToken шифрование и подпись
func EncryptToken(data string, secretKey string) (string, error) {
	src := []byte(data) // данные, которые хотим зашифровать
	// ключ шифрования, будем использовать AES256, создав ключ длиной 32 байта (256 бит)
	key := sha256.Sum256([]byte(secretKey))
	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	// создаём вектор инициализации
	//nonce, err := generateRandom(aesgcm.NonceSize())
	nonce := key[len(key)-aesgcm.NonceSize():]
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	dst := aesgcm.Seal(nil, nonce, src, nil) // зашифровываем
	return fmt.Sprintf("%x", dst), nil
}

func DecryptToken(msg string, secretKey string) (string, error) {
	// 1) получите ключ из password, используя sha256.Sum256
	key := sha256.Sum256([]byte(secretKey))

	// 2) создайте aesblock и aesgcm
	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	// создаём вектор инициализации
	// 3) получите вектор инициализации aesgcm.NonceSize() байт с конца ключа
	nonce := key[len(key)-aesgcm.NonceSize():]

	// 4) декодируйте сообщение msg в двоичный формат
	encrypted, err := hex.DecodeString(msg)
	if err != nil {
		return "", err
	}
	// расшифровываем
	// 5) расшифруйте и выведите данные
	decrypted, err := aesgcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

// Package encryp - middleware, работает с шифрованием cookie аутентификации.
package encryp

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"net/http"

	"github.com/SETTER2000/shorturl/scripts"
)

var x interface{} = "access_token" //прочитать значение можно так: var keyToken string = x.(string)

// Encrypt -.
type Encrypt struct {
	cfg *config.Cookie
}

// EncryptionCookie Compress is a middleware that sets and encrypts authentication cookies.
func EncryptionCookie(cfg *config.Cookie) func(next http.Handler) http.Handler {
	encrypt := NewEncrypt(cfg)
	return encrypt.Handler
}

// NewEncrypt создаёт новый Encrypt, который будет обрабатывать кодированные ответы.
func NewEncrypt(cfg *config.Cookie) *Encrypt {
	return &Encrypt{
		cfg: cfg,
	}
}

// Handler returns a new middleware that will encode the response based on the
// current settings.
func (e *Encrypt) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		en := Encrypt{}
		idUser := ""
		at, err := r.Cookie("access_token")
		if err == http.ErrNoCookie {
			// создать токен
			token, err := en.EncryptToken(e.cfg.SecretKey)
			if err != nil {
				fmt.Printf("Encrypt error: %v\n", err)
			}

			http.SetCookie(w, &http.Cookie{
				Name:  "access_token",
				Path:  "/",
				Value: token,
				//Expires: time.Now().Add(time.Nanosecond * time.Duration(sessionLifeNanos)),
			})

			idUser, err = en.DecryptToken(token, e.cfg.SecretKey)
			if err != nil {
				fmt.Printf(" Decrypt error: %v\n", err)
			}
			ctx = context.WithValue(ctx, x, idUser)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		idUser, err = en.DecryptToken(at.Value, e.cfg.SecretKey)
		if err != nil {
			fmt.Printf("Decrypt token error: %v\n", err)
			// создать токен
			token, err := en.EncryptToken(e.cfg.SecretKey)
			if err != nil {
				fmt.Printf("Encrypt error: %v\n", err)
			}

			http.SetCookie(w, &http.Cookie{
				Name:  "access_token",
				Path:  "/",
				Value: token,
				//Expires: time.Now().Add(time.Nanosecond * time.Duration(sessionLifeNanos)),
			})

			idUser, err = en.DecryptToken(token, e.cfg.SecretKey)
			if err != nil {
				fmt.Printf(" Decrypt error: %v\n", err)
			}
			ctx = context.WithValue(ctx, x, idUser)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		ctx = context.WithValue(ctx, x, idUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// EncryptToken шифрование и подпись
// data - данные для кодирования
// secretKey - пароль/ключ для шифрования,
// из него создаётся ключ с помощью которого можно шифровать и расшифровать данные
// возвращает зашифрованную строку/токен
func (e *Encrypt) EncryptToken(secretKey string) (string, error) {
	if secretKey == "" {
		return "", ErrEncryptToken
	}

	data := scripts.UniqueString()

	src := []byte(data) // данные, которые хотим зашифровать
	// ключ шифрования, будем использовать AES256,
	// создав ключ длиной 32 байта (256 бит)
	key := sha256.Sum256([]byte(secretKey))
	aesBlock, _ := e.cipher(key)
	aesGSM, _ := e.gsm(aesBlock)
	// создаём вектор инициализации
	nonce := key[len(key)-aesGSM.NonceSize():]
	dst := aesGSM.Seal(nil, nonce, src, nil) // зашифровываем
	return fmt.Sprintf("%x", dst), nil
}

func (e *Encrypt) cipher(key [32]byte) (cipher.Block, error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, ErrNewCipherNotCreated
	}
	return block, nil
}

func (e *Encrypt) gsm(aes cipher.Block) (cipher.AEAD, error) {
	aesGSM, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, ErrNewGCMNotCreated
	}
	return aesGSM, nil
}

// DecryptToken расшифровать токен
// data - данные для расшифровки
// secretKey - пароль/ключ для шифрования,
// ключ с помощью которого шифровались данные
// возвращает расшифрованную строку
func (e *Encrypt) DecryptToken(data string, secretKey string) (string, error) {
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
	encrypted, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}

	// расшифровываем
	// 5) расшифруйте и выведите данные
	decrypted, err := aesgcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		fmt.Printf("Chiper фонит!\n")
		return "", err
	}
	return string(decrypted), nil
}

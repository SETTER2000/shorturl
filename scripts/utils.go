// Package scripts - вспомогательный функционал
package scripts

import (
	"fmt"
	"github.com/SETTER2000/shorturl/internal/entity"
	"math/rand"
	"os"
	"time"

	"github.com/SETTER2000/shorturl/config"
)

// GenerateString генерирует случайную строку, можно указать какой длинны.
func GenerateString(n int) string {
	// generate string
	digits := "0123456789"
	//specials := "~=+%^*/()[]{}/!@#$?|"
	specials := "_"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + digits + specials
	length := 3
	if n > length {
		length = n
	}

	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

// UniqueString возвращает уникальную строку.
// Используется как ключ короткого url.
func UniqueString() entity.Slug {
	return entity.Slug(fmt.Sprintf("%v%s", time.Now().UnixNano(), GenerateString(3)))
}

// GetHost формирует короткий URL
func GetHost(cfg config.HTTP, slug entity.Slug) entity.URL {
	return entity.URL(fmt.Sprintf("%s/%s", cfg.BaseURL, slug))
}

// CheckEnvironFlag проверка значения переменной окружения и одноименного флага
// при отсутствие переменной окружения в самой среде или пустое значение этой переменной, проверяется
// значение флага с таким же именем, по сути сама переменная окружение отсутствовать не может в системе,
// идет лишь проверка значения в двух местах в начале в окружение, затем во флаге.
func CheckEnvironFlag(environName string, flagName string) bool {
	dsn, ok := os.LookupEnv(environName)
	if !ok || dsn == "" {
		dsn = flagName
		if dsn == "" {
			fmt.Printf("connect environName string is empty: %v\n", environName)
			return false
		}
	}
	return true
}

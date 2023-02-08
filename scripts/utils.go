package scripts

import (
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"hash/fnv"
	"math/rand"
	"time"
)

func FNV32a(text string) uint32 {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(text))
	return algorithm.Sum32()
}
func GenerateString(n int) string {
	// generate string
	digits := "0123456789"
	//specials := "~=+%^*/()[]{}/!@#$?|"
	specials := "_"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + digits + specials
	length := 3
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
func UniqueString() string {
	return fmt.Sprintf("%v%s", time.Now().UnixNano(), GenerateString(3))
}

func GetHost(cfg config.HTTP, shorturl string) string {
	// Формирует короткий URL
	//_, err := os.LookupEnv("BASE_URL")
	// Если BASE_URL пустой или нет такой переменной окружения, то формирование url
	// происходит из значений, которые стоят по умолчанию в конфиге (порт отдельно, хост отдельно)
	//if err {
	//	return fmt.Sprintf("%s:%s/%s", cfg.BaseURL, cfg.Port, shorturl)
	//}
	// .. в противном случаи т.к. BASE_URL имеет формат составной
	// типа такого "http://$SERVER_HOST:$SERVER_PORT",
	// соответственно он готов к использованию, возвращаем.
	return fmt.Sprintf("%s/%s", cfg.BaseURL, shorturl)

}

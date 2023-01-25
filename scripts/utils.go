package scripts

import (
	"fmt"
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

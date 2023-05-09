package scripts_test

import (
	"fmt"
	"github.com/SETTER2000/shorturl/scripts"
)

func ExampleGenerateString() {
	// Выполняет генерацию случайной строки, используется для идентификации объектов системы.
	s := scripts.GenerateString(5)
	fmt.Println(s)
}

func ExampleUniqueString() {
	s := scripts.UniqueString()
	fmt.Println(s)
}

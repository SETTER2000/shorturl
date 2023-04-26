package scripts_test

import (
	"fmt"
	"github.com/SETTER2000/shorturl/scripts"
)

func ExampleGenerateString() {
	// Выполняет генерацию случайной строки, используется для идентификации объектов системы.
	s := scripts.GenerateString(13)
	fmt.Println(s)
}

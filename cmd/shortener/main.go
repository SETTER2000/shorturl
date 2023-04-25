// Сервис сокращения URL
//
// Приложение чистой архитектуры.
// Сервис сохраняет URL переданный на его API и возвращает укороченный URL,
// пройдя по которому получите тот же результат, что и при переходе по-длинному URL.
package main

import (
	"flag"
	"log"
	_ "net/http/pprof" // подключаем пакет pprof
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/app"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	//flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...

	// Run
	app.Run(cfg)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

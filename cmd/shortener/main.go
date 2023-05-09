// Сервис сокращения URL
//
// Приложение чистой архитектуры.
// Сервис сохраняет URL переданный на его API и возвращает укороченный URL,
// пройдя по которому получите тот же результат, что и при переходе по-длинному URL.

// Copyright 2023 Developer Team. All rights reserved
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file

package main

import (
	"flag"
	"log"
	_ "net/http/pprof" // подключаем пакет pprof
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/SETTER2000/shorturl/internal/app"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
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
	app.Run()

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

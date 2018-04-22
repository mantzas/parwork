package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"

	"github.com/mantzas/parwork"
	"github.com/mantzas/parwork/examples"
)

func main() {

	var l int

	flag.IntVar(&l, "length", 0, "the length of the string to guess")
	flag.Parse()

	if l == 0 {
		fmt.Println("length should be positive")
		flag.Usage()
		os.Exit(1)
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	started := time.Now()
	rand.Seed(started.UnixNano())
	hashed, original := examples.RandStringMD5Bytes(l)

	g := examples.NewValueGenerator(l)
	c := examples.HashCollector{Original: original, Hashed: hashed, Started: started}

	wg := sync.WaitGroup{}

	for {
		wrk := g.Generate()
		if wrk == nil {
			break
		}
		wg.Add(1)
		go func(w parwork.Work) {
			defer wg.Done()
			w.Do()
			c.Collect(w)
		}(wrk)
	}

	wg.Wait()
	fmt.Printf("every combination finished in %s\n", time.Since(started))
}
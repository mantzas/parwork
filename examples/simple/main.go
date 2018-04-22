package main

import (
	"flag"
	"fmt"
	"math/rand"
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

	started := time.Now()
	rand.Seed(started.UnixNano())
	hashed, original := examples.RandStringMD5Bytes(l)

	g := examples.NewValueGenerator(l)
	c := examples.HashCollector{Original: original, Hashed: hashed, Started: started}
	wg := sync.WaitGroup{}
	for {
		w := g.Generate()
		if w == nil {
			break
		}
		wg.Add(1)
		go func(wrk parwork.Work) {
			defer wg.Done()
			wrk.Do()
			c.Collect(wrk)
		}(w)
	}

	wg.Wait()
	fmt.Printf("every combination finished in %s\n", time.Since(started))
}

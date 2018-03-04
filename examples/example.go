package main

import (
	"bytes"
	"crypto/md5"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mantzas/parwork"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type data struct {
	hashed    []byte
	timestamp time.Time
}

// md5Work defines a structure that holds the value to be hashed and the result of the hashing
type md5Work struct {
	result    []byte
	hashed    []byte
	timestamp time.Time
}

// Do calculates the hash of the given value
func (gw *md5Work) Do() {
	gw.result = md5.New().Sum(gw.hashed)
}

// GetError returns nil since the work does not fail
func (gw *md5Work) GetError() error {
	return nil
}

// Result returns the hashed result
func (gw *md5Work) Result() interface{} {
	return gw.result
}

type valueGenerator struct {
	mask []int
}

func newValueGenerator(lenght int) *valueGenerator {
	mask := make([]int, lenght)
	mask[0] = -1
	return &valueGenerator{mask}
}

func (vg *valueGenerator) generate() parwork.Work {

	if vg.maskComplete() {
		return nil
	}

	vg.calcNextMask(0)

	w := md5Work{
		hashed:    *vg.getStringBytes(),
		timestamp: time.Now(),
	}
	fmt.Println(w)

	return &w
}

func (vg *valueGenerator) maskComplete() bool {

	count := 0

	for i := 0; i < len(vg.mask); i++ {

		if vg.mask[i] == len(letters)-1 {
			count++
		}
	}

	return len(vg.mask) == count
}

func (vg *valueGenerator) calcNextMask(index int) {

	if index >= len(vg.mask) {
		return
	}

	if vg.mask[index] < len(letters)-1 {
		vg.mask[index]++
	} else {
		vg.mask[index] = 0
		vg.calcNextMask(index + 1)
	}
}

func (vg *valueGenerator) getStringBytes() *[]byte {
	b := make([]byte, len(vg.mask))

	for i := 0; i < len(vg.mask); i++ {
		b[i] = letters[vg.mask[i]]
	}

	return &b
}

type hashCollector struct {
	original string
	hashed   []byte
}

func (hc *hashCollector) collect(w parwork.Work) {

	r := w.(*md5Work)

	if bytes.Equal(r.result, hc.hashed) {

		fmt.Printf("MATCH %s hash in %d\n", hc.original, time.Since(r.timestamp).Nanoseconds())
	}
}

func randStringMD5Bytes(n int) ([]byte, string) {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return md5.New().Sum(b), string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	var l int

	flag.IntVar(&l, "length", 0, "the lenght of the string to guess")
	flag.Parse()

	if l == 0 {
		fmt.Println("length should be positive")
		flag.Usage()
		os.Exit(1)
	}

	hashed, original := randStringMD5Bytes(l)

	g := newValueGenerator(l)
	c := hashCollector{original, hashed}

	p, err := parwork.New(g.generate, parwork.Collector(c.collect))
	if err != nil {
		fmt.Printf("failed to create processor with %v\n", err)
		os.Exit(1)
	}

	p.Process()
}

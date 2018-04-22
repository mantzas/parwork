package examples

import (
	"crypto/md5"
	"math/rand"
)

// RandStringMD5Bytes creates a random byte slice with length n
func RandStringMD5Bytes(n int) ([]byte, string) {
	b := make([]byte, n)
	for i := range b {
		b[i] = Letters[rand.Intn(len(Letters))]
	}
	return md5.New().Sum(b), string(b)
}

// MD5Work defines a structure that holds the value to be hashed and the result of the hashing
type MD5Work struct {
	Data   []byte
	Hashed []byte
}

// Do calculates the hash of the given value
func (gw *MD5Work) Do() {
	gw.Data = md5.New().Sum(gw.Hashed)
}

// Err returns nil since the work does not fail
func (gw *MD5Work) Err() error {
	return nil
}

// Result returns the hashed result
func (gw *MD5Work) Result() interface{} {
	return gw.Data
}

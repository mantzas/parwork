package examples

import (
	"bytes"
	"fmt"
	"time"

	"github.com/mantzas/parwork"
)

type HashCollector struct {
	Original string
	Hashed   []byte
	Started  time.Time
}

func (hc *HashCollector) Collect(w parwork.Work) {

	r := w.(*MD5Work)

	if bytes.Equal(r.Data, hc.Hashed) {
		fmt.Printf("MATCH %s hash in %s\n", hc.Original, time.Since(hc.Started))
	}
}

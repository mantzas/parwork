package examples

import (
	"github.com/mantzas/parwork"
)

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type valueGenerator struct {
	mask []int
}

func NewValueGenerator(length int) *valueGenerator {
	mask := make([]int, length)
	mask[0] = -1
	return &valueGenerator{mask}
}

func (vg *valueGenerator) Generate() parwork.Work {

	if vg.maskComplete() {
		return nil
	}

	vg.calcNextMask(0)

	w := MD5Work{
		Hashed: *vg.getStringBytes(),
	}

	return &w
}

func (vg *valueGenerator) maskComplete() bool {

	count := 0

	for i := 0; i < len(vg.mask); i++ {

		if vg.mask[i] == len(Letters)-1 {
			count++
		}
	}

	return len(vg.mask) == count
}

func (vg *valueGenerator) calcNextMask(index int) {

	if index >= len(vg.mask) {
		return
	}

	if vg.mask[index] < len(Letters)-1 {
		vg.mask[index]++
	} else {
		vg.mask[index] = 0
		vg.calcNextMask(index + 1)
	}
}

func (vg *valueGenerator) getStringBytes() *[]byte {
	b := make([]byte, len(vg.mask))

	for i := 0; i < len(vg.mask); i++ {
		b[i] = Letters[vg.mask[i]]
	}

	return &b
}

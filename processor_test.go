package parwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		g       WorkGenerator
		options []Option
	}
	p, _ := New(generator)
	tests := []struct {
		name    string
		args    args
		want    *Processor
		wantErr bool
	}{
		{"failure due to nil generator", args{nil, nil}, nil, true},
		{"failure due to error option", args{generator, []Option{Collector(nil)}}, nil, true},
		{"success", args{generator, nil}, p, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.g, tt.args.options...)
			if tt.wantErr {
				assert.Errorf(err, "New() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Equal(got.workers, tt.want.workers)
				assert.Equal(got.queue, tt.want.queue)
				assert.NotNil(got.generator)
				assert.NotNil(got.reporter)
			}
		})
	}
}

func TestProcessor_Process(t *testing.T) {
	assert := assert.New(t)
	max := 100
	gen := testWorkGenerator{max: max}
	col := testCollector{}

	p, err := New(gen.Generate, Collector(col.Collect))

	assert.NoError(err)
	assert.NotPanics(p.Process)
	assert.Len(col.results, max+1)
}

func BenchmarkProcessor_Process(b *testing.B) {

	max := 1000
	gen := testWorkGenerator{max: max}
	col := testCollector{}

	p, _ := New(gen.Generate, Collector(col.Collect))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		p.Process()
	}
}

type testWork struct {
	err      error
	previous int
	current  int
	result   int
}

func (w *testWork) Do() {

	w.result = w.previous + w.current
}

func (w *testWork) Err() error { return w.err }

func (w *testWork) Result() interface{} { return w.current }

type testWorkGenerator struct {
	current int
	max     int
}

func (twg *testWorkGenerator) Generate() Work {

	if twg.current > twg.max {
		return nil
	}

	var w testWork
	if twg.current == 0 {
		w.previous = 0
		w.current = 0
	} else {
		w.previous = twg.current - 1
		w.current = twg.current
	}

	twg.current++

	return &w
}

type testCollector struct {
	results []int
}

func (fc *testCollector) Collect(w Work) {

	fc.results = append(fc.results, w.Result().(int))
}

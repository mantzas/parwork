package processor

import "github.com/mantzas/parwork"

type testWork struct {
	err      error
	previous int
	current  int
	result   int
}

func (w *testWork) Do() {

	w.result = w.previous + w.current
}

func (w *testWork) GetError() error { return w.err }

func (w *testWork) Result() interface{} { return w.current }

type testWorkGenerator struct {
	current int
	max     int
}

func (twg *testWorkGenerator) Generate() parwork.Work {

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

func (fc *testCollector) Collect(w parwork.Work) {

	fc.results = append(fc.results, w.Result().(int))
}

func generator() parwork.Work {
	return &testWork{}
}

func reporter(w parwork.Work) {
}

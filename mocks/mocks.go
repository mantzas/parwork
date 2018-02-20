package mocks

import "github.com/mantzas/parwork"

type MockWork struct{}

func (t MockWork) ID() string      { return "1" }
func (t MockWork) Do()             {}
func (t MockWork) GetError() error { return nil }

func Generator() parwork.Work {
	return MockWork{}
}

func Reporter(w parwork.Work) {
}

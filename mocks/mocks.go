package mocks

import "github.com/mantzas/parwork"

// MockWork definition
type MockWork struct{}

// ID returns a ID
func (t MockWork) ID() string { return "1" }

// Do executes the work
func (t MockWork) Do() {}

// GetError returns the work's error
func (t MockWork) GetError() error { return nil }

// Generator mock for work
func Generator() parwork.Work {
	return MockWork{}
}

// Reporter mock for work
func Reporter(w parwork.Work) {
}

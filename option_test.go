package parwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkers(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		count     int
		want      int
		wantError bool
	}{
		{"success with 1 worker", 1, 1, false},
		{"fails with 0 worker", 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p, err := New(generator, Workers(tt.count))
			if tt.wantError {
				assert.NotNil(err, "Worker error = %v, wantErr %v", err, tt.wantError)
			} else {
				assert.Equal(p.workers, tt.want, "workers = %v, want %v", p.workers, tt.want)
			}
		})
	}
}

func TestQueue(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		length    int
		want      int
		wantError bool
	}{
		{"success with 1 queue length", 1, 1, false},
		{"fails with 0 queue length", 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p, err := New(generator, Queue(tt.length))
			if tt.wantError {
				assert.NotNil(err, "Queue error = %v, wantErr %v", err, tt.wantError)
			} else {
				assert.NotEqual(p.workers, tt.want, "queue length = %v, want %v", p.workers, tt.want)
			}
		})
	}
}

func TestReporter(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		collector WorkCollector
		wantError bool
	}{
		{"success", reporter, false},
		{"fails with nil reporter", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(generator, Collector(tt.collector))
			if tt.wantError {
				assert.NotNil(err, "reporter error = %v, wantErr %v", err, tt.wantError)
			} else {
				assert.Nil(nil, "reporter is nil but wanted not nil")
			}
		})
	}
}

func generator() Work {
	return &testWork{}
}

func reporter(w Work) {
}

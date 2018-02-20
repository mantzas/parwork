package parwork

import (
	"sync"
	"testing"
)

type TestWork struct{}

func (t TestWork) ID() string      { return "1" }
func (t TestWork) Do()             {}
func (t TestWork) GetError() error { return nil }

func testGenerator() Work {
	return TestWork{}
}

func testReporter(w Work) {

}

func TestNew(t *testing.T) {
	type args struct {
		g       WorkGenerator
		options []ProcessorOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"success", args{testGenerator, nil}, false},
		{"failed with nil generator", args{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.g, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWorkers(t *testing.T) {
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

			p, err := New(testGenerator, Workers(tt.count))

			if tt.wantError && err == nil {
				t.Errorf("Worker error = %v, wantErr %v", err, tt.wantError)
			}

			if !tt.wantError {
				if p.workers != tt.want {
					t.Errorf("workers = %v, want %v", p.workers, tt.want)
				}
			}
		})
	}
}

func TestQueue(t *testing.T) {
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

			p, err := New(testGenerator, Queue(tt.length))

			if tt.wantError && err == nil {
				t.Errorf("Queue error = %v, wantErr %v", err, tt.wantError)
			}

			if !tt.wantError {
				if p.queue != tt.want {
					t.Errorf("queue length = %v, want %v", p.workers, tt.want)
				}
			}
		})
	}
}

func TestReporter(t *testing.T) {
	type args struct {
		reporter WorkReporter
	}
	tests := []struct {
		name      string
		reporter  WorkReporter
		wantError bool
	}{
		{"success", testReporter, false},
		{"fails with nil reporter", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p, err := New(testGenerator, Reporter(tt.reporter))

			if tt.wantError && err == nil {
				t.Errorf("reporter error = %v, wantErr %v", err, tt.wantError)
			}

			if !tt.wantError {
				if p.reporter == nil {
					t.Error("reporter is nil but wanted not nil")
				}
			}
		})
	}
}

func TestProcessor_Process(t *testing.T) {
	tests := []struct {
		name string
		p    Processor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Process()
		})
	}
}

func TestProcessor_startWorkers(t *testing.T) {
	type args struct {
		wg   *sync.WaitGroup
		q    <-chan Work
		repQ chan<- Work
	}
	tests := []struct {
		name string
		p    Processor
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.startWorkers(tt.args.wg, tt.args.q, tt.args.repQ)
		})
	}
}

func TestProcessor_startReporter(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
		q  <-chan Work
	}
	tests := []struct {
		name string
		p    Processor
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.startReporter(tt.args.wg, tt.args.q)
		})
	}
}

func TestProcessor_startGenerator(t *testing.T) {
	type args struct {
		q chan<- Work
	}
	tests := []struct {
		name string
		p    Processor
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.startGenerator(tt.args.q)
		})
	}
}

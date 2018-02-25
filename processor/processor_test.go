package processor

import (
	"sync"
	"testing"

	"github.com/mantzas/parwork"
	"github.com/mantzas/parwork/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		g       parwork.WorkGenerator
		options []Option
	}
	p, _ := New(mocks.Generator)
	tests := []struct {
		name    string
		args    args
		want    *Processor
		wantErr bool
	}{
		{"failure due to nil generator", args{nil, nil}, nil, true},
		{"failure due to error option", args{mocks.Generator, []Option{Reporter(nil)}}, nil, true},
		{"success", args{mocks.Generator, nil}, p, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.g, tt.args.options...)
			if tt.wantErr {
				assert.Error(err, "New() error = %v, wantErr %v", err, tt.wantErr)
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
		q    <-chan parwork.Work
		repQ chan<- parwork.Work
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
			tt.p.bootstrapWorkers(tt.args.wg, tt.args.q, tt.args.repQ)
		})
	}
}

func TestProcessor_startReporter(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
		q  <-chan parwork.Work
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
			tt.p.bootstrapReporter(tt.args.wg, tt.args.q)
		})
	}
}

func TestProcessor_startGenerator(t *testing.T) {
	type args struct {
		q chan<- parwork.Work
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
			tt.p.generateWork(tt.args.q)
		})
	}
}

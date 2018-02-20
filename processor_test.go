package parwork

import (
	"reflect"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		g       WorkGenerator
		options []ProcessorOption
	}
	tests := []struct {
		name    string
		args    args
		want    *Processor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.g, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
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

package parwork

import "testing"

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

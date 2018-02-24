package processor

import (
	"testing"

	"github.com/mantzas/parwork"
	"github.com/stretchr/testify/require"

	"github.com/mantzas/parwork/mocks"
)

func TestWorkers(t *testing.T) {
	require := require.New(t)
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

			p, err := New(mocks.Generator, Workers(tt.count))
			if tt.wantError {
				require.NotNil(err, "Worker error = %v, wantErr %v", err, tt.wantError)
			} else {
				require.Equal(p.workers, tt.want, "workers = %v, want %v", p.workers, tt.want)
			}
		})
	}
}

func TestQueue(t *testing.T) {
	require := require.New(t)
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

			p, err := New(mocks.Generator, Queue(tt.length))
			if tt.wantError {
				require.NotNil(err, "Queue error = %v, wantErr %v", err, tt.wantError)
			} else {
				require.NotEqual(p.workers, tt.want, "queue length = %v, want %v", p.workers, tt.want)
			}
		})
	}
}

func TestReporter(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		name      string
		reporter  parwork.WorkReporter
		wantError bool
	}{
		{"success", mocks.Reporter, false},
		{"fails with nil reporter", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := New(mocks.Generator, Reporter(tt.reporter))
			if tt.wantError {
				require.NotNil(err, "reporter error = %v, wantErr %v", err, tt.wantError)
			} else {
				require.Nil(nil, "reporter is nil but wanted not nil")
			}
		})
	}
}

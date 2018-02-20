package processor

import (
	"errors"

	"github.com/mantzas/parwork"
)

// Option defines a option for the processor
type Option func(*Processor) error

// Workers defines a processor option for the workers
func Workers(count int) Option {
	return func(p *Processor) error {
		if count <= 0 {
			return errors.New("worker count must be positive")
		}
		p.workers = count
		return nil
	}
}

// Queue defines a processor option for the queue length
func Queue(length int) Option {
	return func(p *Processor) error {
		if length <= 0 {
			return errors.New("queue length must be positive")
		}
		p.queue = length
		return nil
	}
}

// Reporter defines a processor option for the work reporter
func Reporter(reporter parwork.WorkReporter) Option {
	return func(p *Processor) error {
		if reporter == nil {
			return errors.New("reporter is nil")
		}
		p.reporter = reporter
		return nil
	}
}

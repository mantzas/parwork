package processor

import (
	"errors"
	"fmt"
	"runtime"
	"sync"

	"github.com/mantzas/parwork"
)

// Processor handles the generation, distribution and reporting or work
type Processor struct {
	workers   int
	queue     int
	generator parwork.WorkGenerator
	reporter  parwork.WorkReporter
}

// New returns a new work processor with default worker, queue length and reporter.
// Optional definitions are available through the processor options variadic arguments.
// In case of a
// Workers: Number of CPU
// Queue: Number of CPU * 100
// Reporter: output to stdout
func New(g parwork.WorkGenerator, options ...Option) (*Processor, error) {

	if g == nil {
		return nil, errors.New("generator is nil")
	}

	p := &Processor{
		workers:   runtime.NumCPU(),
		queue:     runtime.NumCPU() * 100,
		generator: g,
		reporter: func(w parwork.Work) {
			fmt.Println(w)
		},
	}

	for _, opt := range options {
		err := opt(p)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

// Process begins the parallel processing of work
func (p Processor) Process() {
	pending := make(chan parwork.Work, p.queue)
	done := make(chan parwork.Work, p.queue)
	workers := sync.WaitGroup{}
	collector := sync.WaitGroup{}

	p.bootstrapWorkers(&workers, pending, done)
	p.bootstrapReporter(&collector, done)
	p.generateWork(pending)

	close(pending)
	workers.Wait()
	close(done)
	collector.Wait()
}

func (p Processor) bootstrapWorkers(wg *sync.WaitGroup, pending <-chan parwork.Work, done chan<- parwork.Work) {
	wCount := 0
	for wCount < p.workers {
		wg.Add(1)
		go func() {
			for work := range pending {
				work.Do()
				done <- work
			}
			wg.Done()
		}()
		wCount++
	}
}

func (p Processor) bootstrapReporter(wg *sync.WaitGroup, done <-chan parwork.Work) {
	wg.Add(1)
	go func() {
		for work := range done {
			p.reporter(work)
		}
		wg.Done()
	}()
}

func (p Processor) generateWork(pending chan<- parwork.Work) {
	for {
		work := p.generator()
		if work == nil {
			break
		} else {
			pending <- work
		}
	}
}

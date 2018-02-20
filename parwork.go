package parwork

import (
	"fmt"
	"runtime"
	"sync"
)

// Work define the interface that each work item has to implement in order to be processed.
// This implementation follows the command pattern.
type Work interface {
	ID() string
	Do()
	GetError() error
}

// WorkGenerator defines a function that generates work.
// Every time the function is called it will return work or nil which signals the end of the work generation.
type WorkGenerator func() Work

// WorkReporter defines a function that handles the reporting of a completed work.
type WorkReporter func(Work)

// Processor handles the generation, distribution and reporting or work
type Processor struct {
	workers   int
	queue     int
	generator WorkGenerator
	reporter  WorkReporter
}

// New returns a new work processor with default worker, queue length and reporter.
// Optional definitions are available through the processor options variadic arguments.
// Workers: Number of CPU
// Queue: Number of CPU * 100
// Reporter: output to stdout
func New(g WorkGenerator, options ...ProcessorOption) *Processor {
	p := &Processor{
		workers:   runtime.NumCPU(),
		queue:     runtime.NumCPU() * 100,
		generator: g,
		reporter: func(w Work) {
			fmt.Println(w)
		},
	}

	for _, opt := range options {
		opt(p)
	}

	return p
}

// ProcessorOption defines a option for the processor
type ProcessorOption func(*Processor)

// Workers defines a processor option for the workers
func Workers(count int) ProcessorOption {
	return func(p *Processor) {
		p.workers = count
	}
}

// Queue defines a processor option for the queue length
func Queue(length int) ProcessorOption {
	return func(p *Processor) {
		p.queue = length
	}
}

// Reporter defines a processor option for the work reporter
func Reporter(reporter WorkReporter) ProcessorOption {
	return func(p *Processor) {
		p.reporter = reporter
	}
}

// Process begins the parallel processing of work
func (p Processor) Process() {
	workerQ := make(chan Work, p.queue)
	reporterQ := make(chan Work, p.queue)
	wgWorker := sync.WaitGroup{}
	wgCollector := sync.WaitGroup{}

	p.startWorkers(&wgWorker, workerQ, reporterQ)
	p.startReporter(&wgCollector, reporterQ)
	p.startGenerator(workerQ)

	close(workerQ)
	wgWorker.Wait()
	close(reporterQ)
	wgCollector.Wait()
}

func (p Processor) startWorkers(wg *sync.WaitGroup, q <-chan Work, repQ chan<- Work) {
	wCount := 0
	for wCount < p.workers {
		wg.Add(1)
		go func() {
			for work := range q {
				work.Do()
				repQ <- work
			}
			wg.Done()
		}()
		wCount++
	}
}

func (p Processor) startReporter(wg *sync.WaitGroup, q <-chan Work) {
	wg.Add(1)
	go func() {
		for work := range q {
			p.reporter(work)
		}
		wg.Done()
	}()
}

func (p Processor) startGenerator(q chan<- Work) {
	for {
		work := p.generator()
		if work == nil {
			break
		} else {
			q <- work
		}
	}
}

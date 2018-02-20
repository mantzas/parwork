package parwork

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
)

// Processor handles the generation, distribution and reporting or work
type Processor struct {
	workers   int
	queue     int
	generator WorkGenerator
	reporter  WorkReporter
}

// New returns a new work processor with default worker, queue length and reporter.
// Optional definitions are available through the processor options variadic arguments.
// In case of a
// Workers: Number of CPU
// Queue: Number of CPU * 100
// Reporter: output to stdout
func New(g WorkGenerator, options ...ProcessorOption) (*Processor, error) {

	if g == nil {
		return nil, errors.New("generator is nil")
	}

	p := &Processor{
		workers:   runtime.NumCPU(),
		queue:     runtime.NumCPU() * 100,
		generator: g,
		reporter: func(w Work) {
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

package parwork

import (
	"sync"
)

// TODO: this might be a command of the command pattern

// Work interface
type Work interface {
	ID() string
	Do()
	GetError() error
}

// WorkGenerator defines a function that generates work.
// Every time the function is called it will return
// work or nil which signals the end of the work generation
type WorkGenerator func() Work

// WorkReporter defines a function that handles the reporting of completed Work
type WorkReporter func(w Work)

// Processor processes work that is generated by
type Processor struct {
	workers   int
	buffer    int
	generator WorkGenerator
	reporter  WorkReporter
}

// New returns a new work processor
func New() *Processor {
	return &Processor{}
}

// Process begins the parallel processing of work
func (c Processor) Process() {
	workerQ := make(chan Work, c.buffer)
	reporterQ := make(chan Work, c.buffer)
	wgWorker := sync.WaitGroup{}
	wgCollector := sync.WaitGroup{}

	c.startWorkers(&wgWorker, workerQ, reporterQ)
	c.startReporter(&wgCollector, reporterQ)
	c.startGenerator(workerQ)

	close(workerQ)
	wgWorker.Wait()
	close(reporterQ)
	wgCollector.Wait()
}

func (c Processor) startWorkers(wg *sync.WaitGroup, workerQ <-chan Work, reporterQ chan<- Work) {

	wCount := 0
	for wCount < c.workers {

		wg.Add(1)
		go func() {

			for work := range workerQ {
				work.Do()
				reporterQ <- work
			}
			wg.Done()
		}()
		wCount++
	}
}

func (c Processor) startReporter(wg *sync.WaitGroup, q <-chan Work) {
	wg.Add(1)
	go func() {
		for work := range q {
			c.reporter(work)
		}
		wg.Done()
	}()
}

func (c Processor) startGenerator(q chan<- Work) {
	for {
		work := c.generator()
		if work == nil {
			break
		} else {
			q <- work
		}
	}
}

package parwork

// Work define the interface that each work item has to implement in order to be processed.
// This implementation follows the command pattern.
type Work interface {
	Do()
	GetError() error
}

// WorkGenerator defines a function that generates work.
// Every time the function is called it will return work or nil which signals the end of the work generation.
type WorkGenerator func() Work

// WorkReporter defines a function that handles the reporting of a completed work.
type WorkReporter func(Work)

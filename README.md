# parwork [![Build Status](https://travis-ci.org/mantzas/parwork.svg?branch=master)](https://travis-ci.org/mantzas/parwork) [![codecov](https://codecov.io/gh/mantzas/parwork/branch/master/graph/badge.svg)](https://codecov.io/gh/mantzas/parwork) [![Go Report Card](https://goreportcard.com/badge/github.com/mantzas/parwork)](https://goreportcard.com/report/github.com/mantzas/parwork) [![GoDoc](https://godoc.org/github.com/mantzas/parwork?status.svg)](https://godoc.org/github.com/mantzas/parwork)

Parallel work processing package

## Description

This package allows work to processed in parallel using a fan-out/fan-in pattern. The implementation relies on goroutines, channels and wait groups.

The implementation allows the configuration of the processor by providing the degree of parallelism which defines how many goroutines will process work from the queues in parallel. By creating only a small number of goroutines, which defaults to the number of cores on the system, we avoid context switching instead of allowing a high number of goroutines to fight for processor resources.

This is the reason why this package makes more sense when used with work items that are CPU bound and do not switch context like waiting for IO.

In order to use the package the user has only to provide the implementation of the following:

### Work interface

```go
type Work interface {
    Do()
    GetError() error
    Result() interface{}
}
```

The work interface defines a method ```Do()``` which contains all the processing logic of the work item. The ```GetError() error``` method can be used to flag the work item as failed and return a error. The ```Result() interface{}``` defines a method which returns the result of the work. Due to the lack of generics the data return has to be cast from ```interface{}``` to the actual result type in order to be usable in the WorkCollector.

### WorkGenerator function

```go
type WorkGenerator func() Work
```

The WorkGenerator function allows the user to provide a implementation that returns on each call a work item to be processed. If the generator returns ```nil``` the generation of work has finished.

### WorkCollector function

```go
type WorkCollector func(Work)
```

The WorkCollector function takes as a argument a completed Work item. It can check for a failure by calling the ```GetError``` or the ```Result``` method of the Work item and handle it appropriately.

### Example

For a example implementation please take a look in the examples folder of the repository. The example implements a brute force method of trying to find the MD5 hash of a string. This is just a example implementation to demonstrate the usage of the package. And it should not be misused to break secrets.

There are 3 example folder:

- linear, which processes each work in a linear fashion
- simple, which processes each work in a separate go routine
- parwork, which processes the work in a pipeline

Running the above gave us the following result (length=5)

- linear, 7m57.544370159s
- simple, 9m23.536487533s
- parwork, 7m48.694807617s

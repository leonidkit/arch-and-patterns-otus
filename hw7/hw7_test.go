package hw7

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type WorkerTestSuite struct {
	suite.Suite
	queue  chan ICommand
	worker *worker
}

func (suite *WorkerTestSuite) SetupTest() {
	suite.queue = make(chan ICommand, 1)
	suite.worker = NewWorker(suite.queue)

	startCommand := NewStartCommand(suite.worker)
	err := startCommand.Execute()
	assertNoError(err)

	if !suite.worker.IsWorking() {
		panic("goroutine should work")
	}
}

func (suite *WorkerTestSuite) TestHardStopCommand() {
	if !suite.worker.IsWorking() {
		panic("goroutine should work")
	}

	hardStopCommand := NewHardStopCommand(suite.worker)
	cmd := &command{
		f: func() error {
			return nil
		},
		duration: 1 * time.Second,
	}
	suite.queue <- hardStopCommand
	suite.queue <- cmd

	time.Sleep(100 * time.Millisecond)

	assertEqual(len(suite.queue), 1)
	assertEqual(suite.worker.IsWorking(), false)
}

func (suite *WorkerTestSuite) TestSoftStopCommand() {
	if !suite.worker.IsWorking() {
		panic("goroutine should work")
	}

	softStopCommand := NewSoftStopCommand(suite.worker)
	cmd := &command{
		f: func() error {
			return nil
		},
		duration: 1 * time.Millisecond,
	}
	suite.queue <- softStopCommand
	suite.queue <- cmd

	time.Sleep(2 * time.Millisecond)

	assertEqual(len(suite.queue), 0)
	assertEqual(suite.worker.IsWorking(), false)
}

func TestFlow(t *testing.T) {
	suite.Run(t, new(WorkerTestSuite))
}

type command struct {
	f        func() error
	duration time.Duration
}

func (c *command) Execute() error {
	select {
	case <-time.After(c.duration):
		return c.f()
	}
}

func assertNoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("err not nil: %s", err))
	}
}

func assertEqual[T comparable](got T, want T) {
	if got != want {
		panic(fmt.Sprintf("not equal: got = %v, want = %v", got, want))
	}
}

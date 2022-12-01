package hw7

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type WorkerTestSuite struct {
	suite.Suite
	queue           chan ICommand
	queueToEvacuate chan ICommand
	worker          *worker
}

func (suite *WorkerTestSuite) SetupTest() {
	suite.queue = make(chan ICommand, 1)
	suite.queueToEvacuate = make(chan ICommand, 100)
	suite.worker = NewWorker(suite.queue, suite.queueToEvacuate)

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

	hardStopCommand := NewHardStopCommand()
	cmd := &command{
		f: func() error {
			return nil
		},
		duration: 1 * time.Second,
	}
	suite.queue <- hardStopCommand
	suite.queue <- cmd

	time.Sleep(400 * time.Millisecond)

	assertEqual(len(suite.queue), 1)
	assertEqual(suite.worker.IsWorking(), false)
}

func (suite *WorkerTestSuite) TestSoftStopCommand() {
	if !suite.worker.IsWorking() {
		panic("goroutine should work")
	}

	softStopCommand := NewSoftStopCommand()
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

func (suite *WorkerTestSuite) TestHardStopCommandMoveTo() {
	if !suite.worker.IsWorking() {
		panic("goroutine should work")
	}

	moveToCommand := &MoveToCommand{}
	hardStopCommand := NewHardStopCommand()
	cmd := &command{
		f: func() error {
			return nil
		},
		duration: 1 * time.Second,
	}
	suite.queue <- moveToCommand
	suite.queue <- hardStopCommand
	suite.queue <- cmd

	time.Sleep(400 * time.Millisecond)

	assertEqual(len(suite.queue), 1)
	assertEqual(suite.worker.IsWorking(), false)
}

func (suite *WorkerTestSuite) TestSoftStopCommandMoveTo() {
	if !suite.worker.IsWorking() {
		panic("goroutine should work")
	}

	moveToCommand := &MoveToCommand{}
	softStopCommand := NewSoftStopCommand()
	cmd := &command{
		f: func() error {
			return nil
		},
		duration: 1 * time.Millisecond,
	}
	suite.queue <- moveToCommand
	suite.queue <- cmd
	suite.queue <- cmd
	suite.queue <- softStopCommand

	time.Sleep(4 * time.Millisecond)

	assertEqual(len(suite.queue), 0)
	assertEqual(len(suite.queueToEvacuate), 2)
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

package hw7

import (
	"fmt"
)

type state interface {
	handle() state
}

type simpleState struct {
	w *worker
}

func (ss *simpleState) handle() state {
	cmd, ok := <-ss.w.queue
	if !ok {
		return nil
	}
	if _, ok := cmd.(*MoveToCommand); ok {
		return ss.w.moveToState
	}
	if _, ok := cmd.(*SoftStopCommand); ok {
		for i := 0; i < len(ss.w.queue); i++ {
			executeCommand(<-ss.w.queue)
		}
		return nil
	}
	if _, ok := cmd.(*HardStopCommand); ok {
		return nil
	}
	executeCommand(cmd)
	return ss
}

type moveToState struct {
	w *worker
}

func (mts *moveToState) handle() state {
	cmd, ok := <-mts.w.queue
	if !ok {
		return nil
	}
	if _, ok := cmd.(*RunCommand); ok {
		return mts.w.simpleState
	}
	if _, ok := cmd.(*SoftStopCommand); ok {
		for i := 0; i < len(mts.w.queue); i++ {
			mts.w.queueToEvacuate <- <-mts.w.queue
		}
		return nil
	}
	if _, ok := cmd.(*HardStopCommand); ok {
		return nil
	}
	mts.w.queueToEvacuate <- cmd
	return mts
}

type worker struct {
	simpleState state
	moveToState state
	currState   state

	working         bool
	softstop        chan struct{}
	hardstop        chan struct{}
	queue           <-chan ICommand
	queueToEvacuate chan<- ICommand
}

func NewWorker(queue <-chan ICommand, queueToEvacuate chan<- ICommand) *worker {
	w := &worker{
		softstop:        make(chan struct{}, 1),
		hardstop:        make(chan struct{}, 1),
		queue:           queue,
		queueToEvacuate: queueToEvacuate,
	}

	ss := &simpleState{w: w}
	ms := &moveToState{w: w}

	w.currState = ss
	w.simpleState = ss
	w.moveToState = ms

	return w
}

func (w *worker) Start() {
	go func() {
		defer func() {
			w.working = false
		}()
		for {
			w.currState = w.currState.handle()
			if w.currState == nil {
				return
			}
		}
	}()
	w.working = true
}

func (w *worker) IsWorking() bool {
	return w.working
}

func executeCommand(cmd ICommand) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic happend: %v", r)
		}
	}()
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("error occured: %s", err)
	}
}

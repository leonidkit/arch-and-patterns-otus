package hw7

import (
	"fmt"
)

type worker struct {
	working  bool
	softstop chan struct{}
	hardstop chan struct{}
	queue    <-chan ICommand
}

func NewWorker(queue <-chan ICommand) *worker {
	return &worker{
		softstop: make(chan struct{}, 1),
		hardstop: make(chan struct{}, 1),
		queue:    queue,
	}
}

func (w *worker) Start() {
	go func() {
		defer func() {
			w.working = false
		}()
		for {
			select {
			case <-w.hardstop:
				return
			default:
			}

			select {
			case <-w.softstop:
				for i := 0; i < len(w.queue); i++ {
					executeCommand(<-w.queue)
				}
				return
			case <-w.hardstop:
				return
			case cmd, ok := <-w.queue:
				if !ok {
					return
				}
				executeCommand(cmd)
			}
		}
	}()
	w.working = true
}

func (w *worker) SoftStop() {
	w.softstop <- struct{}{}
}

func (w *worker) HardStop() {
	w.hardstop <- struct{}{}
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

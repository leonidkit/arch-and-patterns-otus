package hw3

import (
	"fmt"
)

type ICommand interface {
	Execute() error
	Name() string
}

//go:generate mockgen -source=hw3.go -destination=mock.go -package=${GOPACKAGE}
type Queue interface {
	Push(ICommand) error
	Pop() (ICommand, error)
	IsEmpty() bool
}

func QueueProcessing(queue Queue, eh *ErrorHandler) (err error) {
	for !queue.IsEmpty() {
		cmd, err := queue.Pop()
		if err != nil {
			return fmt.Errorf("queue Pop() command error: %w", err)
		}

		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic handled: %w", err)
			}
		}()

		err = cmd.Execute()
		if err != nil {
			err = queue.Push(eh.Handle(cmd, err))
			if err != nil {
				return fmt.Errorf("queue Push() command error: %w", err)
			}
		}
	}
	return nil
}

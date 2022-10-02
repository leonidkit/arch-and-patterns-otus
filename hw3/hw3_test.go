package hw3

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestQueueProcessing(t *testing.T) {
	t.Run("first strategy test", func(t *testing.T) {
		eh := NewErrorHandler()

		eh.Setup(
			CommandName,
			errConnectionError,
			func(cmd ICommand, err error) ICommand {
				return NewRepeatCommand(cmd, err)
			},
		)
		eh.Setup(
			RepeatCommandName,
			errConnectionError,
			func(cmd ICommand, err error) ICommand {
				return NewLogCommand(cmd, err)
			},
		)
		eh.Setup(
			LogCommandName,
			errConnectionError,
			func(cmd ICommand, err error) ICommand {
				return nil
			},
		)

		ctrl := gomock.NewController(t)

		cmdMock := NewMockICommand(ctrl)
		cmdMock.EXPECT().Execute().Return(errConnectionError).Times(2)
		cmdMock.EXPECT().Name().Return(CommandName)

		queueMock := new(queue)
		queueMock.Push(cmdMock)

		err := QueueProcessing(queueMock, eh)
		if err != nil {
			panic("unexpected error")
		}
	})

	t.Run("second strategy test", func(t *testing.T) {
		eh := NewErrorHandler()

		eh.Setup(
			CommandName,
			errConnectionError,
			func(cmd ICommand, err error) ICommand {
				return NewRepeatCommand(cmd, err)
			},
		)
		eh.Setup(
			RepeatCommandName,
			errConnectionError,
			func(cmd ICommand, err error) ICommand {
				return NewRepeatCommandTwo(cmd, err)
			},
		)
		eh.Setup(
			RepeatTwoCommandName,
			errConnectionError,
			func(cmd ICommand, err error) ICommand {
				return NewLogCommand(cmd, err)
			},
		)
		eh.Setup(
			LogCommandName,
			errConnectionError,
			func(cmd ICommand, err error) ICommand {
				return nil
			},
		)

		ctrl := gomock.NewController(t)

		cmdMock := NewMockICommand(ctrl)
		cmdMock.EXPECT().Execute().Return(errConnectionError).Times(3)
		cmdMock.EXPECT().Name().Return(CommandName)

		queueMock := new(queue)
		queueMock.Push(cmdMock)

		err := QueueProcessing(queueMock, eh)
		if err != nil {
			panic("unexpected error")
		}
	})

	t.Run("wrong error strategy test", func(t *testing.T) {
		eh := NewErrorHandler()

		eh.Setup(
			CommandName,
			errConnectionError,
			func(cmd ICommand, err error) ICommand {
				return NewRepeatCommand(cmd, err)
			},
		)

		ctrl := gomock.NewController(t)

		cmdMock := NewMockICommand(ctrl)
		cmdMock.EXPECT().Execute().Return(errors.New("new type of error"))
		cmdMock.EXPECT().Name().Return(CommandName)

		queueMock := new(queue)
		queueMock.Push(cmdMock)

		err := QueueProcessing(queueMock, eh)
		if err == nil {
			panic("expected error")
		}
	})
}

type queue struct {
	head *node
}

type node struct {
	cmd  ICommand
	prev *node
}

func (q *queue) Push(cmd ICommand) error {
	if cmd == nil {
		return nil
	}
	if q.head == nil {
		q.head = &node{cmd: cmd}
	} else {
		n := &node{
			cmd:  cmd,
			prev: q.head,
		}
		q.head = n
	}
	return nil
}

func (q *queue) Pop() (ICommand, error) {
	if q.head == nil {
		return nil, nil
	} else {
		r := q.head.cmd
		q.head = q.head.prev
		return r, nil
	}
}

func (q *queue) IsEmpty() bool {
	return q.head == nil
}

func TestQueue(t *testing.T) {
	var q queue

	cmd := NewCommand()
	logCmd := NewLogCommand(nil, nil)
	err := q.Push(cmd)
	if err != nil {
		panic("unexpected error")
	}
	err = q.Push(logCmd)
	if err != nil {
		panic("unexpected error")
	}

	c, err := q.Pop()
	if err != nil {
		panic("unexpected error")
	}

	if c != logCmd {
		panic("want to be equal")
	}
}

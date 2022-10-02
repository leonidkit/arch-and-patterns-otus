package hw3

import "errors"

var (
	errConnectionError = errors.New("connection error")
	errTimeoutError    = errors.New("timeout error")
)

type ErrorHandler struct {
	handlers map[string]map[string]func(cmd ICommand, err error) ICommand
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		handlers: make(map[string]map[string]func(cmd ICommand, err error) ICommand),
	}
}

func (eh *ErrorHandler) Setup(cmdName string, err error, cmd func(cmd ICommand, err error) ICommand) {
	if _, ok := eh.handlers[cmdName]; !ok {
		eh.handlers[cmdName] = make(map[string]func(cmd ICommand, err error) ICommand)
	}
	eh.handlers[cmdName][err.Error()] = cmd
}

func (eh *ErrorHandler) Handle(cmd ICommand, err error) ICommand {
	return eh.handlers[cmd.Name()][err.Error()](cmd, err)
}

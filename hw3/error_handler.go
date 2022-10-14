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
	if cmd == nil || err == nil {
		panic("invalid arguments")
	}

	cmdName := cmd.Name()
	errStr := err.Error()

	if _, ok := eh.handlers[cmdName]; !ok {
		panic("unknown command name: " + cmdName)
	}
	if _, ok := eh.handlers[cmdName][errStr]; !ok {
		// TODO: логировать ошибки, для дальнейшей аналитики и создания нового обработчика.
		panic("unknown error: " + errStr)
	}

	return eh.handlers[cmdName][errStr](cmd, err)
}

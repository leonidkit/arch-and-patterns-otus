package hw3

import (
	"log"
)

var (
	CommandName          = "command"
	LogCommandName       = "log_command"
	RepeatCommandName    = "repeat_command"
	RepeatTwoCommandName = "repeat_two_command"
)

type Command struct{}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string {
	return CommandName
}

func (c *Command) Execute() error {
	return nil
}

type LogCommand struct {
	err error
	cmd ICommand
}

func NewLogCommand(cmd ICommand, err error) *LogCommand {
	return &LogCommand{
		err: err,
		cmd: cmd,
	}
}

func (lc *LogCommand) Execute() error {
	log.Println(lc.err)
	return nil
}

func (lc *LogCommand) Name() string {
	return LogCommandName
}

type RepeatCommand struct {
	err error
	cmd ICommand
}

func NewRepeatCommand(cmd ICommand, err error) *RepeatCommand {
	return &RepeatCommand{
		err: err,
		cmd: cmd,
	}
}

func (rc *RepeatCommand) Execute() error {
	return rc.cmd.Execute()
}

func (rc *RepeatCommand) Name() string {
	return RepeatCommandName
}

type RepeatCommandTwo struct {
	err error
	cmd ICommand
}

func NewRepeatCommandTwo(cmd ICommand, err error) *RepeatCommandTwo {
	return &RepeatCommandTwo{
		err: err,
		cmd: cmd,
	}
}

func (rc *RepeatCommandTwo) Execute() error {
	rc.cmd.Execute()
	return nil
}

func (rc *RepeatCommandTwo) Name() string {
	return RepeatTwoCommandName
}

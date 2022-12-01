package hw7

type ICommand interface {
	Execute() error
}

type MoveToCommand struct{}

func (mts *MoveToCommand) Execute() error { return nil }

type RunCommand struct{}

func (rc *RunCommand) Execute() error { return nil }

type StartCommand struct {
	worker *worker
}

func NewStartCommand(worker *worker) *StartCommand {
	return &StartCommand{worker}
}

func (sc *StartCommand) Execute() error {
	sc.worker.Start()
	return nil
}

type HardStopCommand struct{}

func NewHardStopCommand() *HardStopCommand {
	return &HardStopCommand{}
}

func (hsc *HardStopCommand) Execute() error {
	return nil
}

type SoftStopCommand struct{}

func NewSoftStopCommand() *SoftStopCommand {
	return &SoftStopCommand{}
}

func (ssc *SoftStopCommand) Execute() error {
	return nil
}

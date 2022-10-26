package hw7

type ICommand interface {
	Execute() error
}

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

type HardStopCommand struct {
	worker *worker
}

func NewHardStopCommand(worker *worker) *HardStopCommand {
	return &HardStopCommand{worker}
}

func (hsc *HardStopCommand) Execute() error {
	hsc.worker.HardStop()
	return nil
}

type SoftStopCommand struct {
	worker *worker
}

func NewSoftStopCommand(worker *worker) *SoftStopCommand {
	return &SoftStopCommand{worker}
}

func (ssc *SoftStopCommand) Execute() error {
	ssc.worker.SoftStop()
	return nil
}

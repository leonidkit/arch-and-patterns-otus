package hw4

import (
	"errors"
	"fmt"
	"math"
)

var (
	errNotEnoughFueld   = errors.New("not enough fuel error")
	errObjectNotMovable = errors.New("not movable object error")
	errExecuteError     = errors.New("command Execute() error")
)

type ICommand interface {
	Execute() error
}

//go:generate mockgen -source=hw4.go -destination=mock.go -package=${GOPACKAGE}
type Fuelable interface {
	Current() int64
	Refuel(int64) error
	Burn() error
}

type Movable interface {
	Position() (int64, int64)
	Velocity() (int64, int64)
	SetPosition(x, y int64) error
	SetVelocity(x, y int64) error
}

type Rotable interface {
	Direction() int64
	AngurlarVelocity() int64
	DirectionNumber() int64
	SetDirection(int64) error
}

type VelocityChangable interface {
	SetVelocity(x, y int64) error
	VelocityValue() int64
	Angle() int64
}

// Macrocommand - реализовать простейшую макрокоманду и тесты к ней.
type Macrocommand struct {
	commands []ICommand
}

func NewMacrocommand(cmd ...ICommand) *Macrocommand {
	return &Macrocommand{
		commands: cmd,
	}
}

func (m *Macrocommand) Execute() error {
	for _, cmd := range m.commands {
		if err := cmd.Execute(); err != nil {
			return err
		}
	}
	return nil
}

// MoveAndBurnFuelCommand - реализовать команду движения по прямой с расходом топлива, используя команды с предыдущих шагов.
func NewMoveAndBurnFuelCommand(m Movable, f Fuelable) *Macrocommand {
	return NewMacrocommand(
		NewMoveCommand(m),
		NewBurnFuelCommand(f),
	)
}

// RotateAndChangeVelocityCommnad - реализовать команду поворота, которая еще и меняет вектор мгновенной скорости, если есть.
func NewRotateAndChangeVelocityCommnad(r Rotable, vc VelocityChangable) *Macrocommand {
	return NewMacrocommand(
		NewRotateCommand(r),
		NewChangeVelocityCommand(vc),
	)
}

type MoveCommand struct {
	Movable
}

func NewMoveCommand(m Movable) *MoveCommand {
	return &MoveCommand{m}
}

func (mc *MoveCommand) Execute() error {
	currX, currY := mc.Position()
	velX, velY := mc.Velocity()
	return mc.SetPosition(currX+velX, currY+velY)
}

type RotateCommand struct {
	Rotable
}

func NewRotateCommand(r Rotable) *RotateCommand {
	return &RotateCommand{r}
}

func (rc *RotateCommand) Execute() error {
	return rc.SetDirection(
		(rc.Direction() + rc.AngurlarVelocity()) % rc.DirectionNumber(),
	)
}

// CheckFuelCommand - реализовать класс CheckFuelComamnd.
type CheckFuelCommand struct {
	Fuelable
}

func NewCheckFuelCommand(f Fuelable) *CheckFuelCommand {
	return &CheckFuelCommand{f}
}

func (cfc *CheckFuelCommand) Execute() error {
	if cfc.Current() == 0 {
		return errNotEnoughFueld
	}
	return nil
}

// BurnFuelCommand - реализовать класс BurnFuelCommand.
type BurnFuelCommand struct {
	Fuelable
}

func NewBurnFuelCommand(f Fuelable) *BurnFuelCommand {
	return &BurnFuelCommand{f}
}

func (bfc *BurnFuelCommand) Execute() error {
	return bfc.Burn()
}

// ChangeVelocityCommand - реализована команда ChangeVelocityCommand.
type ChangeVelocityCommand struct {
	obj VelocityChangable
}

func NewChangeVelocityCommand(obj VelocityChangable) *ChangeVelocityCommand {
	return &ChangeVelocityCommand{obj}
}

func (cvc *ChangeVelocityCommand) Execute() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%w: %s", errExecuteError, r)
		}
	}()

	angl := cvc.obj.Angle()
	vv := cvc.obj.VelocityValue()
	return cvc.obj.SetVelocity(
		int64(float64(vv)*math.Cos(toRadians(angl))),
		int64(float64(vv)*math.Sin(toRadians(angl))),
	)
}

func toRadians(angle int64) float64 {
	return float64(angle) * (math.Pi / float64(180))
}

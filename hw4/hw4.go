package hw4

import (
	"errors"
	"math"
)

var (
	errNotEnoughFueld   = errors.New("not enough fuel error")
	errObjectNotMovable = errors.New("not movable object error")
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
	Position() (float64, float64)
	Velocity() (float64, float64)
	SetPosition(x, y float64) error
	SetVelocity(x, y float64) error
}

type Rotable interface {
	Direction() int64
	AngurlarVelocity() int64
	DirectionNumber() int64
	SetDirection(int64) error
}

type VelocityChangable interface {
	Velocity() (float64, float64)
	SetVelocity(x, y float64) error
	Direction() int64
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
func MoveAndBurnFuelCommand(m Movable, f Fuelable) *Macrocommand {
	return NewMacrocommand(
		NewMoveCommand(m),
		NewBurnFuelCommand(f),
	)
}

// RotateAndChangeVelocityCommnad - реализовать команду поворота, которая еще и меняет вектор мгновенной скорости, если есть.
func RotateAndChangeVelocityCommnad(r Rotable, vc VelocityChangable) *Macrocommand {
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

func (cvc *ChangeVelocityCommand) Execute() error {
	currVelX, currVelY := cvc.obj.Velocity()
	if currVelX == 0 && currVelY == 0 {
		return errObjectNotMovable
	}
	curDir := cvc.obj.Direction()
	return cvc.obj.SetVelocity(
		currVelX*(math.Cos(toRadians(curDir)))-currVelY*(math.Sin(toRadians(curDir))),
		currVelX*(math.Sin(toRadians(curDir))+currVelY*(math.Cos(toRadians(curDir)))),
	)
}

func toRadians(angle int64) float64 {
	return float64(angle) * (math.Pi / float64(180))
}

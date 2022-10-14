package hw2

import "fmt"

//go:generate mockgen -source=hw2.go -destination=mock.go -package=${GOPACKAGE}
type Rotable interface {
	Direction() int64
	AngurlarVelocity() int64
	DirectionNumber() int64
	SetDirection(int64)
}

type Rotate struct {
	Rotable
}

func NewRotate(r Rotable) Rotate {
	return Rotate{r}
}

func (r *Rotate) Execute() (err error) {
	// конструкция для отлавливания паники и конвертации ее в ошибку.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("execute processing error: %v", r)
		}
	}()

	r.SetDirection(
		(r.Direction() + r.AngurlarVelocity()) % r.DirectionNumber(),
	)

	return nil
}

//go:generate mockgen -source=hw2.go -destination=mock.go -package=${GOPACKAGE}
type Movable interface {
	Position() (int64, int64)
	Velocity() (int64, int64)
	SetPosition(x, y int64)
}

type Move struct {
	Movable
}

func NewMove(m Movable) Move {
	return Move{m}
}

func (m *Move) Execute() (err error) {
	// конструкция для отлавливания паники и конвертации ее в ошибку.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("execute processing error: %v", r)
		}
	}()

	currX, currY := m.Position()
	velX, velY := m.Velocity()
	m.SetPosition(currX+velX, currY+velY)

	return nil
}

package hw4

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/golang/mock/gomock"
)

const float64EqualityThreshold = 1e-3

func TestChangeVelocityCommand_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name    string
		movable Movable
		rotable Rotable
		wantErr bool
		want    []float64
		err     error
	}{
		{
			name: "movable and rotable object",
			movable: func() Movable {
				mock := NewMockMovable(ctrl)
				mock.EXPECT().Velocity().Return(float64(1), float64(1))
				mock.EXPECT().SetVelocity(FloatEq(-0.366), FloatEq(1.366))
				return mock
			}(),
			rotable: func() Rotable {
				mock := NewMockRotable(ctrl)
				mock.EXPECT().Direction().Return(int64(60))
				return mock
			}(),
			want: []float64{-0.366, 1.366}, // x = x*cos(alpha)-y*sin(alpha) = 1*0.5-1*0.866 = -0.366
			// y = x*sin(alpha)+y*cos(alpha) = 1*0.866 + 1*0.5 = 1.366
		},
		{
			name: "not movable object",
			movable: func() Movable {
				mock := NewMockMovable(ctrl)
				mock.EXPECT().Velocity().Return(float64(0), float64(0))
				return mock
			}(),
			rotable: func() Rotable {
				mock := NewMockRotable(ctrl)
				return mock
			}(),
			wantErr: true,
			err:     errObjectNotMovable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := struct {
				Movable
				Rotable
			}{tt.movable, tt.rotable}
			cvc := NewChangeVelocityCommand(s)
			err := cvc.Execute()
			if err != nil && !tt.wantErr {
				t.Fatalf("unexpected error: got = %s", err)
			}
			if !errors.Is(err, tt.err) {
				t.Fatalf("got = %s, want = %s", err, tt.err)
			}
		})
	}
}

type floatMatcher struct {
	x any
}

func FloatEq(x any) gomock.Matcher { return floatMatcher{x} }

// Matches returns whether x is a match.
func (fm floatMatcher) Matches(x interface{}) bool {
	switch x.(type) {
	case float64:
		return floatEqual(x.(float64), fm.x.(float64), float64EqualityThreshold)
	}
	return false
}

// String describes what the matcher matches.
func (fm floatMatcher) String() string {
	return "is float equal"
}

func floatEqual(a, b float64, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestCheckFuelCommand_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name     string
		fuelable Fuelable
		wantErr  bool
		err      error
	}{
		{
			name: "fuel is not empty",
			fuelable: func() Fuelable {
				f := NewMockFuelable(ctrl)
				f.EXPECT().Current().Return(int64(12))
				return f
			}(),
		},
		{
			name: "fuel is empty",
			fuelable: func() Fuelable {
				f := NewMockFuelable(ctrl)
				f.EXPECT().Current().Return(int64(0))
				return f
			}(),
			wantErr: true,
			err:     errNotEnoughFueld,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := NewCheckFuelCommand(tt.fuelable)
			err := fc.Execute()
			if err != nil && !tt.wantErr {
				t.Fatalf("unexpected error: got = %s", err)
			}
			if !errors.Is(err, tt.err) {
				t.Fatalf("got = %s, want = %s", err, tt.err)
			}
		})
	}
}

func TestBurnFuelCommand_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	errUnexpected := fmt.Errorf("unexpected error")

	tests := []struct {
		name     string
		fuelable Fuelable
		wantErr  bool
		err      error
	}{
		{
			name: "burn successfully",
			fuelable: func() Fuelable {
				f := NewMockFuelable(ctrl)
				f.EXPECT().Burn().Return(nil)
				return f
			}(),
		},
		{
			name: "burn unsuccessfully",
			fuelable: func() Fuelable {
				f := NewMockFuelable(ctrl)
				f.EXPECT().Burn().Return(errUnexpected)
				return f
			}(),
			wantErr: true,
			err:     errUnexpected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := NewBurnFuelCommand(tt.fuelable)
			err := fc.Execute()
			if err != nil && !tt.wantErr {
				t.Fatalf("unexpected error: got = %s", err)
			}
			if !errors.Is(err, tt.err) {
				t.Fatalf("got = %s, want = %s", err, tt.err)
			}
		})
	}
}

func TestMacrocommand_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("macrocommand ok", func(t *testing.T) {
		cmd1 := NewMockICommand(ctrl)
		cmd1.EXPECT().Execute().Return(nil)
		cmd2 := NewMockICommand(ctrl)
		cmd2.EXPECT().Execute().Return(nil)

		macro := NewMacrocommand(cmd1, cmd2)
		err := macro.Execute()
		if err != nil {
			t.Fatalf("unexpected error: got = %s", err)
		}
	})

	t.Run("macrocommand error", func(t *testing.T) {
		cmd1 := NewMockICommand(ctrl)
		cmd1.EXPECT().Execute().Return(fmt.Errorf("error"))
		cmd2 := NewMockICommand(ctrl)

		macro := NewMacrocommand(cmd1, cmd2)
		err := macro.Execute()
		if nil == err {
			t.Fatalf("expected error: got = nil")
		}
	})
}

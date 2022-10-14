package hw4

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

const float64EqualityThreshold = 1e-3

func TestChangeVelocityCommand_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name              string
		velocityChangable VelocityChangable
		wantErr           bool
		err               error
	}{
		{
			name: "movable and rotable object",
			velocityChangable: func() VelocityChangable {
				mock := NewMockVelocityChangable(ctrl)
				mock.EXPECT().Angle().Return(int64(45))
				mock.EXPECT().VelocityValue().Return(int64(2))
				mock.EXPECT().SetVelocity(int64Eq(int64(1)), int64Eq(int64(1)))
				return mock
			}(),
		},
		{
			name: "not movable object",
			velocityChangable: func() VelocityChangable {
				mock := NewMockVelocityChangable(ctrl)
				mock.EXPECT().Angle().Do(func() {
					panic("not implemented")
				})
				return mock
			}(),
			wantErr: true,
			err:     errExecuteError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cvc := NewChangeVelocityCommand(tt.velocityChangable)
			err := cvc.Execute()
			if err != nil && !tt.wantErr {
				t.Fatalf("unexpected error: got = %s", err)
			}
			if !errors.Is(err, tt.err) {
				t.Fatalf("got = %q, want = %q", err, tt.err)
			}
		})
	}
}

type int64Matcher struct {
	x any
}

func int64Eq(x any) gomock.Matcher { return int64Matcher{x} }

// Matches returns whether x is a match.
func (fm int64Matcher) Matches(x interface{}) bool {
	switch x.(type) {
	case int64, int:
		return x.(int64) == fm.x.(int64)
	}
	return false
}

// String describes what the matcher matches.
func (fm int64Matcher) String() string {
	return fmt.Sprintf("int64 is equal %d", fm.x)
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

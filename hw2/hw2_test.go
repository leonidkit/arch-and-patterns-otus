package hw2

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMove(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name    string
		object  Movable
		wantErr bool
	}{
		{
			name: "simple moving test",
			object: func(ctrl *gomock.Controller) *MockMovable {
				m := NewMockMovable(ctrl)
				m.EXPECT().Position().Return(int64(12), int64(5))
				m.EXPECT().Velocity().Return(int64(-7), int64(3))
				m.EXPECT().SetPosition(int64(5), int64(8))
				return m
			}(ctrl),
		},
		{
			name: "unable get position test",
			object: func(ctrl *gomock.Controller) *MockMovable {
				m := NewMockMovable(ctrl)
				m.EXPECT().Position().Do(func() { panic("Position() not implemented") })
				return m
			}(ctrl),
			wantErr: true,
		},
		{
			name: "unable get velocity test",
			object: func(ctrl *gomock.Controller) *MockMovable {
				m := NewMockMovable(ctrl)
				m.EXPECT().Position().Return(int64(12), int64(5))
				m.EXPECT().Velocity().Do(func() { panic("Velocity() not implemented") })
				return m
			}(ctrl),
			wantErr: true,
		},
		{
			name: "unable call set position method test",
			object: func(ctrl *gomock.Controller) *MockMovable {
				m := NewMockMovable(ctrl)
				m.EXPECT().Position().Return(int64(12), int64(5))
				m.EXPECT().Velocity().Return(int64(-7), int64(3))
				m.EXPECT().SetPosition(gomock.Any(), gomock.Any()).Do(func(x, y int64) { panic("SetPosition() not implemented") })
				return m
			}(ctrl),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMove(tt.object)
			err := m.Execute()
			if err != nil && !tt.wantErr {
				panic(fmt.Errorf("unexpected error: %w", err))
			}
		})
	}
}

func TestRotate(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name    string
		object  Rotable
		wantErr bool
	}{
		{
			name: "simple rotating test",
			object: func(ctrl *gomock.Controller) *MockRotable {
				m := NewMockRotable(ctrl)
				m.EXPECT().Direction().Return(int64(340))
				m.EXPECT().DirectionNumber().Return(int64(360))
				m.EXPECT().AngurlarVelocity().Return(int64(40))
				m.EXPECT().SetDirection(int64(20))
				return m
			}(ctrl),
		},
		{
			name: "unable get direction test",
			object: func(ctrl *gomock.Controller) *MockRotable {
				m := NewMockRotable(ctrl)
				m.EXPECT().Direction().Do(func() { panic("Direction() not implemented") })
				return m
			}(ctrl),
			wantErr: true,
		},
		{
			name: "unable get angular velocity test",
			object: func(ctrl *gomock.Controller) *MockRotable {
				m := NewMockRotable(ctrl)
				m.EXPECT().Direction().Return(int64(340))
				m.EXPECT().AngurlarVelocity().Do(func() { panic("AngurlarVelocity() not implemented") })
				return m
			}(ctrl),
			wantErr: true,
		},
		{
			name: "unable get direction number test",
			object: func(ctrl *gomock.Controller) *MockRotable {
				m := NewMockRotable(ctrl)
				m.EXPECT().Direction().Return(int64(340))
				m.EXPECT().AngurlarVelocity().Return(int64(40))
				m.EXPECT().DirectionNumber().Do(func() { panic("DirectionNumber() not implemented") })
				return m
			}(ctrl),
			wantErr: true,
		},
		{
			name: "unable set new direction test",
			object: func(ctrl *gomock.Controller) *MockRotable {
				m := NewMockRotable(ctrl)
				m.EXPECT().Direction().Return(int64(340))
				m.EXPECT().DirectionNumber().Return(int64(360))
				m.EXPECT().AngurlarVelocity().Return(int64(40))
				m.EXPECT().SetDirection(gomock.Any()).Do(func(int64) { panic("SetDirection() not implemented") })
				return m
			}(ctrl),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRotate(tt.object)
			err := r.Execute()
			if err != nil && !tt.wantErr {
				panic(fmt.Errorf("unexpected error: %w", err))
			}
		})
	}
}

package hw1

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name      string
		a, b, c   float64
		want      []float64
		wantPanic bool
	}{
		{
			name: "no roots",
			a:    1,
			b:    0,
			c:    1,
			want: []float64{},
		},
		{
			name: "two roots",
			a:    1,
			b:    0,
			c:    -1,
			want: []float64{1, -1},
		},
		{
			name: "one root",
			a:    0.00001,
			b:    0.000001,
			c:    0.00001,
			want: []float64{-5.0000000000000005e-12},
		},
		{
			name:      "panic rised",
			a:         0,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						panic(fmt.Sprintf("panic was not expected in the test %q", tt.name))
					}
				}
			}()

			got := Solve(tt.a, tt.b, tt.c)
			if !isEqualSlices(got, tt.want) {
				panic(fmt.Sprintf("test %q failed: got=%v, want=%v", tt.name, got, tt.want))
			}
		})
	}
}

// isEqualSlices returns true if slices is equal independent of order.
func isEqualSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	presentedInFirst := make(map[T]bool, len(a))
	for i := 0; i < len(a); i++ {
		presentedInFirst[a[i]] = false
	}

	for i := 0; i < len(b); i++ {
		if _, ok := presentedInFirst[b[i]]; ok {
			presentedInFirst[b[i]] = true
		}
	}

	for _, v := range presentedInFirst {
		if v == false {
			return false
		}
	}
	return true
}

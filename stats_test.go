package stats

import (
	"testing"
)

type test struct {
	name     string
	data     []float64
	expected float64
}

func TestMode(t *testing.T) {
	tests := []test{
		{
			name:     "Simple mode test",
			data:     []float64{0.0, 2.0, 3.0, 0.0, 0.0, 0.0, 1.0, 2.0, 3.0, 3.0},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mode, err := Mode(tt.data)
			if err != nil {
				t.Errorf("unexpected error received: %v", err)
			}

			if mode != tt.expected {
				t.Errorf("expected mode: %v, got:%v", mode, tt.expected)
			}
		})
	}
}

package stats

import (
	"reflect"
	"testing"
)

type testrv struct {
	name     string
	data     []float64
	expected float64
}

func TestNewRandVar(t *testing.T) {
	data := []float64{1.0, 3.5, 2.2}
	rv := NewRandVar(data)

	if !reflect.DeepEqual(data, rv.data) {
		t.Errorf("Expected %v, but got %v", data, rv.data)
	}

	// Let's test if the data is immutable despite modifying the original
	data[0] = 22.2
	if rv.data[0] == data[0] {
		t.Errorf("A modification on original data has modified the data of the random variable")
	}
}

func TestMean(t *testing.T) {
	tests := []testrv{
		{
			name:     "Only positive numbers",
			data:     []float64{2.0, 0.0, 1.0},
			expected: 1.0,
		},
		{
			name:     "Only negative numbers",
			data:     []float64{-2.5, -0.5, -1.5},
			expected: -1.5,
		},
		{
			name:     "Mixed numbers",
			data:     []float64{-4.0, 3.0, -2.0, 1.0},
			expected: -0.5,
		},
		{
			name:     "Only one number",
			data:     []float64{1.0},
			expected: 1.0,
		},
		{
			name:     "Empty data",
			data:     []float64{},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := NewRandVar(tt.data)
			mean := rv.Mean()

			if mean != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, mean)
			}
		})
	}

}

func TestMedian(t *testing.T) {
	tests := []testrv{
		{
			name:     "Only positive numbers",
			data:     []float64{5.0, 3.0, 1.0},
			expected: 3.0,
		},
		{
			name:     "Only negative numbers",
			data:     []float64{-2.5, -0.5, -1.5},
			expected: -1.5,
		},
		{
			name:     "Mixed numbers",
			data:     []float64{-2.2, 0.4, 1.2, 3.3},
			expected: 0.8,
		},
		{
			name:     "Only one number",
			data:     []float64{1.0},
			expected: 1.0,
		},
		{
			name:     "Empty data",
			data:     []float64{},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := NewRandVar(tt.data)
			median := rv.Median()

			if median != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, median)
			}
		})
	}
}

func TestVariance(t *testing.T) {
	tests := []testrv{
		{
			name:     "Only positive numbers",
			data:     []float64{4.0, 2.0, 1.0, 3.0},
			expected: 1.25,
		},
		{
			name:     "Only negative numbers",
			data:     []float64{-4.0, -2.0, -1.0, -3.0},
			expected: 1.25,
		},
		{
			name:     "Mixed numbers",
			data:     []float64{4.0, 2.0, -1.0, -3.0},
			expected: 7.25,
		},
		{
			name:     "Only one number",
			data:     []float64{1.0},
			expected: 0.0,
		},
		{
			name:     "Empty data",
			data:     []float64{},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := NewRandVar(tt.data)
			variance := rv.Variance()

			if variance != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, variance)
			}
		})
	}
}

func TestSkewness(t *testing.T) {
	tests := []testrv{
		{
			name:     "Only one number",
			data:     []float64{1.0},
			expected: 0.0,
		},
		{
			name:     "Empty data",
			data:     []float64{},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := NewRandVar(tt.data)
			skewness, err := rv.Skewness()
			if err != nil {
				if skewness != tt.expected {
					t.Errorf("Expected %f, got %f", tt.expected, skewness)
				}
			}
		})
	}
}

func TestKurtosis(t *testing.T) {
	tests := []testrv{
		{
			name:     "Only one number",
			data:     []float64{1.0},
			expected: 0.0,
		},
		{
			name:     "Empty data",
			data:     []float64{},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := NewRandVar(tt.data)
			kurtosis, err := rv.Kurtosis()
			if err != nil {
				if kurtosis != tt.expected {
					t.Errorf("Expected %f, got %f", tt.expected, kurtosis)
				}
			}
		})
	}
}

func TestCovariance(t *testing.T) {
	dataX := []float64{50.2, 60.3, 45.23, 55.75, 70.91}
	dataY := []float64{24.6, 41.9, 33.33, 27.9, 44.1}
	x, y := NewRandVar(dataX), NewRandVar(dataY)

	covXY, err := x.Covariance(y)
	if err != nil {
		t.Errorf("Unexpected error :%v", err)
	}

	covYX, err := y.Covariance(x)
	if err != nil {
		t.Errorf("Unexpected error :%v", err)
	}

	if covXY != covYX {
		t.Errorf("Covariance of two random variable are not equal: COV(X, Y):%f, COV(Y, X): %f", covXY, covYX)
	}

	dataX = []float64{50.2, 60.3, 45.23, 70.91}
	dataY = []float64{24.6, 41.9, 33.33, 27.9, 44.1}
	x, y = NewRandVar(dataX), NewRandVar(dataY)
	_, err = x.Covariance(y)
	if err == nil {
		t.Errorf("Length of both random variables must be equal: len(x):%d; len(y):%d", len(x.data), len(y.data))
	}

	_, err = y.Covariance(x)
	if err == nil {
		t.Errorf("Length of both random variables must be equal: len(x):%d; len(y):%d", len(x.data), len(y.data))
	}
}

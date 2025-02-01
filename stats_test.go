package stats

import (
	"math"
	"reflect"
	"testing"
)

type singleValueTest struct {
	name     string
	data     []float64
	expected map[string]float64
	err      error
}

var GeneralTests = []singleValueTest{
	{
		name: "Empty data",
		data: []float64{},
		expected: map[string]float64{"mean": 0.0,
			"mode": 0.0, "median": 0.0,
			"std": 0, "max": 0,
			"min": 0, "sum": 0},
		err: ErrEmptyData,
	},
	{
		name: "Only one number",
		data: []float64{2.0},
		expected: map[string]float64{"mean": 2.0,
			"mode": 2.0, "median": 2.0,
			"std": 0, "max": 2.0,
			"min": 2.0, "sum": 2.0},
		err: nil,
	},
}

func TestMean(t *testing.T) {
	tests := []singleValueTest{
		{
			name:     "Simple mean singleValueTest",
			data:     []float64{0.0, 2.0, -2.0, 1.5, 0.5, 0.0, -3.0, 4.5},
			expected: map[string]float64{"mean": 0.4375},
			err:      nil,
		},
	}
	tests = append(tests, GeneralTests...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mean, err := Mean(tt.data)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if mean != tt.expected["mean"] {
				t.Errorf("expected mean: %v, got:%v", mean, tt.expected["mean"])
			}
		})
	}
}

func TestMedian(t *testing.T) {
	tests := []singleValueTest{
		{
			name:     "Sorted dataset: Odd number of elements",
			data:     []float64{0.0, 1.0, 2.0, 3.0, 4.0},
			expected: map[string]float64{"median": 2.0},
			err:      nil,
		},
		{
			name:     "Sorted dataset: Even number of elements",
			data:     []float64{0.0, 1.0, 2.0, 3.0},
			expected: map[string]float64{"median": 1.5},
			err:      nil,
		},
		{
			name:     "Not sorted dataset: Odd number of elements",
			data:     []float64{4.0, 1.0, 3.0, 2.0, 0.0},
			expected: map[string]float64{"median": 2.0},
			err:      nil,
		},
		{
			name:     "Not sorted dataset: Even number of elements",
			data:     []float64{1.0, 3.0, 0.0, 2.0},
			expected: map[string]float64{"median": 1.5},
			err:      nil,
		},
	}
	tests = append(tests, GeneralTests...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			median, err := Median(tt.data)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if median != tt.expected["median"] {
				t.Errorf("expected median: %v, got:%v", median, tt.expected["median"])
			}
		})
	}
}

func TestMode(t *testing.T) {
	tests := []singleValueTest{
		{
			name:     "Simple mode singleValueTest",
			data:     []float64{0.0, 2.0, 3.0, 0.0, 0.0, 0.0, 1.0, 2.0, 3.0, 3.0},
			expected: map[string]float64{"mode": 0.0},
			err:      nil,
		},
	}
	tests = append(tests, GeneralTests...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mode, err := Mode(tt.data)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if mode != tt.expected["mode"] {
				t.Errorf("expected mode: %v, got:%v", mode, tt.expected["mode"])
			}
		})
	}
}

func TestStdDeviation(t *testing.T) {
	tests := []singleValueTest{
		{
			name:     "Simple standard deviation singleValueTest",
			data:     []float64{2.0, -2.0, 1.0, -1.0, 0.0},
			expected: map[string]float64{"std": math.Sqrt(2)},
			err:      nil,
		},
	}
	tests = append(tests, GeneralTests...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			std, err := StandardDeviation(tt.data)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if std != tt.expected["std"] {
				t.Errorf("expected std: %v, got:%v", std, tt.expected["std"])
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []singleValueTest{
		{
			name:     "Simple maximum finder singleValueTest",
			data:     []float64{2.0, -2.0, 1.0, -1.0, 0.0},
			expected: map[string]float64{"max": 2},
			err:      nil,
		},
		{
			name:     "Only negative numbers",
			data:     []float64{-4.0, -2.0, -3.0, -0.5, -7.0},
			expected: map[string]float64{"max": -0.5},
			err:      nil,
		},
	}
	tests = append(tests, GeneralTests...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			max, err := Max(tt.data)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if max != tt.expected["max"] {
				t.Errorf("expected max: %v, got:%v", max, tt.expected["max"])
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []singleValueTest{
		{
			name:     "Simple minimum finder singleValueTest",
			data:     []float64{2.0, -2.0, 1.0, -1.0, 0.0},
			expected: map[string]float64{"min": -2},
			err:      nil,
		},
		{
			name:     "Only negative numbers",
			data:     []float64{-4.0, -2.0, -3.0, -0.5, -7.0},
			expected: map[string]float64{"min": -7.0},
			err:      nil,
		},
	}
	tests = append(tests, GeneralTests...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			min, err := Min(tt.data)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if min != tt.expected["min"] {
				t.Errorf("expected min: %v, got:%v", min, tt.expected["min"])
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []singleValueTest{
		{
			name:     "Simple sum singleValueTest",
			data:     []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			expected: map[string]float64{"sum": 15.0},
			err:      nil,
		},
		{
			name:     "Only negative numbers",
			data:     []float64{-1.0, -2.0, -3.0, -4.0, -5.0},
			expected: map[string]float64{"sum": -15.0},
			err:      nil,
		},
	}
	tests = append(tests, GeneralTests...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := Sum(tt.data)

			if sum != tt.expected["sum"] {
				t.Errorf("expected sum: %v, got:%v", sum, tt.expected["sum"])
			}
		})
	}
}

type sortTest struct {
	name     string
	data     []float64
	expected map[string][]float64
	err      error
}

var sortTests = []sortTest{
	{
		name: "Empty data",
		data: []float64{},
		expected: map[string][]float64{
			"sort": nil, "reverseSort": nil,
		},
		err: ErrEmptyData,
	},
	{
		name: "Only one element in data",
		data: []float64{2.0},
		expected: map[string][]float64{
			"sort": {2.0}, "reverseSort": {2.0},
		},
	},
}

func TestSort(t *testing.T) {
	tests := []sortTest{
		{
			name: "Simple sort test",
			data: []float64{2.0, -2.0, 1.0, -1.0, 0.0},
			expected: map[string][]float64{
				"sort":        {-2.0, -1.0, 0.0, 1.0, 2.0},
				"reverseSort": {2.0, 1.0, 0.0, -1.0, -2.0},
			},
			err: nil,
		},
		{
			name: "Another sorting test",
			data: []float64{1.0, 1.01, 0.99, 1.001, 1.1},
			expected: map[string][]float64{
				"sort":        {0.99, 1.0, 1.001, 1.01, 1.1},
				"reverseSort": {1.1, 1.01, 1.001, 1.0, 0.99},
			},
			err: nil,
		},
	}
	tests = append(tests, sortTests...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort := Sort(tt.data)
			if len(sort) != len(tt.expected["sort"]) {
				t.Errorf("expected sort: %v, got:%v", sort, tt.expected["sort"])
			}

			for i, v := range sort {
				if v != tt.expected["sort"][i] {
					t.Errorf("expected sort: %v, got:%v", sort, tt.expected["sort"])
				}
			}

			reverseSort := ReverseSort(tt.data)
			if len(reverseSort) != len(tt.expected["reverseSort"]) {
				t.Errorf("expected reverseSort: %v, got:%v", reverseSort, tt.expected["reverseSort"])
			}

			for i, v := range reverseSort {
				if v != tt.expected["reverseSort"][i] {
					t.Errorf("expected reverseSort: %v, got:%v", reverseSort, tt.expected["reverseSort"])
				}
			}
		})
	}
}

type transformationTest struct {
	name     string
	data     []float64
	expected []float64
	err      error
}

var normalizeTest = []transformationTest{
	{
		name:     "Empty data",
		data:     []float64{},
		expected: nil,
		err:      ErrEmptyData,
	},
	{
		name:     "Only a single element",
		data:     []float64{2.0},
		expected: nil,
		err:      ErrNullStdDeviation,
	},
}

func TestNormalization(t *testing.T) {
	tests := []transformationTest{
		{
			name:     "Simple normalization test",
			data:     []float64{2.0, -2.0, 1.0, -1.0, 0.0},
			expected: []float64{2.0 / math.Sqrt(2), -2.0 / math.Sqrt(2), 1.0 / math.Sqrt(2), -1.0 / math.Sqrt(2), 0.0 / math.Sqrt(2)},
			err:      nil,
		},
	}
	tests = append(tests, normalizeTest...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalized, err := Normalize(tt.data)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if !reflect.DeepEqual(normalized, tt.expected) {
				t.Errorf("expected normalized: %v, got:%v", normalized, tt.expected)
			}
		})
	}
}

type ScaleTest struct {
	name     string
	data     []float64
	factor   float64
	expected []float64
	err      error
}

var scaleTests = []ScaleTest{
	{
		name:     "Empty data",
		data:     []float64{},
		factor:   2.0,
		expected: nil,
		err:      ErrEmptyData,
	},
	{
		name:     "Null factor",
		data:     []float64{1.0},
		factor:   0.0,
		expected: nil,
		err:      ErrNullScaleFactor,
	},
	{
		name:     "Only a single element",
		data:     []float64{1.0},
		factor:   2.0,
		expected: []float64{2.0},
		err:      nil,
	},
}

func TestScale(t *testing.T) {
	tests := []ScaleTest{
		{
			name:     "Simple normalization test",
			data:     []float64{2.0, -2.0, 1.0, -1.0, 0.0},
			factor:   2.0,
			expected: []float64{4.0, -4.0, 2.0, -2.0, 0.0},
			err:      nil,
		},
	}
	tests = append(tests, scaleTests...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scaled, err := Scale(tt.data, tt.factor)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if !reflect.DeepEqual(scaled, tt.expected) {
				t.Errorf("expected scaled: %v, got:%v", scaled, tt.expected)
			}
		})
	}
}

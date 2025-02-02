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
				t.Errorf("expected mean: %v, got:%v", tt.expected["mean"], mean)
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
				t.Errorf("expected median: %v, got:%v", tt.expected["median"], median)
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
				t.Errorf("expected mode: %v, got:%v", tt.expected["mode"], mode)
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
				t.Errorf("expected std: %v, got:%v", tt.expected["std"], std)
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
				t.Errorf("expected max: %v, got:%v", tt.expected["max"], max)
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
				t.Errorf("expected min: %v, got:%v", tt.expected["min"], min)
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
				t.Errorf("expected sum: %v, got:%v", tt.expected["sum"], sum)
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
				t.Errorf("expected sort: %v, got:%v", tt.expected["sort"], sort)
			}

			for i, v := range sort {
				if v != tt.expected["sort"][i] {
					t.Errorf("expected sort: %v, got:%v", tt.expected["sort"], sort)
				}
			}

			reverseSort := ReverseSort(tt.data)
			if len(reverseSort) != len(tt.expected["reverseSort"]) {
				t.Errorf("expected reverseSort: %v, got:%v", tt.expected["reverseSort"], reverseSort)
			}

			for i, v := range reverseSort {
				if v != tt.expected["reverseSort"][i] {
					t.Errorf("expected reverseSort: %v, got:%v", tt.expected["reverseSort"], reverseSort)
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
				t.Errorf("expected normalized: %v, got:%v", tt.expected, normalized)
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
				t.Errorf("expected scaled: %v, got:%v", tt.expected, scaled)
			}
		})
	}
}

type comparisonTest struct {
	name     string
	dataA    []float64
	dataB    []float64
	epsilon  float64
	expected bool
}

var comparisonTests = []comparisonTest{
	{
		name:     "Empty data",
		dataA:    []float64{},
		dataB:    []float64{},
		epsilon:  0.0,
		expected: true,
	},
	{
		name:     "Different lengths",
		dataA:    []float64{2.0, 1.2},
		dataB:    []float64{1.2},
		epsilon:  0.0,
		expected: false,
	},
	{
		name:     "Simple test",
		dataA:    []float64{2.0, 1.2, -3.3, 2.5},
		dataB:    []float64{2.0, 1.2, -3.3, 2.5},
		epsilon:  1e-8,
		expected: true,
	},
}

func TestEquals(t *testing.T) {
	for _, tt := range comparisonTests {
		t.Run(tt.name, func(t *testing.T) {
			equal := Equals(tt.dataA, tt.dataB, tt.epsilon)
			if equal != tt.expected {
				t.Errorf("expected equal: %v, got:%v", tt.expected, equal)
			}
		})
	}
}

type unionIntersectionTest struct {
	name     string
	dataA    []float64
	dataB    []float64
	epsilon  float64
	expected []float64
}

var unionTests = []unionIntersectionTest{
	{
		name:     "Empty data",
		dataA:    []float64{},
		dataB:    []float64{},
		expected: []float64{},
	},
	{
		name:     "Empty data A",
		dataA:    []float64{},
		dataB:    []float64{1.0, 2.0},
		expected: []float64{1.0, 2.0},
	},
	{
		name:     "Empty data B",
		dataA:    []float64{1.0, 2.0},
		dataB:    []float64{},
		expected: []float64{1.0, 2.0},
	},
	{
		name:     "Simple test",
		dataA:    []float64{1.0, 2.0, -1.0001, 2.22222222, 3.0, -2.4545454545},
		dataB:    []float64{2.22222222, 3.0, 4.13, -3.3333333, -1.0001},
		expected: []float64{1.0, 2.0, -1.0001, 2.22222222, 3.0, -2.4545454545, 4.13, -3.3333333},
		epsilon:  1e-8,
	},
}

func TestUnion(t *testing.T) {
	for _, tt := range unionTests {
		t.Run(tt.name, func(t *testing.T) {
			union := Union(tt.dataA, tt.dataB, tt.epsilon)

			if !reflect.DeepEqual(union, tt.expected) {
				t.Errorf("expected union: %v, got:%v", tt.expected, union)
			}
		})
	}
}

var intersectionTest = []unionIntersectionTest{
	{
		name:     "Empty data",
		dataA:    nil,
		dataB:    nil,
		expected: nil,
	},
	{
		name:     "Empty data A",
		dataA:    nil,
		dataB:    []float64{1.0, 2.0},
		expected: nil,
	},
	{
		name:     "Empty data B",
		dataA:    []float64{1.0, 2.0},
		dataB:    nil,
		expected: nil,
	},
	{
		name:     "Simple test",
		dataA:    []float64{1.0, 2.0, -1.0001, 2.22222222, 3.0, -2.4545454545},
		dataB:    []float64{2.22222222, 3.0, 4.13, -3.3333333, -1.0001},
		expected: []float64{-1.0001, 2.22222222, 3.0},
		epsilon:  1e-8,
	},
}

func TestIntersection(t *testing.T) {
	for _, tt := range intersectionTest {
		t.Run(tt.name, func(t *testing.T) {
			intersection := Intersection(tt.dataA, tt.dataB, tt.epsilon)

			if !reflect.DeepEqual(intersection, tt.expected) {
				t.Errorf("expected intersection: %v, got:%v", tt.expected, intersection)
			}
		})
	}
}

type percentileTest struct {
	name     string
	data     []float64
	p        float64
	expected float64
	err      error
}

var percentileTests = []percentileTest{
	{
		name:     "Empty data",
		data:     nil,
		p:        50,
		expected: 0,
		err:      ErrEmptyData,
	},
	{
		name:     "Only single data",
		data:     []float64{1.0},
		p:        50,
		expected: 1.0,
		err:      nil,
	},
	{
		name:     "Invalid percentile",
		data:     []float64{1.0, 2.0, 3.0},
		p:        101,
		expected: 0,
		err:      ErrInvalidPercentile,
	},
	{
		name:     "Invalid percentile 2",
		data:     []float64{1.0, 2.0, 3.0},
		p:        -0.5,
		expected: 0,
		err:      ErrInvalidPercentile,
	},
	{
		name:     "Simple data test",
		data:     []float64{50, 55, 60, 62, 65, 70, 72, 75, 80, 85},
		p:        25,
		expected: 58.75,
		err:      nil,
	},
}

func TestPercentile(t *testing.T) {
	for _, tt := range percentileTests {
		t.Run(tt.name, func(t *testing.T) {
			percentile, err := Percentile(tt.data, tt.p)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if percentile != tt.expected {
				t.Errorf("expected percentile: %v, got:%v", tt.expected, percentile)
			}
		})
	}
}

type quantileTest struct {
	name     string
	data     []float64
	qs       float64
	n        uint
	expected float64
	err      error
}

var quantileTests = []quantileTest{
	{
		name:     "Empty data",
		data:     nil,
		expected: 0,
		err:      ErrEmptyData,
	},
	{
		name:     "Only single data",
		data:     []float64{1.0},
		qs:       3,
		n:        10,
		expected: 1.0,
		err:      nil,
	},
	{
		name:     "Invalid quantile index",
		data:     []float64{1.0, 2.0, 3.0},
		qs:       11,
		n:        10,
		expected: 0,
		err:      ErrInvalideQuantile,
	},
	{
		name:     "Invalid quantile index 2",
		data:     []float64{1.0, 2.0, 3.0},
		qs:       -1,
		n:        10,
		expected: 0,
		err:      ErrInvalideQuantile,
	},
	{
		name:     "Simple data test",
		data:     []float64{12, 7, 18, 5, 13, 15, 8, 20, 22, 25},
		qs:       1,
		n:        4,
		expected: 7.75,
		err:      nil,
	},
}

func TestQuantile(t *testing.T) {
	for _, tt := range quantileTests {
		t.Run(tt.name, func(t *testing.T) {
			quantile, err := Quantile(tt.data, tt.qs, tt.n)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if quantile != tt.expected {
				t.Errorf("expected quantile: %v, got:%v", tt.expected, quantile)
			}
		})
	}
}

type kurtosisSkewnessTest struct {
	name     string
	data     []float64
	expected map[string]float64
	err      error
}

var kurtosisSkewnessTests = []kurtosisSkewnessTest{
	{
		name: "Empty data",
		data: nil,
		expected: map[string]float64{
			"kurtosis": 0.0, "skewness": 0.0,
		},
		err: ErrEmptyData,
	},
	{
		name: "Only single data",
		data: []float64{2.0},
		expected: map[string]float64{
			"kurtosis": 0.0, "skewness": 0.0,
		},
		err: ErrNullStdDeviation,
	},
}

func TestSkewnessKurtosis(t *testing.T) {
	tests := []kurtosisSkewnessTest{
		{
			name:     "Simple test",
			data:     []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			expected: map[string]float64{"kurtosis": -1.3, "skewness": 0.0},
			err:      nil,
		},
	}
	tests = append(tests, kurtosisSkewnessTests...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skewness, err := Skewness(tt.data)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if skewness-tt.expected["skewness"] > 1e-8 {
				t.Errorf("expected skewness: %v, got:%v", tt.expected["skewness"], skewness)
			}

			kurtosis, err := Kurtosis(tt.data)
			if err != tt.err {
				t.Errorf("unexpected error received: %v", err)
			}

			if kurtosis-tt.expected["kurtosis"] > 1e-8 {
				t.Errorf("expected kurtosis: %v, got:%v", tt.expected["kurtosis"], kurtosis)
			}
		})
	}
}

package stats

import (
	"math"
	"slices"
)

// RandVar is a struct that represents a simple set of data of a random variable
type RandVar struct {
	data []float64
}

// Creates a new RandVar
func NewRandVar(data []float64) *RandVar {
	d := append([]float64(nil), data...)
	rv := &RandVar{
		data: d,
	}
	return rv
}

// Returns the mean of the data. It will return 0 when data length is 0
func (rv *RandVar) Mean() float64 {
	var sum float64
	n := len(rv.data)
	if n == 0 {
		return 0.0
	}

	for _, v := range rv.data {
		sum += v
	}

	return sum / float64(n)
}

// Returns the median of the data
func (rv *RandVar) Median() float64 {
	n := len(rv.data)

	data := rv.data
	slices.Sort(data)

	if n == 0 {
		return 0
	}

	if n%2 == 1 {
		return data[n/2]
	}

	return (data[n/2-1] + data[n/2]) / 2
}

// Returns the variance of the data. It will return 0 when data length is 0
func (rv *RandVar) Variance() float64 {
	mean := rv.Mean()
	sum := 0.0
	n := len(rv.data)
	if n == 0 {
		return 0.0
	}

	for _, v := range rv.data {
		sum += math.Pow((v - mean), 2)
	}
	return sum / float64(n)
}

// Returns the standard deviation of the data
func (rv *RandVar) StdDev() float64 {
	return math.Sqrt(rv.Variance())
}

// Returns the value of the skewness of the data and an error when the standard deviation is 0
func (rv *RandVar) Skewness() (float64, error) {
	stdDev := rv.StdDev()
	if stdDev == 0 {
		return 0, ErrNullStdDeviation
	}

	n := len(rv.data)
	mean := rv.Mean()
	sum := 0.0
	for v := 0; v < n; v++ {
		sum += math.Pow((rv.data[v] - mean), 3)
	}

	return sum / (float64(n) * math.Pow(stdDev, 3)), nil
}

// Returns the value of the kurtosis of the data and an error when the standard deviation is 0
func (rv *RandVar) Kurtosis() (float64, error) {
	stdDev := rv.StdDev()
	if stdDev == 0 {
		return 0, ErrNullStdDeviation
	}

	n := len(rv.data)
	mean := rv.Mean()
	sum := 0.0
	for v := 0; v < n; v++ {
		sum += math.Pow((rv.data[v] - mean), 4)
	}

	return (sum / (float64(n) * math.Pow(stdDev, 4))), nil
}

// Returns the maximum value of the data
func (rv *RandVar) Max() float64 {
	max := rv.data[0]
	for _, v := range rv.data {
		if v > max {
			max = v
		}
	}
	return max
}

// Returns the minimum value of the data
func (rv *RandVar) Min() float64 {
	min := rv.data[0]
	for _, v := range rv.data {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns the range of the data
func (rv *RandVar) Range() float64 {
	return rv.Max() - rv.Min()
}

// Returns the covariance between two random variables
func (rv *RandVar) Covariance(rv1 *RandVar) (float64, error) {
	if len(rv.data) != len(rv1.data) {
		return 0, ErrDifferentLength
	}

	n := len(rv.data)
	mean := rv.Mean()
	meanRV1 := rv1.Mean()
	cov := 0.0
	for i := 0; i < n; i++ {
		cov += (rv.data[i] - mean) * (rv1.data[i] - meanRV1)
	}

	return cov / float64(n), nil
}

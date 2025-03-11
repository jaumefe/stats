package randvar

import (
	"slices"

	"github.com/jaumefe/stats"
)

// RandVar is a struct that represents a simple set of data of a random variable. All statistical parameters are computed as it was a whole population
type RandVar struct {
	data []float64
}

// Creates a new RandVar
func NewRandVar(data []float64) *RandVar {
	d := slices.Clone(data)
	rv := &RandVar{
		data: d,
	}
	return rv
}

// Returns the data of the random variable
func (rv *RandVar) Data() []float64 {
	return slices.Clone(rv.data)
}

// Returns the mean of the data. It will return 0 when data length is 0
func (rv *RandVar) Mean() (float64, error) {
	return stats.Mean(rv.data)
}

// Returns the median of the data
func (rv *RandVar) Median() (float64, error) {
	return stats.Median(rv.data)
}

func (rv *RandVar) Mode() (float64, error) {
	return stats.Mode(rv.data)
}

// Returns the variance of the data. It will return 0 when data length is 0
func (rv *RandVar) Variance() (float64, error) {
	return stats.Variance(rv.data)
}

// Returns the standard deviation of the data
func (rv *RandVar) StdDev() (float64, error) {
	return stats.StandardDeviation(rv.data)
}

// Returns the maximum value of the data
func (rv *RandVar) Max() (float64, error) {
	return stats.Max(rv.data)
}

// Returns the minimum value of the data
func (rv *RandVar) Min() (float64, error) {
	return stats.Min(rv.data)
}

// Returns the range of the data
func (rv *RandVar) Range() (float64, error) {
	return stats.Range(rv.data)
}

// Returns the sum of the data
func (rv *RandVar) Sum() float64 {
	return stats.Sum(rv.data)
}

// Returns the value of the skewness of the data and an error when the standard deviation is 0
func (rv *RandVar) Skewness() (float64, error) {
	return stats.Skewness(rv.data)
}

// Returns the value of the kurtosis of the data and an error when the standard deviation is 0
func (rv *RandVar) Kurtosis() (float64, error) {
	return stats.Skewness(rv.data)
}

func (rv *RandVar) Sort() []float64 {
	return stats.Sort(rv.data)
}

func (rv *RandVar) ReverseSort() []float64 {
	return stats.Sort(rv.data)
}

func (rv *RandVar) Frequency(epsilon float64) (map[float64]int, error) {
	return stats.Frequency(rv.data, epsilon)
}

// Returns the covariance between two random variables
func (rv *RandVar) Covariance(rv1 *RandVar) (float64, error) {
	if len(rv.data) != len(rv1.data) {
		return 0, stats.ErrDifferentLength
	}

	n := len(rv.data)
	mean, err := rv.Mean()
	if err != nil {
		return 0, err
	}

	meanRV1, err := rv1.Mean()
	if err != nil {
		return 0, err
	}

	cov := 0.0
	for i := 0; i < n; i++ {
		cov += (rv.data[i] - mean) * (rv1.data[i] - meanRV1)
	}

	return cov / float64(n), nil
}

// Returns the Pearson correlation coefficient between two random variables
func (rv *RandVar) Correlation(y *RandVar) (float64, error) {
	cov, err := rv.Covariance(y)
	if err != nil {
		return 0, err
	}

	stdDevX, err := rv.StdDev()
	if err != nil {
		return 0, err
	}

	stdDevY, err := y.StdDev()
	if err != nil {
		return 0, err
	}

	return cov / (stdDevX * stdDevY), nil
}

func (rv *RandVar) SpearmanCorrelation(y *RandVar) (float64, error) {
	if len(rv.data) == 0 || len(y.data) == 0 {
		return 0, stats.ErrEmptyData
	}

	if len(rv.data) != len(y.data) {
		return 0, stats.ErrDifferentLength
	}

	rankX, err := rv.rankVariable()
	if err != nil {
		return 0, err
	}

	rankY, err := y.rankVariable()
	if err != nil {
		return 0, err
	}

	n := len(rankX)
	var d float64
	for i := 0; i < n; i++ {
		d += rankX[i] - rankY[i]
	}

	return 1 - (6*d*d)/(float64(n)*(float64(n*n)-1)), nil
}

func (rv *RandVar) rankVariable() ([]float64, error) {
	sorted := rv.Sort()
	rank := make([]float64, len(sorted))
	freq, err := stats.Frequency(sorted, 1e8)
	if err != nil {
		return nil, err
	}

	for i, v := range sorted {
		if freq[v] == 1 {
			rank[i] = float64(i + 1)
		} else if rank[i] == 0 {
			f := freq[v]

			var r float64
			for j := i; j <= i+f; j++ {
				r += float64(j + 1)
			}

			for j := i; j <= i+f; j++ {
				rank[j] = r / float64(f)
			}
		}
	}

	return rank, nil
}

func (rv *RandVar) Normalize() (*RandVar, error) {
	norm, err := stats.Normalize(rv.data)
	if err != nil {
		return nil, err
	}

	return NewRandVar(norm), nil
}

func (rv *RandVar) Scale(factor float64) (*RandVar, error) {
	scaled, err := stats.Scale(rv.data, factor)
	if err != nil {
		return nil, err
	}

	return NewRandVar(scaled), nil
}

func (rv *RandVar) Entropy(logBase float64) (float64, error) {
	return stats.Entropy(rv.data, logBase)
}

func (rv *RandVar) Add(y *RandVar) (*RandVar, error) {
	if len(rv.data) != len(y.data) {
		return nil, stats.ErrDifferentLength
	}

	n := len(rv.data)
	added := make([]float64, n)
	for i := 0; i < n; i++ {
		added[i] = rv.data[i] + y.data[i]
	}

	return NewRandVar(added), nil
}

func (rv *RandVar) Subtract(y *RandVar) (*RandVar, error) {
	if len(rv.data) != len(y.data) {
		return nil, stats.ErrDifferentLength
	}

	n := len(rv.data)
	subs := make([]float64, n)
	for i := 0; i < n; i++ {
		subs[i] = rv.data[i] - y.data[i]
	}

	return NewRandVar(subs), nil
}

func (rv *RandVar) Product(y *RandVar) (*RandVar, error) {
	if len(rv.data) != len(y.data) {
		return nil, stats.ErrDifferentLength
	}

	n := len(rv.data)
	prod := make([]float64, n)
	for i := 0; i < n; i++ {
		prod[i] = rv.data[i] * y.data[i]
	}

	return NewRandVar(prod), nil
}

func (rv *RandVar) Shift(offset float64) (*RandVar, error) {
	n := len(rv.data)
	if n == 0 {
		return nil, stats.ErrEmptyData
	}

	shifted := make([]float64, n)
	for i := 0; i < n; i++ {
		shifted[i] = rv.data[i] + offset
	}

	return NewRandVar(shifted), nil
}

func (rv *RandVar) Quantile(qs float64, n uint) (float64, error) {
	return stats.Quantile(rv.data, qs, n)
}

func (rv *RandVar) Percentile(p float64) (float64, error) {
	return stats.Percentile(rv.data, p)
}

func (rv *RandVar) IQR() (float64, error) {
	return stats.IQR(rv.data)
}

func (rv *RandVar) Equals(y *RandVar, epsilon float64) bool {
	return stats.Equals(rv.data, y.data, epsilon)
}

func (rv *RandVar) Intersection(y *RandVar, epsilon float64) *RandVar {
	return NewRandVar(stats.Intersection(rv.data, y.data, epsilon))
}

func (rv *RandVar) Union(y *RandVar, epsilon float64) *RandVar {
	return NewRandVar(stats.Union(rv.data, y.data, epsilon))
}

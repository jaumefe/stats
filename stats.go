package stats

import (
	"math"
	"slices"
)

/*
Mean computes the mean value of a []float64 data input.
It returns an error if the data is empty
*/
func Mean(data []float64) (float64, error) {
	n := len(data)
	if len(data) == 0 {
		return 0, ErrEmptyData
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		sum += data[i]
	}

	return sum / float64(n), nil
}

/*
Median provides the median value of a []float64 data input.
Data input does not need to be sorted.
It returns an error if the data is empty
*/
func Median(data []float64) (float64, error) {
	n := len(data)
	if n == 0 {
		return 0, ErrEmptyData
	}

	sorted := Sort(data)
	if n%2 == 1 {
		return sorted[n/2], nil
	}

	return (sorted[n/2-1] + sorted[n/2]) / 2, nil
}

/*
Mode provides the mode value of a []float64 data input.
It returns an error if the data is empty
*/
func Mode(data []float64) (float64, error) {
	n := len(data)
	if n == 0 {
		return 0, ErrEmptyData
	}

	modeValues := make(map[float64]int, 0)
	for _, d := range data {
		if _, ok := modeValues[d]; !ok {
			modeValues[d] = 0
		}
		modeValues[d]++
	}

	mode := 0.0
	modeIdx := 0
	for m, i := range modeValues {
		if i > modeIdx {
			mode = m
			modeIdx = i
		}
	}
	return mode, nil
}

/*
Variance computes the variance value of a []float64 data input.
It returns an error if the data is empty
*/
func Variance(data []float64) (float64, error) {
	mean, err := Mean(data)
	if err != nil {
		return 0, err
	}

	sum := 0.0
	n := len(data)

	for _, v := range data {
		sum += math.Pow((v - mean), 2)
	}
	return sum / float64(n), nil
}

/*
StandardDeviation computes the standard deviation value of a []float64 data input.
It returns an error if the data is empty
*/
func StandardDeviation(data []float64) (float64, error) {
	variance, err := Variance(data)
	if err != nil {
		return 0, err
	}

	std := math.Sqrt(variance)
	return std, nil
}

/*
Max returns the maximum value of a []float64 data input.
It returns an error if the data is empty
*/
func Max(data []float64) (float64, error) {
	n := len(data)
	if n == 0 {
		return 0, ErrEmptyData
	}

	max := data[0]
	for _, d := range data {
		if d > max {
			max = d
		}
	}

	return max, nil
}

/*
Min returns the minimum value of a []float64 data input.
It returns an error if the data is empty
*/
func Min(data []float64) (float64, error) {
	n := len(data)
	if n == 0 {
		return 0, ErrEmptyData
	}

	min := data[0]
	for _, d := range data {
		if d < min {
			min = d
		}
	}

	return min, nil
}

/*
Range computes the difference between the maximum and the minimum value of a []float64 data input.
It returns an error if the data is empty
*/
func Range(data []float64) (float64, error) {
	max, err := Max(data)
	if err != nil {
		return 0, err
	}

	min, err := Min(data)
	if err != nil {
		return 0, err
	}

	return max - min, nil
}

/*
Sum computes the sum of all the elements of a []float64 data input.
*/
func Sum(data []float64) float64 {
	sum := 0.0
	for _, d := range data {
		sum += d
	}

	return sum
}

/*
Sort returns a sorted copy of the input []float64 data.
It returns an empty slice if the input data is empty
*/
func Sort(data []float64) []float64 {
	n := len(data)
	if n == 0 {
		return nil
	}
	sortData := make([]float64, len(data))
	copy(sortData, data)
	slices.Sort(sortData)
	return sortData
}

/*
ReverseSort returns a sorted copy from maximum to minimum value of the input []float64 data.
It returns an empty slice if the input data is empty
*/
func ReverseSort(data []float64) []float64 {
	n := len(data)
	if n == 0 {
		return nil
	}

	reversed := Sort(data)

	for i := 0; i < n/2; i++ {
		reversed[i], reversed[n-1-i] = reversed[n-1-i], reversed[i]
	}

	return reversed
}

/*
Normalize computes the normalized values of a []float64 data input.
It returns an error whether the standard deviation is null and / or the data is empty
*/
func Normalize(data []float64) ([]float64, error) {
	mean, err := Mean(data)
	if err != nil {
		return nil, err
	}

	std, err := StandardDeviation(data)
	if err != nil {
		return nil, err
	}

	if std == 0 {
		return nil, ErrNullStdDeviation
	}

	normalized := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		normalized[i] = (data[i] - mean) / std
	}

	return normalized, nil
}

/*
Scale computes the scaled values of a []float64 data input given a factor.
Factor must be different from 0.
It returns an error if the data is empty or the factor is null
*/
func Scale(data []float64, factor float64) ([]float64, error) {
	n := len(data)
	if n == 0 {
		return nil, ErrEmptyData
	}

	if factor == 0 {
		return nil, ErrNullScaleFactor
	}

	scaled := make([]float64, n)
	for i := 0; i < n; i++ {
		scaled[i] = data[i] * factor
	}
	return scaled, nil
}

/*
Equals returns a boolean value whether two []float64 data inputs are equal within a given epsilon.
Epsilon value can be set to 0 for exact comparison.
*/
func Equals(a, b []float64, epsilon float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if math.Abs(a[i]-b[i]) > epsilon {
			return false
		}
	}

	return true
}

/*
Intersection provides a []float64 with the common elements between two data inputs within a given epsilon.
Epsilon value can be set to 0 for exact comparison.
*/
func Intersection(a, b []float64, epsilon float64) []float64 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	intersection := make([]float64, 0)
	for _, v := range a {
		for _, w := range b {
			if math.Abs(v-w) <= epsilon {
				intersection = append(intersection, v)
				break
			}
		}
	}

	return intersection
}

/*
Union provides a []float64 with all the elements between two data inputs within a given epsilon.
Repeated elements are only included once.
Epsilon value can be set to 0 for exact comparison.
*/
func Union(a, b []float64, epsilon float64) []float64 {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	seen := make(map[float64]bool)
	union := make([]float64, 0)

	addIfNotSeen := func(value float64) {
		for existing := range seen {
			if math.Abs(existing-value) <= epsilon {
				return
			}
		}
		seen[value] = true
		union = append(union, value)
	}

	for _, v := range a {
		addIfNotSeen(v)
	}
	for _, v := range b {
		addIfNotSeen(v)
	}

	return union
}

/*
IQR computes the interquartile range of a []float64 data input.
It returns an error if the data is empty
*/
func IQR(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptyData
	}

	q3, err := Percentile(data, 75)
	if err != nil {
		return 0, err
	}

	q1, err := Percentile(data, 25)
	if err != nil {
		return 0, err
	}

	return q3 - q1, nil
}

/*
Percentile computes the percentile value of a []float64 data input given a percentage(%) value.
It returns an error if the data is empty or the percentage is out of range (0 - 100%)
*/
func Percentile(data []float64, p float64) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptyData
	}

	if p < 0 || p > 100 {
		return 0, ErrInvalidPercentile
	}

	sorted := Sort(data)
	n := len(sorted)
	pos := p * float64(n+1) / 100

	if pos-1 <= 0 {
		return sorted[0], nil
	}

	if pos >= float64(n) {
		return sorted[n-1], nil
	}

	if pos == float64(int(pos)) {
		return sorted[int(pos)-1], nil
	}

	lower := sorted[int(pos)-1]
	upper := sorted[int(pos)]
	weight := pos - float64(int(pos))

	return lower + (upper-lower)*weight, nil
}

/*
Quantile computes the quantile value of a []float64 data input given a quantile index.
- qs: desired quantile index
- n: total amount of quantiles
It returns an error if the data is emptty or the quantile index is out of range (0 - n)
*/
func Quantile(data []float64, qs float64, n uint) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptyData
	}

	if qs < 0 || qs > float64(n) {
		return 0, ErrInvalideQuantile
	}

	q := qs / float64(n)
	sorted := Sort(data)
	pos := q * float64(len(sorted)+1)
	if pos-1 <= 0 {
		return sorted[0], nil
	}

	if pos >= float64(n) {
		return sorted[n-1], nil
	}

	lower := sorted[int(pos)-1]
	upper := sorted[int(pos)]
	weight := pos - float64(int(pos))

	return lower + (upper-lower)*weight, nil
}

/*
Skewness computes the skewness value of a []float64 data input.
It returns an error if the data is empty or the standard deviation is null
*/
func Skewness(data []float64) (float64, error) {
	stdDev, err := StandardDeviation(data)
	if err != nil {
		return 0, err
	}

	if stdDev == 0 {
		return 0, ErrNullStdDeviation
	}

	n := len(data)
	mean, err := Mean(data)
	if err != nil {
		return 0, err
	}

	sum := 0.0
	for v := 0; v < n; v++ {
		sum += math.Pow((data[v] - mean), 3)
	}

	return sum / (float64(n) * math.Pow(stdDev, 3)), nil
}

/*
Kurtosis computes the kurtosis value of a []float64 data input.
It returns an error if the data is empty or the standard deviation is null
*/
func Kurtosis(data []float64) (float64, error) {
	stdDev, err := StandardDeviation(data)
	if err != nil {
		return 0, err
	}

	if stdDev == 0 {
		return 0, ErrNullStdDeviation
	}

	n := len(data)
	mean, err := Mean(data)
	if err != nil {
		return 0, err
	}

	sum := 0.0
	for v := 0; v < n; v++ {
		sum += math.Pow((data[v] - mean), 4)
	}

	return (sum/(float64(n)*math.Pow(stdDev, 4)) - 3), nil
}

/*
Frequency computes the frequency of each value of a []float64 data input within a given epsilon.
It returns an error if the data is empty
*/
func Frequency(data []float64, epsilon float64) (map[float64]int, error) {
	if len(data) == 0 {
		return nil, ErrEmptyData
	}

	freq := make(map[float64]int, 0)
	addIfNotSeen := func(value float64) {
		for existing := range freq {
			if math.Abs(existing-value) <= epsilon {
				freq[existing]++
				return
			}
		}
		freq[value]++
	}

	for _, v := range data {
		addIfNotSeen(v)
	}

	return freq, nil
}

/*
Entropy computes the entropy value of a []float64 data input.
It can be set to a specific log base:
  - logBase = 0: natural entropy

It returns an error if the data is empty or the log base is invalid
*/
func Entropy(data []float64, logBase float64) (float64, error) {
	n := len(data)
	if n == 0 {
		return 0, ErrEmptyData
	}

	if logBase == 1 || logBase < 0 {
		return 0, ErrInvalidLogBase
	}

	freq, err := Frequency(data, 1e-8)
	if err != nil {
		return 0, err
	}

	entropy := 0.0
	for _, f := range freq {
		p := float64(f) / float64(n)
		entropy -= p * math.Log(p) / math.Log(logBase)
	}

	return entropy, nil
}

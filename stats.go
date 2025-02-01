package stats

import (
	"math"
	"slices"
)

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

func StandardDeviation(data []float64) (float64, error) {
	variance, err := Variance(data)
	if err != nil {
		return 0, err
	}

	std := math.Sqrt(variance)
	return std, nil
}

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

func Sum(data []float64) float64 {
	sum := 0.0
	for _, d := range data {
		sum += d
	}

	return sum
}

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

func Percentile(data []float64, p float64) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptyData
	}

	if p < 0 || p > 100 {
		return 0, ErrInvalidPercentile
	}

	sorted := Sort(data)
	n := len(sorted)
	pos := p * float64(n) / 100

	if pos == float64(int(pos)) {
		return sorted[int(pos)-1], nil
	}

	lower := sorted[int(pos)-1]
	upper := sorted[int(pos)]
	weight := pos - float64(lower)

	return lower*(1-weight) + upper*weight, nil
}

func Quantile(data []float64, qs float64, n uint) (float64, error) {
	if len(data) == 0 {
		return 0, ErrEmptyData
	}

	if qs < 0 || qs > float64(n) {
		return 0, ErrInvalideQuantile
	}

	q := qs / float64(n)
	sorted := Sort(data)
	pos := q * float64(len(sorted))

	lower := sorted[int(pos)-1]
	upper := sorted[int(pos)]
	weight := pos - float64(lower)

	return lower*(1-weight) + upper*weight, nil
}

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

	return (sum / (float64(n) * math.Pow(stdDev, 4))), nil
}

func Frequency(data []float64, epsilon float64) (map[float64]int, error) {
	if len(data) == 0 {
		return nil, ErrEmptyData
	}

	freq := make(map[float64]int, 0)
	addIfNotSeen := func(value float64) {
		for existing := range freq {
			if math.Abs(existing-value) <= epsilon {
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

func Entropy(data []float64, logBase float64) (float64, error) {
	n := len(data)
	if n == 0 {
		return 0, ErrEmptyData
	}

	if logBase == 1 && logBase <= 0 {
		return 0, ErrInvalidLogBase
	}

	freq, err := Frequency(data, 1e-8)
	if err != nil {
		return 0, err
	}

	entropy := 0.0
	for _, f := range freq {
		p := float64(f) / float64(n)
		entropy -= p * math.Log(p) / math.Logb(logBase)
	}

	return entropy, nil
}

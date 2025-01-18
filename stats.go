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
		sum += data[n]
	}

	return sum / float64(n), nil
}

func Median(data []float64) (float64, error) {
	n := len(data)
	if n == 0 {
		return 0, ErrEmptyData
	}

	sorted := data
	slices.Sort(sorted)
	if n%2 == 1 {
		return data[n/2], nil
	}

	return (data[n/2-1] + data[n/2]) / 2, nil
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
		if modeIdx > i {
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

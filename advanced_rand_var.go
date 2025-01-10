package stats

import (
	"fmt"
	"math"
)

/*
AdvRandVar represents a random variable with additional features
Optimized for large datasets, so it is not necessary to recalculate
certain statistical parameters if needed for a certain parameter. For example:
When calculating the variance, it is necessary to obtain first the Mean,
so by using this variable, the mean is not recomputed.
*/
type AdvRandVar struct {
	*RandVar
	weight []float64

	min          float64
	max          float64
	rng          float64
	mean         float64
	weightedMean float64
	median       float64
	variance     float64
	stdDev       float64
	skewness     float64
	kurtosis     float64

	*meta
}

/*
Metadata in case the random variable is desired to be identified:
  - name
  - units
  - timestamp
  - src: Data source, such as a sensor, a database...
  - category: Label to classify a variable (it may be useful for multivariable analysis)
*/
type meta struct {
	name  string
	units string

	timestamp string
	src       string
	category  string
}

// Returns a new AdvRandVar. Meta data can be given to define AdvRandVar
func NewAdvRandVar(data []float64) *AdvRandVar {
	arv := &AdvRandVar{
		RandVar: NewRandVar(data),
	}
	return arv
}

// Defines Meta information about a AdvRandVar
func (arv *AdvRandVar) DefineMeta(name, units, timestamp, source, cat string) {
	if name != "" {
		arv.meta.name = name
	}
	if units != "" {
		arv.meta.units = units
	}
	if timestamp != "" {
		arv.meta.timestamp = timestamp
	}
	if source != "" {
		arv.meta.src = source
	}
	if cat != "" {
		arv.meta.category = cat
	}

}

// Computes the mean value of a AdvRandVar
func (arv *AdvRandVar) Mean() {
	arv.mean = arv.RandVar.Mean()
}

// Computes the median of a AdvRandVar
func (arv *AdvRandVar) Median() {
	arv.median = arv.RandVar.Median()
}

// Computes the variance of a AdvRandVar
func (arv *AdvRandVar) Variance() {
	sum := 0.0
	for _, v := range arv.RandVar.data {
		sum += math.Pow((v - arv.mean), 2)
	}
	arv.variance = sum / float64(len(arv.RandVar.data)-1)
}

// Computes the standard deviation of a AdvRandVar
func (arv *AdvRandVar) StdDev() {
	arv.stdDev = math.Sqrt(arv.variance)
}

// Computes the value of the skewness of a AdvRandVar
func (arv *AdvRandVar) Skewness() {
	if arv.stdDev == 0 {
		arv.skewness = 0
		return
	}

	n := len(arv.RandVar.data)
	sum := 0.0
	for v := 0; v < n; v++ {
		sum += math.Pow((arv.RandVar.data[v] - arv.mean), 3)
	}

	arv.skewness = sum / (float64(n) * math.Pow(arv.stdDev, 3))
}

// Computes the value of the kurtosis of a AdvRandVar
func (arv *AdvRandVar) Kurtosis() {
	if arv.stdDev == 0 {
		arv.kurtosis = 0
		return
	}

	n := len(arv.RandVar.data)
	sum := 0.0
	for v := 0; v < n; v++ {
		sum += math.Pow((arv.RandVar.data[v] - arv.mean), 4)
	}

	arv.kurtosis = (sum / (float64(n) * math.Pow(arv.stdDev, 4))) - 3
}

// Computes the maximum value of the dataset of an AdvRandVar
func (arv *AdvRandVar) Max() {
	arv.max = arv.RandVar.Max()
}

// Computes the minimum value of the dataset of an AdvRandVar
func (arv *AdvRandVar) Min() {
	arv.min = arv.RandVar.Min()
}

// Computes the range of the dataset of an AdvRandVar
func (arv *AdvRandVar) Range() {
	arv.rng = arv.max - arv.min
}

// Computes a weighted mean of the dataset of an AdvRandVar
func (arv *AdvRandVar) WeightedMean() error {
	if len(arv.RandVar.data) != len(arv.weight) {
		return fmt.Errorf("data and Weight arrays must have the same length")
	}

	for i, v := range arv.RandVar.data {
		arv.weightedMean += v * arv.weight[i]
	}
	arv.weightedMean = arv.weightedMean / float64(len(arv.RandVar.data))
	return nil
}

func (arv *AdvRandVar) Update() error {
	arv.Mean()
	arv.Median()
	arv.Variance()
	arv.StdDev()
	arv.Skewness()
	arv.Kurtosis()
	arv.Max()
	arv.Min()
	arv.Range()
	arv.WeightedMean()
	return nil
}

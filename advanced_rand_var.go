package stats

import (
	"fmt"
	"log"
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

func (arv *AdvRandVar) SetWeight(w []float64) error {
	if len(w) != len(arv.data) {
		return fmt.Errorf("length of weight and data are different: Weight:%d, Data: %d", len(w), len(arv.data))
	}

	wTemp := append([]float64(nil), w...)
	arv.weight = wTemp
	return nil
}

// Computes the mean value of a AdvRandVar
func (arv *AdvRandVar) updateMean() {
	arv.mean = arv.RandVar.Mean()
}

// Computes the median of a AdvRandVar
func (arv *AdvRandVar) updateMedian() {
	arv.median = arv.RandVar.Median()
}

// Computes the variance of a AdvRandVar
func (arv *AdvRandVar) updateVariance() {
	sum := 0.0
	if len(arv.RandVar.data) == 0 {
		arv.variance = sum
		return
	}

	for _, v := range arv.RandVar.data {
		sum += math.Pow((v - arv.mean), 2)
	}
	arv.variance = sum / float64(len(arv.RandVar.data))
}

// Computes the standard deviation of a AdvRandVar
func (arv *AdvRandVar) updateStdDev() {
	arv.stdDev = math.Sqrt(arv.variance)
}

// Computes the value of the skewness of a AdvRandVar
func (arv *AdvRandVar) updateSkewness() {
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
func (arv *AdvRandVar) updateKurtosis() {
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
func (arv *AdvRandVar) updateMax() {
	arv.max = arv.RandVar.Max()
}

// Computes the minimum value of the dataset of an AdvRandVar
func (arv *AdvRandVar) updateMin() {
	arv.min = arv.RandVar.Min()
}

// Computes the range of the dataset of an AdvRandVar
func (arv *AdvRandVar) updateRange() {
	arv.rng = arv.max - arv.min
}

// Computes a weighted mean of the dataset of an AdvRandVar
func (arv *AdvRandVar) updateWeightedMean() error {
	if arv.weight != nil {
		return fmt.Errorf("weight not defined")
	}

	for i, v := range arv.RandVar.data {
		arv.weightedMean += v * arv.weight[i]
	}
	arv.weightedMean = arv.weightedMean / float64(len(arv.RandVar.data))
	return nil
}

// Type to set exclusions to Update() function
// True: Exclude the field
type OptsExclusionUpdate struct {
	Skewness     bool
	Kurtosis     bool
	Max          bool
	Min          bool
	Range        bool
	WeightedMean bool
}

/*
It can compute the following statistical data:
- Mean
- Median
- Variance
- Standard Deviation
- Skewness
- Kurtosis
- Maximum value
- Minimum value
- Range
- Weighted Mean: If weights is not defined, it will print a log message

Exclusion can be added through parameters to certain fields: setting to `true` excludes from the calculation

The function must be called every time a modification of data is done to recompute the new statistical parameters
*/
func (arv *AdvRandVar) Update(opts *OptsExclusionUpdate) {
	arv.updateMean()
	arv.updateMedian()
	arv.updateVariance()
	arv.updateStdDev()

	if !opts.Skewness {
		arv.updateSkewness()
	}

	if !opts.Kurtosis {
		arv.updateKurtosis()
	}

	if !opts.Max {
		arv.updateMax()
	}

	if !opts.Min {
		arv.updateMin()
	}

	if !opts.Range {
		arv.updateRange()
	}

	if !opts.WeightedMean {
		err := arv.updateWeightedMean()
		if err != nil {
			log.Print(err)
		}
	}
}

// Returns stored mean value of an AdvRandVar
func (arv *AdvRandVar) Mean() float64 {
	return arv.mean
}

// Returns stored median value of an AdvRandVar
func (arv *AdvRandVar) Median() float64 {
	return arv.median
}

// Returns stored variance value of an AdvRandVar
func (arv *AdvRandVar) Variance() float64 {
	return arv.variance
}

// Returns stored standard deviation value of an AdvRandVar
func (arv *AdvRandVar) StdDev() float64 {
	return arv.stdDev
}

// Returns stored skewness of an AdvRandVar
func (arv *AdvRandVar) Skewness() float64 {
	return arv.skewness
}

// Returns stored kurtosis of an AdvRandVar
func (arv *AdvRandVar) Kurtosis() float64 {
	return arv.kurtosis
}

// Returns stored maximum value of data of an AdvRandVar
func (arv *AdvRandVar) Max() float64 {
	return arv.max
}

// Returns stored minimum value of data of an AdvRandVar
func (arv *AdvRandVar) Min() float64 {
	return arv.min
}

// Returns stored range of data of an AdvRandVar
func (arv *AdvRandVar) Range() float64 {
	return arv.rng
}

// Returns stored range of data of an AdvRandVar
func (arv *AdvRandVar) WeightedMean() float64 {
	return arv.weightedMean
}

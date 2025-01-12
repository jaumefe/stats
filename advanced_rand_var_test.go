package stats

import (
	"reflect"
	"testing"
	"time"
)

func TestNewAdvRandVar(t *testing.T) {
	data := []float64{1.0, 3.5, 2.2}
	arv := NewAdvRandVar(data)

	if !reflect.DeepEqual(data, arv.data) {
		t.Errorf("Expected %v, but got %v", data, arv.data)
	}

	// Let's test if the data is immutable despite modifying the original
	data[0] = 22.2
	if arv.data[0] == data[0] {
		t.Errorf("A modification on original data has modified the data of the random variable")
	}
}

func TestDefineMeta(t *testing.T) {
	name := "test"
	units := "u"
	timestamp := time.RFC1123
	src := "src1"
	cat := "label1"

	data := []float64{1.0, 3.5, 2.2}
	arv := NewAdvRandVar(data)
	arv.DefineMeta(name, units, timestamp, src, cat)
	if arv.name != name {
		t.Errorf("expected name: %v, got: %v", name, arv.name)
	}

	if arv.units != units {
		t.Errorf("expected units: %v, got: %v", units, arv.units)
	}

	if arv.timestamp != timestamp {
		t.Errorf("expected timestamp: %v, got: %v", timestamp, arv.timestamp)
	}

	if arv.src != src {
		t.Errorf("expected source: %v, got: %v", src, arv.src)
	}

	if arv.category != cat {
		t.Errorf("expected category: %v, got: %v", cat, arv.category)
	}

	nameOld := name
	name = "newTest"
	if arv.name != nameOld {
		t.Errorf("A modification on original data has modified the data of the random variable")
	}

}

func TestWeightMean(t *testing.T) {
	data := []float64{1.0, 3.5, 2.2}
	arv := NewAdvRandVar(data)

	if err := arv.updateWeightedMean(); err == nil {
		t.Error("not error returned when no weights defined")
	}

	w1 := []float64{0.5, 1.2}
	if err := arv.SetWeight(w1); err == nil {
		t.Error("SetWeight must not allow to work when lenghts are different")
	}

	w2 := []float64{0.5, 1.2, 0.75}
	if err := arv.SetWeight(w2); err != nil {
		t.Error("Sizes are identical, however the error returned is not nil")
	}

	w2[0] = 3.1
	if arv.weight[0] == w2[0] {
		t.Error("A modification on original data has modified the data of the random variable")
	}

	expected := (0.5*1.0 + 1.2*3.5 + 0.75*2.2) / (0.5 + 1.2 + 0.75)
	arv.updateWeightedMean()
	if arv.weightedMean != expected {
		t.Errorf("expected %f, got %f", expected, arv.weightedMean)
	}
}

func TestUpdate(t *testing.T) {
	// Testing full functionality
	data := []float64{0.5, 1.2, 5.3, 7.5, 2.4, 10.0, 9.1, 8.4, 6.6, 5.5, 5.35, 9.75, 2.25, 7.2, 8.4, 6.6, 3.75, 4.25, 6.9, 7.8}
	arv := NewAdvRandVar(data)
	rv := NewRandVar(data)
	opts := &OptsExclusionUpdate{}
	arv.Update(opts)
	e := 1e-12

	if arv.Mean()-rv.Mean() > e {
		t.Errorf("Expected mean: %f, got %f", rv.Mean(), arv.Mean())
	}

	if arv.Median()-rv.Median() > e {
		t.Errorf("Expected median: %f, got %f", rv.Median(), arv.Median())
	}

	if arv.Variance()-rv.Variance() > e {
		t.Errorf("Expected variance: %f, got %f", rv.Variance(), arv.Variance())
	}

	if arv.StdDev()-rv.StdDev() > e {
		t.Errorf("Expected standard deviation: %f, got %f", rv.StdDev(), arv.StdDev())
	}

	skewness, _ := rv.Skewness()
	if arv.Skewness()-skewness > e {
		t.Errorf("Expected skewness: %f, got %f", skewness, arv.Skewness())
	}

	kurtosis, _ := rv.Kurtosis()
	if arv.Kurtosis()-kurtosis > e {
		t.Errorf("Expected kurtosis: %f, got %f", kurtosis, arv.Kurtosis())
	}

	// Testing 1 parameter
	data = []float64{1.1}
	arv = NewAdvRandVar(data)
	arv.Update(opts)

	if arv.Mean() != data[0] {
		t.Errorf("Expected mean: %f, got %f", data[0], arv.Mean())
	}

	if arv.Median() != data[0] {
		t.Errorf("Expected median: %f, got %f", data[0], arv.Median())
	}

	if arv.Variance() != 0 {
		t.Errorf("Expected variance: %f, got %f", 0.0, arv.Variance())
	}

	if arv.Skewness() != 0 {
		t.Errorf("Expected skewness: %f, got %f", 0.0, arv.Skewness())
	}

	if arv.Kurtosis() != 0 {
		t.Errorf("Expected kurtosis: %f, got %f", 0.0, arv.Kurtosis())
	}

	// Testing nil data
	data = []float64{}
	arv = NewAdvRandVar(data)
	arv.Update(opts)

	if arv.Mean() != 0.0 {
		t.Errorf("Expected mean: %f, got %f", 0.0, arv.Mean())
	}

	if arv.Median() != 0.0 {
		t.Errorf("Expected median: %f, got %f", 0.0, arv.Median())
	}

	if arv.Variance() != 0 {
		t.Errorf("Expected variance: %f, got %f", 0.0, arv.Variance())
	}

	if arv.Skewness() != 0 {
		t.Errorf("Expected skewness: %f, got %f", 0.0, arv.Skewness())
	}

	if arv.Kurtosis() != 0 {
		t.Errorf("Expected kurtosis: %f, got %f", 0.0, arv.Kurtosis())
	}

	// Testing opts not to update
	data = []float64{0.5, 1.2, 5.3, 7.5, 2.4, 10.0, 9.1, 8.4, 6.6, 5.5, 5.35, 9.75, 2.25, 7.2, 8.4, 6.6, 3.75, 4.25, 6.9, 7.8}
	arv = NewAdvRandVar(data)
	arv.Update(opts)
	skewness, kurtosis, max, min, rng := arv.Skewness(), arv.Kurtosis(), arv.Max(), arv.Min(), arv.Range()

	opts = &OptsExclusionUpdate{
		Skewness: true,
		Kurtosis: true,
		Max:      true,
		Min:      true,
		Range:    true,
	}
	arv.Update(opts)

	if arv.Skewness() == skewness || arv.Skewness() != 0 {
		t.Errorf("expected null skewness and got: %f", arv.Skewness())
	}

	if arv.Kurtosis() == kurtosis || arv.Kurtosis() != 0 {
		t.Errorf("expected null kurtosis and got: %f", arv.Kurtosis())
	}

	if arv.Max() == max || arv.Max() != 0 {
		t.Errorf("expected null maximum value and got: %f", arv.Max())
	}

	if arv.Min() == min || arv.Min() != 0 {
		t.Errorf("expected null minimum value and got: %f", arv.Min())
	}

	if arv.Range() == rng || arv.Range() != 0 {
		t.Errorf("expected null range value and got: %f", arv.Range())
	}
}

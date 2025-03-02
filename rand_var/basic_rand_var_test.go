package randvar

import (
	"reflect"
	"testing"
)

type testrv struct {
	name     string
	data     []float64
	expected float64
}

func TestNewRandVar(t *testing.T) {
	data := []float64{1.0, 3.5, 2.2}
	rv := NewRandVar(data)

	if !reflect.DeepEqual(data, rv.data) {
		t.Errorf("Expected %v, but got %v", data, rv.data)
	}

	// Let's test if the data is immutable despite modifying the original
	data[0] = 22.2
	if rv.data[0] == data[0] {
		t.Errorf("A modification on original data has modified the data of the random variable")
	}
}

func TestCovariance(t *testing.T) {
	dataX := []float64{50.2, 60.3, 45.23, 55.75, 70.91}
	dataY := []float64{24.6, 41.9, 33.33, 27.9, 44.1}
	x, y := NewRandVar(dataX), NewRandVar(dataY)

	covXY, err := x.Covariance(y)
	if err != nil {
		t.Errorf("Unexpected error :%v", err)
	}

	covYX, err := y.Covariance(x)
	if err != nil {
		t.Errorf("Unexpected error :%v", err)
	}

	if covXY != covYX {
		t.Errorf("Covariance of two random variable are not equal: COV(X, Y):%f, COV(Y, X): %f", covXY, covYX)
	}

	dataX = []float64{50.2, 60.3, 45.23, 70.91}
	dataY = []float64{24.6, 41.9, 33.33, 27.9, 44.1}
	x, y = NewRandVar(dataX), NewRandVar(dataY)
	_, err = x.Covariance(y)
	if err == nil {
		t.Errorf("Length of both random variables must be equal: len(x):%d; len(y):%d", len(x.data), len(y.data))
	}

	_, err = y.Covariance(x)
	if err == nil {
		t.Errorf("Length of both random variables must be equal: len(x):%d; len(y):%d", len(x.data), len(y.data))
	}
}

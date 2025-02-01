package stats

import "errors"

var (
	ErrEmptyData         = errors.New("data contains no values")
	ErrNullScaleFactor   = errors.New("null scale factor given")
	ErrDifferentLength   = errors.New("different lengths on data")
	ErrNullStdDeviation  = errors.New("standard deviation is null")
	ErrInvalidPercentile = errors.New("percentile must be between 0 and 100")
	ErrInvalideQuantile  = errors.New("quantile must be between 0 and maximum quantile number")
	ErrInvalidLogBase    = errors.New("logarithm base must be greater than 0 and not equal to 1")
)

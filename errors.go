package stats

import "errors"

var (
	ErrEmptyData        = errors.New("data contains no values")
	ErrNullScaleFactor  = errors.New("null scale factor given")
	ErrDifferentLength  = errors.New("different lengths on data")
	ErrNullStdDeviation = errors.New("standard deviation is null")
)

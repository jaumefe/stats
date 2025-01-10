package stats

import "errors"

var (
	ErrDifferentLength  = errors.New("different lengths on data")
	ErrNullStdDeviation = errors.New("standard deviation is null")
)

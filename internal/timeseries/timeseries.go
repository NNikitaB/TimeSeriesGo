package timeseries

import (
	"math/big"

	"golang.org/x/exp/constraints"
)

type AcceptBaseType interface {
	~float64 | ~int64
}

type AcceptLongBaseType interface {
	int64 | constraints.Complex | big.Int | big.Float
}

// Genegics for Time Series
type TimeSeries[baseType AcceptBaseType] struct {
	listSeries []baseType
}

// Constructor empty Time Series
func (ts *TimeSeries[baseType]) New() *TimeSeries[baseType] {
	ts.listSeries = make([]baseType, 0)
	return &TimeSeries[baseType]{listSeries: ts.listSeries}
}

// Constructor Time Series
func (ts *TimeSeries[baseType]) NewFrom(t []baseType) *TimeSeries[baseType] {
	ts.listSeries = make([]baseType, len(t))
	copy(ts.listSeries, t)
	return &TimeSeries[baseType]{}
}

// Mean Time Series
func (ts *TimeSeries[baseType]) Mean() (baseType, error) {
	var sum baseType
	for _, val := range ts.listSeries {
		sum += val
	}
	n := len(ts.listSeries)
	if n == 0 {
		return sum, ErrEmptyInput
	}
	return sum / baseType(n), nil
}

// MovingAverage Time Series
func (ts *TimeSeries[baseType]) MovingAverage(windowSize int) ([]baseType, error) {
	var movingAvg []baseType
	var sum baseType
	var zero baseType
	var window = baseType(windowSize)
	if windowSize < 2 || len(ts.listSeries) < 10 {
		return movingAvg, ErrSmallData
	}
	for i := windowSize - 1; i < len(ts.listSeries); i++ {
		sum = zero
		for j := i - (windowSize - 1); j <= i; j++ {
			sum += ts.listSeries[j]
		}
		movingAvg = append(movingAvg, sum/window)
	}
	sum /= window
	for i := 0; i < windowSize; i++ {
		movingAvg = append(movingAvg, sum)
	}
	return movingAvg, nil
}

func (ts *TimeSeries[baseType]) Append(value baseType) {
	ts.listSeries = append(ts.listSeries, value)
}

// Get basic type into Time Series
func (ts *TimeSeries[baseType]) GetSeries() []baseType {
	return ts.listSeries
}

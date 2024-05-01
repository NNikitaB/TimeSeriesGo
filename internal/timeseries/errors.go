package timeseries

type errorTimeSeries struct {
	err string
}

func (s errorTimeSeries) Error() string {
	return s.err
}

func (s errorTimeSeries) String() string {
	return s.err
}

var (
	// ErrEmptyInput Input must not be empty
	ErrEmptyInput = errorTimeSeries{"Input must not be empty."}
	// ErrSmallData Input must be tiny
	ErrSmallData = errorTimeSeries{"Input must be tiny."}
	// ErrNaN Not a number
	ErrNaN = errorTimeSeries{"Not a number."}
	// ErrNegative Must not contain negative values
	ErrZero = errorTimeSeries{"Must not contain zero values."}
	// ErrBounds Input is outside of range
	ErrSize = errorTimeSeries{"Must be the same length."}
	// ErrInfValue Value is infinite
	ErrInfValue = errorTimeSeries{"Value is infinite."}
	// ErrInfValue Type not supported
	ErrTypeNonSupported = errorTimeSeries{"Type not supported."}
)

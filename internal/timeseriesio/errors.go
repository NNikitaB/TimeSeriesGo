package timeseriesio

type errorTSReader struct {
	err string
}

func (s errorTSReader) Error() string {
	return s.err
}

func (s errorTSReader) String() string {
	return s.err
}

var (
	// ErrorPattern Pattern not equel or empty
	ErrorPattern = errorTSReader{"Pattern not equel or empty."}
	// ErrSmallData Input must be tiny
	ErrMethodNonReliased = errorTSReader{"Method Non Reliased."}
)

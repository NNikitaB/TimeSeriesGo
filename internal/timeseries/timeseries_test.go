package timeseries

import (
	"testing"
)

func TestTimeSeries(t *testing.T) {
	ts := TimeSeries[int64]{}
	date_int := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	ts.NewFrom(date_int)
	mean, ok := ts.Mean()
	if ok != nil {
		t.Errorf("Error method Mean: %s", ok)
	}
	if mean != 8 {
		t.Errorf("Error computing Mean")
	}
	ma, ok := ts.MovingAverage(10)
	if ok != nil {
		t.Errorf("Error method MovingAverage: %s", ok)
	}
	if ma == nil {
		t.Errorf("Error computing MovingAverage")
	}
}

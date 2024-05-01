package timeseries

import (
	"testing"
)

func TestError(t *testing.T) {
	err := errorTimeSeries{"test error"}
	if err.Error() != "test error" {
		t.Errorf("Error method message didn't match")
	}
	if err.String() != "test error" {
		t.Errorf("String method message didn't match")
	}
}

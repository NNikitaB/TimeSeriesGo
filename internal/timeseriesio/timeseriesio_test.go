package timeseriesio

import (
	"os"
	"testing"
)

func TestTSReader(t *testing.T) {
	f, err := os.Open("../../examples/GAZP.txt")
	if err != nil {
		t.Errorf("Error Open file: %s", err)
	}
	defer f.Close()
	var r ITSReader = &TSReader{}
	ts, err := r.ReadFromCSV(f, "<CLOSE>")
	if err != nil {
		t.Errorf("Error Open file: %s", err)
	}
	if len(ts) == 0 {
		t.Errorf("Error Reading file may be empty: %s", err)
	}
}

package svgconverter

import (
	"os"
	"testing"
)

func TestTSReader(t *testing.T) {
	f, err := os.Create("../../examples/test_svg.svg")
	if err != nil {
		t.Errorf("Error Open file: %s", err)
	}
	defer f.Close()
	xAxis := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	yAxis := []float64{41.0, 22.0, 33.0, 44.0, 55.0, 16.0, 47.0, 83.0, 19.0, 100.0}
	zAxis := []float64{22.0, 22.0, 22.0, 22.0, 22.0, 22.0, 22.0, 22.0, 22.0, 22.0}
	ConvertToSVG(f, xAxis, yAxis, xAxis, zAxis)
}

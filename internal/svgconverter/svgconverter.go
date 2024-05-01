package svgconverter

import (
	"io"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// Converting X,Y - Chart1 and Xma, Yma - Chart2 to svg
func ConvertToSVG(r io.Writer, X, Y, Xma, Yma []float64) {
	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 20,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    "Base chart",
				XValues: X,
				YValues: Y,
				Style: chart.Style{
					StrokeColor: drawing.ColorGreen,               // will supercede defaults
					FillColor:   drawing.ColorGreen.WithAlpha(64), // will supercede defaults
				},
			},
			chart.ContinuousSeries{
				Name:    "Moving Average",
				XValues: Xma,
				YValues: Yma,
				Style: chart.Style{
					StrokeColor: drawing.ColorRed, // will supercede defaults
				},
			},
		},
		YAxis: chart.YAxis{
			Style:     chart.Shown(),
			NameStyle: chart.Shown(),
			Range:     &chart.LogarithmicRange{},
		},
	}
	graph.Render(chart.SVG, r)
}

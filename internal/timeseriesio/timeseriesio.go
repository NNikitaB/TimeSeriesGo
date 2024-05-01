package timeseriesio

import (
	"encoding/csv"
	"io"
	"strconv"
)

type ITSReader interface {
	ReadFromJSON(reader io.Reader, patern string) ([]float64, error)
	ReadFromCSV(reader io.Reader, patern string) ([]float64, error)
}

type TSReader struct {
}

func (w *TSReader) New() *TSReader {
	return &TSReader{}
}

// Read []float From CSV
func (w *TSReader) ReadFromCSV(reader io.Reader, patern string) ([]float64, error) {
	r := csv.NewReader(reader)
	r.Comma = ';'
	listSeries := make([]float64, 0, 100)
	columnNames, er := r.Read()
	selectIndex := -1
	if er != nil {
		return listSeries, er
	}
	for i := range columnNames {
		if columnNames[i] == patern {
			selectIndex = i
		}
	}
	if selectIndex == -1 {
		return listSeries, ErrorPattern
	}
	for {
		record, e := r.Read()
		if e != nil {
			break
		}
		f, ok := strconv.ParseFloat(record[selectIndex], 64)
		if ok != nil {
			continue
		}
		listSeries = append(listSeries, f)
	}
	return listSeries, nil
}

// Read []float From JSON
func (w *TSReader) ReadFromJSON(reader io.Reader, patern string) ([]float64, error) {
	return []float64{}, ErrMethodNonReliased
}

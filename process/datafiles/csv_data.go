package datafiles

import (
	"os"
	"encoding/csv"
)


type CsvData struct {
	Key string
	File string
	Data [][]string
}

func NewCsvData(key, datafile string) *CsvData {
	return &CsvData{
		Key: key,
		File: datafile,
	}
}

func (d *CsvData) Parse() (*CsvData, error) {
	f, err := os.Open(d.File)
	if err != nil {
		return d, err
	}
	defer f.Close()
	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		return d, err
	}
	d.Data = records

	return d, nil
}


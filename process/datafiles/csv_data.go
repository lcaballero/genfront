package datafiles

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/lcaballero/genfront/process"
	"log"
	"github.com/lcaballero/genfront/cli"
)

type CsvData struct {
	Keyed     cli.DataFile
	Data      [][]string
	Delimiter rune
}

func NewCsvData(keyed cli.DataFile, delimiter rune) *CsvData {
	return &CsvData{
		Keyed: keyed,
		Delimiter: delimiter,
	}
}

func (d *CsvData) Parse() (*CsvData, error) {
	f, err := os.Open(d.Keyed.File)
	if err != nil {
		return d, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true

	reader.Comma = d.Delimiter

	log.Println("Using delimiter " + string(reader.Comma))

	records, err := reader.ReadAll()
	if err != nil {
		return d, err
	}
	d.Data = records

	return d, nil
}

func (d *CsvData) HasData() bool {
	return d.Data != nil && len(d.Data) > 0
}

func (d *CsvData) MapFieldNames() ([]map[string]interface{}, error) {
	if !d.HasData() {
		return nil, fmt.Errorf("Csv does't have data")
	}

	headers := d.Data[0]
	fields := d.Data[1:]

	for i, header := range headers {
		headers[i] = process.ToSymbol(header)
	}

	if len(fields) < 1 {
		return nil, fmt.Errorf("Csv doesn't have any field data.")
	}

	data := make([]map[string]interface{}, 0)

	for _, line := range fields {
		if len(line) != len(headers) {
			return nil, fmt.Errorf("Fields and Headers must agree in length")
		}
		vals := make(map[string]interface{})
		data = append(data, vals)
		for i, header := range headers {
			vals[header] = line[i]
		}
	}
	return data, nil
}

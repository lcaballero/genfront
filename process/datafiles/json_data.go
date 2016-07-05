package datafiles

import (
	"io/ioutil"
	"encoding/json"
	"github.com/lcaballero/genfront/cli"
)


type JsonData struct {
	Keyed cli.DataFile
	Data map[string]interface{}
}


func NewJsonData(keyed cli.DataFile) *JsonData {
	return &JsonData{
		Keyed: keyed,
	}
}

func (d *JsonData) Unmarshal() (interface{}, error) {
	bits, err := ioutil.ReadFile(d.Keyed.File)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(bits, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

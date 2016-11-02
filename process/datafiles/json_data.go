package datafiles

import (
	"encoding/json"
	"io/ioutil"
)

type JsonData struct {
	Key  string
	File string
	Data interface{}
}

func NewJsonData(key, datafile string) *JsonData {
	return &JsonData{
		Key:  key,
		File: datafile,
		Data: []string{},
	}
}

func (j *JsonData) Parse() (*JsonData, error) {
	bits, err := ioutil.ReadFile(j.File)
	if err != nil {
		return j, err
	}

	err = json.Unmarshal(bits, &j.Data)
	if err != nil {
		return j, err
	}
	return j, nil
}

func (j *JsonData) HasData() bool {
	return j.Data != nil
}

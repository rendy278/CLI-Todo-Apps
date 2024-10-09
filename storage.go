package main

import (
	"encoding/json"
	"os"
)

type Storange[To any] struct {
	FileName string
}

func NewStorage[To any](fileName string) *Storange[To] {
	return &Storange[To]{FileName: fileName}
}

func (s *Storange[To]) Save(data To) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storange[To]) Load(data *To) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}

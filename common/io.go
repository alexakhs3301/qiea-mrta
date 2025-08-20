package common

import (
	"encoding/json"
	"os"
)

func LoadInputFromJSON(filename string) (*Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var input Input
	if err := decoder.Decode(&input); err != nil {
		return nil, err
	}

	return &input, nil
}

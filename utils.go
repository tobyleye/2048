package main

import (
	"encoding/json"
	"os"
)

func saveBestScore(file *os.File, score int) error {
	err := file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil
	}
	return json.NewEncoder(file).Encode(score)

}

func readBestScore(file *os.File) int {
	var score int
	json.NewDecoder(file).Decode(&score) // ignore error
	return score
}

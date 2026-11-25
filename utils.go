package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

const (
	oneTrillion = 1_000_000_000_000
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

// suffix for 1000^i.
var suffixes = []string{"", "K", "M", "B"}

// formatTile shortens a tile value so it fits in a cell
func formatTile(tile int) string {
	if tile >= oneTrillion {
		return fmt.Sprintf("2^%d", bits.TrailingZeros(uint(tile)))
	}
	if tile < 10_000 {
		return strconv.Itoa(tile)
	}

	exp := (len(strconv.Itoa(tile)) - 1) / 3
	f := float64(tile) / math.Pow(1000, float64(exp))

	s := strconv.FormatFloat(f, 'f', 6, 64)[:4]
	return strings.TrimSuffix(s, ".") + suffixes[exp]
}

package main

import (
	"encoding/json"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
)

const (
	gridCols   = 4
	gridRows   = 4
	cellHeight = 3
	cellWidth  = cellHeight * 2
	gridGapY   = 1
	gridGapX   = gridGapY * 2

	gridWidth = (cellWidth * gridRows) + (gridGapX * (gridRows + 1))

	gridHeight = (cellHeight * gridCols) + (gridGapY * (gridCols + 1))
)

func openBestScore() (int, *os.File, error) {
	scoreFile, err := os.ReadFile("./score.json")

	var bestScore int
	if err == nil {
		json.Unmarshal(scoreFile, &bestScore)
	}

	f, err := os.OpenFile("./score.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return bestScore, nil, err
	}

	return bestScore, f, nil

}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	defer f.Close()

	bestScore, bestScoreFile, err := openBestScore()

	if err != nil {
		log.Println("error reading best score..", err)
	} else {
		defer bestScoreFile.Close()
	}

	board := createBoard(gridRows, gridCols)

	// initialize with 2 cells
	board.addToRandomCell()
	board.addToRandomCell()

	m := model{
		board:     board,
		bestScore: bestScore,
		scoreFile: bestScoreFile,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}
}

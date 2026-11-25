package main

import (
	"log"
	"os"
	"path"

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

func getBestScorePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// ?choose a path relative to where program is running from
		return ""
	}
	return path.Join(homeDir, "2048_best_score")
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	defer f.Close()

	bestScoreFilePath := getBestScorePath()

	var bestScoreFile *os.File = nil

	if bestScoreFilePath != "" {
		bestScoreFile, err = os.OpenFile(bestScoreFilePath, os.O_CREATE|os.O_RDWR,
			0644)

		if err != nil {
			log.Println(
				"error reading best score",
				err,
			)
		} else {
			defer bestScoreFile.Close()
		}

	}

	board := createBoard(gridRows, gridCols)

	// start with 2 tiles
	board.spawnRandomTile(2)

	m := model{
		board:     board,
		scoreFile: bestScoreFile,
		bestScore: readBestScore(bestScoreFile),
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}
}

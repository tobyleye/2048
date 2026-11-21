package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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

var cellStyle = lipgloss.NewStyle().
	Width(cellWidth).
	Height(cellHeight).
	Background(lipgloss.Green).
	Align(lipgloss.Center, lipgloss.Center)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer f.Close()

	board := createBoard(gridRows, gridCols)

	// initialize with 2 cells
	board.addToRandomCell()
	board.addToRandomCell()

	m := model{
		board: board,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}
}

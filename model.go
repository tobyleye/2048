package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	posX      int
	posY      int
	width     int
	board     Board
	scoreFile *os.File
	score     int
	moves     int
	bestScore int
}

var cellStyle = lipgloss.NewStyle().
	Width(cellWidth).
	Height(cellHeight).
	Background(lipgloss.Green).
	Align(lipgloss.Center, lipgloss.Center)

func renderCell(cell *Cell) string {
	num := fmt.Sprintf("%d", cell.number)
	if cell.isEmpty() {
		return cellStyle.Background(
			lipgloss.Color(emptyCellBg),
		).Render("")

	} else {
		color, ok := cellColors[cell.number]
		if !ok {
			color = cellColors[2]
		}

		return cellStyle.Background(lipgloss.Color(color.bg)).
			Foreground(
				lipgloss.Color(color.fg)).Render(num)

	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) mergeCells(cell1, cell2 *Cell) model {

	cell1.number += cell2.number
	cell2.reset()
	m.score += cell1.number

	if m.score > m.bestScore {
		m.bestScore = m.score
		go func(m model) {
			if m.scoreFile != nil {
				// Wipe the file and reset the cursor to the beginning
				if err := m.scoreFile.Truncate(0); err != nil {
					return
				}
				if _, err := m.scoreFile.Seek(0, 0); err != nil {
					return
				}
				if err := json.NewEncoder(m.scoreFile).Encode(m.score); err != nil {
					log.Println("error saving  best score..", err)
				}
				// if _, err := fmt.Fprintf(m.scoreFile, "%d", m.score); err != nil {
				// }
			}

		}(m)
	}

	return m

}

func (m model) handleDirectionKeyPress(direction string) model {
	generateNewCell := false
	switch direction {
	case "right":

		for _, row := range m.board.grid {
			for i := range slices.Backward(row) {
				for j := i - 1; j >= 0; j-- {

					cell1 := row[i]
					cell2 := row[j]

					if cell2.isEmpty() {
						continue
					}

					if cell1.isEmpty() && !cell2.isEmpty() {
						cell1.number = cell2.number
						cell2.reset()
						generateNewCell = true
					} else if cell2.number == cell1.number {
						m = m.mergeCells(cell1, cell2)
						// m.mergeCells(n, m)
						// n.number *= 2
						// m.reset()
						generateNewCell = true
						break
					} else {
						break
					}

				}

			}
		}

	case "left":
		for _, row := range m.board.grid {
			for i := 0; i < len(row)-1; i++ { //skip last cell
				for j := i + 1; j < len(row); j++ {

					cell1 := row[i]
					cell2 := row[j]

					if cell2.isEmpty() {
						continue
					}

					if cell1.isEmpty() && !cell2.isEmpty() {
						cell1.number = cell2.number
						cell2.reset()
						generateNewCell = true
					} else if cell2.number == cell1.number {
						m = m.mergeCells(cell1, cell2)
						// n.number *= 2
						// m.reset()
						generateNewCell = true
						break
					} else {
						break
					}

				}

			}
		}

	case "up":

		for colIdx := range m.board.cols {
			for i := 0; i < m.board.rows; i++ { //skip last cell

				for j := i + 1; j < m.board.rows; j++ {

					cell1 := m.board.grid[i][colIdx]
					cell2 := m.board.grid[j][colIdx]

					if cell2.isEmpty() {
						continue
					}

					if cell1.isEmpty() && !cell2.isEmpty() {
						cell1.number = cell2.number
						cell2.reset()
						generateNewCell = true
					} else if cell2.number == cell1.number {
						m = m.mergeCells(cell1, cell2)
						// n.number *= 2
						// m.reset()
						generateNewCell = true
						break
					} else {
						break
					}

				}
			}
		}

	case "down":
		for colIdx := range m.board.cols {
			for i := m.board.rows - 1; i > 0; i-- { //skip last cell

				for j := i - 1; j >= 0; j-- {

					cell1 := m.board.grid[i][colIdx]
					cell2 := m.board.grid[j][colIdx]

					if cell2.isEmpty() {
						continue
					}

					if cell1.isEmpty() && !cell2.isEmpty() {
						cell1.number = cell2.number
						cell2.number = 0
						generateNewCell = true
					} else if cell2.number == cell1.number {
						m = m.mergeCells(cell1, cell2)
						// n.number *= 2
						// m.reset()
						generateNewCell = true
						break
					} else {
						break
					}

				}

			}
		}
	}

	if generateNewCell {
		m.board.addToRandomCell()
	}
	return m
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		key := msg.String()

		switch key {
		case "q":
			return m, tea.Quit
		case "right", "left", "up", "down":
			m = m.handleDirectionKeyPress(key)
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
	}

	return m, nil
}

func (m model) View() tea.View {

	boardRows := []string{}

	cellGap := lipgloss.NewStyle().Width(gridGapX).Height(cellHeight).Background(
		lipgloss.Color(boardBg),
	).Render("")

	rowGap := lipgloss.NewStyle().Width(gridGapY).Background(
		lipgloss.Color(boardBg),
	).
		Width(gridWidth).Render("")

	for rowIndex := range m.board.rows {

		rowChildren := []string{}

		for itemIndex := range m.board.cols {
			cell := m.board.grid[rowIndex][itemIndex]
			if itemIndex == 0 {
				rowChildren = append(rowChildren, cellGap)
			}
			rowChildren = append(rowChildren, renderCell(cell))
			rowChildren = append(rowChildren, cellGap)
		}

		row := lipgloss.NewStyle().Background(lipgloss.Color(boardBg)).Render(
			lipgloss.JoinHorizontal(lipgloss.Left, rowChildren...),
		)

		if rowIndex == 0 {
			boardRows = append(boardRows, rowGap)
		}
		boardRows = append(boardRows, row)
		boardRows = append(boardRows, rowGap)

	}

	// gridStyle := lipgloss.NewStyle().
	// 	Background(lipgloss.Color(boardBg))
	// .Width(gridWidth).Height(gridHeight)

	grid := lipgloss.JoinVertical(lipgloss.Top, boardRows...)
	scoreLine := lipgloss.NewStyle().MarginBottom(1).Render(fmt.Sprintf("score %d", m.score))
	bestScoreLine := lipgloss.NewStyle().
		MarginBottom(1).
		Render(fmt.Sprintf("best score %d", m.bestScore))

	return tea.NewView(lipgloss.JoinVertical(lipgloss.Top, scoreLine, bestScoreLine, grid))
}

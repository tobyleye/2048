package main

import (
	"fmt"
	"slices"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	posX  int
	posY  int
	width int
	board Board
	score int
	moves int
}

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

func (m model) mergeCell(cell1, cell2 *Cell) int {
	var score int = 0
	cell1.number += cell2.number
	score = cell1.number

	return score
}

func (m model) handleDirectionKeyPress(direction string) model {
	generateNewCell := false
	switch direction {
	case "right":

		for _, row := range m.board.grid {
			for i := range slices.Backward(row) {
				for j := i - 1; j >= 0; j-- {

					n := row[i]
					m := row[j]

					if m.isEmpty() {
						continue
					}

					if n.isEmpty() && !m.isEmpty() {
						n.number = m.number
						m.reset()
						generateNewCell = true
					} else if m.number == n.number {
						n.number *= 2
						m.reset()
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

					n := row[i]
					m := row[j]

					if m.isEmpty() {
						continue
					}

					if n.isEmpty() && !m.isEmpty() {
						n.number = m.number
						m.reset()
						generateNewCell = true
					} else if m.number == n.number {
						n.number *= 2
						m.reset()
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

					n := m.board.grid[i][colIdx]
					m := m.board.grid[j][colIdx]

					if m.isEmpty() {
						continue
					}

					if n.isEmpty() && !m.isEmpty() {
						n.number = m.number
						m.reset()
						generateNewCell = true
					} else if m.number == n.number {
						n.number *= 2
						m.reset()
						generateNewCell = true
						break
					} else {
						break
					}

				}

			}
		}

	case "down":
		rows := 4
		cols := 4

		for colIdx := range cols {
			for i := rows - 1; i > 0; i-- { //skip last cell

				for j := i - 1; j >= 0; j-- {

					n := m.board.grid[i][colIdx]
					m := m.board.grid[j][colIdx]

					if m.isEmpty() {
						continue
					}

					if n.isEmpty() && !m.isEmpty() {
						n.number = m.number
						m.number = 0
						generateNewCell = true
					} else if m.number == n.number {
						n.number *= 2
						m.reset()
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

	gridStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(boardBg))
		// .Width(gridWidth).Height(gridHeight)

	grid := lipgloss.JoinVertical(lipgloss.Top, boardRows...)

	return tea.NewView(gridStyle.Render(grid))
}

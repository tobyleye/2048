package main

import (
	"fmt"
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
	gameOver  bool
}

var cellStyle = lipgloss.
	NewStyle().
	Width(cellWidth).
	Height(cellHeight).
	Background(lipgloss.Green).
	Align(lipgloss.Center, lipgloss.Center)

func renderCell(cell *Cell, dimmed bool) string {

	if cell.isEmpty() {
		var bg = emptyCellBg
		if dimmed {
			bg = emptyCellBgDimmed
		}
		return cellStyle.
			Background(
				lipgloss.Color(bg),
			).
			Render("")

	} else {
		bg, fg := getTileColor(cell.tile, dimmed)
		return cellStyle.
			Background(lipgloss.Color(bg)).
			Foreground(
				lipgloss.Color(fg),
			).
			Render(formatTile(cell.tile))

	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) mergeTiles(cell1, cell2 *Cell) model {

	cell1.tile += cell2.tile
	cell2.reset()
	m.score += cell1.tile
	if m.score > m.bestScore && m.scoreFile != nil {
		m.bestScore = m.score
		go saveBestScore(m.scoreFile, m.bestScore)
	}

	return m
}

func (m model) handleDirectionKeyPress(direction string) model {
	spawnNewTile := false
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
						cell1.tile = cell2.tile
						cell2.reset()
						spawnNewTile = true
					} else if cell2.tile == cell1.tile {
						m = m.mergeTiles(cell1, cell2)

						spawnNewTile = true
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
						cell1.tile = cell2.tile
						cell2.reset()
						spawnNewTile = true
					} else if cell2.tile == cell1.tile {
						m = m.mergeTiles(cell1, cell2)

						spawnNewTile = true
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
						cell1.tile = cell2.tile
						cell2.reset()
						spawnNewTile = true
					} else if cell2.tile == cell1.tile {
						m = m.mergeTiles(cell1, cell2)

						spawnNewTile = true
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
						cell1.tile = cell2.tile
						cell2.tile = 0
						spawnNewTile = true
					} else if cell2.tile == cell1.tile {
						m = m.mergeTiles(cell1, cell2)
						// n.tile *= 2
						// m.reset()
						spawnNewTile = true
						break
					} else {
						break
					}

				}

			}
		}
	}

	if spawnNewTile {
		m.board.spawnRandomTile(1)
	}
	return m
}

func (m model) checkGameOver() bool {

	for _, cell := range m.board.flattenGrid() {
		if cell.isEmpty() {
			return false
		}
	}

	for i := 0; i < m.board.rows-1; i += 1 {
		row := m.board.grid[i]
		nextRow := m.board.grid[i+1]

		for j := 0; j < m.board.cols-1; j++ {
			if row[j].tile == nextRow[j].tile || row[j].tile == row[j+1].tile {
				return false
			}

		}
	}

	return true
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		key := msg.String()

		switch key {
		case "q":
			return m, tea.Quit
		case "right", "left", "up", "down":
			if !m.gameOver {
				m = m.handleDirectionKeyPress(key)
				m.gameOver = m.checkGameOver()
			}
		case "r":
			if m.gameOver {
				// reset game state
				m.board.reset()
				m.board.spawnRandomTile(2)
				m.gameOver = false
				m.score = 0
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
	}

	return m, nil
}

func (m model) renderBoard(dimmed bool) string {
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
			rowChildren = append(rowChildren, renderCell(cell, dimmed))
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

	board := lipgloss.JoinVertical(lipgloss.Top, boardRows...)

	return board
}

func (m model) View() tea.View {

	board := m.renderBoard(m.gameOver)

	scoreLine :=
		lipgloss.NewStyle().
			// Background(lipgloss.Blue).
			Width(10).Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				lipgloss.NewStyle().Foreground(
					lipgloss.Color("#6C7086"),
				).Render("Score "),

				lipgloss.NewStyle().
					Bold(true).
					Foreground(lipgloss.Color("#F9E2AF")).
					Render(fmt.Sprintf("%d", m.score)),
			),
		)

	bestScoreLine := lipgloss.NewStyle().
		// Background(lipgloss.Yellow).
		Width(10).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				lipgloss.NewStyle().Foreground(
					lipgloss.Color("#6C7086"),
				).Render("Best "),
				lipgloss.NewStyle().Bold(true).Foreground(
					lipgloss.Color("#CDD6F4"),
				).Render(fmt.Sprintf("%d", m.bestScore)),
			))

	middleGap := gridWidth - lipgloss.Width(scoreLine) - lipgloss.Width(bestScoreLine)

	scoreView := lipgloss.NewStyle().MarginBottom(1).Render(
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			scoreLine,
			lipgloss.NewStyle().Width(middleGap).String(),
			bestScoreLine,
		),
	)

	if m.gameOver {
		cardBg := lipgloss.Color(boardBg)
		keyHintStyle := lipgloss.NewStyle().Background(
			lipgloss.Color("#45475A"),
		).Foreground(
			lipgloss.Color("#F9E2AF")).Padding(0, 1)

		gameover := lipgloss.NewStyle().
			Background(cardBg).
			Padding(1, 0).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Bottom,
					lipgloss.NewStyle().MarginBottom(1).Render(
						lipgloss.PlaceHorizontal(
							gridWidth,
							lipgloss.Center,
							lipgloss.NewStyle().
								Foreground(lipgloss.Color("#F38BA8")).
								Background(cardBg).
								Render("game over"),
							lipgloss.WithWhitespaceStyle(
								lipgloss.NewStyle().
									Background(cardBg),
							),
						),
					),
					// lipgloss.PlaceHorizontal(
					// 	gridWidth,
					// 	lipgloss.Center,
					// 	lipgloss.NewStyle().
					// 		Foreground(lipgloss.Color("#A6E3A1")).
					// 		Background(cardBg).
					// 		Render("no moves left"),
					// 	lipgloss.WithWhitespaceStyle(
					// 		lipgloss.NewStyle().
					// 			Background(cardBg),
					// 	),
					// ),

					lipgloss.PlaceHorizontal(
						gridWidth,
						lipgloss.Center,
						lipgloss.JoinHorizontal(
							lipgloss.Left,
							keyHintStyle.
								Render("r"),
							lipgloss.NewStyle().
								Background(cardBg).
								Render(" new game "),
							keyHintStyle.Render("q"),
							lipgloss.NewStyle().Background(cardBg).Render(" quit"),
						),
						lipgloss.WithWhitespaceStyle(
							lipgloss.NewStyle().
								Background(cardBg),
						),
					),
				),
			)

		// game over box positining
		y := (gridHeight - lipgloss.Height(gameover)) / 2
		body := lipgloss.NewCompositor(
			lipgloss.NewLayer(board),
			lipgloss.NewLayer(gameover).X(0).Y(y).Z(1),
		).Render()

		return tea.NewView(lipgloss.JoinVertical(lipgloss.Top, scoreView, body))
	} else {
		return tea.NewView(lipgloss.JoinVertical(lipgloss.Top, scoreView, board))
	}

}

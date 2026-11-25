package main

import "math/rand"

type Board struct {
	rows int
	cols int
	grid [][]*Cell
}

type Cell struct {
	row  int
	col  int
	tile int // 0 when cell is empty
}

type position [2]int

// choose a random number between 2 & 4, with 2 having a higher probability
func randomTile() int {
	randomNo := rand.Float64()

	if randomNo > 0.6 {
		return 2
	} else {
		return 4
	}
}

func (c *Cell) isEmpty() bool {
	return c.tile == 0
}
func (c *Cell) reset() {
	c.tile = 0
}

func (c *Cell) double() {
	if c.tile > 0 {
		c.tile *= 2
	}
}

func (b Board) getEmptyCells() []position {

	positions := []position{}

	for _, row := range b.grid {
		for _, cell := range row {
			if cell.tile == 0 {
				positions = append(
					positions,
					position{cell.row, cell.col},
				)
			}
		}
	}

	return positions
}

func (b Board) getRandomEmptyCell() position {
	emptyPositions := b.getEmptyCells()
	idx := rand.Intn(len(emptyPositions))
	return emptyPositions[idx]
}

func (b Board) addTile(tile int, row, col int) {
	if row >= b.rows {
		row = b.rows - 1
	}

	if col >= b.cols {
		col = b.cols - 1
	}

	if b.grid[row][col].isEmpty() {
		b.grid[row][col].tile = tile
	}
}

func (b Board) getCell(row, col int) *Cell {
	return b.grid[row][col]
}

func (b Board) reset() {
	for _, cell := range b.flattenGrid() {
		cell.reset()
	}
}

func (b Board) flattenGrid() []*Cell {
	var flattedGrid = []*Cell{}
	for _, row := range b.grid {
		flattedGrid = append(flattedGrid, row...)
	}
	return flattedGrid
}

func (b Board) spawnRandomTile(n int) {
	for _ = range n {
		randomCell := b.getRandomEmptyCell()
		b.addTile(
			randomTile(),
			randomCell[0],
			randomCell[1],
		)
	}
}

func createBoard(rows, cols int) Board {

	var grid [][]*Cell

	for r := range gridRows {
		row := []*Cell{}

		for j := range gridCols {
			row = append(row, &Cell{
				row: r,
				col: j,
			})
		}

		grid = append(grid, row)
	}

	board := Board{
		rows: rows,
		cols: cols,
		grid: grid,
	}
	return board

}

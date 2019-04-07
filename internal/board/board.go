package board

import "fmt"

// Board holds a game board
type Board struct {
	Size    int
	Values  [][]int  // These two variables are organized as a coluimn of rows
	Flipped [][]bool // so y then x coordinate
}

// Generate created a board corresponding to the given level
func Generate(level int) *Board {
	dim := (2*(level/8) + 5)
	b := &Board{
		Size: dim * dim,
	}

	scoreTiles := genScoreTiles(level)
	voltTiles := genVoltorbs(level, b.Size, scoreTiles)

	for i := 0; i < dim; i++ {
		for j:= 0, j < dim; i++ {
			2

			
		}
	}
	return b
}

// Print displays the board, formatted into a grid
func (b *Board) Print() {
	for _, row := range b.Values {
		// print some formatting boundaries
		for _, elem := range row {
			fmt.Printf(" %d |", elem)
		}
		fmt.Print("\n")
	}
}

// Flip markes a given space of the board as flipped. It returns the value stored
// at that space
func (b *Board) Flip(y, x int) int {
	b.Flipped[y][x] = true
	return b.Values[y][x]
}

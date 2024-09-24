package functions

import (
	"fmt"
)

var Count int

// Tetromino represents a single tetromino piece
type Tetromino [][]string

// Point represents a position on the grid
type Point struct {
	X, Y int
}

// ? FitTetrominos tries to fit all tetrominos into the smallest possible square
// ? Starting from the given minimum size
func FitTetrominos(tetrominos []Tetromino, minSize int) [][]string {
	Count = 0
	maxAttempts := 3 // Limit the number of attempts to avoid infinite loop
	for size := minSize; size < minSize+maxAttempts; size++ {
		// fmt.Printf("Trying with square size: %d\n", size)
		grid := CreateSquare(size)
		//If the backtrack function returns true, that means all tetrominos got placed successfully
		if backtrack(grid, tetrominos, 0) {
			return grid
		}
	}
	fmt.Println("Unable to find a solution within reasonable attempts")
	return nil
}

// ? Backtrack tries to place all tetrominos recursively
func backtrack(grid [][]string, tetrominos []Tetromino, index int) bool {
	if index == len(tetrominos) {
		return true // All tetrominos placed successfully
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			if canPlace(grid, tetrominos[index], Point{X: x, Y: y}) {
				place(grid, tetrominos[index], Point{X: x, Y: y})
				if backtrack(grid, tetrominos, index+1) {
					return true
				}
				remove(grid, tetrominos[index], Point{X: x, Y: y})
			}
		}
	}
	return false // Couldn't place this tetromino
}

// ? Function that check if it's possible to place the tetromino
func canPlace(grid [][]string, tetromino Tetromino, pos Point) bool {
	//Ranging over the tetrominos
	for y := 0; y < len(tetromino); y++ {
		for x := 0; x < len(tetromino[y]); x++ {
			//Checking if the tetrominos position isn't an empty space
			if tetromino[y][x] != "  " && tetromino[y][x] != "." {
				gridY, gridX := pos.Y+y, pos.X+x
				//Checking if the x and y position are out of grid range
				if gridY >= len(grid) || gridX >= len(grid) || (grid[gridY][gridX] != "." && grid[gridY][gridX] != "  ") {
					return false
				}
			}
		}
	}
	return true
}

// ? Function that will replace empty blocks on the grid by the blocks of the given tetromino
func place(grid [][]string, tetromino Tetromino, pos Point) {
	//Ranging over our current tetromino
	for y := 0; y < len(tetromino); y++ {
		for x := 0; x < len(tetromino[y]); x++ {
			//If the current value of the tetromino is block, then we replace the grid value
			if tetromino[y][x] != "  " && tetromino[y][x] != "." {
				grid[pos.Y+y][pos.X+x] = tetromino[y][x]
			}
		}
	}
}

// ? Function that will remove a tetromino just like the place function
func remove(grid [][]string, tetromino Tetromino, pos Point) {
	Count++
	//We range over that tetromino to remove all his blocks
	for y := 0; y < len(tetromino); y++ {
		for x := 0; x < len(tetromino[y]); x++ {
			//If the current position isn't an empty space, we make it one
			if tetromino[y][x] != "  " && tetromino[y][x] != "." {
				grid[pos.Y+y][pos.X+x] = "  "
			}
		}
	}
}

// ? Function that will create a square of the size given
func CreateSquare(size int) [][]string {
	grid := make([][]string, size)
	//Range over all the values of that created square
	for i := range grid {
		grid[i] = make([]string, size)
		for j := range grid[i] {
			//Giving them empty spaces as values
			grid[i][j] = "  "
		}
	}
	//Return a two dimensionnal array of the given size
	return grid
}

// ? Function that will replace "." with "  " for a better output result
func FormatResult(grid [][]string) [][]string {
	//Ranging over the grid
	for i := range grid {
		for j := range grid[i] {
			//Checking for empty spots (".")
			if grid[i][j] == "." {
				//Replace with "  "
				grid[i][j] = "  "
			}
		}
	}
	//Return the modified grid
	return grid
}

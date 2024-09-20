package functions

import (
	"fmt"
)

// Tetromino represents a single tetromino piece
type Tetromino [][]string

// Point represents a position on the grid
type Point struct {
	X, Y int
}

// FitTetrominos tries to fit all tetrominos into the smallest possible square
// starting from the given minimum size
func FitTetrominos(tetrominos []Tetromino, minSize int) [][]string {
	maxAttempts := 10 // Limit the number of attempts to avoid infinite loop
	for size := minSize; size < minSize+maxAttempts; size++ {
		// fmt.Printf("Trying with square size: %d\n", size)
		grid := CreateSquare(size)
		if backtrack(grid, tetrominos, 0) {
			return grid
		}
	}
	fmt.Println("Unable to find a solution within reasonable attempts")
	return nil
}

// backtrack tries to place all tetrominos recursively
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

// canPlace checks if a tetromino can be placed at the given position
func canPlace(grid [][]string, tetromino Tetromino, pos Point) bool {
	for y := 0; y < len(tetromino); y++ {
		for x := 0; x < len(tetromino[y]); x++ {
			if tetromino[y][x] != "  " && tetromino[y][x] != "." {
				gridY, gridX := pos.Y+y, pos.X+x
				if gridY >= len(grid) || gridX >= len(grid) || (grid[gridY][gridX] != "." && grid[gridY][gridX] != "  ") {
					return false
				}
			}
		}
	}
	return true
}

// place puts a tetromino on the grid at the given position
func place(grid [][]string, tetromino Tetromino, pos Point) {
	for y := 0; y < len(tetromino); y++ {
		for x := 0; x < len(tetromino[y]); x++ {
			if tetromino[y][x] != "  " && tetromino[y][x] != "." {
				grid[pos.Y+y][pos.X+x] = tetromino[y][x]
			}
		}
	}
}

// remove takes a tetromino off the grid at the given position
func remove(grid [][]string, tetromino Tetromino, pos Point) {
	for y := 0; y < len(tetromino); y++ {
		for x := 0; x < len(tetromino[y]); x++ {
			if tetromino[y][x] != "  " && tetromino[y][x] != "." {
				grid[pos.Y+y][pos.X+x] = "  "
			}
		}
	}
}

// CreateSquare creates a square grid of the given size
func CreateSquare(size int) [][]string {
	grid := make([][]string, size)
	for i := range grid {
		grid[i] = make([]string, size)
		for j := range grid[i] {
			grid[i][j] = "  "
		}
	}
	return grid
}

// FormatResult replaces "." with "  " in the result grid
func FormatResult(grid [][]string) [][]string {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "." {
				grid[i][j] = "  "
			}
		}
	}
	return grid
}

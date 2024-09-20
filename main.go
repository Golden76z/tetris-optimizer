package main

import (
	"fmt"
	"os"
	"tetris-optimizer/functions"
)

func main() {
	t := functions.NewTimer()
	//Start the timer
	t.Start()

	//Checking arguments length
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	//Storing our filename in a variable
	filename := os.Args[1]

	//Variable that will store all our tetrominos
	var multitetroarray [][][]string

	//Checking if the tetrominos are valid or not
	tetronumber, tetroarray, errorcheck := functions.Check(filename)
	if tetronumber < 1 {
		return
	} else if errorcheck != 0 {
		return
	}
	fmt.Printf("          Number of valid tetrominos: %d\n", tetronumber)
	multitetroarray = functions.Chunk(tetroarray)
	multitetroarray = functions.CheckLines(multitetroarray)

	//Getting rid of the empty colums to the left and right
	for i := 0; i < 4; i++ {
		multitetroarray = functions.CheckColumnsLeft(multitetroarray)
		multitetroarray = functions.CheckColumnsRight(multitetroarray)
	}

	//Replacing all the # by colors, each color is unique
	multitetroarray = functions.ReplaceWithColor(multitetroarray)

	// Printing all my tetrominos
	println("<--------------------------------------------------->")
	for i := 0; i < len(multitetroarray); i++ {
		for j := 0; j < len(multitetroarray[i]); j++ {
			for l := 0; l < len(multitetroarray[i][j]); l++ {
				if multitetroarray[i][j][l] != "⬛" {
					fmt.Print(multitetroarray[i][j][l])
				} else {
					fmt.Print("  ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
	println("<--------------------------------------------------->")

	//Get the minimum square size
	minSquareSize := functions.SquareSize(tetronumber)
	fmt.Printf("              Minimum square size: %d\n", minSquareSize)
	println("<--------------------------------------------------->")

	//Convert multitetroarray to []functions.Tetromino
	var tetrominos []functions.Tetromino
	for _, tetro := range multitetroarray {
		tetrominos = append(tetrominos, functions.Tetromino(tetro))
	}

	//Fit the tetrominos into the smallest possible square, starting from the minimum size
	result := functions.FitTetrominos(tetrominos, minSquareSize)

	//Print the result
	if result != nil {
		formattedResult := functions.FormatResult(result)
		for i := 0; i < len(formattedResult); i++ {
			for j := 0; j < len(formattedResult[i]); j++ {
				fmt.Print(formattedResult[i][j])
			}
			fmt.Println()
		}
	} else {
		fmt.Println("No solution found")
	}

	println("<--------------------------------------------------->")
	//Stop the timer and get the elapsed time
	elapsed := t.ElapsedSeconds()
	fmt.Printf("          Program took %.2f seconds to finish\n", elapsed)
	fmt.Println()
}

package functions

import (
	"bufio"
	"fmt"
	"os"
)

func Check(s string) (int, []string, int) {
	//Variable that will check for errors
	errorcheck := 0

	//Opening the file
	file, err := os.Open("./test/" + s + ".txt")
	if err != nil {
		return 0, []string{}, errorcheck
	}
	//Array that will contain all the file content
	tetrominos := []string{}

	count := 0
	//Storing all the file content into the tetrominos array
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		count++
		if count == 5 && len(fileScanner.Text()) != 0 {
			println("\033[31m" + "<--------------------------------------------------->")
			println("               File has a bad format")
			println("<--------------------------------------------------->" + "\033[0m")
			return 0, tetrominos, errorcheck
		}
		if len(fileScanner.Text()) != 0 {
			tetrominos = append(tetrominos, fileScanner.Text())
		}
	}

	//Numbers of tetrominos on the file
	tetronumber := 1
	//Variable for the numbers of links
	count = 0
	//Number of lines on the file
	loopnumber := 0
	//Numbers of blocs contained on 4 lines
	blocnumber := 0

	//Checking for the variable values
	for i := 0; i < len(tetrominos); i++ {
		loopnumber++
		for j := 0; j < len(tetrominos[i]); j++ {
			//Checking for tetrominos links
			if tetrominos[i][j] == '#' {
				blocnumber++
				if i != 0 && tetrominos[i-1][j] == '#' {
					count++
				}
				if i != len(tetrominos)-1 && tetrominos[i+1][j] == '#' {
					count++
				}
				if j != 0 && tetrominos[i][j-1] == '#' {
					count++
				}
				if j != len(tetrominos[i])-1 && tetrominos[i][j+1] == '#' {
					count++
				}
			}
		}
		if loopnumber == 4 {
			//Error checking for links inferior to 6
			if count < 6 {
				fmt.Println(count)
				println("\033[31m" + "<--------------------------------------------------->")
				println("            Tetrominos number", tetronumber, "is invalid")
				println("             ", count, "Link instead of 6 minimum ")
				println("<--------------------------------------------------->" + "\033[0m")
				errorcheck++
				return tetronumber - 1, tetrominos, errorcheck
				//Error checking for numbers of block (aka different from 4)
			} else if blocnumber != 4 {
				println("\033[31m" + "<--------------------------------------------------->")
				println("            Too many block on tetrominos:", tetronumber)
				println("             ", blocnumber, "Block instead of 4.")
				println("<--------------------------------------------------->" + "\033[0m")
				errorcheck++
				return tetronumber - 1, tetrominos, errorcheck
			} else {
				//If no errors, reset values and increment the tetronumber
				loopnumber = 0
				count = 0
				blocnumber = 0
				tetronumber++
			}
		}
	}
	//If no error on the full array, all tetrominos are valid
	println("\033[32m" + "<--------------------------------------------------->")
	println("            All Tetrominos are valid! ")
	println("<--------------------------------------------------->" + "\033[0m")
	return tetronumber - 1, tetrominos, errorcheck
}

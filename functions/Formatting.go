package functions

// ? Function that will store every tetrominos into a 2 dimensional array
func Chunk(tab []string) [][][]string {
	//Array containing all the tetrominos
	finaltab := [][][]string{}
	//Array containing a single tetrominos
	tetrotab := [][]string{}
	//Array containing a single line of a tetrominos
	line := []string{}
	count := 0
	for i := 0; i < len(tab); i++ {
		count++
		//Append every value to the array
		for j := 0; j < len(tab[i]); j++ {
			line = append(line, string(tab[i][j]))
		}
		//We then append that array to the two dimensionnal array
		tetrotab = append(tetrotab, line)
		line = []string{}
		//When the count is at 4, we have our complete tetrominos
		if count == 4 {
			//We then append it to the 3 dimensionnal array and reset the tetrominos array
			finaltab = append(finaltab, tetrotab)
			tetrotab = [][]string{}
			//We reset the count aswell
			count = 0
		}
	}
	return finaltab
}

// ? Function that will erase the useless lines
func CheckLines(tab [][][]string) [][][]string {
	//Array containing all the tetrominos
	finaltab := [][][]string{}
	//Array containing a single tetrominos
	temptab := [][]string{}
	//Array containing a single line of a tetrominos
	line := []string{}
	blocknumber := 0
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			//We iterate over a line
			for k := 0; k < len(tab[i][j]); k++ {
				//If we encounter a block, we increment the blocknumber variable
				if tab[i][j][k] != "." {
					blocknumber++
				}
				line = append(line, string(tab[i][j][k]))
			}
			//If the blocknumber variable isn't 0, then we keep that line
			if blocknumber != 0 {
				temptab = append(temptab, line)
				blocknumber = 0
			}
			//We reset the line array before iterating over a new one
			line = []string{}
		}
		//We then append the modified version of the tetrominos to the final array
		finaltab = append(finaltab, temptab)
		temptab = [][]string{}
	}
	return finaltab
}

// ? Function that will get rid of the empty column to the left
func CheckColumnsLeft(tab [][][]string) [][][]string {
	finaltab := [][][]string{}
	temptab := [][]string{}
	line := []string{}
	blocknumber := 0
	i := 0
	j := 0
	for i < len(tab) {
		//If the current value isn't a block, we iterate over the next lines
		if tab[i][j][0] == "." {
			//We then iterate over the 4 lines of the tetrominos
			for l := 0; l < len(tab[i]); l++ {
				//Increment the blocknumber variable
				if tab[i][j+l][0] != "." {
					blocknumber++
				}
			}
			//If that variable is 0, then the column is empty
			if blocknumber == 0 {
				//We then proceed to remove that column
				for l := 0; l < len(tab[i]); l++ {
					for k := 0; k < len(tab[i][l]); k++ {
						line = append(line, tab[i][l][k])
					}
					temptab = append(temptab, line[1:])
					line = []string{}
				}
			} else {
				//Otherwise we keep that column
				for l := 0; l < len(tab[i]); l++ {
					for k := 0; k < len(tab[i][l]); k++ {
						line = append(line, tab[i][l][k])
					}
					temptab = append(temptab, line)
					line = []string{}
				}
			}
		} else {
			for l := 0; l < len(tab[i]); l++ {
				for k := 0; k < len(tab[i][l]); k++ {
					line = append(line, tab[i][l][k])
				}
				temptab = append(temptab, line)
				line = []string{}
			}
		}
		finaltab = append(finaltab, temptab)
		temptab = [][]string{}
		blocknumber = 0
		if i < len(tab) {
			i++
		} else {
			return finaltab
		}
	}
	return finaltab
}

// ? Function that will get rid of the empty column to the right
func CheckColumnsRight(tab [][][]string) [][][]string {
	finaltab := [][][]string{}
	temptab := [][]string{}
	line := []string{}
	blocknumber := 0
	i := 0
	j := 0
	for i < len(tab) {
		if tab[i][j][len(tab[i][j])-1] == "." {
			for l := 0; l < len(tab[i]); l++ {
				if tab[i][j+l][len(tab[i][j])-1] != "." {
					blocknumber++
				}
			}
			if blocknumber == 0 {
				for l := 0; l < len(tab[i]); l++ {
					for k := 0; k < len(tab[i][l]); k++ {
						line = append(line, tab[i][l][k])
					}
					temptab = append(temptab, line[:len(tab[i][j])-1])
					line = []string{}
				}
			} else {
				for l := 0; l < len(tab[i]); l++ {
					for k := 0; k < len(tab[i][l]); k++ {
						line = append(line, tab[i][l][k])
					}
					temptab = append(temptab, line)
					line = []string{}
				}
			}
		} else {
			for l := 0; l < len(tab[i]); l++ {
				for k := 0; k < len(tab[i][l]); k++ {
					line = append(line, tab[i][l][k])
				}
				temptab = append(temptab, line)
				line = []string{}
			}
		}
		finaltab = append(finaltab, temptab)
		temptab = [][]string{}
		blocknumber = 0
		if i < len(tab) {
			i++
		} else {
			return finaltab
		}
	}
	return finaltab
}

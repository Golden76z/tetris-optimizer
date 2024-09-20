package functions

// ? Function that replace all the # with colors
func ReplaceWithColor(tab [][][]string) [][][]string {
	//Array that will store all our color values
	colors := []string{
		"🟨", // Yellow
		"🟧", // Orange
		"🟥", // Red
		"🟪", // Purple
		"🟦", // Blue
		"🟩", // Green
		"🟫", // Brown
		"🈹", // Smiley1
		"🈴", // Smiley2
		"🈯", // Smiley3
		"🈳", // Smiley4
		"🆗", // White
		// "⬛",
	}
	finaltab := [][][]string{}
	temptab := [][]string{}
	line := []string{}
	index := 0
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			for k := 0; k < len(tab[i][j]); k++ {
				//If the value is a block, we give it a color
				if tab[i][j][k] == "#" {
					line = append(line, colors[index])
				} else {
					//Otherwise we give it a black square
					line = append(line, "  ")
				}
			}
			temptab = append(temptab, line)
			line = []string{}
		}
		//Everytime we iterate over a tetrominos, we increment the index to change the color
		finaltab = append(finaltab, temptab)
		temptab = [][]string{}
		index++
	}
	return finaltab
}

// 🆚
// 🆙
// 🆘
// 🆗

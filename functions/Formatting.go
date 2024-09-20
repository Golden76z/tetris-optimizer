package functions

func Chunk(tab []string) [][][]string {
	finaltab := [][][]string{}
	tetrotab := [][]string{}
	line := []string{}
	count := 0
	for i := 0; i < len(tab); i++ {
		count++
		for j := 0; j < len(tab[i]); j++ {
			line = append(line, string(tab[i][j]))
		}
		tetrotab = append(tetrotab, line)
		line = []string{}
		if count == 4 {
			finaltab = append(finaltab, tetrotab)
			tetrotab = [][]string{}
			count = 0
		}
	}
	return finaltab
}

func CheckLines(tab [][][]string) [][][]string {
	finaltab := [][][]string{}
	temptab := [][]string{}
	line := []string{}
	blocknumber := 0
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			for k := 0; k < len(tab[i][j]); k++ {
				if tab[i][j][k] != "." {
					blocknumber++
				}
				line = append(line, string(tab[i][j][k]))
			}
			if blocknumber != 0 {
				temptab = append(temptab, line)
				blocknumber = 0
			}
			line = []string{}
		}
		finaltab = append(finaltab, temptab)
		temptab = [][]string{}
	}
	return finaltab
}

func CheckColumnsLeft(tab [][][]string) [][][]string {
	finaltab := [][][]string{}
	temptab := [][]string{}
	line := []string{}
	blocknumber := 0
	i := 0
	j := 0
	for i < len(tab) {
		if tab[i][j][0] == "." {
			for l := 0; l < len(tab[i]); l++ {
				if tab[i][j+l][0] != "." {
					blocknumber++
				}
			}
			if blocknumber == 0 {
				for l := 0; l < len(tab[i]); l++ {
					for k := 0; k < len(tab[i][l]); k++ {
						line = append(line, tab[i][l][k])
					}
					temptab = append(temptab, line[1:])
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

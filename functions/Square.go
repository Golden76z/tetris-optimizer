package functions

// ? Checking the minimum square size needed to fit all our tetrominos
func SquareSize(n int) int {
	//n is the number of tetrominos
	for i := 2; i < 10; i++ {
		//n*4 because every tetrominos has exactly 4 blocs
		if n*4 <= i*i {
			return i
		}
	}
	return 0
}

// ? Function that will take the number of tetrominos and create a minimum
// ? square that can potentially fit all of them
func CreateSquare2(n int) [][]string {
	finaltab := make([][]string, n)
	for i := range finaltab {
		finaltab[i] = make([]string, n)
		for j := range finaltab[i] {
			finaltab[i][j] = "  "
		}
	}
	return finaltab
}

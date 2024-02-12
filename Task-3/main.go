package main

import "fmt"

var board [9][9]byte = [9][9]byte{

	// Valid and already solved Sudoku matrix
	// {5, 3, 4, 6, 7, 8, 9, 1, 2},
	// {6, 7, 2, 1, 9, 5, 3, 4, 8},
	// {1, 9, 8, 3, 4, 2, 5, 6, 7},
	// {8, 5, 9, 7, 6, 1, 4, 2, 3},
	// {4, 2, 6, 8, 5, 3, 7, 9, 1},
	// {7, 1, 3, 9, 2, 4, 8, 5, 6},
	// {9, 6, 1, 5, 3, 7, 2, 8, 4},
	// {2, 8, 7, 4, 1, 9, 6, 3, 5},
	// {3, 4, 5, 2, 8, 6, 1, 7, 9},

	// Valid but not solved Sudoku matrix
	{5, 3, 0, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 0, 9, 5, 3, 0, 8},
	{1, 0, 8, 3, 4, 0, 5, 6, 7},
	{8, 5, 9, 7, 0, 1, 4, 0, 3},
	{4, 2, 0, 8, 5, 3, 0, 9, 1},
	{7, 0, 3, 9, 2, 0, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 0, 4},
	{2, 8, 0, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 0, 7, 9},

	// Unsolvable Sudoku matrix
	// {5, 3, 4, 6, 7, 8, 0, 1, 2},
	// {6, 7, 2, 1, 0, 5, 9, 4, 8},
	// {1, 9, 8, 3, 4, 2, 5, 6, 7},
	// {8, 5, 9, 7, 6, 1, 4, 2, 3},
	// {4, 2, 6, 8, 5, 3, 7, 9, 1},
	// {7, 1, 3, 9, 2, 4, 8, 5, 6},
	// {9, 6, 1, 5, 3, 7, 2, 8, 4},
	// {2, 8, 7, 4, 1, 9, 6, 3, 5},
	// {3, 4, 5, 2, 8, 6, 1, 7, 9},
}

// Print the sudoku board
func printBoard() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
}

// Swap numbers used for matrix tranposation
func swap(x, y byte) (byte, byte) {
	return y, x
}

// Transpose the sudoku matrix for accessing columns
func transposeBoard() {
	for i := 0; i < 9; i++ {
		for j := i; j < 9; j++ {
			board[i][j], board[j][i] = swap(board[i][j], board[j][i])
		}
	}
}

// Checking that row or column has unique digits or not
func isRowColumnValid(row [9]byte) bool {
	// Map for checking that each value from 1-9 is coming one time only
	rowMap := make(map[byte]byte)

	for _, v := range row {

		// v != 0 means there is a digit and not 0 for user input
		if v != 0 {
			// Increase the occurrence of that value using map
			rowMap[v]++

			// If the value occurred more than one time, return false
			if rowMap[v] > 1 {
				return false
			}
		}
	}
	// If values in row or column is unique then return true
	return true
}

// This is generate 3X3 grid and
func subGrid(i, j int) {
	boxMap := make(map[byte]byte)

	for start := i; start < i+3; start++ {
		for end := j; end < j+3; end++ {
			if board[start][end] != 0 {
				boxMap[board[start][end]]++
			}
		}
	}

	for _, v := range boxMap {
		if v > 1 {
			fmt.Println("Grid has Error!")
		}
	}
}

// Checking that sudoku is valid or not
func isValidSudoku() {

	// Checking that all element is unique or not in row
	for i := 0; i < 9; i++ {
		if !isRowColumnValid(board[i]) {
			fmt.Print("Unvalid Value at [ ", i+1, ",")
		}
	}

	// Transpose the sudoku matrix for accessing column
	transposeBoard()

	// Checking that is all element is unique or not in column
	for i := 0; i < 9; i++ {
		if !isRowColumnValid(board[i]) {
			fmt.Println(i+1, "]")
		}
	}
	// Transpose the sudoku matrix for accessing column
	transposeBoard()

	for i := 0; i < 7; i += 3 {
		for j := 0; j < 7; j += 3 {
			subGrid(i, j)
		}
	}
}

// This wil check that sudoku is solvable or not
func isSolvableSudoku() bool {
	// Iterate through all Sudoku board to encounter 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			// If 0 found, then it will try to place all 1 to 9 digit and check that its fitting or not
			if board[row][col] == 0 {
				for i := 1; i <= 9; i++ {
					if isSafePlace(row, col, byte(i)) {
						board[row][col] = byte(i)
						if isSolvableSudoku() {
							return true
						} else {
							board[row][col] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

// This is checking that val is fitting at that place or not
func isSafePlace(row int, col int, val byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
		if board[i][col] == val {
			return false
		}
	}

	x := 3 * (row / 3)
	y := 3 * (col / 3)
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println("Sudoku Solver")
	fmt.Println(":: Unsolved Sudoku ::")
	printBoard()

	if isSolvableSudoku() {
		fmt.Println("\nSudoku is Solvable")
		fmt.Println(":: Solved Sudoku ::")
		printBoard()
	} else {
		fmt.Println("\nSudoku is NOT Solvable")
	}

	isValidSudoku()
}

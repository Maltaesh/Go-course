package main

import (
	utils "example/project/mypackage"
	"fmt"
	"strconv"
)

type puzzType [][]int

// Outputs the current board to screen
func displayBoard(puzz puzzType) {
	puzz_cols := len(puzz[0])

	// Cycles through rows in the puzz
	for i := 0; i < len(puzz); i++ {
		// Cycles through columns in the puzz
		for j := 0; j < puzz_cols; j++ {
			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
		}
		utils.Pl()
	}
}

// Return a valid row and column for an empty space
func getEmptySpace(puzz puzzType) (int, int) {
	puzz_cols := len(puzz[0])

	// Cycles through rows in the puzz
	for i := 0; i < len(puzz); i++ {
		// Cycles through columns in the puzz
		for j := 0; j < puzz_cols; j++ {
			if puzz[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isNumValid(puzz puzzType, guess, row, column int) bool {
	// Check rows
	for index := range puzz {
		if puzz[row][index] == guess && column != index {
			return false
		}
	}

	// Check columns
	for index := range puzz {
		if puzz[index][column] == guess && column != index {
			return false
		}
	}

	// Cheeck box 3x3
	// Row 0 & column 3
	// Row - (Row % 3) + value for cycling (0-2)
	// 0 - (0 % 3) + 0 = 0 - 0 -

	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			if puzz[row-(row%3)+k][column-(column%3)+l] == guess && (row-(row%3)+k != row || column-(column%3)+l != column) {
				return false
			}
		}
	}

	return true
}

func main() {
	puzz := puzzType{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}

	displayBoard(puzz)

	row, _ := getEmptySpace(puzz)
	if row != -1 {
		utils.Pl(getEmptySpace(puzz))
	} else {
		utils.Pl("Puzzle is solved!")
	}

	utils.Pl(isNumValid(puzz, 7, 4, 0))
}

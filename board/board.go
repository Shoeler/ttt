package board

import "fmt"

func CheckWin(board [3][3]int, numMoves int) int {
	//The idea here is to check every row for all of one player move and then
	//check the two diagonals and if any of them have that the player is a winner
	b := board
	for i := 0; i < 3; i++ {
		if b[i][0] != 0 && b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			return b[i][0]
		}
		if b[0][i] != 0 && b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return b[0][i]
		}
	}
	// Check diagonals
	if b[0][0] != 0 && b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return b[0][0]
	}
	if b[0][2] != 0 && b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		return b[0][2]
	}
	if numMoves == 9 {
		return 3
	}
	return 0 // no winner
}

func Print(board [3][3]int) {
	fmt.Println("Current Board:")
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
		for j := 0; j < 3; j++ {
			symbol := " "
			switch board[i][j] {
			case 0:
				symbol = "-"
			case 1:
				symbol = "X"
			case 2:
				symbol = "O"
			}
			fmt.Print(symbol)
			if j < 2 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
	fmt.Println()
}

package main

import (
	"fmt"
	"ttt/board"
	"ttt/player"
)

func main() {
	var myBoard [3][3]int // Initialize the board to zeros
	var winner, row, col, numMoves, playerLetter, value int
	var playerLetterArry = [3]string{"draw", "x", "o"}
	fmt.Println(" New Game ")
	board.Print(myBoard)
	playerLetter = player.GetLetter()
	for winner = 0; winner <= 0; winner = board.CheckWin(myBoard, numMoves) {
		row, col, numMoves = player.GetMove(numMoves)
		fmt.Printf("inside loop\n")
		myBoard[row][col] = playerLetter
		fmt.Printf("Values are %d %d %d \n", row, col, value)
		fmt.Printf("The value from the array is %d \n", myBoard[row][col])
		board.Print(myBoard)
	}
	if winner != 3 {
		fmt.Printf("Congratulations to player %s !\n", playerLetterArry[winner])
	} else {
		fmt.Printf("The game is a draw.\n")
	}
}

func computerMove(board [3][3]int, newMoves int) (row, col, outMoves int) {

	return row, col, outMoves
}

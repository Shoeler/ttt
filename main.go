package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"ttt/board"
)

func main() {
	var myBoard [3][3]int // Initialize the board to zeros
	var winner, row, col, numMoves, playerLetter, value int
	var playerLetterArry = [3]string{"draw", "x", "o"}
	fmt.Println(" New Game ")
	board.Print(myBoard)
	playerLetter = getLetter()
	for winner = 0; winner <= 0; winner = board.CheckWin(myBoard, numMoves) {
		row, col, numMoves = getMove(numMoves)
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

func getMove(newMoves int) (row, col, outMoves int) {
	var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your move in the format 'row,col' in the range 1-3 (e.g., 1,2):")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		log.Println("invalid input format, please use 'row,col'")
	}
	row, err = strconv.Atoi(parts[0])
	if err != nil {
		log.Printf("invalid row number: %s\n", parts[0])
	}
	col, err = strconv.Atoi(parts[1])
	if err != nil {
		log.Printf("invalid column number: %s\n", parts[1])
	}
	newMoves++
	row-- // row is 1-based on input but 0-based on storage
	col-- // col too
	return row, col, newMoves
}

func computerMove(board [3][3]int, newMoves int) (row, col, outMoves int) {

	return row, col, outMoves
}

func getLetter() int {
	var err error
	var input string
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Select either x or o:")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Println(`Error, please only enter an x or an o`)
			break
		}
		input = strings.TrimSpace(input)
		switch input {
		case "x":
			return 1
		case "o":
			return 2
		default:
			fmt.Println(`Error, please only enter an x or an o`)
		}
	}
	return 0
}

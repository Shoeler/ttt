package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var myBoard [3][3]int // Initialize the board to zeros
	var winner, row, col, value, numMoves int
	var playerLetter = [3]string{"draw", "X", "O"}
	var selectedLetter string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" New Game ")
	selectedLetter = getLetter()
	input, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Println(err)
	// }
	// input = strings.TrimSpace(input)
	// if input == "x" {
	// 	selectedLetter = "x"
	// } else if selectedLetter == "o" {
	// 	selectedLetter = "o"
	// } else {
	// 	fmt.Println("error - you must select x or o")
	// }
	for true {
		fmt.Println("Select either x or o:")
		input, err = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "x":
			value = 1
			myBoard[row][col] = 1
			// return row, col, value, numMoves
			break
		case "o":
			value = 2
			myBoard[row][col] = 2
			// return row, col, value, numMoves
			break
		default:
			fmt.Println(`Error, please only enter an x or an o`)
			// break
		}
	}
	printBoard(myBoard)
	for winner = 0; winner <= 0; winner = checkWin(myBoard, numMoves) {
		row, col, value, numMoves = getMove(myBoard, numMoves)
		fmt.Printf("inside loop\n")
		myBoard[row][col] = value
		fmt.Printf("Values are %d %d %d \n", row, col, value)
		fmt.Printf("The value from the array is %d \n", myBoard[row][col])
		printBoard(myBoard)
	}
	if winner != 3 {
		fmt.Printf("Congratulations to player %s !\n", playerLetter[winner])
	} else {
		fmt.Printf("The game is a draw.\n")
	}
}

func getMove(board [3][3]int, newMoves int, letter string) (row, col, value, outMoves int) {
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
	row++
	col++
	return row, col, value, newMoves
}

func printBoard(board [3][3]int) {
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

func checkWin(board [3][3]int, numMoves int) int {
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

func computerMove(board [3][3]int) {

}

func getLetter() string {
	var err error
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Select either x or o")
	for input == "NaN" {
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		input = strings.TrimSpace(input)
		if input == "x" || input == "o" {
			log.Printf("Selected letter is %s\n", input)
			return input
		}
	}
	return input
}

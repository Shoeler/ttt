package board

import (
	"context"
	"fmt"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
)

func CheckWin(ctx context.Context, tracer trace.Tracer, myBoard [3][3]int) (int, [][2]int) {
	_, childSpan := tracer.Start(ctx, "CheckWin")
	defer childSpan.End()

	//The idea here is to check every row for all of one player move and then
	//check the two diagonals and if any of them have that the player is a winner
	for i := range 3 {
		if myBoard[i][0] != 0 && myBoard[i][0] == myBoard[i][1] && myBoard[i][1] == myBoard[i][2] {
			childSpan.SetAttributes(attribute.Int("winner", myBoard[i][0]))
			return myBoard[i][0], [][2]int{{i, 0}, {i, 1}, {i, 2}}
		}
		if myBoard[0][i] != 0 && myBoard[0][i] == myBoard[1][i] && myBoard[1][i] == myBoard[2][i] {
			childSpan.SetAttributes(attribute.Int("winner", myBoard[0][i]))
			return myBoard[0][i], [][2]int{{0, i}, {1, i}, {2, i}}
		}
	}
	// Check diagonals
	if myBoard[0][0] != 0 && myBoard[0][0] == myBoard[1][1] && myBoard[1][1] == myBoard[2][2] {
		childSpan.SetAttributes(attribute.Int("winner", myBoard[0][0]))
		return myBoard[0][0], [][2]int{{0, 0}, {1, 1}, {2, 2}}
	}
	if myBoard[0][2] != 0 && myBoard[0][2] == myBoard[1][1] && myBoard[1][1] == myBoard[2][0] {
		childSpan.SetAttributes(attribute.Int("winner", myBoard[0][2]))
		return myBoard[0][2], [][2]int{{0, 2}, {1, 1}, {2, 0}}
	}
	childSpan.SetAttributes(attribute.Int("winner", 0))
	return 0, nil
}

func CheckDraw(ctx context.Context, tracer trace.Tracer, myBoard [3][3]int) bool {
	_, childSpan := tracer.Start(ctx, "CheckDraw")
	var fullRows int = 0
	defer childSpan.End()
	arryStr := "[" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(myBoard)), ","), "[]") + "]"
	for i := range 3 {
		if myBoard[i][0] != 0 && myBoard[i][1] != 0 && myBoard[i][2] != 0 { // any zeros, not a tie
			fullRows++
		}
	}
	if fullRows == 3 {
		childSpan.SetAttributes(attribute.String("myBoard", arryStr))
		return true
	} else {
		childSpan.SetAttributes(attribute.String("myBoard", arryStr))
		return false
	}
}

func PrintBoard(ctx context.Context, tracer trace.Tracer, myBoard [3][3]int, highlight [][2]int) {
	_, childSpan := tracer.Start(ctx, "PrintBoard")
	defer childSpan.End()

	highlightMap := map[[2]int]bool{}
	for _, coord := range highlight {
		highlightMap[coord] = true
	}

	fmt.Println("Current Board:")
	for i := range 3 {
		fmt.Print(" ")
		for j := range 3 {
			symbol := "-"
			color := ""

			if myBoard[i][j] == 1 {
				symbol = "X"
			} else if myBoard[i][j] == 2 {
				symbol = "O"
			}

			if highlightMap[[2]int{i, j}] {
				color = ColorGreen // highlight winner with green
			}

			fmt.Print(color + symbol + ColorReset)
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

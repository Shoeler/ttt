package board

import (
	"context"
	"fmt"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func CheckWin(ctx context.Context, tracer trace.Tracer, myBoard [3][3]int) int {
	ctx, childSpan := tracer.Start(ctx, "CheckWin")
	defer childSpan.End()

	//The idea here is to check every row for all of one player move and then
	//check the two diagonals and if any of them have that the player is a winner
	for i := 0; i < 3; i++ {
		if myBoard[i][0] != 0 && myBoard[i][0] == myBoard[i][1] && myBoard[i][1] == myBoard[i][2] {
			childSpan.SetAttributes(attribute.Int("winner", myBoard[i][0]))
			return myBoard[i][0]
		}
		if myBoard[0][i] != 0 && myBoard[0][i] == myBoard[1][i] && myBoard[1][i] == myBoard[2][i] {
			childSpan.SetAttributes(attribute.Int("winner", myBoard[0][i]))
			return myBoard[0][i]
		}
	}
	// Check diagonals
	if myBoard[0][0] != 0 && myBoard[0][0] == myBoard[1][1] && myBoard[1][1] == myBoard[2][2] {
		childSpan.SetAttributes(attribute.Int("winner", myBoard[0][0]))
		return myBoard[0][0]
	}
	if myBoard[0][2] != 0 && myBoard[0][2] == myBoard[1][1] && myBoard[1][1] == myBoard[2][0] {
		childSpan.SetAttributes(attribute.Int("winner", myBoard[0][2]))
		return myBoard[0][2]
	}
	childSpan.SetAttributes(attribute.Int("winner", 0))
	return 0 // no winner
}

func CheckDraw(ctx context.Context, tracer trace.Tracer, myBoard [3][3]int) bool {
	ctx, childSpan := tracer.Start(ctx, "CheckDraw")
	var fullRows int = 0
	defer childSpan.End()
	arryStr := "[" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(myBoard)), ","), "[]") + "]"
	for i := 0; i < 3; i++ {
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

func PrintBoard(ctx context.Context, tracer trace.Tracer, myBoard [3][3]int) {
	ctx, childSpan := tracer.Start(ctx, "PrintBoard")
	defer childSpan.End()
	fmt.Println("Current Board:")
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
		for j := 0; j < 3; j++ {
			symbol := " "
			switch myBoard[i][j] {
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

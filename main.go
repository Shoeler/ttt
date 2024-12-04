package main

import (
	"context"
	"fmt"
	"log"
	"ttt/board"
	"ttt/computer"
	"ttt/player"

	"github.com/honeycombio/otel-config-go/otelconfig"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

const (
	PlayerX = 1
	PlayerO = 2
	Empty   = 0
)

// var myBoard [3][3]int // Initialize the board to zeros
var myBoard [3][3]int = [3][3]int{{2, 0, 0}, {1, 1, 2}, {0, 0, 0}}
var winner, row, col, playerLetterNum int

func main() {
	ctx := context.Background()
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry()
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}
	defer otelShutdown()

	tracer := otel.Tracer("tic-tac-toe")
	ctx, parentSpan := tracer.Start(ctx, "parent-span")
	defer parentSpan.End()

	type Move struct {
		Row int
		Col int
	}
	var computerLetterNum int
	var playerLetterArry = [3]string{"-", "x", "o"}
	fmt.Println(" New Game ")
	board.Print(ctx, tracer, myBoard)
	playerLetterNum = player.GetLetter(ctx, tracer)
	parentSpan.SetAttributes(attribute.Int("playerLetterNum", playerLetterNum))
	if playerLetterNum == 1 {
		parentSpan.SetAttributes(attribute.Int("computerLetterNum", 2))
		computerLetterNum = 2
	} else if playerLetterNum == 2 {
		parentSpan.SetAttributes(attribute.Int("computerLetterNum", 1))
		computerLetterNum = 1
	}
	for winner = 0; winner <= 0; winner = board.CheckWin(ctx, tracer, myBoard) { // This is the per-turn loop
		row, col = player.GetMove(ctx, tracer, myBoard)
		if board.CheckDraw(ctx, tracer, myBoard) {
			winner = 3
		} else if board.CheckWin(ctx, tracer, myBoard) != 0 {
			winner = playerLetterNum
		}
		if myBoard[row][col] == 0 { // Make sure the cell is empty
			myBoard[row][col] = playerLetterNum
		} else {
			log.Println("Error - This should not happen")
			parentSpan.SetStatus(codes.Error, "Tried to change a cell already full")
		}

		row, col = computer.GetBestMove(ctx, tracer, myBoard, computerLetterNum, playerLetterNum)
		myBoard[row][col] = computerLetterNum
		board.Print(ctx, tracer, myBoard)
	}
	if winner == 1 || winner == 2 {
		fmt.Printf("Congratulations to player %s !\n", playerLetterArry[winner])
	} else if winner == 3 {
		fmt.Printf("The game is a draw.\n")
	} else {
		fmt.Println("Error - This should not happen, check honeycomb")
		parentSpan.SetStatus(codes.Error, "This should not happen")
	}
}

// To-do:  Get x being first, minimax working, write tests
// Bugs: Full board does not generate draw

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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

var myBoard [3][3]int // Initialize the board to zeros
// var myBoard [3][3]int = [3][3]int{{2, 0, 0}, {1, 1, 2}, {0, 0, 0}} //debug to simplify the board
var row, col, playerLetterNum int
var gameNotEnded int = 0
var computerTurn = false

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	ctx := context.Background()
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry()
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}
	defer otelShutdown()

	tracer := otel.Tracer("tic-tac-toe")
	ctx, parentSpan := tracer.Start(ctx, "parent-span")
	defer parentSpan.End()
	go func() {
		<-c
		parentSpan.End()
		os.Exit(1)
	}()

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
		computerTurn = false
	} else if playerLetterNum == 2 {
		parentSpan.SetAttributes(attribute.Int("computerLetterNum", 1))
		computerLetterNum = 1
		computerTurn = true
		fmt.Sprintln("Computer will move first")
	}

	for gameNotEnded == 0 { // This is the per-turn loop
		if computerTurn == false {
			row, col = player.GetMove(ctx, tracer, myBoard)
			if myBoard[row][col] == 0 { // Make sure the cell is empty
				myBoard[row][col] = playerLetterNum
			} else {
				log.Println("Error - This should not happen, tried to change cell already full")
				parentSpan.SetStatus(codes.Error, "Tried to change a cell already full")
				parentSpan.SetAttributes(attribute.Bool("computerTurn", computerTurn))
			}
			if board.CheckDraw(ctx, tracer, myBoard) {
				gameNotEnded = 3
			} else if board.CheckWin(ctx, tracer, myBoard) != 0 {
				gameNotEnded = playerLetterNum
			}

			computerTurn = true
		} else if computerTurn {
			row, col = computer.GetBestMove(ctx, tracer, myBoard, computerLetterNum, playerLetterNum)
			// myBoard[row][col] = computerLetterNum
			if myBoard[row][col] == 0 { // Make sure the cell is empty
				myBoard[row][col] = computerLetterNum
				fmt.Printf("Computer move is %d,%d\n\n", row, col)
			} else {
				log.Println("Error - This should not happen, tried to change cell already full")
				parentSpan.SetStatus(codes.Error, "Tried to change a cell already full")
				parentSpan.SetAttributes(attribute.Bool("computerTurn", computerTurn))
			}
			if board.CheckDraw(ctx, tracer, myBoard) {
				gameNotEnded = 3
			} else if board.CheckWin(ctx, tracer, myBoard) != 0 {
				gameNotEnded = computerLetterNum
			}
			computerTurn = false
		}

		board.Print(ctx, tracer, myBoard)
	}
	if gameNotEnded == 1 || gameNotEnded == 2 {
		fmt.Printf("Congratulations to player %s !\n", playerLetterArry[gameNotEnded])
	} else if gameNotEnded == 3 {
		fmt.Printf("The game is a draw.\n")
	} else {
		fmt.Println("Error - This should not happen, check honeycomb")
		parentSpan.SetStatus(codes.Error, "This should not happen")
	}
}

// To-do:  Get x being first, minimax working, write tests

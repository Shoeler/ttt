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
var row, col, playerLetterNum int
var gameEnded bool = false
var winner int = 0
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
	var computerLetterNum int
	var playerLetterArry = [3]string{"-", "X", "O"}
	fmt.Println(" New Game ")
	board.PrintBoard(ctx, tracer, myBoard)
	playerLetterNum = player.GetLetter(ctx, tracer)
	parentSpan.SetAttributes(attribute.Int("playerLetterNum", playerLetterNum))
	if playerLetterNum == PlayerX {
		parentSpan.SetAttributes(attribute.Int("computerLetterNum", PlayerO))
		computerLetterNum = PlayerO
		computerTurn = false
	} else if playerLetterNum == PlayerO {
		parentSpan.SetAttributes(attribute.Int("computerLetterNum", PlayerX))
		computerLetterNum = PlayerX
		computerTurn = true
		fmt.Sprintln("Computer will move first")
	}

	for gameEnded == false { // This is the per-turn loop
		if computerTurn == false {
			row, col = player.GetMove(ctx, tracer, myBoard)
			if myBoard[row][col] == Empty { // Make sure the cell is empty
				myBoard[row][col] = playerLetterNum
			} else {
				log.Println("Error - This should not happen, tried to change cell already full")
				parentSpan.SetStatus(codes.Error, "Tried to change a cell already full")
				parentSpan.SetAttributes(attribute.Bool("computerTurn", computerTurn))
			}
			if board.CheckDraw(ctx, tracer, myBoard) {
				gameEnded = true
				winner = 3
				break
			} else if board.CheckWin(ctx, tracer, myBoard) != 0 {
				gameEnded = true
				winner = playerLetterNum
				break
			}
			computerTurn = true
		} else if computerTurn {
			row, col = computer.GetBestMove(ctx, tracer, myBoard, computerLetterNum, playerLetterNum)
			if myBoard[row][col] == Empty { // Make sure the cell is empty
				myBoard[row][col] = computerLetterNum
				fmt.Printf("Computer move is %d,%d\n\n", row+1, col+1)
			} else {
				log.Println("Error - This should not happen, tried to change cell already full")
				parentSpan.SetStatus(codes.Error, "Tried to change a cell already full")
				parentSpan.SetAttributes(attribute.Bool("computerTurn", computerTurn))
			}
			if board.CheckDraw(ctx, tracer, myBoard) {
				gameEnded = true
				winner = 3
				break
			} else if board.CheckWin(ctx, tracer, myBoard) != 0 {
				gameEnded = true
				winner = computerLetterNum
				break
			}
			computerTurn = false
		}

		board.PrintBoard(ctx, tracer, myBoard)
	}
	board.PrintBoard(ctx, tracer, myBoard) // Always print at end
	switch winner {
	case 1, 2:
		fmt.Printf("Congratulations to player %s!\n", playerLetterArry[winner])
	case 3:
		fmt.Println("The game is a draw.")
		parentSpan.SetAttributes(attribute.String("gameEnd", "draw"))
	default:
		fmt.Println("Error - This should not happen, check honeycomb")
		parentSpan.SetStatus(codes.Error, "This should not happen")
	}
}

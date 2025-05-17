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
var gameEnded bool
var winner int
var winLine [][2]int
var computerTurn bool

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
	board.PrintBoard(ctx, tracer, myBoard, nil)
	playerLetterNum = player.GetLetter(ctx, tracer)
	parentSpan.SetAttributes(attribute.Int("playerLetterNum", playerLetterNum))

	if playerLetterNum == PlayerX {
		computerLetterNum = PlayerO
		computerTurn = false
	} else {
		computerLetterNum = PlayerX
		computerTurn = true
		fmt.Println("Computer will move first")
	}
	parentSpan.SetAttributes(attribute.Int("computerLetterNum", computerLetterNum))

	for !gameEnded {
		if !computerTurn {
			row, col = player.GetMove(ctx, tracer, myBoard)
			if myBoard[row][col] == Empty { // Make sure the cell is empty
				myBoard[row][col] = playerLetterNum
			} else {
				log.Println("Error - This should not happen, tried to change cell already full")
				parentSpan.SetStatus(codes.Error, "Tried to change a cell already full")
				parentSpan.SetAttributes(attribute.Bool("computerTurn", computerTurn))
			}
		} else {
			row, col = computer.GetBestMove(ctx, tracer, myBoard, computerLetterNum, playerLetterNum)
			if myBoard[row][col] == Empty { // Make sure the cell is empty
				myBoard[row][col] = computerLetterNum
				fmt.Printf("Computer move is %d,%d\n\n", row+1, col+1)
			} else {
				log.Println("Error - This should not happen, tried to change cell already full")
				parentSpan.SetStatus(codes.Error, "Tried to change a cell already full")
				parentSpan.SetAttributes(attribute.Bool("computerTurn", computerTurn))
			}
		}

		if board.CheckDraw(ctx, tracer, myBoard) {
			gameEnded = true
			winner = 3
			break
		}

		var result int
		result, winLine = board.CheckWin(ctx, tracer, myBoard)
		if result != 0 {
			gameEnded = true
			winner = result
			break
		}

		computerTurn = !computerTurn
		board.PrintBoard(ctx, tracer, myBoard, nil)
	}

	board.PrintBoard(ctx, tracer, myBoard, winLine)
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

package player

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func GetMove(ctx context.Context, tracer trace.Tracer, board [3][3]int) (row, col int) {
	ctx, childSpan := tracer.Start(ctx, "GetMove")
	defer childSpan.End()
	reader := bufio.NewReader(os.Stdin)
	for validMove := 0; validMove != 1; {
		fmt.Println("Enter your move in the format 'row,col' in the range 1-3 (e.g., 1,2):")
		input, err := reader.ReadString('\n')
		if err != nil {
			childSpan.AddEvent("Input error on row,col")
			childSpan.SetStatus(codes.Error, "input broke")
			childSpan.RecordError(err)
			childSpan.End()
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		parts := strings.Split(input, ",")
		if len(parts) != 2 {
			log.Println("invalid input format, please use 'row,col'")
			continue
		}
		row, err = strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("invalid row number: %s\n", parts[0])
		}
		col, err = strconv.Atoi(parts[1])
		if err != nil {
			log.Printf("invalid column number: %s\n", parts[1])
		}
		if (row > 0 && row < 4) && (col > 0 && col < 4) {
			row-- // row is 1-based on input but 0-based on storage
			col-- // col too
		} else {
			log.Printf("Error - rows and columns must be between 1 and 3.  You sent row %d and col %d\n", row, col)
			continue
		}

		if board[row][col] == 0 {
			validMove = 1
		} else {
			log.Println("That cell is already used - try again")
		}
	}
	childSpan.SetAttributes(attribute.Int("humanRow", row+1))
	childSpan.SetAttributes(attribute.Int("humanCol", col+1))
	return row, col
}

func GetLetter(ctx context.Context, tracer trace.Tracer) int {
	ctx, childSpan := tracer.Start(ctx, "GetLetter")
	defer childSpan.End()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Select either x or o:")
		input, err := reader.ReadString('\n')
		if err != nil {
			childSpan.AddEvent("Something broke in GetLetter")
			childSpan.SetStatus(codes.Error, "GetLetter input broke")
			childSpan.RecordError(err)
			childSpan.End()
			log.Fatal(err)
		}
		input = strings.ToLower(strings.TrimSpace(input))
		switch input {
		case "x":
			childSpan.SetAttributes(attribute.Int("chosenLetter", 1))
			return 1
		case "o":
			childSpan.SetAttributes(attribute.Int("chosenLetter", 2))
			return 2
		default:
			fmt.Println(`Error, please only enter an x or an o`)
		}
	}
}

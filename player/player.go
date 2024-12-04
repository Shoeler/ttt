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

func GetMove(ctx context.Context, tracer trace.Tracer) (row, col int) {
	var err error
	ctx, childSpan := tracer.Start(ctx, "GetMove")
	defer childSpan.End()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your move in the format 'row,col' in the range 1-3 (e.g., 1,2):")
	input, err := reader.ReadString('\n')
	if err != nil {
		childSpan.AddEvent("Input error on row,col")
		childSpan.SetStatus(codes.Error, "input broke")
		childSpan.RecordError(err)
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
	row-- // row is 1-based on input but 0-based on storage
	col-- // col too
	childSpan.SetAttributes(attribute.Int("row", row))
	childSpan.SetAttributes(attribute.Int("col", col))
	fmt.Printf("Row is %d and Col is %d\n", row, col)
	childSpan.End()
	return row, col
}

func GetLetter(ctx context.Context, tracer trace.Tracer) int {
	ctx, childSpan := tracer.Start(ctx, "GetLetter")
	defer childSpan.End()
	var err error
	var input string
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Select either x or o:")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Println(`Please only enter an x or an o`)
			childSpan.AddEvent("Incorrect letter input")
			childSpan.SetStatus(codes.Error, "input broke")
			childSpan.RecordError(err)
			break
		}
		input = strings.TrimSpace(input)
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
	childSpan.End()
	return 0
}

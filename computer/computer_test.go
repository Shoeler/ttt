package computer

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

func newMockTracer() trace.Tracer {
	tp := noop.NewTracerProvider()
	return tp.Tracer("mock-tracer")
}

func TestGetBestMove(t *testing.T) {
	tracer := newMockTracer()
	ctx := context.Background()
	type tests struct {
		name        string
		board       [3][3]int
		computer    int
		human       int
		rowExpected int
		colExpected int
	}

	var listTests = []tests{
		{
			name: "Computer(o) should block human (x) win",
			board: [3][3]int{
				{2, 0, 1},
				{0, 1, 0},
				{0, 0, 0},
			},
			human:       1,
			computer:    2,
			rowExpected: 3,
			colExpected: 1,
		},
		{
			name: "Computer(x) should block human (o) win",
			board: [3][3]int{
				{2, 0, 0},
				{1, 1, 0},
				{0, 0, 0},
			},
			human:       1,
			computer:    2,
			rowExpected: 2,
			colExpected: 3,
		},
		{
			name: "Brandon's test - computer should pick best next move to win not blocking",
			board: [3][3]int{
				{1, 2, 0},
				{1, 2, 1},
				{0, 0, 0},
			},
			human:       1,
			computer:    2,
			rowExpected: 3,
			colExpected: 2,
		},
	}

	for _, test := range listTests {
		t.Run(test.name, func(t *testing.T) {
			row, col := GetBestMove(ctx, tracer, test.board, test.computer, test.human) //returns 0-based
			row++                                                                       // make 1-based
			col++                                                                       //make 1-based
			if row != test.rowExpected || col != test.colExpected {
				t.Errorf("For test %s expected row %d col %d, got row = %d col = %d", test.name, test.rowExpected, test.colExpected, row, col)
			}
		})
	}
}

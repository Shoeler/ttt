package board

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

func TestCheckWin(t *testing.T) {
	tracer := newMockTracer()
	ctx := context.Background()

	tests := []struct {
		name     string
		board    [3][3]int
		expected int
	}{
		{
			name: "Row win for X",
			board: [3][3]int{
				{1, 1, 1},
				{0, 2, 0},
				{2, 0, 0},
			},
			expected: 1,
		},
		{
			name: "Column win for O",
			board: [3][3]int{
				{1, 2, 1},
				{0, 2, 0},
				{1, 2, 0},
			},
			expected: 2,
		},
		{
			name: "Diagonal win for X",
			board: [3][3]int{
				{1, 2, 0},
				{2, 1, 0},
				{0, 2, 1},
			},
			expected: 1,
		},
		{
			name: "Diagonal win for player 0",
			board: [3][3]int{
				{2, 1, 0},
				{1, 2, 0},
				{0, 1, 2},
			},
			expected: 2,
		},
		{
			name: "Draw",
			board: [3][3]int{
				{1, 2, 1},
				{2, 1, 2},
				{2, 1, 2},
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CheckWin(ctx, tracer, test.board)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestCheckDraw(t *testing.T) {
	tracer := newMockTracer()
	ctx := context.Background()

	tests := []struct {
		name     string
		board    [3][3]int
		expected bool
	}{
		{
			name: "Draw",
			board: [3][3]int{
				{1, 2, 1},
				{2, 1, 2},
				{2, 1, 2},
			},
			expected: true,
		},
		{
			name: "Not a draw - empty spaces",
			board: [3][3]int{
				{1, 2, 1},
				{2, 1, 0},
				{2, 1, 2},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CheckDraw(ctx, tracer, test.board)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestPrintBoard(t *testing.T) {
	tracer := newMockTracer()
	ctx := context.Background()

	board := [3][3]int{
		{1, 2, 1},
		{2, 0, 2},
		{1, 2, 1},
	}

	t.Run("Print board test", func(t *testing.T) {
		PrintBoard(ctx, tracer, board)
	})
}

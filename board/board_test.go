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
		name         string
		board        [3][3]int
		expected     int
		expectedLine [][2]int
	}{
		{
			name: "Row win for X",
			board: [3][3]int{
				{1, 1, 1},
				{0, 2, 0},
				{2, 0, 0},
			},
			expected:     1,
			expectedLine: [][2]int{{0, 0}, {0, 1}, {0, 2}},
		},
		{
			name: "Column win for O",
			board: [3][3]int{
				{1, 2, 1},
				{0, 2, 0},
				{1, 2, 0},
			},
			expected:     2,
			expectedLine: [][2]int{{0, 1}, {1, 1}, {2, 1}},
		},
		{
			name: "Diagonal win for X",
			board: [3][3]int{
				{1, 2, 0},
				{2, 1, 0},
				{0, 2, 1},
			},
			expected:     1,
			expectedLine: [][2]int{{0, 0}, {1, 1}, {2, 2}},
		},
		{
			name: "Diagonal win for player 0",
			board: [3][3]int{
				{0, 1, 2},
				{1, 2, 0},
				{2, 1, 0},
			},
			expected:     2,
			expectedLine: [][2]int{{0, 2}, {1, 1}, {2, 0}},
		},		
		{
			name: "Draw",
			board: [3][3]int{
				{1, 2, 1},
				{2, 1, 2},
				{2, 1, 2},
			},
			expected:     0,
			expectedLine: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			winner, winLine := CheckWin(ctx, tracer, test.board)
			if winner != test.expected {
				t.Errorf("expected winner %d, got %d", test.expected, winner)
			}
			if len(winLine) != len(test.expectedLine) {
				t.Errorf("expected win line %v, got %v", test.expectedLine, winLine)
			} else {
				for i := range winLine {
					if winLine[i] != test.expectedLine[i] {
						t.Errorf("expected win line %v, got %v", test.expectedLine, winLine)
						break
					}
				}
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

	t.Run("Print board without highlight", func(t *testing.T) {
		board := [3][3]int{
			{1, 2, 1},
			{2, 0, 2},
			{1, 2, 1},
		}
		PrintBoard(ctx, tracer, board, nil)
	})

	t.Run("Print board with highlighted win", func(t *testing.T) {
		board := [3][3]int{
			{2, 2, 2},
			{1, 1, 0},
			{0, 0, 0},
		}
		highlight := [][2]int{{0, 0}, {0, 1}, {0, 2}}
		PrintBoard(ctx, tracer, board, highlight)
	})
}

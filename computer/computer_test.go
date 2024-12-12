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

	tests := []struct {
		name     string
		board    [3][3]int
		computer int
		human    int
		expected [2]int
	}{
		{
			name: "Computer(o) should block human (x) win",
			board: [3][3]int{
				{2, 0, 1},
				{0, 1, 0},
				{0, 0, 0},
			},
			human:    1,
			computer: 2,
			expected: [2]int{3, 1},
		},
		{
			name: "Computer(x) should block human (o) win",
			board: [3][3]int{
				{2, 1, 1},
				{2, 1, 0},
				{0, 0, 0},
			},
			human:    1,
			computer: 2,
			expected: [2]int{3, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			row, col := GetBestMove(ctx, tracer, test.board, test.computer, test.human)
			if row+1 != test.expected[0] && col+1 != test.expected[1] {
				t.Errorf("expected %d, got row = %d col = %d", test.expected, row, col)
			}
		})
	}
}

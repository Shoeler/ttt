package computer

import (
	"context"
	"ttt/board"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// GetBestMove calculates and returns the best move for the AI player.
func GetBestMove(ctx context.Context, tracer trace.Tracer, gameBoard [3][3]int, aiPlayer int, humanPlayer int) (row, col int) {
	ctx, childSpan := tracer.Start(ctx, "GetBestMove")
	defer childSpan.End()
	bestScore := -10
	var nodesEvaluated int
	for i := range 3 {
		for j := range 3 {
			if gameBoard[i][j] == 0 {
				gameBoard[i][j] = aiPlayer
				score := minimax(ctx, tracer, gameBoard, 0, false, aiPlayer, humanPlayer, &nodesEvaluated)
				gameBoard[i][j] = 0
				if score > bestScore {
					bestScore = score
					row, col = i, j
				}
			}
		}
	}
	childSpan.SetAttributes(
		attribute.Int("compRow", row+1),
		attribute.Int("compCol", col+1),
		attribute.Int("bestScore", bestScore),
		attribute.Int("nodesEvaluated", nodesEvaluated),
	)
	return row, col
}

func minimax(ctx context.Context, tracer trace.Tracer, gameBoard [3][3]int, depth int, isMaximizing bool, aiPlayer, humanPlayer int, nodesEvaluated *int) int {
	*nodesEvaluated++
	win, _ := board.CheckWin(ctx, tracer, gameBoard)
	if win == humanPlayer {
		return -10 + depth
	}
	if win == aiPlayer {
		return 10 - depth
	}
	if board.CheckDraw(ctx, tracer, gameBoard) {
		return 0
	}
	if isMaximizing {
		bestScore := -10
		for i := range 3 {
			for j := range 3 {
				if gameBoard[i][j] == 0 {
					gameBoard[i][j] = aiPlayer
					score := minimax(ctx, tracer, gameBoard, depth+1, false, aiPlayer, humanPlayer, nodesEvaluated)
					gameBoard[i][j] = 0
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := 10
		for i := range 3 {
			for j := range 3 {
				if gameBoard[i][j] == 0 {
					gameBoard[i][j] = humanPlayer
					score := minimax(ctx, tracer, gameBoard, depth+1, true, aiPlayer, humanPlayer, nodesEvaluated)
					gameBoard[i][j] = 0
					if score < bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	}
}

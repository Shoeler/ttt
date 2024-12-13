package computer

import (
	"context"
	"math"
	"ttt/board"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// GetBestMove calculates and returns the best move for the AI player.
func GetBestMove(ctx context.Context, tracer trace.Tracer, gameBoard [3][3]int, aiPlayer int, humanPlayer int) (row, col int) {
	ctx, childSpan := tracer.Start(ctx, "GetBestMove")
	defer childSpan.End()
	bestScore := math.MinInt
	// iterate through the board and check every empty space to see if it's the max score for the computer move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gameBoard[i][j] == 0 {
				gameBoard[i][j] = aiPlayer
				score := minimax(ctx, tracer, gameBoard, 0, false, aiPlayer, humanPlayer)
				gameBoard[i][j] = 0
				if score > bestScore {
					bestScore = score
					row, col = i, j
				}
			}
		}
	}
	childSpan.SetAttributes(attribute.Int("compRow", row+1))
	childSpan.SetAttributes(attribute.Int("compCol", col+1))
	return row, col
}

func minimax(ctx context.Context, tracer trace.Tracer, gameBoard [3][3]int, depth int, isMaximizing bool, aiPlayer, humanPlayer int) int {
	ctx, childSpan := tracer.Start(ctx, "minimax")
	defer childSpan.End()
	win := board.CheckWin(ctx, tracer, gameBoard)
	if win == humanPlayer {
		return math.MinInt32 + depth
	}
	if win == aiPlayer {
		return math.MaxInt32 - depth
	}
	if board.CheckDraw(ctx, tracer, gameBoard) {
		return 0
	}
	if isMaximizing {
		bestScore := math.MinInt32
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if gameBoard[i][j] == 0 {
					gameBoard[i][j] = aiPlayer
					score := minimax(ctx, tracer, gameBoard, depth+1, false, aiPlayer, humanPlayer)
					gameBoard[i][j] = 0
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		childSpan.SetAttributes(attribute.Int("bestScore", bestScore))
		return bestScore
	} else {
		bestScore := math.MaxInt32
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if gameBoard[i][j] == 0 {
					gameBoard[i][j] = humanPlayer
					score := minimax(ctx, tracer, gameBoard, depth+1, true, aiPlayer, humanPlayer)
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

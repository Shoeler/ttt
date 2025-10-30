package computer

import (
	"context"
	"ttt/board"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// GetBestMove calculates and returns the best move for the AI player.
//
// It uses the minimax algorithm to determine the optimal move.
//
// Parameters:
//   - ctx: The context for OpenTelemetry tracing.
//   - tracer: The tracer for OpenTelemetry.
//   - gameBoard: A 3x3 integer array representing the game board.
//   - aiPlayer: The integer representing the AI player.
//   - humanPlayer: The integer representing the human player.
//
// Returns:
//   - The row and column of the best move.
func GetBestMove(ctx context.Context, tracer trace.Tracer, gameBoard [3][3]int, aiPlayer int, humanPlayer int) (row, col int) {
	ctx, childSpan := tracer.Start(ctx, "GetBestMove")
	defer childSpan.End()
	bestScore := -10
	// iterate through the board and check every empty space to see if it's the max score for the computer move
	for i := range 3 {
		for j := range 3 {
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

// minimax implements the minimax algorithm to determine the best move.
//
// It recursively explores the game tree to find the move that maximizes the
// AI's score while minimizing the human player's score.
//
// Parameters:
//   - ctx: The context for OpenTelemetry tracing.
//   - tracer: The tracer for OpenTelemetry.
//   - gameBoard: A 3x3 integer array representing the game board.
//   - depth: The current depth in the game tree.
//   - isMaximizing: A boolean indicating whether the current player is maximizing
//     their score.
//   - aiPlayer: The integer representing the AI player.
//   - humanPlayer: The integer representing the human player.
//
// Returns:
//   - The score of the best move.
func minimax(ctx context.Context, tracer trace.Tracer, gameBoard [3][3]int, depth int, isMaximizing bool, aiPlayer, humanPlayer int) int {
	ctx, childSpan := tracer.Start(ctx, "minimax")
	defer childSpan.End()
	win, _ := board.CheckWin(ctx, tracer, gameBoard)
	if win == humanPlayer {
		winValue := -10 + depth
		childSpan.SetAttributes(attribute.Int64("minimax_return", int64(winValue)))
		return winValue
	}
	if win == aiPlayer {
		winValue := 10 - depth
		childSpan.SetAttributes(attribute.Int64("minimax_return", int64(winValue)))
		return winValue
	}
	if board.CheckDraw(ctx, tracer, gameBoard) {
		childSpan.SetAttributes(attribute.Int64("minimax_return", 0))
		return 0
	}
	if isMaximizing {
		bestScore := -10
		for i := range 3 {
			for j := range 3 {
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
		bestScore := 10
		for i := range 3 {
			for j := range 3 {
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
		childSpan.SetAttributes(attribute.Int("bestScore", bestScore))
		return bestScore
	}
}

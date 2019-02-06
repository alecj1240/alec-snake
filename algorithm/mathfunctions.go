package algorithm

import (
	"github.com/alecj1240/astart/api"
)

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhatten(pointA api.Coord, pointB api.Coord) int {
	var manhattenX = Abs(pointB.X - pointA.X)
	var manhattenY = Abs(pointB.Y - pointA.Y)
	var manhattenDistance = Abs(manhattenX + manhattenY)
	return manhattenDistance
}

func onBoard(square api.Coord, boardHeight int, boardWidth int) bool {
	if square.X > 0 || square.X < boardWidth {
		return true
	}

	if square.Y > 0 || square.Y < boardHeight {
		return true
	}

	return false
}

func squareBlocked(point api.Coord, board api.Board) bool {
	for i := 0; i < len(board.Snakes); i++ {
		for j := 0; j < len(board.Snakes[i].Body); j++ {
			if board.Snakes[i].Body[j].X == point.X && board.Snakes[i].Body[j].Y == point.Y {
				return true
			}
		}
	}

	return false
}

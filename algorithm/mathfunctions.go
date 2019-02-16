package algorithm

import (
	"github.com/alecj1240/astart/api"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhatten(pointA api.Coord, pointB api.Coord) int {
	var manhattenX = abs(pointB.X - pointA.X)
	var manhattenY = abs(pointB.Y - pointA.Y)
	var manhattenDistance = manhattenX + manhattenY
	return manhattenDistance
}

func onBoard(square api.Coord, boardHeight int, boardWidth int) bool {
	if square.X >= 0 || square.X < boardWidth {
		return true
	}

	if square.Y >= 0 || square.Y < boardHeight {
		return true
	}

	return false
}

/*
	Determines if a snake is blocking the square
*/
func squareBlocked(point api.Coord, Snakes []api.Snake) bool {
	for i := 0; i < len(Snakes); i++ {
		for j := 0; j < len(Snakes[i].Body); j++ {
			if Snakes[i].Body[j].X == point.X && Snakes[i].Body[j].Y == point.Y {
				return true
			}
		}
	}

	return false
}

/*
Determines the heading between two points - must be side by side
*/
func Heading(startingPoint api.Coord, headingPoint api.Coord) string {
	if headingPoint.X > startingPoint.X {
		return "right"
	}
	if headingPoint.X < startingPoint.X {
		return "left"
	}
	if headingPoint.Y > startingPoint.Y {
		return "down"
	}
	if headingPoint.Y < startingPoint.Y {
		return "up"
	}

	return "up"
}

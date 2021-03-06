package algorithm

import (
	"github.com/alecj1240/astart/api"
)

// ChaseTail returns the coordinate of the position behind my tail
func ChaseTail(You []api.Coord) api.Coord {
	return You[len(You)-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Manhatten gives the difference between two points
func Manhatten(pointA api.Coord, pointB api.Coord) int {
	var manhattenX = abs(pointB.X - pointA.X)
	var manhattenY = abs(pointB.Y - pointA.Y)
	var manhattenDistance = manhattenX + manhattenY
	return manhattenDistance
}

// determines if the square is actually on the board
func OnBoard(square api.Coord, boardHeight int, boardWidth int) bool {
	if square.X >= 0 && square.X < boardWidth && square.Y >= 0 && square.Y < boardHeight {
		return true
	}

	return false
}

// determines if the square is blocked by a snake
func SquareBlocked(point api.Coord, Snakes []api.Snake) bool {
	for i := 0; i < len(Snakes); i++ {
		for j := 0; j < len(Snakes[i].Body); j++ {
			if Snakes[i].Body[j].X == point.X && Snakes[i].Body[j].Y == point.Y {
				if len(Snakes[i].Body)-1 == j {
					return false
				}

				return true
			}
		}
	}

	return false
}

// Heading determines the direction between two points - must be side by side
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

// NearestFood finds the closest food to the head of my snake
func NearestFood(FoodCoords []api.Coord, You api.Coord) api.Coord {
	var nearestFood = FoodCoords[0]
	var nearestFoodF = Manhatten(FoodCoords[0], You)

	for i := 0; i < len(FoodCoords); i++ {
		if Manhatten(FoodCoords[i], You) < nearestFoodF {
			nearestFood = FoodCoords[i]
			nearestFoodF = Manhatten(FoodCoords[i], You)
		}
	}

	return nearestFood
}

// HeadOnCollision determines the nearest snake on the board based on the head of the snake
func HeadOnCollision(Destination api.Coord, Snakes []api.Snake, You api.Snake) bool {
	destinationAdjacents := GetAdjacentCoords(Destination)

	for i := 0; i < len(Snakes); i++ {
		for j := 0; j < len(destinationAdjacents); j++ {
			if Snakes[i].Body[0] == destinationAdjacents[j] && Snakes[i].ID != You.ID {
				return true
			}
		}
	}

	return false
}

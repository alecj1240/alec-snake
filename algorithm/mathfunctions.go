package algorithm

import (
	"github.com/alecj1240/astart/api"
)

func ChaseTail(You []api.Coord) api.Coord {
	return You[len(You)-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Manhatten(pointA api.Coord, pointB api.Coord) int {
	var manhattenX = abs(pointB.X - pointA.X)
	var manhattenY = abs(pointB.Y - pointA.Y)
	var manhattenDistance = manhattenX + manhattenY
	return manhattenDistance
}

func onBoard(square api.Coord, boardHeight int, boardWidth int) bool {
	if square.X >= 0 && square.X < boardWidth && square.Y >= 0 && square.Y < boardHeight {
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
				if len(Snakes[i].Body)-1 == j {
					return false
				}

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

// func NearestSnake(Snakes []api.Snake, You api.Snake) api.Coord {
// 	var nearestSnake = Snakes[0]
// 	var nearestSnakeDistance = Manhatten(Snakes[0].Body[0], You.Body[0])

// 	if Snakes[0].Body[0] == You.Body[0] {
// 		var nearestSnake = Snakes[1]
// 		var nearestSnakeDistance = Manhatten(Snakes[1].Body[0], You.Body[0])
// 	}

// 	for i := 0; i < len(Snakes); i++ {
// 		if Manhatten(Snakes[i].Body[0], You.Body[0]) < nearestSnakeDistance {
// 			if Snakes[i].Body[0] == You.Body[0] {
// 				continue
// 			}

// 			nearestSnake = Snakes[i]
// 			nearestSnakeDistance = Manhatten(Snakes[i].Body[0], You.Body[0])
// 		}
// 	}

// 	return nearestSnake.Body[0]
// }

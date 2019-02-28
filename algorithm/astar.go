package algorithm

import (
	"github.com/alecj1240/astart/api"
)

/*
	G - the amount of steps its taken to get to that Square
	H - the heuristic - estimation from this Square to the destination
	F - the sum of G and H
	ParentCoords - the coordinates of the previous step (Square)
*/

// Square holds: coordinates, G, H, F, parent coords
type Square struct {
	Coord        api.Coord
	G            int
	H            int
	F            int
	ParentCoords api.Coord
}

func getAdjacentCoords(Location api.Coord) []api.Coord {
	var adjacentCoords = make([]api.Coord, 0)
	adjacentCoords = append(adjacentCoords, api.Coord{X: Location.X + 1, Y: Location.Y})
	adjacentCoords = append(adjacentCoords, api.Coord{X: Location.X - 1, Y: Location.Y})
	adjacentCoords = append(adjacentCoords, api.Coord{X: Location.X, Y: Location.Y + 1})
	adjacentCoords = append(adjacentCoords, api.Coord{X: Location.X, Y: Location.Y - 1})

	return adjacentCoords
}

func removeFromOpenList(removalSquare Square, openList []Square) []Square {
	var newOpenList = make([]Square, 0)
	for i := 0; i < len(openList); i++ {
		if removalSquare.Coord != openList[i].Coord {
			newOpenList = append(newOpenList, openList[i])
		}
	}
	return newOpenList
}

func appendList(appendingSquare Square, Snakes []api.Snake, List []Square, BoardHeight int, BoardWidth int) []Square {
	if squareBlocked(appendingSquare.Coord, Snakes) == false && onBoard(appendingSquare.Coord, BoardHeight, BoardWidth) {
		List = append(List, appendingSquare)
	}

	return List
}

// reverseCoords reverses the path of coordinates so it's in chronological order
func reverseCoords(path []api.Coord) []api.Coord {
	for a := 0; a < len(path)/2; a++ {
		b := len(path) - a - 1
		path[a], path[b] = path[b], path[a]
	}
	return path
}

// Astar determines the best path to a point on the board.
func Astar(BoardHeight int, BoardWidth int, MySnake api.Snake, Snakes []api.Snake, Destination api.Coord) []api.Coord {
	closedList := make(map[api.Coord]bool)
	openList := make([]Square, 0)
	pathTracker := make(map[api.Coord]api.Coord)

	myHead := Square{Coord: MySnake.Body[0], G: 0, H: 0, F: 0}
	openList = append(openList, myHead)

	for len(openList) > 0 {

		// find the Square the least F on the open list
		var closeSquare = openList[0]
		for _, openItem := range openList {
			if openItem.F < closeSquare.F {
				closeSquare = openItem
			}
		}

		// put it on the closed list
		closedList[closeSquare.Coord] = true
		openList = removeFromOpenList(closeSquare, openList)
		// loop through leastSquares's adjacent tiles -- call them T
		closeNeighbours := getAdjacentCoords(closeSquare.Coord)

		for _, neighbour := range closeNeighbours {

			// 1. If T on the closed list, ignore it
			if closedList[neighbour] {
				continue
			}

			if neighbour == Destination {

				closedList[neighbour] = true

				path := make([]api.Coord, 0)
				path = append(path, Destination)
				path = append(path, neighbour)
				current := closeSquare.Coord
				path = append(path, current)

				_, pathway := pathTracker[current]

				for ; pathway; _, pathway = pathTracker[current] {
					current = pathTracker[current]
					path = append(path, current)
				}

				return reverseCoords(path)
			}

			// 2. If T is not on the open list add it
			for _, item := range openList {
				if neighbour == item.Coord {
					if squareBlocked(neighbour, Snakes) == false && onBoard(neighbour, BoardHeight, BoardWidth) {
						if (closeSquare.G+1)+Manhatten(neighbour, Destination) < item.F {
							item.F = (closeSquare.G + 1) + Manhatten(neighbour, Destination)
							item.G = closeSquare.G + 1
							item.H = Manhatten(neighbour, Destination)
							item.ParentCoords = neighbour

							pathTracker[item.Coord] = closeSquare.Coord
						}
					}
				}
			}

			var openSquare = Square{
				Coord:        neighbour,
				G:            closeSquare.G + 1,
				H:            Manhatten(neighbour, Destination),
				F:            (closeSquare.G + 1) + (Manhatten(neighbour, Destination)),
				ParentCoords: closeSquare.Coord,
			}

			pathTracker[neighbour] = closeSquare.Coord
			openList = appendList(openSquare, Snakes, openList, BoardHeight, BoardWidth)

		}

	}

	return nil

}

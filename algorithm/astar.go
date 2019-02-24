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

// create struct that will hold: coordinates, G, H, F, parent coords
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

func ReversePath(path []api.Coord) []api.Coord {
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - i - 1
		path[i], path[j] = path[j], path[i]
	}
	return path
}

/*
the a star algorithm
*/
func Astar(BoardHeight int, BoardWidth int, MySnake api.Snake, Snakes []api.Snake, Destination api.Coord) []api.Coord {
	closedList := make(map[api.Coord]bool)
	openList := make([]Square, 0)

	cameFrom := make(map[api.Coord]api.Coord)

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

				_, exists := cameFrom[current]

				for ; exists; _, exists = cameFrom[current] {
					current = cameFrom[current]
					path = append(path, current)
				}

				return ReversePath(path)
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

							cameFrom[item.Coord] = closeSquare.Coord
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

			cameFrom[neighbour] = closeSquare.Coord
			openList = appendList(openSquare, Snakes, openList, BoardHeight, BoardWidth)

		}

	}

	return nil

}

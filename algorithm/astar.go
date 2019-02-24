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

/*
the a star algorithm
*/
func Astar(BoardHeight int, BoardWidth int, MySnake api.Snake, Snakes []api.Snake, Destination api.Coord) []Square {
	var closedList = make([]Square, 0)
	var openList = make([]Square, 0)

	myHead := Square{Coord: MySnake.Body[0], G: 0, H: 0, F: 0}
	closedList = append(closedList, myHead)

	// get the adjacent coords
	var adjacentHead = getAdjacentCoords(MySnake.Body[0])

	for i := 0; i < len(adjacentHead); i++ {
		var newSquare = Square{
			Coord:        adjacentHead[i],
			G:            1,
			H:            Manhatten(adjacentHead[i], Destination),
			F:            1 + Manhatten(adjacentHead[i], Destination),
			ParentCoords: MySnake.Body[0],
		}

		openList = appendList(newSquare, Snakes, openList, BoardHeight, BoardWidth)
	}

	// TODO: LOOP THROUGH
	for len(openList) > 0 {
		// find the Square the least F on the open list
		var leastSquare = openList[0]
		for i := 0; i < len(openList); i++ {
			if openList[i].F < leastSquare.F {
				leastSquare = openList[i]
			}
		}

		if leastSquare.Coord == Destination {
			closedList = append(closedList, leastSquare)
			openList = removeFromOpenList(leastSquare, openList)
			break
		}
		// put it on the closed list
		closedList = append(closedList, leastSquare)

		openList = removeFromOpenList(leastSquare, openList)

		// loop through leastSquares's adjacent tiles -- call them T
		var leastSquareAdjacents = getAdjacentCoords(leastSquare.Coord)

		for i := 0; i < len(leastSquareAdjacents); i++ {
			if leastSquareAdjacents[i] == Destination {
				var closedSquare = Square{
					Coord:        leastSquareAdjacents[i],
					G:            leastSquare.G + 1,
					H:            Manhatten(leastSquareAdjacents[i], Destination),
					F:            (leastSquare.G + 1) + Manhatten(leastSquareAdjacents[i], Destination),
					ParentCoords: leastSquare.Coord,
				}

				closedList = appendList(closedSquare, Snakes, closedList, BoardHeight, BoardWidth)

				break
			}
		}

		for i := 0; i < len(leastSquareAdjacents); i++ {
			// 1. If T on the closed list, ignore it
			var onClosedList = false
			for j := 0; j < len(closedList); j++ {
				if leastSquareAdjacents[i] == closedList[j].Coord {
					onClosedList = true
				}
			}

			if onClosedList == true {
				continue
			}
			// 2. If T is not on the open list add it
			var onOpenList = false
			var SquaresOnOpenList = make([]Square, 0)

			for j := 0; j < len(openList); j++ {
				if leastSquareAdjacents[i] == openList[j].Coord {
					onOpenList = true
					SquaresOnOpenList = append(SquaresOnOpenList, openList[j])

					if (leastSquare.G+1)+Manhatten(leastSquareAdjacents[i], Destination) < openList[j].F {
						openList[j].F = (leastSquare.G + 1) + Manhatten(leastSquareAdjacents[i], Destination)
						openList[j].G = leastSquare.G + 1
						openList[j].H = Manhatten(leastSquareAdjacents[i], Destination)
						openList[j].ParentCoords = leastSquareAdjacents[i]
					}
				}
			}

			// 3. if not on the open list, add it
			if onOpenList == false {
				if squareBlocked(leastSquareAdjacents[i], Snakes) == false && onBoard(leastSquareAdjacents[i], BoardHeight, BoardWidth) {
					var openSquare = Square{
						Coord:        leastSquareAdjacents[i],
						G:            leastSquare.G + 1,
						H:            Manhatten(leastSquareAdjacents[i], Destination),
						F:            (leastSquare.G + 1) + (Manhatten(leastSquareAdjacents[i], Destination)),
						ParentCoords: leastSquare.Coord,
					}

					openList = appendList(openSquare, Snakes, openList, BoardHeight, BoardWidth)
				}
			}

		}
	}
	// TODO: END OF LOOP
	return closedList // [1].Coord
}

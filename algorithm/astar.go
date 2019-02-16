package algorithm

import (
	"github.com/alecj1240/astart/api"
)

/*
	G - the amount of steps its taken to get to that square
	H - the heuristic - estimation from this square to the destination
	F - the sum of G and H
	ParentCoords - the coordinates of the previous step (square)
*/

// create struct that will hold: coordinates, G, H, F, parent coords
type square struct {
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

func removeFromOpenList(removalSquare square, openList []square) []square {
	var newOpenList = make([]square, 0)
	for i := 0; i < len(openList); i++ {
		if removalSquare.Coord != openList[i].Coord {
			newOpenList = append(newOpenList, openList[i])
		}
	}
	return newOpenList
}

/*
the a star algorithm
*/
func Astar(BoardHeight int, BoardWidth int, MySnake api.Snake, Snakes []api.Snake, Destination api.Coord) []square {
	var closedList = make([]square, 0)
	var openList = make([]square, 0)

	myHead := square{Coord: MySnake.Body[0], G: 0}
	closedList = append(closedList, myHead)

	// get the adjacent coords
	var adjacentHead = getAdjacentCoords(MySnake.Body[0])

	for i := 0; i < len(adjacentHead); i++ {
		if squareBlocked(adjacentHead[i], Snakes) == false && onBoard(adjacentHead[i], BoardHeight, BoardWidth) == true {
			openList = append(openList, square{
				Coord:        adjacentHead[i],
				G:            1,
				H:            manhatten(adjacentHead[i], Destination),
				F:            1 + manhatten(adjacentHead[i], Destination),
				ParentCoords: MySnake.Body[0],
			})
		}
	}

	// TODO: LOOP THROUGH
	for {

		// find the square the least F on the open list
		var leastSquare = openList[0]
		for i := 0; i < len(openList); i++ {
			if openList[i].F < leastSquare.F {
				leastSquare = openList[i]
			}
		}

		if leastSquare.Coord == Destination {
			break
		}

		// put it on the closed list
		closedList = append(closedList, leastSquare)

		openList = removeFromOpenList(leastSquare, openList)

		// loop through leastSquares's adjacent tiles -- call them T
		var leastSquareAdjacents = getAdjacentCoords(leastSquare.Coord)

		// blocked or off grid
		var checkingAdjacents = make([]api.Coord, 0)
		for i := 0; i < len(leastSquareAdjacents); i++ {
			if squareBlocked(leastSquareAdjacents[i], Snakes) == false && onBoard(leastSquareAdjacents[i], BoardHeight, BoardWidth) == true {
				checkingAdjacents = append(checkingAdjacents, leastSquareAdjacents[i])
			}
		}

		leastSquareAdjacents = checkingAdjacents

		var destinationFound = false
		for i := 0; i < len(leastSquareAdjacents); i++ {
			if leastSquareAdjacents[i] == Destination {
				destinationFound = true
			}
		}

		if destinationFound == true {
			break
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
			var squaresOnOpenList = make([]square, 0)

			for j := 0; j < len(openList); j++ {
				if leastSquareAdjacents[i] == openList[j].Coord {
					onOpenList = true
					squaresOnOpenList = append(squaresOnOpenList, openList[j])

					if (leastSquare.G+1)+manhatten(leastSquareAdjacents[i], Destination) < openList[j].F {
						openList[j].F = (leastSquare.G + 1) + manhatten(leastSquareAdjacents[i], Destination)
						openList[j].G = leastSquare.G + 1
						openList[j].H = manhatten(leastSquareAdjacents[i], Destination)
						openList[j].ParentCoords = leastSquareAdjacents[i]
					}
				}
			}

			// 3. if not on the open list, add it
			if onOpenList == false {
				if squareBlocked(leastSquareAdjacents[i], Snakes) == false && onBoard(leastSquareAdjacents[i], BoardHeight, BoardWidth) == true {
					openList = append(openList, square{
						Coord:        leastSquareAdjacents[i],
						G:            leastSquare.G + 1,
						H:            manhatten(leastSquareAdjacents[i], Destination),
						F:            (leastSquare.G + 1) + (manhatten(leastSquareAdjacents[i], Destination)),
						ParentCoords: leastSquare.Coord,
					})
				}
			}

		}

	}
	// TODO: END OF LOOP

	return closedList // [1].Coord
}

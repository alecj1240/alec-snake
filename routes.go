package main

import (
	"log"
	"net/http"

	"github.com/alecj1240/astart/algorithm"
	"github.com/alecj1240/astart/api"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Battlesnake documentation can be found at <a href=\"https://docs.battlesnake.io\">https://docs.battlesnake.io</a>."))
}

func Start(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}
	dump(decoded)

	respond(res, api.StartResponse{
		Color: "#2ecc71",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	totalSnakeLength := 0
	for _, snake := range decoded.Board.Snakes {
		totalSnakeLength += len(snake.Body)
	}
	averageSnakeLength := totalSnakeLength / len(decoded.Board.Snakes)

	var moveCoord []api.Coord
	// if there is no food chase your tail
	if decoded.You.Health > 30 && ((len(decoded.You.Body) >= averageSnakeLength && len(decoded.You.Body) >= 4) || len(decoded.Board.Food) == 0) {
		moveCoord = algorithm.Astar(decoded.Board.Height, decoded.Board.Width, decoded.You, decoded.Board.Snakes, algorithm.ChaseTail(decoded.You.Body))
	} else {
		moveCoord = algorithm.Astar(decoded.Board.Height, decoded.Board.Width, decoded.You, decoded.Board.Snakes, algorithm.NearestFood(decoded.Board.Food, decoded.You.Body[0]))
		if moveCoord == nil {
			moveCoord = algorithm.Astar(decoded.Board.Height, decoded.Board.Width, decoded.You, decoded.Board.Snakes, algorithm.ChaseTail(decoded.You.Body))
		}
	}

	var finalMove = algorithm.Heading(decoded.You.Body[0], moveCoord[1])

	// if algorithm.HeadOnCollision(moveCoord[1], decoded.Board.Snakes, decoded.You) {
	// 	adjacentHead := algorithm.GetAdjacentCoords(decoded.You.Body[0])

	// 	for i := 0; i < len(adjacentHead); i++ {
	// 		if algorithm.OnBoard(adjacentHead[i], decoded.Board.Height, decoded.Board.Width) && adjacentHead[i] != moveCoord[1] && algorithm.SquareBlocked(decoded.You.Body[0], decoded.Board.Snakes) == false {
	// 			finalMove = algorithm.Heading(decoded.You.Body[0], adjacentHead[i])
	// 		}
	// 	}
	// }

	respond(res, api.MoveResponse{
		Move: finalMove,
	})
}

func End(res http.ResponseWriter, req *http.Request) {
	return
}

func Ping(res http.ResponseWriter, req *http.Request) {
	return
}

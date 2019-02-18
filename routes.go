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
		Color: "#75CEDD",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	var moveCoord []algorithm.Square

	// if there is no food chase your tail
	if len(decoded.Board.Food) < 1 {
		moveCoord = algorithm.Astar(decoded.Board.Height, decoded.Board.Width, decoded.You, decoded.Board.Snakes, decoded.You.Body[len(decoded.You.Body)-1])
	} else {
		moveCoord = algorithm.Astar(decoded.Board.Height, decoded.Board.Width, decoded.You, decoded.Board.Snakes, algorithm.NearestFood(decoded.Board.Food, decoded.You.Body[0]))
	}

	var finalMove string

	// if the food is right next to you, just head to that
	if algorithm.Manhatten(decoded.You.Body[0], algorithm.NearestFood(decoded.Board.Food, decoded.You.Body[0])) == 1 {
		finalMove = algorithm.Heading(decoded.You.Body[0], algorithm.NearestFood(decoded.Board.Food, decoded.You.Body[0]))
	} else {
		finalMove = algorithm.Heading(decoded.You.Body[0], moveCoord[1].Coord)
	}

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

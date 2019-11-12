package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// {
//   game": {
//     "id": "game-id-string"
//   },
//   "turn": 1,
//   "board": {
//     "height": 11,
//     "width": 11,
//     "food": [{
//       "x": 1,
//       "y": 3
//     }],
//     "snakes": [{
//       "id": "snake-id-string",
//       "name": "Sneky Snek",
//       "health": 100,
//       "body": [{
//         "x": 1,
//         "y": 3
//       }]
//     }]
//   },
//   "you": {
//     "id": "snake-id-string",
//     "name": "Sneky Snek",
//     "health": 100,
//     "body": [{
//       "x": 1,
//       "y": 3
//     }]
//   }
// }

type Point struct {
	X uint32 `json:"x"`
	Y uint32 `json:"y"`
}

type Snake struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Health uint32  `json:"health"`
	Body   []Point `json:"body"`
}

type Board struct {
	Height uint32  `json:"height"`
	Width  uint32  `json:"width"`
	Food   []Point `json:"food"`
	Snakes []Snake `json:"snakes"`
}

type Game struct {
	Id string `json:"id"`
}

type Turn struct {
	Game  Game   `json:"game"`
	Turn  uint32 `json:"turn"`
	Board Board  `json:"board"`
	You   Snake  `json:"you"`
}

type Move struct {
	Move string `json:"move"`
}

type SnakeConfig struct {
	Color    string `json:"color"`
	HeadType string `json:"headType"`
	TailType string `json:"tailType"`
}

func HttpError(w http.ResponseWriter, err error) {

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func StartHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	body, berr := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if berr != nil {
		HttpError(w, err)
		return
	}
	if err = r.Body.Close(); err != nil {
		HttpError(w, err)
		return
	}

	var turn Turn
	if err := json.Unmarshal(body, &turn); err != nil {
		HttpError(w, err)
		return
	}

	fmt.Println(turn)

	sc := SnakeConfig{
		Color:    "#ff00ff",
		HeadType: "bendr",
		TailType: "pixel",
	}

	var js []byte
	js, err = json.Marshal(sc)
	if err != nil {
		HttpError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func EndHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	body, berr := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if berr != nil {
		HttpError(w, err)
		return
	}
	if err = r.Body.Close(); err != nil {
		HttpError(w, err)
		return
	}

	var turn Turn
	if err := json.Unmarshal(body, &turn); err != nil {
		HttpError(w, err)
		return
	}

	fmt.Println(turn)
	fmt.Println("Game Over")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Pong")
}

func MoveHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	body, berr := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if berr != nil {
		HttpError(w, err)
		return
	}
	if err = r.Body.Close(); err != nil {
		HttpError(w, err)
		return
	}

	var turn Turn
	if err := json.Unmarshal(body, &turn); err != nil {
		HttpError(w, err)
		return
	}

	fmt.Println(turn)

	move := MakeMove(turn)

	var js []byte
	js, err = json.Marshal(move)
	if err != nil {
		HttpError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

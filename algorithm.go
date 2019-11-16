package main

import (
	"fmt"
)

func (p Point) Equal(p2 Point) bool {

	if p.X == p2.X && p.Y == p2.Y {
		return true
	} else {
		return false
	}
}

func (p Point) ToIndex(w uint32) uint32 {
	return w*p.X + p.Y
}

func MakeBoard(turn Turn) []uint8 {
	// An Internal representation of the game board
	// useful for checking the content on board

	// AIR := 0
	// MYSNAKE := 1
	// FOOD := 2
	// OTHERSNAKE := 3

	w, h := BoardDims(turn)

	board := make([]uint8, w*h)

	for _, p := range turn.Board.Food {
		board[p.ToIndex(w)] = 0
	}

	for _, s := range turn.Board.Snakes {
		for _, p := range s.Body {
			board[p.ToIndex(w)] = 3
		}
	}

	for _, p := range turn.You.Body {
		board[p.ToIndex(w)] = 1
	}

	return board
}

func CurrentPos(turn Turn) (uint32, uint32) {

	x := turn.You.Body[0].X
	y := turn.You.Body[0].Y

	return x, y
}

func CurrentPosInt(turn Turn) (int32, int32) {

	x := turn.You.Body[0].X
	y := turn.You.Body[0].Y

	return int32(x), int32(y)
}

func BoardDims(turn Turn) (uint32, uint32) {

	width := turn.Board.Width
	height := turn.Board.Height

	return width, height
}

func GetNeighbours(turn Turn, point Point) map[string]Point {

	x := point.X
	y := point.Y

	w, h := BoardDims(turn)

	neighbours := make(map[string]Point)

	if x > 0 {
		neighbours["left"] = Point{X: x - 1, Y: y}
	}

	if x < w-1 {
		neighbours["right"] = Point{X: x + 1, Y: y}
	}

	if y > 0 {
		neighbours["up"] = Point{X: x, Y: y - 1}
	}

	if y < h-1 {
		neighbours["down"] = Point{X: x, Y: y + 1}
	}

	return neighbours
}

func CheckIfEdge(turn Turn) bool {

	w, h := BoardDims(turn)
	x, y := CurrentPos(turn)

	if x > 0 && y > 0 && x < w-1 && y < h-1 {
		return false
	} else {
		return true
	}
}

func AvoidEdgeAndCorners(turn Turn) string {

	// if the snake is heading to an edge, avoid it
	// top edge -> take left or right
	// left edge -> take top or bottom
	// right edge -> take top or bottom
	// bottom edge -> take left or right

	x, y := CurrentPos(turn)
	w, h := BoardDims(turn)

	heading := GetCurrentHeading(turn)

	// check for corner cases
	if x <= 0 && y <= 0 {
		// top left corner
		if heading == "up" {
			return "right"
		} else if heading == "left" {
			return "down"
		} else {
			return "right"
		}
	}

	if x <= 0 && y >= h-1 {
		// bottom left corner
		if heading == "down" {
			return "left"
		} else if heading == "left" {
			return "up"
		} else {
			return "left"
		}
	}

	if x >= w-1 && y <= 0 {
		// top right corner
		if heading == "right" {
			return "down"
		} else if heading == "up" {
			return "left"
		} else {
			return "left"
		}
	}

	if x >= w-1 && y >= h-1 {
		// bottom right corner
		if heading == "down" {
			return "left"
		} else if heading == "right" {
			return "up"
		} else {
			return "left"
		}
	}

	// check for edge cases
	// the escape sequence is in clockwise
	if x <= 0 && heading == "left" {
		// left edge
		return "up"
	} else if y <= 0 && heading == "up" {
		// top edge
		return "right"
	} else if x >= w-1 && heading == "right" {
		// right edge
		return "down"
	} else if y >= h-1 && heading == "down" {
		// bottom edge
		return "left"
	}

	// we are in the middle
	return ""
}

func GetCurrentHeading(turn Turn) string {

	// if snake is of size 1, then default empty string
	// need atleast size 2 to

	x1, y1 := CurrentPosInt(turn)

	x2 := int32(turn.You.Body[1].X)
	y2 := int32(turn.You.Body[1].Y)

	xdiff := x1 - x2
	ydiff := y1 - y2

	if xdiff == 0 {
		if ydiff > 0 {
			return "down"
		} else {
			return "up"
		}
	}

	if ydiff == 0 {
		if xdiff > 0 {
			return "right"
		} else {
			return "left"
		}
	}

	return "up"
}

func GetFoodDirection(turn Turn) string {

	if len(turn.Board.Food) == 0 {
		return GetCurrentHeading(turn)
	}

	w, h := BoardDims(turn)

	path := make([]string, w*h)
	for _, f := range turn.Board.Food {
		_path := AStar(turn, turn.You.Body[0], f)

		if len(_path) < len(path) {
			path = _path
		}
	}

	return path[0]

	// fx := path[0].X
	// fy := path[0].Y
	//
	// x, y := CurrentPos(turn)
	//
	// if fx > x {
	// 	return "right"
	// } else if fx < x {
	// 	return "left"
	// } else {
	// 	if fy > y {
	// 		return "down"
	// 	} else {
	// 		return "up"
	// 	}
	// }
}

func Opposite(move string) string {

	switch move {
	case "right":
		return "left"
	case "left":
		return "right"
	case "up":
		return "down"
	case "down":
		return "up"
	}

	return move
}

func CollisionAware(turn Turn, chosenMove string) string {

	// if any heuristics give result without considering the self collision
	// then it is averted here

	board := MakeBoard(turn)
	w, _ := BoardDims(turn)

	nb := GetNeighbours(turn, turn.You.Body[0])
	nextTile := nb[chosenMove]

	if board[nextTile.ToIndex(w)] != 0 {
		fmt.Println("Chosen move is not AIR")

		for k, v := range nb {
			if board[v.ToIndex(w)] == 0 {
				return k
			}
		}
	}

	return chosenMove
}

func MakeMove(turn Turn) Move {

	var move string
	// move = AvoidEdgeAndCorners(turn)
	if move != "" {
		fmt.Println("At the Edge")
		// move = AvoidEdgeAndCorners(turn)
	} else {
		move = CollisionAware(turn, GetFoodDirection(turn))
	}

	fmt.Println("Heading: ", GetCurrentHeading(turn))
	fmt.Println("Chosen move: ", move)
	return Move{Move: move}
}

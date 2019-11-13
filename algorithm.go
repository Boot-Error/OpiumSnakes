package main

import (
	"fmt"
)

func CurrentPos(turn Turn) (uint32, uint32) {

	x := turn.You.Body[0].X
	y := turn.You.Body[0].Y

	return x, y
}

func BoardDims(turn Turn) (uint32, uint32) {

	width := turn.Board.Width
	height := turn.Board.Height

	return width, height
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

	x1, y1 := CurrentPos(turn)

	if len(turn.You.Body) > 1 {

		x2 := turn.You.Body[1].X
		y2 := turn.You.Body[1].Y

		xdiff := x1 - x2
		ydiff := y1 - y2

		if xdiff == 0 {
			if ydiff > 0 {
				return "up"
			} else {
				return "down"
			}
		}

		if ydiff == 0 {
			if xdiff > 0 {
				return "right"
			} else {
				return "left"
			}
		}
	}

	return "up"
}

func GetFoodDirection(turn Turn) string {

	if len(turn.Board.Food) == 0 {
		return GetCurrentHeading(turn)
	}

	fx := turn.Board.Food[0].X
	fy := turn.Board.Food[0].Y

	x, y := CurrentPos(turn)

	if fx > x {
		return "right"
	} else if fx < x {
		return "left"
	} else {
		if fy > y {
			return "down"
		} else {
			return "up"
		}
	}
}

func MakeMove(turn Turn) Move {

	var move string
	if CheckIfEdge(turn) {

		fmt.Println("At the Edge")
		move = AvoidEdgeAndCorners(turn)

	} else {
		move = GetFoodDirection(turn)
	}

	fmt.Println("Heading: ", GetCurrentHeading(turn))
	fmt.Println("Chosen move: ", move)
	return Move{Move: move}
}

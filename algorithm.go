package main

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

	width := turn.Board.Width
	height := turn.Board.Height

	x, y := CurrentPos(turn)

	if x > 0 && y > 0 && x < width && y < height {
		return false
	} else {
		return true
	}
}

func AvoidEdge(turn Turn) string {

	// if the snake is heading to an edge, avoid it
	// top edge -> take left or right
	// left edge -> take top or bottom
	// right edge -> take top or bottom
	// bottom edge -> take left or right

	x, y := CurrentPos(turn)
	w, h := BoardDims(turn)

	if x == 0 {
		return "right"
	} else if y == 0 {
		return "up"
	} else if x == w-1 {
		return "down"
	} else if y == h-1 {
		return "left"
	}

	return "up"
}

func GetCurrentHeading(turn Turn) string {

	// if snake is of size 1, then default empty string
	// need atleast size 2 to

	x1, y1 := CurrentPos(turn)

	if len(turn.You.Body) == 1 {
		return ""
	} else {

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
				return "left"
			} else {
				return "right"
			}
		}
	}

	return ""
}

func MakeMove(turn Turn) Move {

	var move string
	if CheckIfEdge(turn) {
		move = AvoidEdge(turn)
	} else {
		move = "up"
	}

	return Move{Move: move}
}

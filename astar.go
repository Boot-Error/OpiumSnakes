package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Item struct {
	Point    Point
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[i].Priority
}

func (pq PriorityQueue) Swap(i, j int) {

	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {

	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {

	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, point Point, priority int) {

	item.Point = point
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func (pq PriorityQueue) membership(point Point) *Item {

	for _, i := range pq {
		if point.X == i.Point.X && point.Y == i.Point.Y {
			return i
		}
	}
	return nil
}

func PointToItem(point Point, priority int) *Item {
	return &Item{
		Point:    point,
		Priority: priority,
		Index:    0,
	}
}

func (p Point) ManhattanDistance(other Point) int {

	return int(math.Abs(float64(p.X-other.X)) + math.Abs(float64(p.Y-other.Y)))
}

type GridItem struct {
	TileType uint8
	Gcost    int
	Hcost    int
}

func MakeGrid(board []uint8) []GridItem {

	grid := make([]GridItem, len(board))
	for _, i := range board {
		grid[i] = GridItem{
			TileType: board[i],
			Gcost:    0,
			Hcost:    0,
		}
	}

	return grid
}

// General A* Implementation
func AStar(turn Turn, start Point, goal Point) []string {

	// OPEN = priority queue containing START
	// CLOSED = empty set
	// while lowest rank in OPEN is not the GOAL:
	//   current = remove lowest rank item from OPEN
	//   add current to CLOSED
	//   for neighbors of current:
	//     cost = g(current) + movementcost(current, neighbor)
	//     if neighbor in OPEN and cost less than g(neighbor):
	//       remove neighbor from OPEN, because new path is better
	//     if neighbor in CLOSED and cost less than g(neighbor): ⁽²⁾
	//       remove neighbor from CLOSED
	//     if neighbor not in OPEN and neighbor not in CLOSED:
	//       set g(neighbor) to cost
	//       add neighbor to OPEN
	//       set priority queue rank to g(neighbor) + h(neighbor)
	//       set neighbor's parent to current
	//
	// reconstruct reverse path from goal to start
	// by following parent pointers

	var path []string

	w, h := BoardDims(turn)

	board := MakeBoard(turn)
	grid := MakeGrid(board)

	OPEN := make(PriorityQueue, len(board))
	OPEN.Push(PointToItem(start, 0))

	grid[start.ToIndex(w)].Hcost = start.ManhattanDistance(goal)

	CLOSED := make([]bool, w*h)

	for len(OPEN) > 0 {

		current := OPEN.Pop().(*Item)

		if current.Point.Equal(goal) {
			return path
		}

		CLOSED[current.Point.ToIndex(w)] = true

		for d, v := range GetNeighbours(turn, current.Point) {

			// Gcost := grid[current.Point.ToIndex(w)].Gcost + 1
			Hcost := v.ManhattanDistance(goal)

			if !CLOSED[v.ToIndex(w)] {
				OPEN.Push(PointToItem(v, Hcost))

				CLOSED[v.ToIndex(w)] = true
				path = append(path, d)
			}
		}
	}

	fmt.Println("A*")
	return path
}

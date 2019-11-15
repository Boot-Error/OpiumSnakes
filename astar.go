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

func PointToItem(point Point, priority int) Item {
	return Item{
		Point:    point,
		Priority: priority,
		Index:    0,
	}
}

func (p Point) ManhattanDistance(other Point) int {

	return int(math.Abs(float64(p.X-other.X)) + math.Abs(float64(p.Y-other.Y)))
}

func CostFunction(start Point, current Point, goal Point) int {

	G := start.ManhattanDistance(current)
	H := current.ManhattanDistance(goal)

	return G + H
}

// General A* Implementation
func AStar(board []Point, start Point, goal Point) []Point {

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

	var path []Point

	frontier := make(PriorityQueue, len(board))
	frontier.Push(PointToItem(start, 0))

	for len(frontier) > 0 {

		current := frontier.Pop().(*Item)

		if current.Point.Equal(goal) {
			return path
		}

		// for k, v := range GetNeighbours(current.Point) {
		//
		// 	fmt.Println(k, v)
		// }
	}

	fmt.Println("A*")
	return []Point{Point{X: 2, Y: 3}}
}

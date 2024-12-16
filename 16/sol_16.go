package main

import (
	"aoc/parseutil"
	bh "github.com/emirpasic/gods/trees/binaryheap"
	"log"
	"slices"
)

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

const Penalty = 1000

type Point [2]int

var Moves = map[Direction]Point{
	N: {-1, 0},
	E: {0, 1},
	S: {1, 0},
	W: {0, -1},
}

type State struct {
	Pos  Point
	Dir  Direction
	Cost int
}

func (st *State) Move(d Direction) State {
	dCost := int(min((st.Dir-d+4)%4, (d-st.Dir+4)%4)*Penalty + 1)
	return State{
		Cost: st.Cost + dCost,
		Dir:  d,
		Pos:  Point{st.Pos[0] + Moves[d][0], st.Pos[1] + Moves[d][1]},
	}
}

func (st *State) Key() [3]int {
	return [3]int{st.Pos[0], st.Pos[1], int(st.Dir)}
}

func main() {
	cmp := func(a, b interface{}) int {
		return a.(State).Cost - b.(State).Cost
	}

	var start, end Point
	board := [][]int{}
	minPq := bh.NewWith(cmp)
	seen := map[[3]int]int{}

	for _, line := range parseutil.ReadInputLines() {
		row := []int{}
		for _, token := range line {
			var t int
			switch token {
			case 'S':
				start = Point{len(board), len(row)}
			case 'E':
				end = Point{len(board), len(row)}
			case '#':
				t = 1
			}
			row = append(row, t)
		}
		board = append(board, row)
	}

	part1 := func() (int, map[[3]int][][3]int) {
		minPq.Push(State{Pos: start, Dir: E, Cost: 0})
		seen[[3]int{start[0], start[1], int(E)}] = 0
		from := map[[3]int][][3]int{}
		best := 1 << 31
		for !minPq.Empty() {
			x, _ := minPq.Pop()
			cur := x.(State)
			if seen[cur.Key()] < cur.Cost || cur.Pos == end && cur.Cost > best {
				continue
			}
			if cur.Pos == end {
				best = cur.Cost
				continue
			}
			for d := range 4 {
				nS := cur.Move(Direction(d))
				if cost, ok := seen[nS.Key()]; board[nS.Pos[0]][nS.Pos[1]] == 0 && (!ok || cost >= nS.Cost) {
					if cost > nS.Cost {
						from[nS.Key()] = from[nS.Key()][:0]
					} else if slices.Contains(from[nS.Key()], cur.Key()) {
						continue
					}
					minPq.Push(nS)
					seen[nS.Key()] = nS.Cost
					from[nS.Key()] = append(from[nS.Key()], cur.Key())
				}
			}
		}
		// makes it simpler to initiate path trace in part 2
		for s, c := range seen {
			if c > best {
				delete(from, s)
			}
		}
		return best, from
	}

	part2 := func(from map[[3]int][][3]int) int {
		onPath := map[Point]bool{}
		q := [][3]int{}
		for i := range 4 {
			if src := from[[3]int{end[0], end[1], i}]; len(src) > 0 {
				q = append(q, src...)
				onPath[end] = true
			}
		}
		for len(q) > 0 {
			next := [][3]int{}
			for len(q) > 0 {
				cur := q[len(q)-1]
				onPath[Point{cur[0], cur[1]}] = true
				next = append(next, from[cur]...)
				q = q[:len(q)-1]
			}
			q = next
		}
		return len(onPath)
	}

	tot, from := part1()
	log.Println("Part 1", tot)
	log.Println("Part 2", part2(from))
}

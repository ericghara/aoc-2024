package main

import (
	"aoc/parseutil"
	bh "github.com/emirpasic/gods/trees/binaryheap"
	"log"
	"strconv"
)

type (
	Stroke [2]rune
	Point  [2]int

	PadState struct {
		Memo   map[Stroke]int
		Pad    [][]rune
		Coords map[rune]Point
	}

	DjState struct {
		Pos  Point
		Cost int
		Last rune
		Done bool
	}
)

var (
	Moves = map[rune]Point{'v': {1, 0}, '<': {0, -1}, '>': {0, 1}, '^': {-1, 0}}

	NumPad = [][]rune{{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{0, '0', 'A'}}

	DirPad = [][]rune{{0, '^', 'A'},
		{'<', 'v', '>'}}
)

func (dj *DjState) Key() [3]int {
	return [3]int{dj.Pos[0], dj.Pos[1], int(dj.Last)}
}

func djComp(a, b interface{}) int {
	return a.(DjState).Cost - b.(DjState).Cost
}

func validMove(move Point, pad [][]rune) bool {
	return move[0] >= 0 && move[0] < len(pad) && move[1] >= 0 &&
		move[1] < len(pad[0]) && pad[move[0]][move[1]] != 0
}

func recurse(stroke Stroke, states []PadState) int {
	if len(states) == 0 {
		return 1
	}
	state := &states[0]
	if cost, ok := state.Memo[stroke]; ok {
		return cost
	}
	start := DjState{
		Pos:  state.Coords[stroke[0]],
		Last: 'A',
	}
	end := state.Coords[stroke[1]]
	costs := map[[3]int]int{start.Key(): 0}
	q := bh.NewWith(djComp)
	q.Push(start)
	for cur, ok := q.Pop(); ok; cur, ok = q.Pop() {
		curState := cur.(DjState)
		if curState.Done {
			state.Memo[stroke] = curState.Cost
			break
		}
		if curState.Pos == end {
			q.Push(DjState{
				Pos:  end,
				Cost: curState.Cost + recurse(Stroke{curState.Last, 'A'}, states[1:]),
				Last: 'A',
				Done: true,
			})
			continue
		}
		for button, move := range Moves {
			nPos := Point{move[0] + curState.Pos[0], move[1] + curState.Pos[1]}
			if validMove(nPos, state.Pad) {
				nCost := curState.Cost + recurse(Stroke{curState.Last, button}, states[1:])
				nState := DjState{
					Pos:  nPos,
					Cost: nCost,
					Last: button,
				}
				if oldCost, ok := costs[nState.Key()]; !ok || oldCost > nCost {
					q.Push(nState)
					costs[nState.Key()] = nCost
				}
			}
		}
	}
	return state.Memo[stroke]
}

func main() {
	codes := map[[4]rune]int{}

	for _, line := range parseutil.ReadInputLines() {
		numStr := line[0:3]
		num, err := strconv.ParseInt(line[0:3], 10, 64)
		if err != nil {
			log.Panic("unable to parse", numStr)
		}
		tokens := [4]rune{}
		for i, r := range line {
			tokens[i] = r
		}
		codes[tokens] = int(num)
	}

	getCoords := func(pad [][]rune) map[rune]Point {
		coords := map[rune]Point{}
		for r := range len(pad) {
			for c := range len(pad[r]) {
				coords[pad[r][c]] = Point{r, c}
			}
		}
		return coords
	}

	initStates := func(num int) []PadState {
		stack := make([]PadState, num)
		stack[0] = PadState{
			Memo:   map[Stroke]int{},
			Pad:    NumPad,
			Coords: getCoords(NumPad),
		}
		for i := 1; i < len(stack); i++ {
			stack[i] = PadState{
				Memo:   map[Stroke]int{},
				Pad:    DirPad,
				Coords: getCoords(DirPad),
			}
		}
		return stack
	}

	var p1, p2 int
	states1, states2 := initStates(3), initStates(26)

	for code, factor := range codes {
		last := 'A'
		for _, c := range code {
			p1 += recurse(Stroke{last, c}, states1) * factor
			p2 += recurse(Stroke{last, c}, states2) * factor
			last = c
		}
	}
	log.Println("Part1", p1)
	log.Println("Part2", p2)
}

package main

import (
	"aoc/parseutil"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point [2]int

type UF struct {
	weights, edgeTo []int
	dim             int
}

func (uf *UF) encode(p Point) int {
	return p[0]*(uf.dim+1) + p[1]
}

func (uf *UF) Find(i int) int {
	if uf.edgeTo[i] == i {
		return i
	}
	uf.edgeTo[i] = uf.Find(uf.edgeTo[i])
	return uf.edgeTo[i]
}

func (uf *UF) Union(a, b Point) bool {
	aI, bI := uf.Find(uf.encode(a)), uf.Find(uf.encode(b))
	if aI == bI {
		return false
	}
	if uf.weights[aI] < uf.weights[bI] {
		uf.edgeTo[aI] = bI
	}
	if uf.weights[aI] >= uf.weights[bI] {
		uf.edgeTo[bI] = aI
		if uf.weights[aI] == uf.weights[bI] {
			uf.weights[aI]++
		}
	}
	return true
}

func NewUF(dim int) UF {
	uf := UF{
		weights: make([]int, (dim+1)*(dim+1)),
		edgeTo:  make([]int, (dim+1)*(dim+1)),
		dim:     dim,
	}
	for i := 0; i < len(uf.edgeTo); i++ {
		uf.edgeTo[i] = i
	}
	return uf
}

func main() {
	moves := [4]Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	if len(os.Args) != 4 {
		log.Panic("Usage: input file, dimensions, numBytes (pt1)")
	}
	dim, _ := strconv.ParseInt(os.Args[2], 10, 64)
	numBytes, _ := strconv.ParseInt(os.Args[3], 10, 64)
	addrs := []Point{}
	board := [][]int{}

	for _, line := range parseutil.ReadInputLines() {
		split := strings.Split(line, ",")
		a, errA := strconv.ParseInt(split[0], 10, 64)
		b, errB := strconv.ParseInt(split[1], 10, 64)
		if errA != nil || errB != nil {
			log.Panic("unable to parse", line)
		}
		addrs = append(addrs, Point{int(b), int(a)})
	}

	initBoard := func(numBytes int) {
		board = [][]int{}
		for r := 0; r <= int(dim); r++ {
			board = append(board, make([]int, int(dim)+1))
		}

		for _, a := range addrs[:numBytes] {
			board[a[0]][a[1]] = -1
		}
	}

	inBounds := func(p Point) bool {
		return p[0] >= 0 && p[0] < len(board) &&
			p[1] >= 0 && p[1] < len(board[p[0]])
	}

	part1 := func(p Point) (int, map[Point]bool) {
		q := []Point{p}
		seen := map[Point]bool{p: true}
		board[0][0] = 1
		step := 1 // step offset by 1
		for len(q) > 0 {
			next := []Point{}
			step++
			for _, cur := range q {
				for _, move := range moves {
					nMove := Point{move[0] + cur[0], move[1] + cur[1]}
					if inBounds(nMove) && board[nMove[0]][nMove[1]] == 0 {
						board[nMove[0]][nMove[1]] = step
						next = append(next, nMove)
						seen[nMove] = true
					}
				}
			}
			q = next
		}
		return board[dim][dim] - 1, seen
	}

	part2 := func() string {
		uf := NewUF(int(dim))
		var last Point

		for r := range int(dim) + 1 {
			for c := range int(dim) + 1 {
				if board[r][c] == 0 {
					_, seen := part1(Point{r, c})
					pts := []Point{}
					pts = slices.AppendSeq(pts, maps.Keys(seen))
					for i := 1; i < len(pts); i++ {
						uf.Union(pts[0], pts[i])
					}
				}
			}
		}

		for i := len(addrs) - 1; uf.Find(0) != uf.Find(int((dim+1)*(dim+1)-1)); i-- {
			pt := addrs[i]
			for _, move := range moves {
				nPt := Point{pt[0] + move[0], pt[1] + move[1]}
				if inBounds(nPt) {
					uf.Union(pt, nPt)
				}
			}
			last = pt
		}
		return fmt.Sprintf("%d,%d", last[1], last[0])
	}

	initBoard(int(numBytes))
	steps, _ := part1(Point{0, 0})
	log.Println("Part 1:", steps-1)

	initBoard(len(addrs))
	log.Println("Part 2:", part2())
}

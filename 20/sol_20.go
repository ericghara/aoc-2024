package main

import (
    "aoc/parseutil"
    "log"
    "math"
)

func main() {
    type Point [2]int
    moves := [4]Point{{0,1},{0,-1},{1,0},{-1,0}}

    board := [][]int{}
    var start, end Point

    for _, line := range parseutil.ReadInputLines() {
        row := make([]int, len(line))
        for c, t := range line {
            switch t {
            case '#': row[c] = -1
            case 'S': 
                start = Point{len(board), c}
                row[c] = math.MaxInt64
            case 'E': 
                end = Point{len(board), c}
                row[c] = math.MaxInt64
            default:
                row[c] = math.MaxInt64
            }
        }
        board = append(board, row)
        //log.Println(row)
    }

    inBounds := func(p Point) bool {
        return p[0] >= 0 && p[0] < len(board) && 
            p[1] >= 0 && p[1] < len(board[p[0]])
    }

    bfs := func() {
        q := []Point{start}
        board[start[0]][start[1]] = 0
        step := 0
        for len(q) > 0 {
            nextQ := []Point{}
            step++
            for _, cur := range q {
                for _, move := range moves {
                    nMove := Point{move[0]+cur[0], move[1]+cur[1]}
                    if board[nMove[0]][nMove[1]] == math.MaxInt64 {
                        board[nMove[0]][nMove[1]] = step
                        nextQ = append(nextQ, nMove)
                    }
                }
            }
            q = nextQ
        }
    }

    optimize := func(minSave int) int {
        numSave := 0

        for r := range(len(board)) {
            for c := range(len(board[r])) {
                if board[r][c] != -1 {
                    continue
                }
                lo, hi := math.MaxInt64, math.MinInt64
                for _, move := range moves {
                    nMove := Point{move[0]+r, move[1]+c}
                    if !inBounds(nMove) {
                        continue
                    }
                    if val := board[nMove[0]][nMove[1]]; val >= 0 && val <= board[end[0]][end[1]] {
                        lo = min(lo, val)
                        hi = max(hi, val)
                    }
                }
                if savings := hi-lo-2; savings >= minSave {
                    numSave++
                    //log.Println(r, c)
                }
            }
            log.Println(board[r])
        }
        return numSave
    }

    bfs()
    log.Println(end, board[end[0]][end[1]])
    log.Println(start, board[start[0]][start[1]])
    log.Println("Part 1:", optimize(100))
}

package main

import (
    "aoc/parseutil"
    "log"
    "os"
    "strings"
    "strconv"
)

func main() {
    type Point [2]int
    moves := [4]Point{{0,1},{1,0},{0,-1},{-1,0}}

    if len(os.Args) != 4 {
        log.Panic("Usage: input file, dimensions")
    }
    dim, _ := strconv.ParseInt(os.Args[2], 10, 64)
    numBytes, _ := strconv.ParseInt(os.Args[3], 10, 64)
    addrs := []Point{}

    for _, line := range parseutil.ReadInputLines() {
        split := strings.Split(line, ",")
        a, errA := strconv.ParseInt(split[0], 10, 64)
        b, errB := strconv.ParseInt(split[1], 10, 64)
        if errA != nil || errB != nil {
            log.Panic("unable to parse", line)
        }
        addrs = append(addrs, Point{int(b),int(a)})
    }

    board := [][]int{}
    for r := 0; r <= int(dim); r++ {
        board = append(board, make([]int, int(dim)+1))
    }
    
    log.Println(addrs)
    for _, a := range addrs[:numBytes] {
        board[a[0]][a[1]] = -1
    }

    inBounds := func(p Point) bool {
        return p[0] >= 0 && p[0] < len(board) && 
            p[1] >= 0 && p[1] < len(board[p[0]])
    }

    part1 := func() int {
        q := []Point{{0,0}}
        board[0][0] = 1
        step := 1 // step offset by 1
        for len(q) > 0 {
            next := []Point{}
            step++
            for _, cur := range q {
                for _, move := range moves {
                    nMove := Point{move[0]+cur[0], move[1]+cur[1]}
                    if inBounds(nMove) && board[nMove[0]][nMove[1]] == 0 {
                        board[nMove[0]][nMove[1]] = step
                        next = append(next, nMove)
                    }
                }
            }
            q = next
        }
        return board[dim][dim]-1
    }

    log.Println("Part 1:", part1())
//    for _, row := range board {
//        log.Println(row)
//    }


}

package main


import (
    "log"
    "aoc/parseutil"
    "fmt"
)

func main() {

    type Point [2]int

    moveDecoder := map[rune]Point{'v': {1,0}, '^': {-1,0}, '<': {0,-1}, '>': {0,1}}
    sections := parseutil.ReadInputSections(`^\s*$`)
    var start Point

    board := [][]rune{}
    cmds := []Point{} 

    printBoard := func() {
        for _, row := range board {
            for _, c := range row {
                fmt.Print(string(c));
            }
            fmt.Println()
        }
    }

    tryMove := func(cur, move Point) Point {
        next := Point{cur[0]+move[0], cur[1]+move[1]}
        if board[next[0]][next[1]] != '.' {
            end := next
            for board[end[0]][end[1]] == 'O' {
                end[0] += move[0]
                end[1] += move[1]
            }
            if board[end[0]][end[1]] == '#' {
                return cur
            }
            board[next[0]][next[1]] = '.'
            board[end[0]][end[1]] = 'O'
        }
        board[next[0]][next[1]] = '@'
        board[cur[0]][cur[1]] = '.'
        return next
    }

    score := func() int {
        var tot int
        for r, row := range board {
            for c, token := range row {
                if token == 'O' {
                    tot += r*100+c
                }
            }
        }
        return tot
    }

    for _, line := range sections[0] {
        row := []rune{}
        for _, token := range line {
            if token == '@' {
                start = Point{len(board), len(row)}
            }
            row = append(row, token)
        }
        board = append(board, row)
    }

    for _, section := range sections[1:] {
        for _, line := range section {
            for _, token := range line {
                cmds = append(cmds, moveDecoder[token])
            }
        }
    }
    
    cur := start
    for _, move := range cmds {
        cur = tryMove(cur, move)
    }



    log.Println(start)
    printBoard()
    log.Println(cmds)
    log.Println("part 1: ", score())
}

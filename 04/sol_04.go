package main

import (
    "os"
    "log"
    "strings"
)

func main() {
    fileName := "./input"
    if len(os.Args) > 1 {
        fileName = os.Args[1]
    }

    b, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal("unable to open: ", fileName)
    }

    board := make([][]rune, 0)
    xs := make([][]int, 0)
    for _, strs := range strings.Split(string(b), "\n") {
        row := make([]rune, 0)
        for _, c := range strs {
            row = append(row, c)
        }
        board = append(board, row)
        xs = append(xs, make([]int, len(row)))
    }

    moves := [][]int{{1,0},{1,1},{0,1},{-1,1},{-1,0},{-1,-1},{0,-1},{1,-1}}
    target := []rune{'X', 'M','A','S'}

    xmas := func(r,c,dr,dc int) bool {
        for i := range len(target) {
            if r < 0 || r >= len(board) || c < 0 || c >= len(board[r]) || target[i] != board[r][c] {
                return false
            } 
            r += dr
            c += dc
        }
        return true
    }

    mas := func(r,c,dr,dc int) bool {
        for i := range len(target)-1 {
            if r < 0 || r >= len(board) || c < 0 || c >= len(board[r]) || target[i+1] != board[r][c] {
                return false
            }
            r += dr
            c += dc
        }
        return true
    }
    
    var count int64 
    var xcount int64

    for r := range len(board) {
        for c := range len(board[r]) {
            for _, dir := range moves {
                if xmas(r,c,dir[0],dir[1]) {
                    count++
                }
                if dir[0] != 0 && dir[1] != 0 && mas(r,c,dir[0],dir[1]) {
                    if xs[r+dir[0]][c+dir[1]]++; xs[r+dir[0]][c+dir[1]] == 2 {
                        xcount++
                    }
                }
            }
        }
    } 
    log.Println("Count:", count)
    log.Println("XCount:", xcount)
}

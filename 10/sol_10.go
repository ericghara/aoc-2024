package main

import (
    "aoc/parseutil"
    "log"
)

func main() {
    
    type Point [2]int
    lines := parseutil.ReadInputLines()
    
    board := make([][]int, 0)
    score := make([][][]bool, 0)
    topos := make([][]Point, 10)

    for r, l := range lines {
        row := make([]int, 0)
        for c, v := range l {
            elv := int(v)-'0'
            row = append(row, elv)
            topos[elv] = append(topos[elv], Point{r,c})
        }
        board = append(board, row)
        log.Println(row)
        score = append(score, make([][]bool, len(row)))
    }

    findTrails := func() {
        moves := [4][2]int{{0,1},{0,-1},{1,0},{-1,0}} 

        for r := range len(score) {
            for c := range len(score[r]) {
                score[r][c] = make([]bool, len(topos[9]))
            }
        }

        for id, p := range topos[9] {
            score[p[0]][p[1]][id] = true
        }

        for elv := 8; elv >= 0; elv-- {
            for _, pt := range topos[elv] {
                for _, move := range moves {
                    nR, nC := pt[0]+move[0], pt[1]+move[1]
                    if nR >= 0 && nR < len(board) && nC >= 0 && nC < len(board[nR]) && 
                        board[nR][nC] == elv+1 {
                        for id := range len(score[nR][nC]) {
                            score[pt[0]][pt[1]][id] = score[pt[0]][pt[1]][id] || score[nR][nC][id]
                        }
                    }
                }
            }
        }
    }

    findTrails()
    numTrail := 0

    log.Println()
    for _, row := range score {
        log.Println(row)
    }

    for _, pt := range topos[0] {
        for _, reachable := range score[pt[0]][pt[1]] {
            if reachable {
                numTrail++
            }
        }
    }

    log.Println("Num Trails (pt1)", numTrail)

}

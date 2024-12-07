package main

import (
    "os"
    "bufio"
    "log"
    "strings"
)

func main() {
    fileName := "./input"
    if len(os.Args) > 1 {
        fileName = os.Args[1]
    }

    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal("unable to open: ", fileName)
    }
    defer file.Close()

    headings := [4][2]int{{-1,0},{0,1},{1,0},{0,-1}}
    board := make([][]int, 0)

    var start [2]int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        row := make([]int, 0)
        for _, c := range strings.TrimSpace(scanner.Text()) {
            var token int
            if c == '^' {
                start = [2]int{len(board), len(row)}
            } else if c == '#' {
                token = -1
            }
            row = append(row, token)
        }
        board = append(board, row)
    }

    canVisit := func() int {
        visited := map[[2]int]bool{start: true}
        heading := 0
        states := map[[3]int]bool{{heading, start[0], start[1]}: true}
        pos := start
        for true {
            nR, nC := pos[0]+headings[heading][0], pos[1]+headings[heading][1]
            if nR < 0 || nR >= len(board) || nC < 0 || nC >= len(board) {
                break
            }
            if board[nR][nC] == -1 {
                heading = (heading+1)%len(headings)
                continue
            }
            pos[0], pos[1] = nR, nC
            state := [3]int{heading, pos[0], pos[1]}
            if states[state] {
                return -1
            }
            states[state] = true
            visited[pos] = true
        }
        return len(visited) 
    }

    numObs := 0
    numVisit := 0 

    for r := range len(board) {
        for c := range len(board) {
            if [2]int{r,c} == start {
                numVisit = canVisit()
            } else if board[r][c] != -1 {
                board[r][c] = -1
                if canVisit() == -1 {
                    numObs++
                }
                board[r][c] = 0
            }
        }
    }
    log.Println("Num Visited:", numVisit)
    log.Println("Num Obs", numObs)
}

package main


import (
    "log"
    "aoc/parseutil"
    "slices"
    "maps"
)

func main() {

    type Point [2]int

    moveDecoder := map[rune]Point{'v': {1,0}, '^': {-1,0}, '<': {0,-1}, '>': {0,1}}
    sections := parseutil.ReadInputSections(`^\s*$`)
    var cur Point

    board := [][]rune{}
    cmds := []Point{} 

    hMove := func(cur Point, dir int) Point {
        next := Point{cur[0], cur[1]+dir}
        endC := next[1]
        row := board[next[0]]
        for row[endC] != '#' && row[endC] != '.' {
            endC += dir
        }
        if row[endC] == '#' {
            return cur
        }
        for r := endC; r != cur[1]; r-=dir {
            row[r] = row[r-dir]
        }
        row[cur[1]] = '.'
        board[cur[0]] = row
        return next
    }

    var vTest func(cur Point, dir int, seen map[Point]int, clock int) bool
        
    vTest = func(cur Point, dir int, seen map[Point]int, clock int) bool {
        if seen[cur] > 0 || board[cur[0]][cur[1]] == '.' {
            return true
        } else if board[cur[0]][cur[1]] == '#' {
            return false
        }
        seen[cur] = clock
        clock++
        res := true
        if board[cur[0]][cur[1]] == '[' {
            res = vTest(Point{cur[0], cur[1]+1}, dir, seen, clock)
        } else if board[cur[0]][cur[1]] == ']'{
             res = vTest(Point{cur[0], cur[1]-1}, dir, seen, clock)
        }
        return res && vTest(Point{cur[0]+dir, cur[1]}, dir, seen, clock+1)
    }

    vMove := func(seen map[Point]int, dir int) {
        cmp := func(a,b Point) int {
            return seen[b]-seen[a]
        }
        points := slices.Collect(maps.Keys(seen))
        slices.SortFunc(points, cmp)
        for _, p := range points {
            board[p[0]+dir][p[1]] = board[p[0]][p[1]]
            board[p[0]][p[1]] = '.'
        }
    }

    score := func() int {
        var tot int
        for r, row := range board {
            for c, token := range row {
                if token == '[' || token == 'O'{
                    tot += r*100+c
                }
            }
        }
        return tot
    }

    simulate := func() {
        for _, move := range cmds {
            if move[1] != 0 {
                cur = hMove(cur, move[1])
            } else if seen := map[Point]int{}; vTest(cur, move[0], seen, 1) {
                vMove(seen, move[0])
                cur[0] += move[0]
            }
        }
    }

 

    part2 := func() {
        board = [][]rune{}
        for _, line := range sections[0] {
            row := []rune{}
            for _, token := range line {
                stride := []rune{'.','.'}
                if token == '@' {
                    cur = Point{len(board), len(row)}
                    stride[0] = '@'
                } else if token == 'O' {
                    stride = []rune{'[',']'}
                } else if token == '#' {
                    stride = []rune{'#','#'}
                }
                row = append(row, stride[0], stride[1])
            }
            board = append(board, row)
        }
        simulate()
    }

    part1 := func() {
        board = [][]rune{}
        for _, line := range sections[0] {
            row := []rune{}
            for _, token := range line {
                if token == '@' {
                    cur = Point{len(board), len(row)}
                }
                row = append(row, token)
            }
            board = append(board, row)
        }
        simulate()
    }

   for _, section := range sections[1:] {
        for _, line := range section {
            for _, token := range line {
                cmds = append(cmds, moveDecoder[token])
            }
        }
    }
    
    part1()
    log.Println("part 1:", score())
    part2()
    log.Println("part 2:", score())
}

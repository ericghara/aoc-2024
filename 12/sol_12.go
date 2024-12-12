package main

import (
    "log"
    "aoc/parseutil"
)

func main() {
    
    input := parseutil.ReadInputLines()
    garden := make([][]rune, len(input))
    seen := make([][]int, len(input))

    for _, line := range input {
        garden = append(garden, parseutil.ToRunes(line))
        seen = append(seen, make([]int, len(garden[len(seen)])))
    }

    var measure func(int, int, int, rune, int) (p int, a int);

    moves := [4][2]int{{0,1},{1,0},{0,-1},{-1,0}}
    boundaries := []map[[2]int]int{{},{},{},{}} // by direction

    measure = func(r, c, d int, plant rune, id int) (int,int) {
        if r < 0 || r >= len(garden) || c < 0 || c >= len(garden[r]) {
            boundaries[d][[2]int{r,c}] = id
            return 1, 0
        }
        if seen[r][c] == id {
            return 0, 0
        }
        if seen[r][c] != 0 || garden[r][c] != plant {
            boundaries[d][[2]int{r,c}] = id
            return 1, 0
        }
        seen[r][c] = id
        p, a := 0, 1
        for nD, move := range moves {
            dP, dA := measure(r+move[0], c+move[1], nD, plant, id);
            p += dP
            a += dA
        }
        return p, a
    }

    countSides := func() int {
        sides := 0
        for i := range len(boundaries) {
            for k, id := range boundaries[i] {
                sides++
                dir, start := moves[(i+1)%4], k
                for boundaries[i][start] == id {
                    start[0]+=dir[0]
                    start[1]+=dir[1]
                }
                dir[0] *= -1
                dir[1] *= -1
                start[0] += dir[0]
                start[1] += dir[1]
                for boundaries[i][start] == id {
                    delete(boundaries[i], start)
                    start[0] += dir[0]
                    start[1] += dir[1]
                }
            }
        }
        return sides
    }

    fullCost, discountCost := 0, 0

    for r := range len(garden) {
        for c := range len(garden[r]) {
            if seen[r][c] != 0 {
                continue
            }
            id := r*len(garden[r])+c+1
            plant := garden[r][c]
            p, a := measure(r,c,-1,plant, id)
            fullCost += p * a
            s := countSides()
            discountCost += s * a
        }
    }

    log.Println("Full cost:", fullCost)
    log.Println("Discount cost:", discountCost)
}

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

    var measure func(int, int, rune, int) (p int, a int);

    moves := [4][2]int{{0,1},{0,-1},{1,0},{-1,0}}

    measure = func(r, c int, plant rune, id int) (int,int) {
        if r < 0 || r >= len(garden) || c < 0 || c >= len(garden[r]) {
            return 1, 0
        }
        if seen[r][c] == id {
            return 0, 0
        }
        if seen[r][c] != 0 || garden[r][c] != plant {
            return 1, 0
        }
        seen[r][c] = id
        p, a := 0, 1
        for _, move := range moves {
            dP, dA := measure(r+move[0], c+move[1], plant, id);
            p += dP
            a += dA
        }
        return p, a
    }

    cost := 0

    for r := range len(garden) {
        for c := range len(garden[r]) {
            if seen[r][c] != 0 {
                continue
            }
            id := r*len(garden[r])+c
            plant := garden[r][c]
            p, a := measure(r,c,plant, id)
            cost += p * a
        }
    }

    log.Println("Fence cost: ", cost)
}

package main

import (
    "aoc/parseutil"
    "log"
    "regexp"
    "strconv"
    "math"
)


type Point [2]int


func main() {

    sections := parseutil.ReadInputSections(`^\s*$`)
    matcher := regexp.MustCompile(`[^:]*:[^\d]*(\d+)[^\d]*(\d*)`)
    costA, costB := 3, 1

    movesTo := func(dist, button Point) (int, bool) {
        a,b := dist[0]%button[0], dist[1]%button[1]
        if a != 0 || b != 0 {
            return math.MaxInt64, false
        }
        return dist[0]/button[0], true
    }


    solve2 := func(target, a, b Point) int {
        best := math.MaxInt64
        needed := target
        for bCost := 0; needed[0] >= 0 && needed[1] >= 0; bCost+=costB {
            moves, ok := movesTo(needed, a)
            if ok {
                best = bCost + costA * moves
                break
            }
            needed[0] -= b[0]
            needed[1] -= b[1]
        }
        needed = target
        for aCost := 0; needed[0] >= 0 && needed[1] >= 0; aCost += costA {
            moves, ok := movesTo(needed, b)
            if ok {
                best = min(best, aCost + costB*moves)
            } 
            needed[0] -= a[0]
            needed[1] -= a[1]
        }
        return best
    }

    tot := 0
    for _, lines := range sections {
        var data [3]Point
        for i, line := range lines {
            m := matcher.FindAllStringSubmatch(line, -1)
            data[i][0], _ = strconv.Atoi(m[0][1])
            data[i][1], _ = strconv.Atoi(m[0][2])
        }
        // log.Println(data)
        data[2][0]+= 10000000000000
        data[2][1]+= 10000000000000
        found := solve2(data[2], data[0], data[1])
        if found < math.MaxInt64 {
            log.Println(found)
            tot += found
        }   
    }

    log.Println("part 1:", tot)


}

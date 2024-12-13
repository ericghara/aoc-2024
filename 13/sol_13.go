package main

import (
    "aoc/parseutil"
    "log"
    "regexp"
    "strconv"
    "math"
)

func main() {
    type Point [2]int

    sections := parseutil.ReadInputSections(`^\s*$`)
    matcher := regexp.MustCompile(`[^:]*:[^\d]*(\d+)[^\d]*(\d*)`)
    presses := 100
    costA, costB := 3, 1

    solve1 := func(target, a, b Point) int {
        curA := [2]int{0, 0}
        best := math.MaxInt32
        for aI := 0; aI <= presses; aI++ {
           curB := curA
            for bI := 0; bI <= presses; bI++ {
                if curB == target {
                    best = min(best, aI*costA+bI*costB)
                }
                curB[0] += b[0]
                curB[1] += b[1]
            }
            curA[0] += a[0]
            curA[1] += a[1]
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
        found := solve1(data[2], data[0], data[1])
        if found < math.MaxInt32 {
            // log.Println(found)
            tot += found
        }   
    }

    log.Println("part 1:", tot)


}

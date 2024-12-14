package main

import (
    "aoc/parseutil"
    "log"
    "regexp"
    "fmt"
)

func main() {

    type Point [2]int
    type Velocity [2]int

    lines := parseutil.ReadInputLines()
    matcher := regexp.MustCompile(`p=([-\d]+),([-\d]+)\sv=([-\d]+),([-\d]+)`)

    starts := []Point{}
    speeds := []Velocity{}

    for _, line := range lines {
        m := matcher.FindAllStringSubmatch(line, -1)
        ints := parseutil.ParseInts([]string{m[0][2], m[0][1], m[0][4], m[0][3]})
        starts = append(starts, Point{int(ints[0]), int(ints[1])})
        speeds = append(speeds, Velocity{int(ints[2]),int(ints[3])})
    }

    simulate := func(starts []Point, speeds []Velocity, sec, r, c int) []Point {
        ends := []Point{}
        for i, p := range starts {
            end := Point{}
            end[0] = (((speeds[i][0] * sec + p[0])%r)+r)%r
            end[1] = (((speeds[i][1] * sec + p[1])%c)+c)%c
            ends = append(ends, end)
        }
        return ends
    }

    getQuad := func(p Point, r, c int) (int, bool) {
        mR, mC := r/2., c/2.
        pR, pC := p[0], p[1]
        if pR < mR && pC < mC {
            return 0, true
        }
        if pR < mR && pC > mC {
            return 1, true
        }
        if pR > mR && pC < mC {
            return 2, true
        }
        if pR > mR && pC > mC {
            return 3, true
        }
        return -1, false
    }

    quads := [4]int{}
    maxR, maxC := 103, 101
    ends := simulate(starts, speeds, 100, maxR, maxC)
    log.Println(ends)
    for _, p := range ends {
        q, ok := getQuad(p, maxR, maxC)
        if ok {
            quads[q]++
        }
    }

    printGrid := func(points map[Point]bool) bool {
        for r := range maxR {
            for c := range maxC {
                s := " "
                if points[Point{r,c}] {
                    s = "X"
                }
                fmt.Print(s)
            }
            fmt.Printf("\n")
        } 
        return true
    }
    
    var floodFill func(cur Point, seen, points map[Point]bool) int

    floodFill = func(cur Point, seen, points map[Point]bool) int {
        if seen[cur] || !points[cur] {
            return 0
        }
        seen[cur] = true
        found := 1
        for _, move := range [4]Point{{0,1},{0,-1},{1,0},{-1,0}} {
            next := Point{cur[0]+move[0], cur[1]+move[1]}
            found += floodFill(next, seen, points)
        }
        return found
    }

    log.Println("Part 1", quads[0]*quads[1]*quads[2]*quads[3])
    ends = starts
    for i := range 10000 {
        seen, points := map[Point]bool{}, map[Point]bool{}

        for _, p := range ends {
            points[p] = true
        }

        for p, _ := range points {
            if floodFill(p, seen, points) > 20 {
                printGrid(points)
                log.Println("Move:", i)
            }
        }
        printGrid(points)
        ends = simulate(ends, speeds, 1, maxR, maxC)
    }
}

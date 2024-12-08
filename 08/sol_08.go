package main

import (
    "aoc/parseutil"
    "log"
)

func main() {
    lines := parseutil.ReadInputLines()
    maxR, maxC := len(lines), len(lines[0])

    dist := func(p0, p1 [2]int) [2]int {
        return [2]int{p0[0]-p1[0], p0[1]-p1[1]}
    }

    inBounds := func(p [2]int) bool {
        return 0 <= p[0] && p[0] < maxR && 0 <= p[1] && p[1] < maxC
    }

    var gcd func(int, int) int;

    gcd = func(a,b int) int {
        if b == 0 {
            return max(a, -a)
        }
        return gcd(b, a%b)
    }
    
    freqs := map[rune][][2]int{}

    for r, row := range lines {
        for c, a := range row {
            if a != '.' {
                freqs[a] = append(freqs[a], [2]int{r,c})
            }
        }
    }

    doubles := map[[2]int]bool{}
    colinear := map[[2]int]bool{}

    for _, ants := range freqs {
        for i := 0; i < len(ants); i++ {
            for j:= 0; j < len(ants); j++ {
                a, b := ants[i], ants[j]
                if a == b {
                    continue
                }
                delta := dist(a, b)
                node := [2]int{delta[0]+a[0], delta[1]+a[1]}
                if inBounds(node) {
                    doubles[node] = true
                }
                f := gcd(delta[0], delta[1])
                delta[0], delta[1] = delta[0]/f, delta[1]/f
                for inBounds(a) {
                    colinear[a] = true
                    a[0] += delta[0]
                    a[1] += delta[1]
                }
            }
        }
    }

    log.Println("AntiNodes:", len(doubles))
    log.Println("AnyNodes:", len(colinear))
}

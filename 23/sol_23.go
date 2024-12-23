package main

import (
    "aoc/parseutil"
    "log"
    "strings"
    "sort"
    "slices"
    "maps"
)

var (
    edgeToSet = map[string]map[string]bool{}
    edgeTo = map[string][]string{}
    seen = map[string]bool{}
)

func main() {

    for _, line := range parseutil.ReadInputLines() {
        split := strings.Split(line, "-")
        if edgeToSet[split[0]] == nil {
            edgeToSet[split[0]] = map[string]bool{}
            edgeTo[split[0]] = []string{split[0]} // add self edge for p2
        }
        if edgeToSet[split[1]] == nil {
            edgeToSet[split[1]] = map[string]bool{}
            edgeTo[split[1]] = []string{split[1]}
        }
        edgeToSet[split[0]][split[1]] = true
        edgeTo[split[0]] = append(edgeTo[split[0]], split[1])
        edgeToSet[split[1]][split[0]] = true
        edgeTo[split[1]] = append(edgeTo[split[1]], split[0])
    }

    part1 := func() int {
        seen := map[[3]string]bool{}
        for n0, neighs := range edgeToSet {
            if !strings.HasPrefix(n0, "t") {
                continue;
            }
            nList := slices.Collect(maps.Keys(neighs))
            for i := 0; i < len(nList); i++ {
                n1 := nList[i]
                for j := i+1; j < len(nList); j++ {
                    n2 := nList[j] 
                    if edgeToSet[n2][n1] {
                        key := [3]string{n1, n2, n0}
                        keySlice := key[:]
                        sort.Strings(keySlice)
                        seen[key] = true
                    }
                }
            }
        }
        return len(seen)
    }

    var canClique func(sI, aI int, selected, all []string) bool

    canClique = func(sI, aI int, selected, all []string) bool {
        if sI == len(selected) {
            for i := 0; i < len(selected); i++ {
                a := selected[i]
                for j := i+1; j < len(selected); j++ {
                    b := selected[j]
                    if !edgeToSet[a][b] {
                        return false
                    }
                }
            }
            return true
        }
        for nI := aI; nI <= len(all) - len(selected) + sI; nI++ {
            selected[sI] = all[nI]
            if canClique(sI+1, nI+1, selected, all) {
                return true
            }
        }
        return false
    }

    part2 := func() string {
        var b, e int
        for _, neighs := range edgeTo {
            sort.Strings(neighs) // make returning results simpler
            e = max(len(neighs), e)
        }

        password := []string{}
        for b <= e {
            mid := (e-b)/2+b
            selected := make([]string, mid)
            for _, neighs := range edgeTo {
                if canClique(0, 0, selected, neighs) {
                    password = selected
                    break
                }
            }
            if len(password) == mid  {
                b = mid+1
            } else {
                e = mid-1
            }
        }
        return strings.Join(password,",")
    }

    log.Println("Part 1", part1())
    log.Println("Part 2", part2())
}

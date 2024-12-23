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
    edgeTo = map[string]map[string]bool{}
    seen = map[string]bool{}
)

//func dfs(cur string) []string {
//    if seen[cur] {
//        return []string{}
//    }
//    seen[cur] = true
//    found := []string{cur}
//    for _, next := range edgeTo[cur] {
//        log.Println(next)
//       found = append(found, dfs(next)...)
//    }
//    return found
//}

func main() {

    for _, line := range parseutil.ReadInputLines() {
        split := strings.Split(line, "-")
        if edgeTo[split[0]] == nil {
            edgeTo[split[0]] = map[string]bool{}
        }
        if edgeTo[split[1]] == nil {
            edgeTo[split[1]] = map[string]bool{}
        }
        edgeTo[split[0]][split[1]] = true
        edgeTo[split[1]][split[0]] = true
    }
//    log.Println(edgeTo)

//    var part1 int
//    for node, _ := range edgeTo {
//        visited := dfs(node)
//        if len(visited) >= 3 {
//            for _, name := range visited {
//                if strings.HasPrefix(name, "t") {
//                    for i := 0; i < len(visited); i++ {
//                        for j := i+1; j < len(visited); j++ {
//                            for k := j+1; k < len(visited); k++ {
//                                if strings.HasPrefix(visited[i], "t") || strings.HasPrefix(visited[j], "t") || strings.HasPrefix(visited[k], "t") {
//                                    part1 ++
//                                }
//                            }
//                        }
//                    }
//                    log.Println("broke")
//                    break
//                }
//            }
//        }
//        log.Println(visited)
//
//    }

    seen := map[[3]string]bool{}
    for n0, neighs := range edgeTo {
        if !strings.HasPrefix(n0, "t") {
            continue;
        }
        nList := slices.Collect(maps.Keys(neighs))
        for i := 0; i < len(nList); i++ {
            n1 := nList[i]
            for j := i+1; j < len(nList); j++ {
                n2 := nList[j] 
                if edgeTo[n2][n1] {
                    key := [3]string{n1, n2, n0}
                    keySlice := key[:]
                    sort.Strings(keySlice)
                    seen[key] = true
                }
            }
//            for _, n2 := range edgeTo[n1] {
//                if n2 == node {
//                    continue
//                }
//                key := [3]string{n1, n2, node}
//                keySlice := key[:]
//                sort.Strings(keySlice)
//                seen[key] = true
//            }
        }
    }

    log.Println(seen)
    log.Println("Part 1", len(seen))
}

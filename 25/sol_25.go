package main

import (
    "log"
    "aoc/parseutil"
)

func main() {
    locks, keys := map[[5]int]int{}, map[[5]int]int{}
    
    for _, section := range parseutil.ReadInputSections(`^\s*$`) {
        data := [5]int{-1,-1,-1,-1,-1}
        for _, line := range section {
            for i := range len(line) {
                if line[i] == '#' {
                    data[i]++
                }
            }
        }
        if section[0] == "#####" {
                locks[data]++
        } else {
            keys[data]++
        }
    }

    var part1 int
    for data, cnt := range locks {
        for keyData, keyCnt := range keys {
            ok := true
            for i := range len(data) {
                if data[i] + keyData[i] > 5 {
                    ok = false
                    break
                }
            }
            if ok {
                part1 += min(keyCnt, cnt)
            }
        }
    }
    log.Println(part1)
}

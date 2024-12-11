package main

import (
    "aoc/parseutil"
    "log"
    "strings"
    "fmt"
    "strconv"
)


func main() {
    
    input := parseutil.ParseInts(strings.Split(parseutil.ReadInputLines()[0], " "))

    simulate := func(nums []int64, rounds int) []int64 {
            last := nums
            for ;rounds > 0; rounds-- {
                next := []int64{}
                for _, n := range last {
                    if n == 0 {
                        next = append(next, 1)
                    } else if numStr := fmt.Sprint(n); len(numStr) % 2 == 0 {
                        a, err0 := strconv.ParseInt(numStr[:len(numStr)/2], 10, 64)
                        b, err1 := strconv.ParseInt(numStr[len(numStr)/2:], 10, 64)
                        if (err0 != nil || err1 != nil) {
                            log.Println("HMM")
                        }
                        next = append(next, a, b)
                    } else {
                        next = append(next, n * 2024)
                        if n*2024 < 0 {
                            log.Panic("OVERFLOW")
                        }
                    }
                }
                last = next
            }
            return last
        }

    endState := simulate(input, 25)
    //log.Println(endState)
    log.Println("Num Stones", len(endState))


}

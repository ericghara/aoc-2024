package main

import (
    "aoc/parseutil"
    "log"
    "strings"
    "strconv"
)

func main() {
    input := parseutil.ParseInts(strings.Split(parseutil.ReadInputLines()[0], " "))
    memo := map[[2]int64]int64{}

    var simulate func(int64, int64, int64) int64

    simulate = func(num, round, endRound int64) int64 {
        roundLeft := endRound-round
        if roundLeft == 0 {
            return 1
        }
        if stones, ok := memo[[2]int64{num,roundLeft}]; ok {
            return stones
        }
        var numStones int64
        if num == 0 {
            numStones = simulate(1, round+1, endRound)
        } else if numStr := strconv.FormatInt(num, 10); len(numStr) % 2 == 0 {
            a, _ := strconv.ParseInt(numStr[:len(numStr)/2], 10, 64)
            b, _ := strconv.ParseInt(numStr[len(numStr)/2:], 10, 64)
            numStones = simulate(a, round+1, endRound) + simulate(b, round+1, endRound)
        } else {
            numStones = simulate(num * 2024, round+1, endRound)
        }
        memo[[2]int64{num,roundLeft}] = numStones
        return numStones
    }

    var totStones1, totStones2, rounds1, rounds2 int64
    rounds1, rounds2 = 25, 75

    for _, num := range input {
        totStones1 += simulate(num, 0, rounds1)
        totStones2 += simulate(num, 0, rounds2)
    }

    log.Println("pt1 Num Stones:", totStones1)
    log.Println("pt2 Num Stones:", totStones2)
}

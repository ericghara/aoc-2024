package main

import (
	"aoc/parseutil"
	"log"
	"strconv"
)

var Q int = 16777216

func mix(a, b int) int {
	return a ^ b
}

func prune(a, q int) int {
	return a % q
}

func calcNext(cur int) int {
	op := cur << 6
	next := prune(mix(cur, op), Q)
	op = next >> 5
	next = prune(mix(next, op), Q)
	op = next << 11
	next = prune(mix(next, op), Q)
	return next
}

func main() {
	starts := []int{}

	for _, line := range parseutil.ReadInputLines() {
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Panic("unable to read line", line)
		}
		starts = append(starts, int(num))
	}

	var part1, part2 int

	globSale := map[[4]int]int{}

	for _, num := range starts {
		last := num
		pattern := []int{}
		curSale := map[[4]int]bool{}

		for i := 0; i < 2000; i++ {
			next := calcNext(last)
			delta := next%10 - last%10
			last = next
			pattern = append(pattern, delta)
			if len(pattern) > 4 {
				pattern = pattern[1:]
			}
			if len(pattern) == 4 {
				key := [4]int{pattern[0], pattern[1], pattern[2], pattern[3]}
				if !curSale[key] {
					curSale[key] = true
					globSale[key] += (next % 10)
					part2 = max(part2, globSale[key])
				}
			}
		}
		part1 += last
	}

	log.Println("Part 1:", part1)
	log.Println("Part 2:", part2)
}

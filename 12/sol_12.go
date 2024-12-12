package main

import (
	"aoc/parseutil"
	"log"
)

func main() {

	input := parseutil.ReadInputLines()
	garden := make([][]rune, len(input))
	seen := make([][]bool, len(input))

	for _, line := range input {
		garden = append(garden, parseutil.ToRunes(line))
		seen = append(seen, make([]bool, len(garden[len(seen)])))
	}

	var measure func(int, int, int, rune) (p int, a int)

	moves := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	boundaries := map[[3]int]bool{} // by direction

	measure = func(r, c, d int, plant rune) (int, int) {
		if r < 0 || r >= len(garden) || c < 0 || c >= len(garden[r]) {
			boundaries[[3]int{d, r, c}] = true
			return 1, 0
		}
		if garden[r][c] != plant {
			boundaries[[3]int{d, r, c}] = true
			return 1, 0
		}
		if seen[r][c] {
			return 0, 0
		}
		seen[r][c] = true
		p, a := 0, 1
		for nD, move := range moves {
			dP, dA := measure(r+move[0], c+move[1], nD, plant)
			p += dP
			a += dA
		}
		return p, a
	}

	countSides := func() int {
		sides := 0
		for k, _ := range boundaries {
			sides++
			dir, start := moves[(k[0]+1)%4], k
			for boundaries[start] {
				start[1] += dir[0]
				start[2] += dir[1]
			}
			dir[0] *= -1
			dir[1] *= -1
			start[1] += dir[0]
			start[2] += dir[1]
			for boundaries[start] {
				delete(boundaries, start)
				start[1] += dir[0]
				start[2] += dir[1]
			}
		}
		return sides
	}

	fullCost, discountCost := 0, 0

	for r := range len(garden) {
		for c := range len(garden[r]) {
			if seen[r][c] {
				continue
			}
			plant := garden[r][c]
			p, a := measure(r, c, -1, plant)
			fullCost += p * a
			s := countSides()
			discountCost += s * a
		}
	}

	log.Println("Full cost:", fullCost)
	log.Println("Discount cost:", discountCost)
}

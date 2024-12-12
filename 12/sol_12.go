package main

import (
	"aoc/parseutil"
	"log"
)

func main() {

	input := parseutil.ReadInputLines()
	garden := make([][]rune, len(input))
	seen := make([][]int, len(input))

	for _, line := range input {
		garden = append(garden, parseutil.ToRunes(line))
		seen = append(seen, make([]int, len(garden[len(seen)])))
	}

	var measure func(int, int, int, rune, int) (p int, a int)

	moves := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	boundaries := map[[3]int]int{} // by direction

	measure = func(r, c, d int, plant rune, id int) (int, int) {
		if r < 0 || r >= len(garden) || c < 0 || c >= len(garden[r]) {
			boundaries[[3]int{d, r, c}] = id
			return 1, 0
		}
		if seen[r][c] == id {
			return 0, 0
		}
		if seen[r][c] != 0 || garden[r][c] != plant {
			boundaries[[3]int{d, r, c}] = id
			return 1, 0
		}
		seen[r][c] = id
		p, a := 0, 1
		for nD, move := range moves {
			dP, dA := measure(r+move[0], c+move[1], nD, plant, id)
			p += dP
			a += dA
		}
		return p, a
	}

	countSides := func() int {
		sides := 0
		for k, id := range boundaries {
			sides++
			dir, start := moves[(k[0]+1)%4], k
			for boundaries[start] == id {
				start[1] += dir[0]
				start[2] += dir[1]
			}
			dir[0] *= -1
			dir[1] *= -1
			start[1] += dir[0]
			start[2] += dir[1]
			for boundaries[start] == id {
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
			if seen[r][c] != 0 {
				continue
			}
			id := r*len(garden[r]) + c + 1
			plant := garden[r][c]
			p, a := measure(r, c, -1, plant, id)
			fullCost += p * a
			s := countSides()
			discountCost += s * a
		}
	}

	log.Println("Full cost:", fullCost)
	log.Println("Discount cost:", discountCost)
}

package main

import (
	"aoc/parseutil"
	"log"
	"strings"
)

func main() {
	sections := parseutil.ReadInputSections(`^\s*$`)
	patterns := strings.Split(sections[0][0], ", ")
	designs := sections[1]

	planDesign := func(design string) int {
		reachable := make([]int, len(design)+1)
		reachable[0] = 1

		for i, cnt := range reachable {
			if cnt == 0 {
				continue
			}
			for _, p := range patterns {
				if i+len(p) > len(design) {
					continue
				}
				if design[i:i+len(p)] == p {
					reachable[i+len(p)] += reachable[i]
				}
			}
		}

		return reachable[len(design)]
	}

	numPossible := 0
	numWays := 0
	for _, d := range designs {
		if w := planDesign(d); w > 0 {
			numPossible++
			numWays += w
		}
	}

	log.Println("Part 1", numPossible)
	log.Println("Part 2", numWays)
}

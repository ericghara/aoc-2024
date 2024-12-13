package main

import (
	"aoc/parseutil"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize/convex/lp"
	"log"
	"math"
	"regexp"
	"strconv"
)

func main() {

	type Point [2]float64

	sections := parseutil.ReadInputSections(`^\s*$`)
	matcher := regexp.MustCompile(`[^:]*:[^\d]*(\d+)[^\d]*(\d*)`)

	c := []float64{3, 1}
	offset := 10000000000000.0 // 0 for part 1

	almostEqual := func(a, b float64) bool {
		epsilon := 0.1
		return a - epsilon < b && a + epsilon > b
	}

	tot := 0
	for _, lines := range sections {
		var data [3]Point
		for i, line := range lines {
			m := matcher.FindAllStringSubmatch(line, -1)
			data[i][0], _ = strconv.ParseFloat(m[0][1], 64)
			data[i][1], _ = strconv.ParseFloat(m[0][2], 64)
		}
		data[2][0] += offset
		data[2][1] += offset

		A := mat.NewDense(2, 2, []float64{data[0][0], data[1][0], data[0][1], data[1][1]})
		b := []float64{data[2][0], data[2][1]}
		opt, x, err := lp.Simplex(c, A, b, 0, nil)
		if err != nil {
			//log.Printf("Unfeasible")
			continue
		}
		pressA, pressB := math.Round(x[0]), math.Round(x[1])
		foundX := pressA*data[0][0] + pressB*data[1][0]
		foundY := pressA*data[0][1] + pressB*data[1][1]
		if almostEqual(data[2][0], foundX) && almostEqual(data[2][1], foundY) {
			tot += int(math.Round(opt))
		}
	}

	log.Println("Tot:", tot)
}

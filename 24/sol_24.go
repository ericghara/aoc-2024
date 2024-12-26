package main

import (
	"aoc/parseutil"
	"fmt"
	"log"
	"math/bits"
	"regexp"
	"sort"
	"strings"
)

type Gate struct {
	Inputs [2]string
	Output string
	eval   func(g *Gate, vals map[string]uint)
}

func and(g *Gate, vals map[string]uint) {
	inputs := g.Known(vals)
	vals[g.Output] = vals[inputs[0]] & vals[inputs[1]]
}

func or(g *Gate, vals map[string]uint) {
	inputs := g.Known(vals)
	vals[g.Output] = vals[inputs[0]] | vals[inputs[1]]
}

func xor(g *Gate, vals map[string]uint) {
	inputs := g.Known(vals)
	vals[g.Output] = vals[inputs[0]] ^ vals[inputs[1]]
}

func (g *Gate) Eval(vals map[string]uint) {
	g.eval(g, vals)
}

func (g *Gate) Known(vals map[string]uint) []string {
	ret := []string{}
	for _, input := range g.Inputs {
		_, ok := vals[input]
		if ok {
			ret = append(ret, input)
		}
	}
	return ret
}

var (
	gatePs     = []*Gate{}
	edgeTo     = map[string][]*Gate{}
	numBits    = 0
	xs, ys, zs []string
)

func main() {
	sections := parseutil.ReadInputSections(`^\s*$`)
	m0 := regexp.MustCompile(`([^:]*):\s*([01])`)
	initVals := map[string]uint{}

	for _, line := range sections[0] {
		match := m0.FindStringSubmatch(line)
		if len(match) == 0 {
			log.Panic("Unable to parse line", line)
		}
		val := match[2][0] - '0'
		initVals[match[1]] = uint(val)
	}

	m1 := regexp.MustCompile(`([^\s]*)\s+(\w*)\s+([^\s]+)\s+->\s+([^\s]*)`)

	for _, line := range sections[1] {
		match := m1.FindStringSubmatch(line)
		if len(match) == 0 {
			log.Panic("Unable to parse line", line)
		}
		var fn func(*Gate, map[string]uint)
		switch match[2] {
		case "AND":
			fn = and
		case "OR":
			fn = or
		case "XOR":
			fn = xor
		default:
			log.Panic("invalid token", match[2])
		}
		gatePs = append(gatePs,
			&Gate{
				Inputs: [2]string{match[1], match[3]},
				Output: match[4],
				eval:   fn,
			})
		cur := gatePs[len(gatePs)-1]
		for _, reg := range [2]string{match[1], match[3]} {
			edgeTo[reg] = append(edgeTo[reg], cur)
		}
		if strings.HasPrefix(cur.Output, "z") {
			numBits++
		}
	}

	for i := range numBits {
		xs = append(xs, fmt.Sprintf("x%02d", i))
		ys = append(ys, fmt.Sprintf("y%02d", i))
		zs = append(zs, fmt.Sprintf("z%02d", i))
	}

	sum := func(x, y uint) uint {
		vals := map[string]uint{}
		q := []string{}
		for i := range numBits {
			vals[xs[i]] = x & 1
			vals[ys[i]] = y & 1
			x >>= 1
			y >>= 1
			q = append(q, xs[i], ys[i])
		}

		for len(q) > 0 {
			next := []string{}
			for _, reg := range q {
				for _, gate := range edgeTo[reg] {
					if len(gate.Known(vals)) == 2 {
						gate.Eval(vals)
						next = append(next, gate.Output)
					}
				}
			}
			q = next
		}
		var tot uint
		for i, zStr := range zs {
			tot |= (vals[zStr] << i)
		}
		return tot
	}

	var test func(uint) uint
	numThread := 14
	in, out := make(chan [2]uint, numThread), make(chan uint, numThread)
	defer close(in)
	defer close(out)

	for i := 0; i < numThread; i++ {
		go func() {
			for nums := range in {
				res := sum(nums[0], nums[1])
				out <- (nums[0] + nums[1]) ^ res
			}
		}()
	}

	test = func(bad uint) uint {
		jobs := [][2]uint{{0, 0}}
		for i := range numBits - 1 { // ignore MSB (overflow)
			mask := uint(1 << i)
			if mask&bad == 0 {
				continue
			}
			beforeMask := uint(mask >> 1)
			adds := [][2]uint{{mask, mask}, {0, mask}, {mask, 0}, {mask | beforeMask, beforeMask}}
			for _, nums := range adds {
				jobs = append(jobs, nums)
			}
		}

		go func() {
			for _, j := range jobs {
				in <- j
			}
		}()

		var newBad uint
		for i := 0; i < len(jobs); i++ {
			newBad |= <-out
		}
		return newBad
	}

	var repair func(i, swap int, bad uint) []*Gate

	repair = func(i, swaps int, bad uint) []*Gate {
		if bad == 0 {
			return []*Gate{}
		}
		if swaps == 4 || i >= len(gatePs) {
			return nil
		}
		if ret := repair(i+1, swaps, bad); ret != nil || gatePs[i] == nil {
			return ret
		}
		for j := i + 1; j < len(gatePs); j++ {
			if gatePs[j] == nil {
				continue
			}
			gatePs[i].Output, gatePs[j].Output = gatePs[j].Output, gatePs[i].Output
			tmp := gatePs[j]
			gatePs[j] = nil
			newBad := test(bad)
			if bits.OnesCount(newBad)+1 < bits.OnesCount(bad) { // >= 2 bits better
				log.Println(swaps, newBad, tmp.Output, gatePs[i].Output)
				if ret := repair(i+1, swaps+1, newBad); ret != nil {
					return append(ret, gatePs[i], tmp)
				}
			}
			gatePs[j] = tmp
			gatePs[i].Output, gatePs[j].Output = gatePs[j].Output, gatePs[i].Output
		}
		return nil
	}

	var x, y uint
	for i := range numBits {
		x |= (initVals[xs[i]] << i)
		y |= (initVals[ys[i]] << i)
	}

	log.Println("Part 1", sum(x, y))
	bad := test(0xffffffffffffffff)
	log.Printf("Tot Bits: %d, Bad Bits %d, Start Mask %d\n", numBits, bits.OnesCount(bad), bad)
	part2 := []string{}
	for _, gate := range repair(0, 0, test(bad)) {
		part2 = append(part2, gate.Output)
	}
	sort.Strings(part2)
	log.Println("Part 2: ", strings.Join(part2, ","))
}

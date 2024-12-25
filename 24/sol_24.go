package main

import (
    "aoc/parseutil"
    "log"
    "regexp"
    "strings"
    "strconv"
    "fmt"
    "math/bits"
    "math"
)

func and(g* Gate) {
    inputs := g.Known()
    vals[g.Output] = vals[inputs[0]] & vals[inputs[1]]
}

func or(g* Gate) {
    inputs := g.Known()
    vals[g.Output] = vals[inputs[0]] | vals[inputs[1]]
}

func xor(g* Gate) {
    inputs := g.Known()
    vals[g.Output] = vals[inputs[0]] ^ vals[inputs[1]]
} 

type Gate struct {
    Inputs [2]string
    Output string
    eval func(g* Gate)
}

func (g* Gate) Eval() {
    g.eval(g)
}

func (g* Gate) Known() []string {
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
    vals = map[string]int{}
    gatePs = []*Gate{}
    edgeTo = map[string][]*Gate{}
    numBits = 0
    xs = []string{}
    ys = []string{}
    zs = []string{}
)

func main() {
    sections := parseutil.ReadInputSections(`^\s*$`)

    m0 := regexp.MustCompile(`([^:]*):\s*([01])`)

    for _, line := range sections[0] {
        match := m0.FindStringSubmatch(line)
        if len(match) == 0 {
            log.Panic("Unable to parse line", line)
        }
        val := match[2][0] - '0'
        vals[match[1]] = int(val)
    }
    
    m1 := regexp.MustCompile(`([^\s]*)\s+(\w*)\s+([^\s]+)\s+->\s+([^\s]*)`)
    for _, line := range sections[1] {
        match := m1.FindStringSubmatch(line)
        if len(match) == 0 {
            log.Panic("Unable to parse line", line)
        }
        var fn func(*Gate)
        switch match[2] {
        case "AND": fn = and
        case "OR": fn = or
        case "XOR": fn = xor
        default: log.Panic("invalid token", match[2])
        }
        gatePs = append(gatePs, 
            &Gate{
                Inputs:[2]string{match[1], match[3]},
                Output: match[4],
                eval: fn,
        })
        cur := gatePs[len(gatePs)-1]
        for _, reg := range [2]string{match[1], match[3]} {
            edgeTo[reg] = append(edgeTo[reg], cur)
            fmt.Printf("cur %p edgeTo %v\n",cur, edgeTo[reg])
        }
        if strings.HasPrefix(cur.Output, "z") {
            bits, err := strconv.Atoi(cur.Output[1:])
            if err != nil {
                log.Panic("unable to parse", cur.Output)
            }
            numBits = max(numBits, int(bits+1))
        }
    }

    for i := range numBits {
        xs = append(xs, fmt.Sprintf("x%02d", i))
        ys = append(ys, fmt.Sprintf("y%02d", i))
        zs = append(zs, fmt.Sprintf("z%02d", i))
    }

    sum := func(x, y int) int {
        vals = map[string]int{}
        q := []string{}
        for i := range numBits {
            vals[xs[i]] = x&1
            vals[ys[i]] = y&1
            x >>= 1
            y >>= 1
            q = append(q, xs[i], ys[i])
        }

        for len(q) > 0 {
            next := []string{}
            for _, reg := range q {
                //log.Println(reg)
                for _, gate := range edgeTo[reg] {
                    if len(gate.Known()) == 2 {
                        gate.Eval()
                        //log.Println(gate.Output)
                        next = append(next, gate.Output)
                    }
                }
            }
            q = next
        }
        var tot int
        for i, zStr := range zs {
            tot |= (vals[zStr]  << i)
        }
        return tot 
    }

    var test func(uint) uint
    test = func(bad uint) uint {
        bitMask := uint((1 << numBits) - 1)
        newBad := uint(sum(0,0))
        for i := range numBits-1 { // test me
            mask := uint(1 << i)
            if mask & bad == 0 || newBad & mask > 0 {
                continue
            }
            beforeMask := uint(mask >> 1)
            adds := [][2]uint{{mask,mask},{0,mask},{mask,0}, {mask|beforeMask,beforeMask}}
            for _, nums := range adds {
                if found := sum(int(nums[0]), int(nums[1])); (nums[0] + nums[1]) & bitMask != uint(found) {
                   //log.Println(nums[0], nums[1], found, (nums[0]+nums[1])&bitMask )
                    newBad |= mask
                    break
                    //log.Println(i)
                }
            }
        }
        if bits.OnesCount(newBad) < bits.OnesCount(bad) && bad != math.MaxUint {
            return test(math.MaxUint)
        }
        return newBad
    }


    var bruteForce func(i, swap int, bad uint) []*Gate

    bruteForce = func(i, swaps int, bad uint) []*Gate {
        if bad == 0 {
            return []*Gate{}
        }
        if swaps == 4 || i >= len(gatePs) {
            return nil
        }
        if ret := bruteForce(i+1, swaps, bad); ret != nil {
            return ret
        }
        if gatePs[i] == nil {
            return nil
        }
        for j := i+1; j < len(gatePs); j++ {
            if gatePs[j] == nil {
                continue
            }
            gatePs[i].Output, gatePs[j].Output = gatePs[j].Output, gatePs[i].Output
            tmp := gatePs[j]
            gatePs[j] = nil
            newBad := test(bad)
            if bits.OnesCount(newBad) < bits.OnesCount(bad) {
                log.Println(swaps, newBad, tmp.Output, gatePs[i].Output)
                if ret := bruteForce(i+1, swaps+1, newBad); ret != nil {
                    return append(ret, gatePs[i], tmp)
                }
            }
            gatePs[j] = tmp
            gatePs[i].Output, gatePs[j].Output = gatePs[j].Output, gatePs[i].Output
        }
        return nil
    }

    var x, y int
    for i := range numBits {
        x |= (vals[xs[i]] << i)
        y |= (vals[ys[i]] << i)
    }
    
    log.Println("Part 1", sum(x, y))
    log.Println(numBits)
    bad := test(0xffffffffffffffff)
    log.Println("start", bad)
    log.Println("test", bits.OnesCount(bad))
    for _, gate := range bruteForce(0, 0, test(bad)) {
        log.Println(gate.Output)
    }
} 

package main

import (
    "aoc/parseutil"
    "log"
    "regexp"
    "strings"
    "sort"
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
    gates = []Gate{}
    edgeTo = map[string][]*Gate{}
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
    q := []string{}
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
        gates = append(gates, 
            Gate{
                Inputs:[2]string{match[1], match[3]},
                Output: match[4],
                eval: fn,
        })
        cur := &gates[len(gates)-1]
        for _, reg := range [2]string{match[1], match[3]} {
            edgeTo[reg] = append(edgeTo[reg], cur)
            q = append(q, reg)
        }
    }

    for len(q) > 0 {
        next := []string{}
        for _, reg := range q {
            for _, gate := range  edgeTo[reg] {
                log.Println("Known:", gate.Known())
                if len(gate.Known()) == 2 {
                    log.Println(gate)
                    gate.Eval()
                    next = append(next, gate.Output)
                }
            }
        }
        q = next
        log.Println("step")
    }

    log.Println(vals)
    zees := []string{}
    for key, _ := range vals {
        if strings.HasPrefix(key, "z") {
            zees = append(zees, key)
        }
    }
    sort.Sort(sort.Reverse(sort.StringSlice(zees)))
    var part1 int
    for i := range len(zees) {
        val, _ := vals[zees[i]]
        part1 <<= 1
        part1 |= val
    }

    log.Println("Part 1", part1)
}

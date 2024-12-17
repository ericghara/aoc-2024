package main

import (
    "aoc/parseutil"
    "log"
    "regexp"
    "strconv"
    "strings"
    "slices"
)

var (
    regs = [3]int64{}
    ins = []int64{}
    insPtr = int64(0)
    outs = []int64{}
)

func adv(combo int64) {
    xdv(&regs[0], &regs[0], combo)
}

func bxl(lit int64) {
    regs[1] ^= lit
}

func bst(combo int64) {
    cVal := comboDecode(combo)
    regs[1] = cVal & 0b111
}

func jnz(lit int64) {
    if regs[0] == 0 {
        return 
    }
    insPtr = lit-2
}

func bxc(noop int64) {
    regs[1] = regs[1]^regs[2]
}

func out(combo int64) {
    cVal := comboDecode(combo)
    outs = append(outs, cVal & 0b111)
}

func bdv(combo int64) {
    xdv(&regs[0], &regs[1], combo)
}

func cdv(combo int64) {
    xdv(&regs[0], &regs[2], combo)
}

func xdv(inReg *int64, outReg *int64, combo int64) {
    shift := comboDecode(combo)
    if shift >= 63 {
        log.Println("overflow")
        *outReg = 0
    } else {
        *outReg = *inReg/(1 << comboDecode(combo))
    }
}

func comboDecode(combo int64) int64 {
    switch combo {
    case 0,1,2,3:
        return combo
    case 4,5,6:
        return regs[combo-4]
    default:
        log.Panic("Unrecognized operand: ", combo)
    }
    return -1
}

func main() {
    sects := parseutil.ReadInputSections(`^\s*$`)
    matcher := regexp.MustCompile(`[^:]*:\s((\d+,*)+)`)
    for i, line := range sects[0] {
        numStr := matcher.FindAllStringSubmatch(line, -1)
        n, err := strconv.ParseInt(numStr[0][1], 10, 64) 
        if err != nil {
            log.Panic("parse error", numStr)
        }
        regs[i] = n
    }
    ins = parseutil.ParseInts(strings.Split(matcher.FindAllStringSubmatch(sects[1][0], -1)[0][1], ","))

    execute := func() {
        for insPtr >= 0 && insPtr < int64(len(ins)) {
            opCode := ins[insPtr]
            operand := ins[insPtr+1]
            var operator func(int64)
            switch opCode {
            case 0: operator = adv
            case 1: operator = bxl
            case 2: operator = bst
            case 3: operator = jnz
            case 4: operator = bxc
            case 5: operator = out
            case 6: operator = bdv
            case 7: operator = cdv
            default: log.Panic("Unknown opcode", opCode)
            }
            operator(operand)
            insPtr+=2
        }
    }

    crack := func() int64 {
        q := []int64{0}
        for tokI := len(ins)-1; tokI >= 0; tokI-- {
            next := []int64{}
            target := ins[tokI:]
            for _, num := range q {
                for i := range 8 {
                    protoA := (num << 3) | int64(i)
                    regs[0] = protoA
                    outs = outs[:0]
                    insPtr = 0
                    execute()
                    if slices.Equal(outs, target) {
                        next = append(next, protoA)
                    }
                }
            }
            q = next
        }
        return slices.Min(q)
    }
    
    execute()
    outStrs := []string{}
    for _, out := range outs {
        outStr := strconv.FormatInt(out, 10)
        outStrs = append(outStrs, outStr)
    }
    log.Println("Part 1:", strings.Join(outStrs, ","))
    log.Println("Part 2:", crack())
}

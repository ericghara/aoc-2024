package main

import (
    "log"
    "aoc/parseutil"
    "strconv"
    bh "github.com/emirpasic/gods/trees/binaryheap"
    //"strings"
)

type Sweep [2]rune
type Point [2]int
var moves = map[rune]Point{'v': {1,0}, '<': {0,-1}, '>': {0,1},'^': {-1,0}}


type State struct {
    memo map[Sweep]int
    pad [][]rune
    coords map[rune]Point
}

type DjState struct {
    Pos Point
    Cost int
    Last rune
    Done bool
}

func djComp(a,b interface{}) int {
    return a.(DjState).Cost - b.(DjState).Cost
}


func main() {
    codes := map[[4]rune]int{}
    
    numPad := [][]rune{  {'7','8','9'},
                         {'4','5','6'},
                         {'1','2','3'},
                         { 0, '0','A'}}

    dirPad := [][]rune{{0 ,'^','A'},
                       {'<','v','>'}}

    for _, line := range parseutil.ReadInputLines() {
        numStr := line[0:3]
        num, err := strconv.ParseInt(line[0:3], 10, 64)
        if err != nil {
            log.Panic("unable to parse", numStr)
        }
        tokens := [4]rune{}
        for i, r := range line {
            tokens[i] = r
        }
        codes[tokens] = int(num)
    }

    getCoords := func(pad [][]rune) map[rune]Point {
        coords := map[rune]Point{}
        for r := range len(pad) {
            for c:= range len(pad[r]) {
                coords[pad[r][c]] = Point{r,c}
            }
        }
        return coords
    }

    numCoords, dirCoords := getCoords(numPad), getCoords(dirPad)

    initStack := func() []State {
        stack := make([]State, 26)
        stack[0] = State{
            memo: map[Sweep]int{},
            pad: numPad,
            coords: numCoords,
        }
        for i := 1; i < len(stack); i++ {
            stack[i] = State {
            memo: map[Sweep]int{},
            pad: dirPad,
            coords: dirCoords,
            }
        }
        return stack
    }

    states := initStack()

    validMove := func(move Point, pad[][]rune) bool {
        return  move[0] >= 0 && move[0] < len(pad) && move[1] >= 0 && 
            move[1] < len(pad[0]) && pad[move[0]][move[1]] != 0
    }
    
    var recurse func(Sweep, int) int

    recurse = func(sweep Sweep, stateI int) int {
        if stateI == len(states) {
            return 1
        }
        state := &states[stateI]
        if cost, ok := state.memo[sweep]; ok {
            return cost
        }
        start := DjState {
            Pos: state.coords[sweep[0]],
            Cost: 0,
            Last: 'A',
        }
        end := state.coords[sweep[1]]
        costs := map[rune]int{sweep[0]: 0}
        q := bh.NewWith(djComp)
        q.Push(start)
        for cur, ok := q.Pop(); ok; cur, ok = q.Pop() {
            curState := cur.(DjState)
            if curState.Done {
                state.memo[sweep] = curState.Cost
                break
            }
            if curState.Pos == end {
                q.Push(DjState{
                    Pos: end,
                    Cost: curState.Cost + recurse(Sweep{curState.Last, 'A'}, stateI+1),
                    Last: 'A',
                    Done: true,
                })
                continue
            }
            for button , move := range moves {
                nPos := Point{move[0]+curState.Pos[0], move[1]+curState.Pos[1]}
                if validMove(nPos, state.pad) {
                    nCost := curState.Cost + recurse(Sweep{curState.Last, button}, stateI+1)
                    nRune := state.pad[nPos[0]][nPos[1]]
                    if oldCost, ok := costs[nRune]; !ok || oldCost >= nCost {
                        q.Push(DjState{
                            Pos: nPos,
                            Cost: nCost,
                            Last: button,
                        })
                        costs[nRune] = nCost
                    }
                }
            }
        }
        cost, ok := state.memo[sweep]
        if !ok {
            log.Panic("unable to find presses")
        }
        return cost
    }


    res1 := 0

    for code, factor := range codes {
        last := 'A'
        tot := 0
        for _, c := range code {
            tot += recurse(Sweep{last,c}, 0)
            last = c
        }
        log.Println(tot, factor)
        res1 += tot * factor
    }

    log.Println("Part1", res1)






//    getDists := func(pad [][]rune) map[[2]rune]Point {
//        dists := map[[2]rune]Point{}
//        for oR := range len(pad) {
//            for oC := range len(pad[oR]) {
//                origin := pad[oR][oC]
//                for dR := oR; dR < len(pad); dR++ {
//                    for dC := oC; dC < len(pad); dC++ {
//                        dest := pad[dR][dC]
//                        dists[[2]rune{dest, origin}] = Point{dR-oR, dC-oC}
//                        dists[[2]rune{origin, dest}] = Point{oR-dR, oC-dC}
//                    }
//                }
//            }
//        }
//        return dists
//    }
//
//    numDist := getDists(numPad)
//    dirDist := getDists(dirPad)
//
//    getPress := func(dist Point, map[[2]rune]Point cost) int {
//        nextDist := Point{0, 0}
//        if dist[0] > 0 {
//            nextDist[0] += cost[[2]rune{'A','^'}] * dist[0]
//            nextDist[0] += cost[[2]rune{'^','A'}] * dist[0]
//        } else {
//            nextDist[0] += cost[[2]rune{'A','v'}] * -dist[0]
//            nextDist[0] += cost[[2]rune{'v','A'}] * -dist[0] 
//        }
//        if dist[1] > 0 {
//            nextDist[1] += cost[[2]rune{'A','>'}] * dist[1]
//            nextDist[1] += cost[[2]rune{'>',''}]
//        }
//        
//
//
//
//
//    }


//
//    getDists := func(pad [][]rune) map[rune]map[rune]int {
//        dists := map[rune]map[rune]int{}
//
//        for r := range len(numPad) {
//            for c, num := range numPad[r] {
//                if pad[r][c] == 0 {
//                    continue
//                }
//                seen := map[rune]int{num: 1, 0: -1}
//                step := 1
//                last := []Point{{r,c}}
//                for len(last) > 0 {
//                    step++
//                    next := []Point{}
//                    for _, p := range last {
//                        for _, delta :=  range moves {
//                            nMove := Point{delta[0]+p[0], delta[1]+p[1]}
//                            if validMove(nMove, pad) && seen[pad[nMove[0]][nMove[1]]] == 0 {
//                                seen[pad[nMove[0]][nMove[1]]] = step
//                                next = append(next, nMove)
//                            } 
//                        }
//                    }
//                }
//                dists[num] = seen
//            }
//        }
//        return dists
//    }




    log.Println(codes)
}

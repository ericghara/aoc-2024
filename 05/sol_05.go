package main

import (
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
)

func main() {

    fileName := "./input"
    if len(os.Args) > 1 {
        fileName = os.Args[1]
    }

    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal("unable to open: ", fileName)
    }
    defer file.Close()

    deps := map[int64][]int64{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if len(line) == 0 {
            break
        }
        tokens := strings.Split(line, "|")
        if len(tokens) != 2 {
            log.Fatal("Unable to parse line")
        }
        a, errA := strconv.ParseInt(tokens[0], 10, 64)
        b, errB := strconv.ParseInt(tokens[1], 10, 64) 
        if errA != nil || errB != nil {
            log.Fatal("Unable to parse token")
        }
        arr, ok := deps[a]
        if !ok {
            arr = make([]int64, 0)
        }
        deps[a] = append(arr, b)
    }

    topo := func(nums map[int64]bool) int64 {
        inDegree := map[int64]int{}
        for n, _ := range nums {
            inDegree[n] = 0
        }
        for v, _ := range nums { 
            for _, u := range deps[v] {
                if nums[u] {
                    inDegree[u]++
                } 
            }
        }
        order := make([]int64, 0)
        next := make([]int64, 0)
        for v, in := range inDegree {
            if in == 0 {
                next = append(next, v)
                inDegree[v]++
            }
        }
        for len(next) > 0 {
            cur := next
            next = make([]int64, 0)
            for _, v := range cur {
                order = append(order, v)
                for _, u := range deps[v] {
                    if inDegree[u]--; inDegree[u] == 0 {
                        next = append(next, u)
                    }
                }
            }
        }
        return order[len(order)/2]
    }

    var sumValid int64
    var sumInvalid int64

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        tokens := strings.Split(line, ",")
        seen := map[int64] bool{}
        validLine := true
        for _, token := range tokens {
            cur, err := strconv.ParseInt(token, 10, 64)
            if err != nil {
                log.Fatal("Unable to parse token (part 1)")
            }
            seen[cur] = true
            curDeps, ok := deps[cur]
            if !ok {
                continue
            }
            for _, d := range curDeps {
                if seen[d] {
                    validLine = false
                }
            }
        }
        if validLine {
            mid, _ := strconv.ParseInt(tokens[len(tokens)/2], 10, 64)
            log.Println(mid)
            sumValid += mid
        } else {
            sumInvalid += topo(seen)
        }
    }
    log.Println("Mid Valid:", sumValid)
    log.Println("Mid invalid:", sumInvalid)
}

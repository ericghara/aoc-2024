package main

import (
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
    "sort"
    "math"
)

func main() {
    file, err := os.Open("./input")
    if err != nil {
        log.Fatal("Could not open")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    var a []int
    var b []int
    m := make(map[int]int)

    for scanner.Scan() {
        line := scanner.Text()
        line = strings.TrimSpace(line)
        numStrs := strings.Split(line, "   ")
        if len(numStrs) != 2 {
            log.Printf("Could not parse line: %v.", line)
            continue
        }
        i,err := strconv.Atoi(numStrs[0])
        if err != nil {
            log.Panic("Could not parse line: %v", line)
        }
        a = append(a, i)

        i, err = strconv.Atoi(numStrs[1])
        if err != nil {
            log.Panic("could not parse line: %v", line)
        }

        b = append(b, i)
        _, ok := m[i]
        if !ok {
            m[i] = 1
        } else {
            m[i]++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    sort.Ints(a)
    sort.Ints(b)

    dif := 0.
    sim := 0

    for i := range len(a) {
        dif += math.Abs(float64(a[i]-b[i]))
        cnt, ok := m[a[i]]
        if ok {
            sim += cnt * a[i] 
        }
    }
    log.Printf("res1: %d\n", int64(dif))
    log.Printf("res2: %d\n", sim)
}

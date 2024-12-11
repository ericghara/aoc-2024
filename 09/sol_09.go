package main

import (
    "log"
    "aoc/parseutil"
    bh "github.com/emirpasic/gods/trees/binaryheap"
    "fmt"
)


func main() {

    disk := parseutil.ToRunes(parseutil.ReadInputLines()[0])
    fmt.Println(len(disk))
    for i := range len(disk) {
       disk[i] = disk[i]-'0' 
    }

    free := [10]bh.Heap{}
    fs := []int{}
    
    buildFs := func() {
        for i := range 10 {
            free[i] = *bh.NewWithIntComparator()
        }
        fs = []int{}

        for i, size := range disk {
            var token int
            if (i&1) == 0 {
                token = i/2
            } else {
                token = -1
                free[int(size)].Push(len(fs))
            }
            for j := 0; j < int(size); j++ {
                fs = append(fs, token)
            }
        }
    }

    allocate := func(i, size int) int {
        pos, blkSize := i, -1

        for j := size; j < len(free); j++ {
            if val, ok := free[j].Peek(); ok && val.(int) < pos {
                blkSize = j
                pos = val.(int)
            }
        }

        if pos < i {
            free[blkSize].Pop()
            reSize := blkSize - size
            free[reSize].Push(pos+size)
        }
        return pos
    }

    score := func() int {
        hash := 0
        for i, v := range fs {
            if v > 0 {
                hash += i * v
            }
        }
        return hash
    }

    partOne := func() {
        for i := len(fs)-1; i >= 0; i-- {
            if fs[i] < 0 {
                continue
            }
            pos := allocate(i, 1)
            if pos == i {
                break
            }
            fs[pos] = fs[i]
            fs[i] = -2
        } 
    }

    partTwo := func() {
        for i := len(fs)-1; i >= 0; {
            if fs[i] < 0 {
                i--
                continue
            }
            id := fs[i]
            j := i
            for ; j >= 0 && fs[j] == id; j-- {}

            pos := allocate(i, i-j)
            if pos == i {
               i = j 
            }
            for i > j  {
                fs[pos] = id
                fs[i] = -2
                pos++
                i--
            }
        }
    }

    buildFs()
    partOne()
    log.Println("Part 1:", score())

    buildFs()
    partTwo()
    log.Println("Part 2:", score())
}

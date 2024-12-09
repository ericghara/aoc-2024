package main

import (
    "log"
    "aoc/parseutil"
    "fmt"
)

func main() {

    disk := parseutil.ToRunes(parseutil.ReadInputLines()[0])
    fmt.Println(len(disk))
    for i := range len(disk) {
       disk[i] = disk[i]-'0' 
    }

    var hash int
    parity := 0
    r := len(disk) - 2 + (len(disk)&1)
    toCompact := int(disk[r]) //remaining

    log.Println(r)
    pos := 0
    
    for l :=0; l < r; l++ {
        if parity == 0 {
            for i := 0; i < int(disk[l]); i++ {
                hash += pos * l/2 
                fmt.Print(l/2)
                pos++
            }
        } else {
            free := int(disk[l])
            for free > 0 && l < r {
                if toCompact == 0 {
                    r -= 2
                    toCompact = int(disk[r])
                } else {
                    for toCompact > 0 && free > 0 {
                        hash += pos * r/2 
                        toCompact--
                        free--
                        fmt.Print(r/2)
                        pos++
                    }
                }

            }
        }
        parity ^= 1
    }
    for ;toCompact > 0; toCompact-- {
        hash += pos * r/2
        pos++
        fmt.Print(r/2)
    }
    fmt.Println()

    log.Println("hash", hash)
}

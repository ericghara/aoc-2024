package main

import (
    "os"
    "log"
    "regexp"
    "strconv"
)

func main() {
    fileName := "./input"
    if len(os.Args) > 1 {
        fileName = os.Args[1]
    }

    b, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal("unable to open: ", fileName)
    }
    data := string(b)
    matcher := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)
    matches := matcher.FindAllStringSubmatch(data, -1)
    var tot int64
    var condTot int64
    do := true
    for _, match := range matches {
        if match[0] == "do()" {
            do = true
        } else if match[0] == "don't()" {
            do = false
        } else {
            a, errA := strconv.ParseInt(match[1], 10, 64)
            b, errB := strconv.ParseInt(match[2], 10, 64)
            if errA != nil || errB != nil {
                log.Println("Unable to parse line: ", match[0])
                continue
            }
            tot += a*b
            if do {
                condTot += a*b
            }
        }
    }
    log.Println("Tot", tot)
    log.Println("Cond Tot", condTot)
}

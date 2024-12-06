package main

import (
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
    "math"
    "slices"
)

func isSafe(nums []int) (bool, int) {
    mn := math.MaxInt64
    mx := math.MinInt64
    last := nums[0]
    for i, num := range nums[1:] {
        d := num - last
        mn = min(mn, d)
        mx = max(mx, d)
        if !(mx <= -1 && mn >= -3 || mn >= 1 && mx <= 3) {
            return false, i+1
        }
        last = num
    }
    return true, -1
}

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

    scanner := bufio.NewScanner(file)
    safe := 0
    dSafe := 0
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        numStrs := strings.Split(line, " ")
        var nums []int
        for _, numStr := range numStrs {
            num, err := strconv.Atoi(numStr)
            if err != nil {
                log.Print("Unable to parse: ", numStr)
                continue;
            }
            nums = append(nums, num) 
        }
        good, i := isSafe(nums)
        if good {
            safe++
            dSafe++
        } else if good, _ = isSafe(slices.Concat(nums[:i], nums[i+1:])); good {
            dSafe++
        } else if good, _ = isSafe(slices.Concat(nums[:i-1], nums[i:])); good {
            dSafe++
        } else if good, _ = isSafe(nums[1:]); good {
            dSafe++
        }
    }
    log.Println("Num Safe:", safe)
    log.Println("Num Damper Safe:", dSafe)
}

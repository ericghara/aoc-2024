package main

import (
    "log"
    "strconv"
    "bufio"
    "regexp"
    "strings"
    "fmt"
    "aoc/parseutil"
)

func main() {

    file := parseutil.OpenInput()
    defer file.Close()

    var isValid, isValidConcat func(int64, int64, []int64) bool

    isValid = func(target, curVal int64, nums []int64) bool {
       if len(nums) == 0 {
            return curVal == target
        }
        if curVal > target {
            return false
        }
        return isValid(target, curVal+nums[0], nums[1:]) || isValid(target, curVal*nums[0], nums[1:])
    }

    seen := map[[2]int64]bool{}

    isValidConcat = func(target, curVal int64, nums []int64) bool {
       if len(nums) == 0 {
            return curVal == target
        }
        state := [2]int64{int64(len(nums)), curVal}
        if curVal > target || seen[state] {
            return false
        }
        seen[state] = true
        concat := curVal
        for i, num := range nums {
            concat, _ = strconv.ParseInt(fmt.Sprintf("%v%v", concat, num), 10, 64)
            if isValidConcat(target, concat, nums[i+1:]) {
                return true
            }
        }
        return isValidConcat(target, curVal+nums[0], nums[1:]) || isValidConcat(target, curVal*nums[0], nums[1:])
    }

    scanner := bufio.NewScanner(file)
    var totValid, totValidConcat int64
    splitter :=  regexp.MustCompile(`:?\s`)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        splits := splitter.Split(line, -1)
        if (len(splits) < 2) {
            log.Fatal("Unable to parse line.")
        }
        ints := parseutil.ParseInts(splits)
        if isValid(ints[0], ints[1], ints[2:]) {
            totValid += ints[0]
            totValidConcat += ints[0]
        } else if clear(seen); isValidConcat(ints[0], ints[1], ints[2:]) {
            totValidConcat += ints[0]
        }
    }
    log.Println("Tot. Valid", totValid)
    log.Println("Tot. Valid Concat", totValidConcat)
}

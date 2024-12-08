package parseutil 

import (
    "log"
    "os"
    "strconv"
    "bufio"
    "strings"
    "regexp"
)

func OpenInput() *os.File {
    fileName := "./input"
    if len(os.Args) > 1 {
        fileName = os.Args[1]
    }

    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal("unable to open: ", fileName)
    }
    return file
}

func ReadInputLines() []string {
    file := OpenInput()
    defer file.Close()
    scanner := bufio.NewScanner(file)
    lines := make([]string, 0)
    
    for scanner.Scan() {
        lines = append(lines, strings.TrimSpace(scanner.Text()))
    }
 
    if len(lines) == 0 {
        log.Println("Empty input.")
    } else if len(lines[len(lines)-1]) == 0 {
        log.Println("Last line is blank.")
    }
    return lines
}

// ReadInputSections reads from the input using [OpenInput].
// Each slice of the output corresponds to a section where each section is
// composed of a slice of lines in the section.  A section break is detected by
// a regex match without any pre-processing to `breakRegex`; the section break line is discarded.
// All lines in the section are trimmed of leading and trailing whitespace.
func ReadInputSections(breakRegex string) [][]string {
    file := OpenInput()
    defer file.Close()
    scanner := bufio.NewScanner(file)
    sections := make([][]string, 0)
    sections = append(sections, make([]string, 0))

    breakMatch := regexp.MustCompile(breakRegex)

    for scanner.Scan() {
        line := scanner.Text()
        if breakMatch.MatchString(line) {
            sections = append(sections, make([]string, 0))
        } else {
            sections[len(sections)-1] = append(sections[len(sections)-1], strings.TrimSpace(scanner.Text()))
        }
    }
    return sections
}

func ParseInts(intStrs []string) []int64 {
    ints := make([]int64, 0)
    for _, str := range intStrs {
        num, err := strconv.ParseInt(str, 10, 64)
        if err != nil {
            log.Fatal("Unable to parse: ", str)
        }
        ints = append(ints, num)
    }
    return ints
}

func ToRunes(s string) []rune {
    runes := make([]rune, 0)
    for _, r := range s {
        runes = append(runes, r)
    }
    return runes
}

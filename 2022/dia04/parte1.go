package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"aocutils"
)

const (
    first = 0
    second = iota
)
type assignment struct {
    start int
    end int
}

func is_overlapping(a, b assignment) bool {
    return a.start <= b.start && a.end >= b.end
}

func main() {
    file, err := os.Open("./biglist.txt")
    if err != nil {
        panic(err)
    }

    // aocutils.Hello_aoc()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var pairs []string
    for scanner.Scan() {
	pairs = append(pairs, scanner.Text())
    }
    file.Close()

    r, _ := regexp.Compile("\\d+")

    overlapping := 0
    var sections [][]assignment

    for i := range pairs {
	var pair []assignment

	extract := r.FindAllString(pairs[i], -1)

	var start int
	var end int

	start, _ = strconv.Atoi(extract[0])
	end, _ = strconv.Atoi(extract[1])

	pair = append(pair, assignment{start, end})

	start, _ = strconv.Atoi(extract[2])
	end, _ = strconv.Atoi(extract[3])

	pair = append(pair, assignment{start, end})

	sections = append(sections, pair)
    }

    for i := range sections {
	if is_overlapping(sections[i][first], sections[i][second]) ||
	is_overlapping(sections[i][second], sections[i][first]) {
	    overlapping += 1
	}
    }
    fmt.Println(overlapping)

}

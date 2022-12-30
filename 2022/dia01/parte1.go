package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    file, err := os.Open("./calories.txt")
    if err != nil {
	panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    sum := 0
    max := 0

    for scanner.Scan() {
	line := scanner.Text()
	if line == "" {
	    sum = 0
	    continue
	}
	num, err := strconv.Atoi(line)
	if err != nil {
	    panic(err)
	}

	sum += num
	if sum > max {
	    max = sum
	}
    }
    fmt.Println(max)
}

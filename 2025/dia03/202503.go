package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func f(line string, n int) int {
	for len(line) > n {
		for i := range line {
			if line[i] < line[i+1] {
				line = line[:i] + line[i+1:]
				break
			}
			if i == len(line)-2 {
				line = line[:n]
				break
			}
		}
	}
	int, _ := strconv.Atoi(line)
	return int
}

func main() {
	b, err := os.ReadFile("03.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	pt1, pt2 := 0, 0
	for _, line := range lines {
		pt1 += f(line, 2)
		pt2 += f(line, 12)
	}
	fmt.Println(pt1)
	fmt.Println(pt2)
}

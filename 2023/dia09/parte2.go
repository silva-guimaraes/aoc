package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func recur(a []int) int {
    zeros := true
    for _, b := range a {
        if b != 0 {
            zeros = false
            break
        }
    }
    if zeros {
        return 0
    }

    c := make([]int, len(a) - 1)

    for i := 0; i < len(a)-1; i++ {
        c[i] = a[i] - a[i+1]
    }

    return a[0] + recur(c)
}

func main() {
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var lines [][]int
    for scanner.Scan() {
        var line []int
        for _, a := range strings.Split(scanner.Text(), " ") {
            b, _ := strconv.Atoi(a)
            line = append(line, b)
        }
        lines = append(lines, line)
    }
    file.Close()

    sum := 0
    for _, l := range lines  {
        sum += recur(l)
    }
    fmt.Println(sum)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
    file, err := os.Open("./instructions.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    // ler linhas como strings
    var lines []string
    for scanner.Scan() {
	lines = append(lines, scanner.Text())
    }
    file.Close()

    r, _ := regexp.Compile("-?\\d+")

    ctr := make([][]byte, 6)
    for i := range ctr {
	ctr[i] = make([]byte, 40)
    }

    cycle := 1
    register := 1
    var register_values []int
    // signal_strength_sum := 0

    for i := range lines {
	register_values = append(register_values, register)
	if lines[i] == "noop" {
	    cycle++

	} else {
	    ret, _ := strconv.Atoi(r.FindString(lines[i]))
	    cycle++
	    register_values = append(register_values, register)
	    register += ret
	    cycle++
	}

    }

    for i := range register_values {
	if i % 40 == register_values[i] || i % 40 == register_values[i] + 1 || i % 40 == register_values[i] - 1 {
	    ctr[int64(math.Floor(float64(i) / 40))][i % 40] = byte('#') // Go não é pra preguiçoso!!!!!

	} else {
	    ctr[int64(math.Floor(float64(i) / 40))][i % 40] = byte('.')
	}
    }

    for i := range ctr {
	for j := range ctr[i] {
	    fmt.Print(string(ctr[i][j]))
	}
	fmt.Println()
    }

}

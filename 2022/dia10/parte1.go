package main

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
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

    cycle := 1
    register := 1
    signal_strength_sum := 0
    for i := range lines {
	if (cycle - 20) % 40 == 0 {
	    signal_strength_sum += register * cycle
	    fmt.Println(register * cycle)
	}
	if lines[i] == "noop" {
	    cycle++
	} else {
	    ret, _ := strconv.Atoi(r.FindString(lines[i]))
	    cycle++
	    if (cycle - 20) % 40 == 0 {
		signal_strength_sum += register * cycle
		fmt.Println(register * cycle)
	    }
	    register += ret
	    cycle++
	}

    }
    fmt.Println(signal_strength_sum)
}

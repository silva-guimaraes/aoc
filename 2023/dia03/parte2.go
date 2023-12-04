package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pos struct {
    i, j int
}

type num struct {
    // pos []pos
    num string
    visited bool
}

func main() {
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var lines [][]byte
    for scanner.Scan() {
        lines = append(lines, []byte(scanner.Text()))
    }
    file.Close()

    var numbers []*num
    lookup := make(map[pos]int)

    for i := range lines {
        for j := 0; j < len(lines[i]); j++ {
            c := lines[i][j]
            if c >= '0' && c <= '9' {
                n := num{}
                for k := j; k < len(lines[i]); k++ {
                    c = lines[i][k]
                    if c >= '0' && c <= '9'{
                        pos := pos{i, k}
                        // n.pos = append(n.pos, pos)
                        n.num += string(c)
                        lookup[pos] = len(numbers)
                        j++
                    } else {
                        break
                    }
                }
                numbers = append(numbers, &n)
            }
        }
    }

    sum := 0
    for i := range lines {
        for j, c := range lines[i] {
            if c != '*' {
                continue
            }
            var parts []*num
            for i2 := i - 1; i2 <= i+1; i2++ {
                for j2 := j - 1; j2 <= j+1; j2++ {
                    if i2 < 0 || i2 >= len(lines) || j2 < 0 || j2 >= len(lines[0]) {
                        continue
                    }
                    p1 := pos{i2, j2}
                    p2i, ok := lookup[p1]
                    if !ok {
                        continue
                    }
                    p2 := numbers[p2i]
                    if p2.visited {
                        continue
                    }
                    p2.visited = true
                    parts = append(parts, p2)
                } 
            } 
            if len(parts) == 2 {
                num1, _ := strconv.Atoi(parts[0].num)
                num2, _ := strconv.Atoi(parts[1].num)

                sum += num1 * num2
            }
            for _, p := range parts {
                p.visited = false
            }
        }
    }
    fmt.Println(sum)
}

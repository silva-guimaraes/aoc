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
    pos []pos
    num string
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

    var numbers []num

    for i := range lines {
        for j := 0; j < len(lines[i]); j++ {
            c := lines[i][j]
            if c >= '0' && c <= '9' {
                n := num{}
                for k := j; k < len(lines[i]); k++ {
                    c = lines[i][k]
                    if c >= '0' && c <= '9'{
                        n.pos = append(n.pos, pos{i, k})
                        n.num += string(c)
                        j++
                    } else {
                        break
                    }
                }
                numbers = append(numbers, n)
            }
        }
    }

    sum := 0
    for i := range lines {
        for j, c := range lines[i] {
            if c != '*' {
                continue
            }
            numbers2 := make([]num, len(numbers))
            copy(numbers2, numbers)
            var parts []num
            for i2 := i - 1; i2 <= i+1; i2++ {
                for j2 := j - 1; j2 <= j+1; j2++ {
                    if i2 < 0 || i2 >= len(lines) || j2 < 0 || j2 >= len(lines[0]) {
                        continue
                    }
                    p := pos{i2, j2}
                    for k := 0; k < len(numbers2); k++ {
                        for _, p2 := range numbers2[k].pos {
                            if  p == p2 {
                                parts = append(parts, numbers2[k])
                                numbers2[k] = numbers2[len(numbers2)-1]
                                numbers2 = numbers2[:len(numbers2)-1]
                                break;
                            }
                        }
                    }
                } 
            } 
            if len(parts) == 2 {
                num1, _ := strconv.Atoi(parts[0].num)
                num2, _ := strconv.Atoi(parts[1].num)

                sum += num1 * num2
            }
        }
    }
    fmt.Println(sum)
}

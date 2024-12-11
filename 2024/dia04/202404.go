package main

import (
	"os"
    "fmt"
	"strings"
)

type pos struct {
    i, j int
}

type direction []pos


// demorei muito mais do que deveria

func part1(lines []string) {

    // péssimo
    directions := []direction {
        { {0,0},{-1,0},{-2,0},{-3,0} },
        { {-0,0},{-1,1},{-2,2},{-3,3} },
        { {0,0},{0,1},{0,2},{0,3}, },
        { {0,0},{1,1},{2,2},{3,3}, },
        { {0,0},{1,0},{2,0},{3,0}, },
        { {0,0},{1,-1},{2,-2},{3,-3}, },
        { {0,0},{0,-1},{0,-2},{0,-3}, },
        { {0,0},{-1,-1},{-2,-2},{-3,-3}, },
    }
    sum := 0
    for i, line := range lines {
        for j := range line {
            for _, cardinality := range directions {
                d := make([]byte, 4)
                for k, diff := range cardinality {
                    di := i+diff.i
                    dj := j+diff.j
                    if di < 0 || di >= len(lines) || dj < 0 || dj >= len(lines[0]) {
                        continue
                    }
                    d[k] = lines[di][dj]
                }
                if string(d) == "XMAS" {
                    sum += 1
                }
            }

        }
    }
    fmt.Println(sum)

}

func part2(lines []string) {
    directions := []direction {
        {{-1,-1},{0,0},{1,1}},
        {{-1,1},{0,0},{1,-1}},
    }
    sum := 0
    for i, line := range lines {
        for j := range line {
            mas := 0
            for _, cardinality := range directions {
                d := make([]byte, 3)
                for k, diff := range cardinality {
                    di := i+diff.i
                    dj := j+diff.j
                    if di < 0 || di >= len(lines) || dj < 0 || dj >= len(lines[0]) {
                        continue
                    }
                    d[k] = lines[di][dj]
                }
                if string(d) == "MAS" || string(d) == "SAM" {
                    mas += 1
                }
            }
            if mas == 2 {
                sum += 1
            }
        }
    }
    fmt.Println(sum)

}

func main() {
    bytes, err := os.ReadFile("./04.txt")
    if err != nil {
        panic(err)
    }
    input := strings.TrimSpace(string(bytes))
    lines := strings.Split(input, "\n")

    part1(lines)
    part2(lines)
}

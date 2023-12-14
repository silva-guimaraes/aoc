package main


import (
    "fmt"
    "os"
    "strings"
    "slices"
)

type pos struct {
    i, j int
}

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(file), "\n")


    var roundRocks []pos
    var cubeRocks []pos

    for i := range lines {
        for j, c := range lines[i] {
            if c == 'O' {
                roundRocks = append(roundRocks, pos{i, j})
            }
            if c == '#' {
                cubeRocks = append(cubeRocks, pos{i, j})
            }
        }
    }

    for i, p := range roundRocks {
        if p.i == 0 {
            continue
        }

        for {
            p.i--

            idx1 := slices.Index(roundRocks, p)
            idx2 := slices.Index(cubeRocks, p)
            if idx1 > -1 || idx2 > -1 {
                roundRocks[i] = pos{p.i+1, p.j}
                break
            }

            if p.i == 0 {
                roundRocks[i] = p
                break
            }
            if p.i == -1 {
                roundRocks[i] = pos{i+i, p.j}
                break
            }
        }
        
    }

    sum := 0
    for _, p := range roundRocks {
        sum += len(lines) - 1 - p.i

    }

    fmt.Println(sum)
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    dirs := map[byte][2]int{
        'R': { 0,  1},
        'L': { 0, -1},
        'U': {-1,  0},
        'D': { 1,  0},
    }

    cubes := [][2]int{
        {0, 0},
    }

    for _, line := range strings.Split(strings.TrimSpace(string(file)), "\n") {

        split := strings.Split(line, " ")

        dir := dirs[split[0][0]]
        times, err := strconv.Atoi(split[1])
        if err != nil {
            panic(err)
        }

        for i := 0; i < times; i++ {
            last := cubes[len(cubes)-1]
            cubes = append(cubes, [2]int{last[0] + dir[0], last[1] + dir[1]})
        }
    }

    maxJ, maxI, minJ, minI := 0, 0, 0, 0
    for _, c := range cubes {
        maxJ = max(maxJ, c[1])
        maxI = max(maxI, c[0])
        minJ = min(minJ, c[1])
        minI = min(minI, c[0])
    }

    grid := make([][]byte, maxI+1 - minI)
    for i := range grid {
        grid[i] = make([]byte, maxJ+1 - minJ)
        for j := range grid[i] {
            grid[i][j] = '.'
        }
    }
    for _, c := range cubes {
        grid[c[0] - minI][c[1] - minJ] = '#'
    }

    inside := false
    count := 0
    for i := range grid {
        for j, c := range grid[i] {
            if i > 0 && c == '#' && grid[i-1][j] == '#' {
                inside = !inside
            }

            if c == '#' || inside {
                count++
            }
        }
    }

    fmt.Println(count)
}

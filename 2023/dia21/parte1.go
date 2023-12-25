package main

import (
    "fmt"
    "os"
    "bytes"
    "slices"
)

type pos struct {
    i, j int
}

type path struct {
    pos pos
    count int
}

var directions []pos = []pos{
    {1, 0},
    {-1, 0},
    {0, 1},
    {0, -1},
}

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    grid := bytes.Split(bytes.TrimSpace(file), []byte("\n"))

    var elf path

    for i := range  grid {
        for j, c := range grid[i] {
            if c == 'S' {
                elf = path{pos{i, j}, 0}
                grid[i][j] = '.'
                goto end
            }
        }
    }
    end:

    queue := []path{elf}
    visited := map[pos]bool{}

    for len(queue) > 0 {

        a := queue[0]
        queue = queue[1:]

        visited[pos{a.pos.i, a.pos.j}] = true

        if a.count == 64 {
            continue
        }

        for _, d := range directions {

            n := pos{a.pos.i + d.i, a.pos.j + d.j}

            if n.i < 0 || n.i >= len(grid) || n.j < 0 || n.j >= len(grid[0]) {
                continue
            }
            if grid[n.i][n.j] == '#' {
                continue
            }
            if _, exists := visited[n]; exists {
                continue
            }
            idx := slices.IndexFunc(queue, func(x path) bool {
                return x.pos == n
            })
            if idx > -1 {
                continue
            }
            queue = append(queue, path{n, a.count+1})
        }
    }

    count := 0
    for i := range grid {
        for j := range grid[i] {
            if (i + j) % 2 != 0 {
                continue
            }
            if _, ok := visited[pos{i, j}]; ok {
                count++
            }
        }
    }
    fmt.Println(count)

}

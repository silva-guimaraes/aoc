package main

import (
	// "cmp"
	"fmt"
	"math"
	"os"
	// "slices"
	"strings"
)

type pos struct {
    i, j int
}

type path struct {
    pos pos
    count int
    prev pos
}

var directions []pos = []pos{
    {0, 1},
    {0, -1},
    {1, 0},
    {-1, 0},
}

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    grid := strings.Split(strings.TrimSpace(string(file)), "\n")

    end := pos{len(grid)-1, len(grid[0])-2}
    start := pos{0, 1}

    score := make([][]float64, len(grid))
    for i := range score {
        score[i] = make([]float64, len(grid[0]))
        for j := range score[i] {
            score[i][j] = math.Inf(-1)
        }
    }

    score[start.i][start.j] = 0

    queue := []path{
        {start, 0, start},
    }

    for len(queue) > 0 {
        a := queue[0]
        queue = queue[1:]

        for _, d := range directions {
            n := pos{d.i + a.pos.i, d.j + a.pos.j}
            if n.i < 0 || n.i >= len(grid) || n.j < 0 || n.j >= len(grid[0]) {
                continue
            }

            if n == a.prev {
                continue
            }

            path := path{n, a.count+1, a.pos}
            
            switch grid[n.i][n.j] {
            case '#':
                continue
            case '>':
                if d != directions[0] {
                    continue
                }
            case '<':
                if d != directions[1] {
                    continue
                }
            case 'v':
                if d != directions[2] {
                    continue
                }
            case '^':
                if d != directions[3] {
                    continue
                }
            }

            try := score[a.pos.i][a.pos.j]+1
            if try > score[n.i][n.j] {
                score[n.i][n.j] = try
                queue = append(queue, path)
            }

        }
    }

    fmt.Println(score[end.i][end.j])

}

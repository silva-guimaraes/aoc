package main

import (
	"fmt"
	"os"
	"strings"
)

type pos struct {
    i, j int
}

var directions []pos = []pos{
    {-1, 0},
    {0, -1},
    {0, 1},
    {1, 0},
}

type direction int

const (
    north direction = iota
    west
    east
    south
)

type beam struct {
    direction direction
    pos pos
}

func addPos(a pos, d direction) pos {
    return pos{a.i + directions[d].i, a.j + directions[d].j}
}

func simulate(grid []string, b beam) int {
    queue := []beam{ b }

    visited := make(map[beam]bool)

    for len(queue) > 0 {
        a := queue[0]
        queue = queue[1:]

        if a.pos.i < 0 || a.pos.i >= len(grid) || 
        a.pos.j < 0 || a.pos.j >= len(grid[0]) {
            continue
        }

        if _, ok := visited[a]; ok {
            continue
        }

        switch grid[a.pos.i][a.pos.j]  {
        case '.':
            queue = append(queue, beam{a.direction, addPos(a.pos, a.direction)})
        case '|':
            if a.direction == west || a.direction == east {
                queue = append(queue, beam{north, addPos(a.pos, north)})
                queue = append(queue, beam{south, addPos(a.pos, south)})
            } else {
                queue = append(queue, beam{a.direction, addPos(a.pos, a.direction)})
            }
        case '-':
            if a.direction == north || a.direction == south {
                queue = append(queue, beam{east, addPos(a.pos, east)})
                queue = append(queue, beam{west, addPos(a.pos, west)})
            } else {
                queue = append(queue, beam{a.direction, addPos(a.pos, a.direction)})
            }
        case '\\':
            switch a.direction {
            case east:
                queue = append(queue, beam{south, addPos(a.pos, south)})
            case south:
                queue = append(queue, beam{east, addPos(a.pos, east)})
            case west:
                queue = append(queue, beam{north, addPos(a.pos, north)})
            case north:
                queue = append(queue, beam{west, addPos(a.pos, west)})
            default:
                panic(a.direction)
            }
        case '/':
            switch a.direction {
            case east:
                queue = append(queue, beam{north, addPos(a.pos, north)})
            case south:
                queue = append(queue, beam{west, addPos(a.pos, west)})
            case west:
                queue = append(queue, beam{south, addPos(a.pos, south)})
            case north:
                queue = append(queue, beam{east, addPos(a.pos, east)})
            default:
                panic(a.direction)
            }
        }

        visited[a] = true
    }

    count := make(map[pos]bool)

    for i := range visited {
        count[i.pos] = true
    }

    return len(count)

}

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    grid := strings.Split(strings.TrimSpace(string(file)), "\n")

    // cache := map[beam]int

    countMax := 0
    for j := range grid[0] {
        countMax = max(countMax, simulate(grid, beam{south, pos{0, j}}))
        countMax = max(countMax, simulate(grid, beam{north, pos{len(grid)-1, j}}))
    }
    for i := range grid {
        countMax = max(countMax, simulate(grid, beam{east, pos{i, 0}}))
        countMax = max(countMax, simulate(grid, beam{west, pos{i, len(grid[0])-1}}))
    }
    fmt.Println(countMax)
}

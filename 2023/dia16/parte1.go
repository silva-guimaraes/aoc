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

var directions []pos = []pos{
    {-1, 0}, // norte
    {0, -1}, // oeste
    {0, 1}, // leste
    {1, 0}, // sul
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

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    grid := strings.Split(strings.TrimSpace(string(file)), "\n")

    queue := []beam{
        {east, pos{0, 0}},
    }

    var visited []beam

    for len(queue) > 0 {
        a := queue[0]
        queue = queue[1:]

        if a.pos.i < 0 || a.pos.i >= len(grid) || 
        a.pos.j < 0 || a.pos.j >= len(grid[0]) {
            continue
        }

        if slices.Contains(visited, a) {
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

        visited = append(visited, a)
    }

    count := make(map[pos]bool)

    for _, b := range visited {
        count[b.pos] = true
    }

    fmt.Println(len(count))
    // fmt.Println(visited)
}

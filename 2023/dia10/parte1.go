package main


import (
	"bufio"
	"fmt"
	"os"
    "slices"
)

type pos struct {
    i, j int
}

type path struct {
    current pos
    number int
    before pos
}

type connections []pos


func main() {
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    file.Close()

    lookup := make([][]connections, len(lines))
    for i := range lookup {
        lookup[i] = make([]connections, len(lines[0]))
    }
    var s pos
    for i := range lines {
        for j, x := range lines[i] {

            switch x {
            case '|':
                lookup[i][j] = connections{
                    pos{i+1, j},
                    pos{i-1, j},
                }
            case '-':
                lookup[i][j] = connections{
                    pos{i, j+1},
                    pos{i, j-1},
                }
            case 'L':
                lookup[i][j] = connections{
                    pos{i-1, j},
                    pos{i, j+1},
                }
            case 'J':
                lookup[i][j] = connections{
                    pos{i-1, j},
                    pos{i, j-1},
                }
            case '7':
                lookup[i][j] = connections{
                    pos{i, j-1},
                    pos{i+1, j},
                }
            case 'F':
                lookup[i][j] = connections{
                    pos{i, j+1},
                    pos{i+1, j},
                }
            case '.':
                lookup[i][j] = nil
            case 'S':
                s = pos{i, j}
            default:
                panic(x)
            }
        }
    }

    queue := []path{
        {pos{s.i, s.j+1}, 1, s},
        {pos{s.i, s.j-1}, 1, s},
        {pos{s.i-1, s.j}, 1, s},
        {pos{s.i+1, s.j}, 1, s},
    }

    visited := []pos{ s }
    for len(queue) > 0 {
        a := queue[0]
        queue = queue[1:]
        if a.current.i < 0 || a.current.i >= len(lines) ||
        a.current.j < 0 || a.current.j >= len(lines[0]) {
            continue
        }
        pipes := lookup[a.current.i][a.current.j]
        if pipes == nil {
            continue
        }

        i := slices.Index(pipes, a.before)

        if i == -1 {
            continue
        }

        for _, b := range visited {
            if a.current == b {
                fmt.Println(a.number)
                return
            }

        }

        next := pipes[1 - i]
        queue = append(queue, path{next, a.number+1, a.current})

        visited = append(visited, a.current)
    }
}

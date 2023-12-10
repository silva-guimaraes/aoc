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
    // directions int
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

        for _, b := range visited {
            if a.current == b {
                goto end
            }

        }

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


        next := pipes[1 - i]
        queue = append(queue, path{next, a.number+1, a.current})

        visited = append(visited, a.current)
    }
    end:
    var loop, temp []pos
    for i := 0; i < len(visited); i += 2 {
        loop = append(loop, visited[i])
    }
    for i := 1; i < len(visited); i += 2 {
        temp = append(temp, visited[i])
    }

    slices.Reverse(temp)
    loop = append(loop, temp...)
    loop = append(loop, s)

    directions := make([]pos, len(loop) - 1)

    for i := 0; i < len(loop) -1 ; i++ {
        directions[i] = pos{ loop[i+1].i - loop[i].i, loop[i+1].j - loop[i].j }
    }

    loop = loop[:len(loop)-1]

    sumi, sumj := 0, 0
    for i := 0; i < len(directions) - 1; i++ {
        sumi += directions[i].i
        sumj += directions[i].j
    }
    // sumj == 1 esquerda -1 == direita

    var innerTiles []pos
    sumj = -sumj

    for i := range loop {
        switch directions[i] {
        case pos{ 1,  0}:
            innerTiles = append(innerTiles, pos{loop[i].i, loop[i].j + sumj})
        case pos{-1,  0}:
            innerTiles = append(innerTiles, pos{loop[i].i, loop[i].j - sumj})
        case pos{ 0, -1}:
            innerTiles = append(innerTiles, pos{loop[i].i + sumj, loop[i].j})
        case pos{ 0,  1}:
            innerTiles = append(innerTiles, pos{loop[i].i - sumj, loop[i].j})
        default:
            panic(directions[i])
        }
    }

    for i := 0; i < len(innerTiles); i++ {

        if lines[innerTiles[i].i][innerTiles[i].j] != '.' ||
        innerTiles[i].i < 0 || innerTiles[i].i >= len(lines) ||
        innerTiles[i].j < 0 || innerTiles[i].j >= len(lines[0]) {
            innerTiles[i] = innerTiles[len(innerTiles)-1]
            innerTiles = innerTiles[:len(innerTiles)-1]
            i--
        }
    }
    fuckYouRob := make(map[pos]byte)
    for i := range innerTiles {
        fuckYouRob[innerTiles[i]] = 0
    }

    innerTiles = []pos{}
    var bleed []pos
    for k := range fuckYouRob {
        bleed = append(bleed, k)
    }

    for len(bleed) > 0 {
        a := bleed[0]
        bleed = bleed[1:]

        if i := slices.Index(loop, a); i > -1 {
            continue
        }

        if i := slices.Index(innerTiles, a); i > -1 {
            continue
        }

        fmt.Println(a)

        bleed = append(bleed, pos{a.i+1, a.j})
        bleed = append(bleed, pos{a.i, a.j+1})
        bleed = append(bleed, pos{a.i-1, a.j})
        bleed = append(bleed, pos{a.i, a.j-1})

        innerTiles = append(innerTiles, a)
    }
    fmt.Println(len(innerTiles))
    // fmt.Println(innerTiles)
    // fmt.Println(sumj)

}

package main


import (
	"bufio"
	"fmt"
	"os"
    "slices"
    "os/exec"
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

    // sumi, sumj := 0, 0
    // for i := 0; i < len(directions) - 1; i++ {
    //     sumi += directions[i].i
    //     sumj += directions[i].j
    // }
    // sumj == 1 esquerda -1 == direita
    sumj := 1
    // sumj := -1

    var innerTiles []pos

    for i := range loop {
        // p1 := pos{67, 59}
        var p pos

        switch directions[i] {
        case pos{ 1,  0}:
            p = pos{loop[i].i, loop[i].j + sumj}
        case pos{-1,  0}:
            p = pos{loop[i].i, loop[i].j - sumj}
        case pos{ 0, -1}:
            p = pos{loop[i].i + sumj, loop[i].j}
        case pos{ 0,  1}:
            p = pos{loop[i].i - sumj, loop[i].j}
        default:
            panic(directions[i])
        }

        innerTiles = append(innerTiles, p)
    }

    // gambiarra adiante!
    fuckYouRob := make(map[pos]byte)
    for i := range innerTiles {
        fuckYouRob[innerTiles[i]] = 0
    } 
    innerTiles = []pos{}
    var bleed []pos
    for k := range fuckYouRob {
        bleed = append(bleed, k)
    }
    // innerTiles = bleed

    for len(bleed) > 0 {
        a := bleed[0]
        bleed = bleed[1:]

        if  a.i < 0 || a.i >= len(lines) || a.j < 0 || a.j >= len(lines[0]) {
            continue
        }

        if i := slices.Index(loop, a); i > -1 {
            continue
        }

        if i := slices.Index(innerTiles, a); i > -1 {
            continue
        }


        bleed = append(bleed, pos{a.i+1, a.j})
        bleed = append(bleed, pos{a.i, a.j+1})
        bleed = append(bleed, pos{a.i-1, a.j})
        bleed = append(bleed, pos{a.i, a.j-1})

        innerTiles = append(innerTiles, a)
    }
    fmt.Println(len(innerTiles))

    return


    cmd := exec.Command(
        "ffmpeg", 
        "-hide_banner",
        "-y",
        "-f", "rawvideo",
        "-pix_fmt", "rgba",
        "-s", fmt.Sprintf("%dx%d", len(lines[0]), len(lines)),
        // "-r", fmt.Sprint(FPS),
        // "-an",
        "-i", "-", 
        // "-vf", "scale=400:-1",
        // "-c:v", "libx264",
        "foobar.bmp",
    )
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    stdin, err := cmd.StdinPipe()
    if err != nil {
        panic(err)
    }

    writer := bufio.NewWriter(stdin)

    fmt.Println(cmd.String())
    cmd.Start()

    buf := make([]byte, len(lines) * len(lines[0]) * 4)

    for i := range lines {
        for j := range lines[i] {
            off := i*len(lines[0])*4 + j*4
            // buf[off + 3] = 255
            buf[off + 0] = 0
            buf[off + 1] = 0
            buf[off + 2] = 100
            buf[off + 3] = 255
        }
    }
    for _, p := range loop {
        off := p.i*len(lines[0])*4 + p.j*4

        buf[off + 0] = 0
        buf[off + 1] = 255
        buf[off + 2] = 0


        // switch lines[p.i][p.j] {
        // case 'F':
        // buf[off + 0] = 255
        // buf[off + 1] = 100
        // buf[off + 2] = 0
        // case 'J':
        // buf[off + 0] = 100
        // buf[off + 1] = 255
        // buf[off + 2] = 100
        // case 'L':
        // buf[off + 0] = 100
        // buf[off + 1] = 100
        // buf[off + 2] = 255
        // case '7':
        // buf[off + 0] = 255
        // buf[off + 1] = 100
        // buf[off + 2] = 255
        // case '|':
        // buf[off + 0] = 255
        // buf[off + 1] = 255
        // buf[off + 2] = 100
        // case '-':
        // buf[off + 0] = 100
        // buf[off + 1] = 255
        // buf[off + 2] = 255
        //
        // }

    }
    for _, p := range innerTiles {
        off := p.i*len(lines[0])*4 + p.j*4

        buf[off + 0] = 255
        buf[off + 1] = 0
        buf[off + 2] = 0

    }

    nn, err := writer.Write(buf); 
    fmt.Println(nn)
    if err != nil {
        panic(err)
    }


    stdin.Close()

    if cmd.Wait() != nil {
        panic(err)
    }

}

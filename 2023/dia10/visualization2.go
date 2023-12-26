package main

import (
	"bufio"
	"fmt"
	// "image/color"
	"os"
	"os/exec"
	"slices"
	// "unsafe"
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

type pixel struct {
    r, g, b, a byte
}

type pass struct {
    pos pos
    status bool
}

type frame [][]pixel

func sendFrame(ffmpeg *bufio.Writer, buf frame) {
    outBuffer := make([]byte, len(buf) * len(buf[0]) * 4)

    k := 0
    for i := range buf {
        for _, b := range buf[i] {
            outBuffer[k + 0] = b.r
            outBuffer[k + 1] = b.g
            outBuffer[k + 2] = b.b
            outBuffer[k + 3] = b.a
            k += 4
        }
    }
    if _, err := ffmpeg.Write(outBuffer); err != nil {
        panic(err)
    }
}

func colourTile(lines []string, buf [][]pixel, v pos, c pixel) {
    switch lines[v.i][v.j] {
    case '-':
        buf[v.i*3+1][v.j*3+0] = c
        buf[v.i*3+1][v.j*3+1] = c
        buf[v.i*3+1][v.j*3+2] = c
    case '|':
        buf[v.i*3+0][v.j*3+1] = c
        buf[v.i*3+1][v.j*3+1] = c
        buf[v.i*3+2][v.j*3+1] = c
    case 'L':
        buf[v.i*3+0][v.j*3+1] = c
        buf[v.i*3+1][v.j*3+1] = c
        buf[v.i*3+1][v.j*3+2] = c
    case 'J':
        buf[v.i*3+0][v.j*3+1] = c
        buf[v.i*3+1][v.j*3+1] = c
        buf[v.i*3+1][v.j*3+0] = c
    case 'F':
        buf[v.i*3+1][v.j*3+2] = c
        buf[v.i*3+1][v.j*3+1] = c
        buf[v.i*3+2][v.j*3+1] = c
    case '7':
        buf[v.i*3+1][v.j*3+0] = c
        buf[v.i*3+1][v.j*3+1] = c
        buf[v.i*3+2][v.j*3+1] = c
    default:
    }
}

func getCountFrame(lines []string, visited []pos, s pos, passes []pass) frame {
    buf := make([][]pixel, len(lines) * 3)
    for i := range buf {
        buf[i] = make([]pixel, len(lines[0]) * 3)
    }
    for i := range buf {
        for j := range buf[i] {
            buf[i][j] = pixel{0, 0, 70, 255}
        }
    }

    primary := pixel{0, 100, 0, 255}
    secondary := pixel{0, 0, 150, 255}

    for i := range lines {
        for j := range lines[i] {
            colourTile(lines, buf, pos{i, j}, secondary)
        }
    }

    for _, v := range visited {
        colourTile(lines, buf, v, primary)
    }

    for _, p := range passes {

        v := p.pos

        var c pixel
        if idx := slices.Index(visited, v); idx > -1 {
            c = pixel{0, 130, 0, 255}
        } else if p.status { 
            c = pixel{255, 255, 0, 255} 
        } else { 
            c = pixel{240, 0, 0, 255} 
        }

        colourTile(lines, buf, v, c)
    }

    buf[s.i*3][s.j*3] = pixel{255, 255, 0, 255}

    return buf

}

func getFrame(lines []string, visited []pos, s pos) frame {
    buf := make([][]pixel, len(lines) * 3)
    for i := range buf {
        buf[i] = make([]pixel, len(lines[0]) * 3)
    }
    for i := range buf {
        for j := range buf[i] {
            buf[i][j] = pixel{0, 0, 70, 255}
        }
    }

    primary := pixel{0, 240, 0, 255}
    secondary := pixel{0, 0, 255, 255}

    for i := range lines {
        for j := range lines[i] {
            colourTile(lines, buf, pos{i, j}, secondary)
        }
    }

    for _, v := range visited {
        colourTile(lines, buf, v, primary)
    }


    buf[s.i*3][s.j*3] = pixel{255, 255, 0, 255}

    return buf

}


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

    cmd := exec.Command(
        "ffmpeg", 
        "-hide_banner",
        "-y",
        "-f", "rawvideo",
        "-pix_fmt", "rgba",
        "-s", fmt.Sprintf("%dx%d", len(lines[0]) * 3, len(lines) * 3),
        "-i", "-", 
        // "-vf", "scale=600:-1",
        "-vf", "scale=600:-1,tpad=stop_mode=clone:stop_duration=2",
        "-r", "60",
        "-sws_flags", "neighbor",
        "foobar.mp4",
    )
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    stdin, err := cmd.StdinPipe()
    if err != nil {
        panic(err)
    }

    ffmpeg := bufio.NewWriter(stdin)

    fmt.Println(cmd.String())
    cmd.Start()

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

    sendFrame(ffmpeg, getFrame(lines, visited, s))

    // frames = append(frames, )
    for len(queue) > 0 {
        a := queue[0]
        queue = queue[1:]

        if i := slices.Index(visited, a.current); i > -1 {
            break
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
        if len(visited) % 30 == 0 {
            sendFrame(ffmpeg, getFrame(lines, visited, s))
        }
    }
    sendFrame(ffmpeg, getFrame(lines, visited, s))

    count := 0
    var passes []pass
    for i := range lines {
        isInside := false
        for j := range lines[i] {
            c := lines[i][j]
            p := pos{i, j}


            if idx := slices.Index(visited, p); idx > -1 {
                if c == 'L' || c == '|' || c == 'J' || c == 'S' {
                    isInside = !isInside
                    passes = append(passes, pass{pos: p, status: isInside})
                    continue
                }
                passes = append(passes, pass{pos: p, status: isInside})
                continue
            }

            passes = append(passes, pass{pos: p, status: isInside})
            if isInside {
                count++
            }

            if j % 21 == 0 {
                sendFrame(ffmpeg, getCountFrame(lines, visited, s, passes))
            }
        }
    }

    sendFrame(ffmpeg, getCountFrame(lines, visited, s, passes))
    fmt.Println(count)

    stdin.Close()

    if cmd.Wait() != nil {
        panic(err)
    }

}

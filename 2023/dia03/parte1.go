package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "os/exec"
    // "os"
)

type pos struct {
    i, j int
}

type num struct {
    // pos []pos
    num string
    visited bool
}

func main() {
    file, err := os.Open("./input.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var lines [][]byte
    for scanner.Scan() {
        lines = append(lines, []byte(scanner.Text()))
    }
    file.Close()

    var numbers []*num
    lookup := make(map[pos]int)

    for i := range lines {
        for j := 0; j < len(lines[i]); j++ {
            c := lines[i][j]
            if c >= '0' && c <= '9' {
                n := num{}
                for k := j; k < len(lines[i]); k++ {
                    c = lines[i][k]
                    if c >= '0' && c <= '9'{
                        pos := pos{i, k}
                        // n.pos = append(n.pos, pos)
                        n.num += string(c)
                        lookup[pos] = len(numbers)
                        j++
                    } else {
                        break
                    }
                }
                numbers = append(numbers, &n)
            }
        }
        // fmt.Print("\r", i+1, "/", len(lines))
    }

    sum := 0
    for i := range lines {
        for j, c := range lines[i] {
            if (c >= '0' && c <= '9') || c == '.' {
                continue
            }
            // fmt.Print("\r", i, j, " -> ", string(c))
            for i2 := i - 1; i2 <= i+1; i2++ {
                for j2 := j - 1; j2 <= j+1; j2++ {
                    if i2 < 0 || i2 >= len(lines) || j2 < 0 || j2 >= len(lines[0]) {
                        continue
                    }
                    p1 := pos{i2, j2}
                    p2i, ok := lookup[p1]
                    if !ok {
                        continue
                    }
                    p2 := numbers[p2i]
                    if p2.visited {
                        continue
                    }
                    num, _ := strconv.Atoi(p2.num)
                    p2.visited = true
                    sum += num
                } 
            } 
        }
    }
    fmt.Println(sum)

    return

    // isso daqui salvou a minha pele

    // cmd := exec.Command(
    //     "ffmpeg", 
    //     "-hide_banner",
    //     "-y",
    //     "-f", "rawvideo",
    //     "-pix_fmt", "rgba",
    //     "-s", fmt.Sprintf("%dx%d", len(lines[0]), len(lines)),
    //     // "-r", fmt.Sprint(FPS),
    //     // "-an",
    //     "-i", "-", 
    //     "-vf", "scale=400:-1",
    //     // "-c:v", "libx264",
    //     "foobar.bmp",
    // )
    // cmd.Stdout = os.Stdout
    // cmd.Stderr = os.Stderr
    //
    // stdin, err := cmd.StdinPipe()
    // if err != nil {
    //     panic(err)
    // }
    //
    // writer := bufio.NewWriter(stdin)
    //
    // fmt.Println(cmd.String())
    // cmd.Start()
    //
    // buf := make([]byte, len(lines) * len(lines[0]) * 4)
    //
    // for i := range lines {
    //     for j, c := range lines[i] {
    //         off := i*len(lines[0])*4 + j*4
    //         buf[off + 3] = 255
    //         if c == '.' {
    //             buf[off + 0] = 0
    //             buf[off + 1] = 0
    //             buf[off + 2] = 0
    //             buf[off + 3] = 0
    //         } else if c >= '0' && c <= '9' {
    //             buf[off + 0] = 0
    //             buf[off + 1] = 255
    //             buf[off + 2] = 0
    //         } else {
    //             buf[off + 0] = 0
    //             buf[off + 1] = 0
    //             buf[off + 2] = 255
    //         }
    //     }
    // }
    // for i := range numbers {
    //     for _, p := range numbers[i].pos {
    //         off := p.i*len(lines[0])*4 + p.j*4
    //         buf[off + 0] = 255
    //         buf[off + 1] = 0
    //         buf[off + 2] = 0
    //         buf[off + 3] = 255
    //     }
    //
    // }
    //
    // if _, err = writer.Write(buf); err != nil {
    //     panic(err)
    // }
    //
    //
    // stdin.Close()
    //
    // if cmd.Wait() != nil {
    //     panic(err)
    // }


}

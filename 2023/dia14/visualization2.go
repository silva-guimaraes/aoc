package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	// "slices"
	"os/exec"
)

type pos struct {
    i, j int
}

type frame [][]pixel

type pixel struct {
    r, g, b, a byte
}

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

func getFrame(lines [][]byte, follow pos) frame {
    buf := make([][]pixel, len(lines))
    for i := range buf {
        buf[i] = make([]pixel, len(lines[0]))
    }
    for i := range buf {
        for j := range buf[i] {
            buf[i][j] = pixel{200, 200, 200, 255}
        }
    }

    for i := range lines {
        for j, c := range lines[i] {
            if c == '#' {
                buf[i][j] = pixel{0, 0, 0, 255}
            } else if c == 'O' {
                buf[i][j] = pixel{200, 200, 0, 255}
            }
        }
    }

    buf[follow.i][follow.j] = pixel{255, 0, 0, 255}
    return buf
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

    cmd := exec.Command(
        "ffmpeg", 
        "-hide_banner",
        "-y",
        "-f", "rawvideo",
        "-pix_fmt", "rgba",
        "-s", fmt.Sprintf("%dx%d", len(lines[0]), len(lines)),
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
    cache := make(map[string][][]byte)

    var keys []string
    var rocks []pos

    for i := range lines {
        for j, c := range lines[i] {
            if c == 'O' {
                rocks = append(rocks, pos{i, j}) }

        }
    }

    r := rand.New(rand.NewSource(0))
    
    follow := rocks[r.Int31n(int32(len(rocks)))]

    for times := 0; times < 10; times++ {

        var join []string
        for i := range lines {
            join = append(join, string(lines[i]))
        }
        key := strings.Join(join, "")

        keys = append(keys, key)

        if ret, ok := cache[key]; ok {
            lines = ret
            break
        }

        copia := make([][]byte, len(lines))
        for i := range copia {
            copia[i] = make([]byte, len(lines[0]))
            copy(copia[i], lines[i])
        }

        cache[key] = copia

        // norte
        for {
            changed := false
            for ai := range lines {
                for aj, c := range lines[ai] {
                    if c == '.' || c == '#' { continue }
                    if ai == 0 { continue }
                    if lines[ai-1][aj] == '.' {
                        changed = true
                        lines[ai-1][aj] = 'O'
                        lines[ai][aj] = '.'

                        p := pos{ai, aj} 
                        if p == follow {
                            follow = pos{ai-1, aj}
                        }
                    } 
                }
            }
            sendFrame(ffmpeg, getFrame(lines, follow))
            if !changed {
                break
            }
        }

        // oeste
        for {
            changed := false
            for aj := range lines[0] {
                for ai := range lines {
                    c := lines[ai][aj]
                    if c == '.' || c == '#' { continue }
                    if aj == len(lines[0])-1 { continue }
                    if lines[ai][aj+1] == '.' {
                        changed = true
                        lines[ai][aj+1] = 'O'
                        lines[ai][aj] = '.'

                        p := pos{ai, aj} 
                        if p == follow {
                            follow = pos{ai, aj+1}
                        }
                    } 
                }
            }
            sendFrame(ffmpeg, getFrame(lines, follow))
            if !changed {
                break
            }
        }

        // sul
        for {
            changed := false
            for ai := range lines {
                for aj, c := range lines[ai] {
                    if c == '.' || c == '#' { continue }
                    if ai == len(lines)-1 { continue }
                    if lines[ai+1][aj] == '.' {
                        changed = true
                        lines[ai+1][aj] = 'O'
                        lines[ai][aj] = '.'

                        p := pos{ai, aj} 
                        if p == follow {
                            follow = pos{ai+1, aj}
                        }
                    } 
                }
            }
            sendFrame(ffmpeg, getFrame(lines, follow))
            if !changed {
                break
            }
        }

        // leste
        for {
            changed := false
            for aj := range lines[0] {
                for ai := range lines {
                    c := lines[ai][aj]
                    if c == '.' || c == '#' { continue }
                    if aj == 0  { continue }
                    if lines[ai][aj-1] == '.' {
                        changed = true
                        lines[ai][aj-1] = 'O'
                        lines[ai][aj] = '.'

                        p := pos{ai, aj} 
                        if p == follow {
                            follow = pos{ai, aj-1}
                        }
                    } 
                }
            }
            sendFrame(ffmpeg, getFrame(lines, follow))
            if !changed {
                break
            }
        }
    }

    sendFrame(ffmpeg, getFrame(lines, follow))

    stdin.Close()

    if cmd.Wait() != nil {
        panic(err)
    }
}

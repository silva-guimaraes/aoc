package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pos struct {
    i, j int
}

func main() {
    file, err := os.Open("./teste.txt")
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

    cache := make(map[string][][]byte)
    // cache["foo"] = [][]byte{}

    // for times := 0; times < 1000000000; times++ {

        // fmt.Print("\r", times)
        var join []string
        for i := range lines {
            join = append(join, string(lines[i]))
        }
        key := strings.Join(join, "")

        if ret, ok := cache[key]; ok {
            lines = ret
            panic("repeat")
            // continue
        }

        for ai := range lines {
            for aj, c := range lines[ai] {
                if c == '.' || c == '#' { continue }
                if ai == 0 { continue }
                lines[ai][aj] = '.'
                for i := ai-1; i >= 0; i-- {
                    if i == -1 {
                        lines[i+1][aj] = 'O'
                        break
                    }
                    ac := lines[i][aj]
                    if ac == '#' || ac == 'O' {
                        lines[i+1][aj] = 'O'
                        break
                    }
                    if i == 0 {
                        lines[i][aj] = 'O'
                        break
                    }
                }
            }
        }

        // for aj := range lines[0] {
        //     for ai := range lines {
        //         c := lines[ai][aj]
        //         if c == '.' || c == '#' { continue }
        //         if aj == 0 { continue }
        //         lines[ai][aj] = '.'
        //         for j := aj-1; j >= 0; j-- {
        //             if j == -1 {
        //                 lines[ai][j+1] = 'O'
        //                 break
        //             }
        //             ac := lines[ai][j]
        //             if ac == '#' || ac == 'O' {
        //                 lines[ai][j+1] = 'O'
        //                 break
        //             }
        //             if j == 0 {
        //                 lines[ai][j] = 'O'
        //                 break
        //             }
        //         }
        //     }
        // }



        // for aj := len(lines[0]) - 1; aj >= 0; aj-- {
        //     for ai := range lines {
        //         c := lines[ai][aj]
        //         if c == '.' || c == '#' { continue }
        //         if aj == len(lines[0]) - 1 { continue }
        //         lines[ai][aj] = '.'
        //         for j := aj+1; j < len(lines[0]); j++ {
        //             if j == len(lines[0]) {
        //                 lines[ai][j-1] = 'O'
        //                 break
        //             }
        //             ac := lines[ai][j]
        //             if ac == '#' || ac == 'O' {
        //                 lines[ai][j-1] = 'O'
        //                 break
        //             }
        //             if j == len(lines[0])-1 {
        //                 lines[ai][j] = 'O'
        //                 break
        //             }
        //         }
        //     }
        // }

        // for ai := len(lines) - 1; ai >= 0; ai-- {
        //     for aj, c := range lines[ai] {
        //         if c == '.' || c == '#' { continue }
        //         if ai == len(lines) - 1 { continue }
        //         lines[ai][aj] = '.'
        //         for i := ai+1; i >= 0; i++ {
        //             if i == len(lines) {
        //                 lines[i-1][aj] = 'O'
        //                 break
        //             }
        //             ac := lines[i][aj]
        //             if ac == '#' || ac == 'O' {
        //                 lines[i-1][aj] = 'O'
        //                 break
        //             }
        //             if i == len(lines)-1 {
        //                 lines[i][aj] = 'O'
        //                 break
        //             }
        //         }
        //     }
        // }

        copia := make([][]byte, len(lines))
        for i := range copia {
            copia[i] = make([]byte, len(lines[0]))
            copy(copia[i], lines[i])
        }

        cache[key] = copia
    // }
    fmt.Println()

    sum := 0
    for i := range lines {
        for _, c := range lines[i] {
            fmt.Print(string(c)) 
            if c == 'O' {
                sum += len(lines) - i
            }
        }
        // fmt.Print("  ")
        // for j, c := range lines[i] {
        //     fmt.Print(string(copia[i][j])) 
        //     if c == 'O' {
        //         sum += len(lines) - i
        //     }
        // }
        fmt.Println()
    }

    fmt.Println(sum)
}

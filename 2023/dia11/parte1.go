package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

type path struct {
    current pos
    number int
}


type pos struct {
    i, j int
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

    var galaxies []pos
    var columns, rows []int
    rows = append(rows, 0)
    columns = append(rows, 0)

    for i := range lines {
        for j := range lines[i] {
            if lines[i][j] == '#' {
                galaxies = append(galaxies, pos{i, j})
                rows = append(rows, i)
                columns = append(columns, j)
            }
        }
    }
    slices.Sort(columns)
    var rowdiff []int
    for i := 1; i < len(rows); i++ {
        rowdiff = append(rowdiff, rows[i] - rows[i-1])
    }
    var coldiff []int
    for i := 1; i < len(columns); i++ {
        coldiff = append(coldiff, columns[i] - columns[i-1])
    }

    rows = rows[1:]
    columns = columns[1:]

    for i, g := range galaxies {
        count := 0
        for j, r := range rows {
            if r > g.i { break }
            if rowdiff[j] > 1 {
                count += rowdiff[j] - 1
            }
        }
        galaxies[i].i += count
        count = 0
        for j, c := range columns {
            if c > g.j { break }
            if coldiff[j] > 1 {
                count += coldiff[j] - 1
            }
        }
        galaxies[i].j += count
    }

    sum := 0
    for i := 0; i < len(galaxies); i++ {
        for j := 0 + i; j < len(galaxies); j++ {
            g1 := galaxies[i]
            g2 := galaxies[j]
            sum += int(math.Abs(float64(g1.i - g2.i)) + math.Abs(float64(g1.j - g2.j)))
        }
    }
    fmt.Println(sum)
}

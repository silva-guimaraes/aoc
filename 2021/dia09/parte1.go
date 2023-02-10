package main

import (
	"bufio"
	"fmt"
	"os"
)

type ij struct { i, j int }

func main() {

    file, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    if err != nil { panic(err) }


    var input [][]int
    for i := 0; scanner.Scan(); i++ {
        text := scanner.Text()
        input = append(input, make([]int, len(text)))
        for j := range text {
            input[i][j] = int(text[j] - '0')
        }
    }

    low := make(map[ij]int) 

    for i := range input {
        for j := range input[i] {

            current := ij{i, j}

            up :=       ij{current.i - 1,   current.j}
            down :=     ij{current.i + 1,   current.j}
            left :=     ij{current.i,       current.j - 1}
            right :=    ij{current.i,       current.j + 1}

            if up.i >= 0 && 
            input[up.i][up.j] <= input[current.i][current.j] {
                continue
            } 
            if down.i <= len(input) - 1 && 
            input[down.i][down.j] <= input[current.i][current.j] {
                continue
            } 
            if left.j >= 0 && 
            input[left.i][left.j] <= input[current.i][current.j] {
                continue
            } 
            if right.j <= len(input[0]) - 1 && 
            input[right.i][right.j] <= input[current.i][current.j] {
                continue
            }

            low[current] = input[current.i][current.j]
        }
    }

    sum := 0
    for _, v := range low {
        sum += v + 1
    }

    fmt.Println(sum)
}

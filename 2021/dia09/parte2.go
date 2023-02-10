package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

            for {
                if input[current.i][current.j] == 9 { break }

                up :=       ij{current.i - 1,   current.j}
                down :=     ij{current.i + 1,   current.j}
                left :=     ij{current.i,       current.j - 1}
                right :=    ij{current.i,       current.j + 1}

                lowest := current

                if up.i >= 0 && 
                input[up.i][up.j] < input[current.i][current.j] {
                    lowest = up
                } 
                if down.i < len(input) && 
                input[down.i][down.j] < input[current.i][current.j] {
                    lowest = down
                } 
                if left.j >= 0 && 
                input[left.i][left.j] < input[current.i][current.j] {
                    lowest = left
                } 
                if right.j < len(input[0]) && 
                input[right.i][right.j] < input[current.i][current.j] {
                    lowest = right
                }

                if lowest == current {
                    low[current]++
                    break
                } else {
                    current = lowest
                }
            }
        }
    }

    var order []int

    for _, v := range low {
        order = append(order, v)
    }
    sort.Slice(order, func (i, j int) bool {
        return order[i] >  order[j]
    })

    fmt.Println(order[0] * order[1] * order[2])
}

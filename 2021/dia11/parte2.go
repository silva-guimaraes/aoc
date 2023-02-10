package main

import (
	"bufio"
	"fmt"
	"os"
)

func is_sync(input [][]int) bool {
    for i := range input {
        for j := range input[i] {
            if input[i][j] != 0 {
                return false
            }
        }
    }
    return true
}


func fire_octo(input [][]int, i, j int) {

    input[i][j] = -1

    var ai, aj, bi, bj int

    if i == 0 { 
        ai = i 
    } else {    
        ai = i - 1 
    }
    if j == 0 { 
        aj = j 
    } else { 
        aj = j - 1 
    }

    if i == len(input) - 1 { 
        bi = i 
    } else { 
        bi = i + 1 
    }
    if j == len(input[0]) - 1 { 
        bj = j 
    } else { 
        bj = j + 1 
    }

    // if i == 4 && j == 8 {
    //     fmt.Println(ai, aj, bi, bj)
    // }

    for k := ai; k <= bi; k++ {
        for m := aj; m <= bj; m++ {

            if input[k][m] > -1 {
                input[k][m]++
            }

            if input[k][m] > 9 {
                fire_octo(input, k, m)
            }
        }
    }
}

func main() {

    file, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)


    var input [][]int
    for i := 0; scanner.Scan(); i++ {
        text := scanner.Text()
        input = append(input, make([]int, len(text)))
        for j := range text {
            input[i][j] = int(text[j] - '0')
        }
    }

    count := 0

    for step := 0;; step++ {
        for i := range input {
            for j := range input[i] {
                input[i][j]++
            }
        }
        for i := range input {
            for j := range input[i] {
                if input[i][j] > 9 {
                    fire_octo(input, i, j)
                }
            }
        }
        for i := range input {
            for j := range input[i] {
                if input[i][j] == -1 {
                    input[i][j] = 0
                    count++
                }
            }
        }
        if is_sync(input) {
            fmt.Println(step + 1)
            return
        }
        // fmt.Println(step + 1)
        // for i := range input {
        //     for j := range input {
        //         fmt.Print(input[i][j])
        //     }
        //     fmt.Println()
        // }
        // fmt.Println()
    }
}

package main

import (
	"bufio"
	"fmt"
	"strconv"

	// "math"
	"os"
	"regexp"
	// "strconv"
)


type fold struct {
    axis byte
    at int
}

// func print_origami(origami [][]byte) {
//     for i := range origami {
//         for j := range origami[i] {
//             if origami[i][j] > 0 {
//                 fmt.Print("# ")
//             } else {
//                 fmt.Print(". ")
//             }
//         }
//         fmt.Println()
//     }
// }

func main() {

    file, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    find_numb, err := regexp.Compile("\\d+")
    if err != nil { panic(err) }

    var input []string
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }
    var fold_input []string
    for i := range input {
        if input[i] == ""{
            fold_input = input[i+1:]
            input = input[:i]
            break
        }
    }

    var converted_nums []int
    for i := range input {
        ext := find_numb.FindAllString(input[i], -1)

        for j := range ext {
            conv, err := strconv.Atoi(ext[j])
            if err != nil { panic(err) }

            converted_nums = append(converted_nums, conv)
        }
    }

    xmax, ymax := converted_nums[0], converted_nums[1]
    for i := 0; i < len(converted_nums); i += 2 {
        if converted_nums[i] > xmax { 
            xmax = converted_nums[i] 
        }
        if converted_nums[i+1] > ymax { 
            ymax = converted_nums[i+1] 
        }
    }

    origami := make([][]byte, ymax + 1)
    for i := range origami {
        origami[i] = make([]byte, xmax + 1)
    }
    for i := 0; i < len(converted_nums); i += 2 {
        origami[converted_nums[i+1]][converted_nums[i]] = '#'
    }

    var folds []fold
    for i := range fold_input {
        ext := find_numb.Find([]byte(fold_input[i]))
        conv, err := strconv.Atoi(string(ext))
        if err != nil { panic(err) }

        folds = append(folds, fold{fold_input[i][11], conv})

    }
    fold := folds[0]

    fmt.Println(fold_input)

    // print_origami(origami)
    fmt.Println()
    for i := range origami {
        for j := range origami[i] {
            if j > fold.at && origami[i][j] == '#' {
                origami[i][j] = 0
                origami[i][fold.at - (j - fold.at)] = '#'
            }
        }
    }
    // print_origami(origami)

    count := 0
    for i := range origami {
        for j := range origami[i] {
            if origami[i][j] == '#' { count++ }
        }
    }

    fmt.Println(fold)
    fmt.Println(count)

}


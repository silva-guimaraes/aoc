package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"regexp"
	// "strconv"
)

func main() {

    file, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    find_chars, err := regexp.Compile("\\w+")
    if err != nil { panic(err) }

    var input [][]string
    for scanner.Scan() {
        ext := find_chars.FindAllString(scanner.Text(), -1)
        input = append(input, ext)
    }

    count := 0
    const OUTPUT int = 10
    for i := range input {
        for j := range input[i][OUTPUT:] {
            if length := len(input[i][OUTPUT + j]); length == 2 || length == 4 || 
            length == 3 || length == 7 {
                count++
                }
        }
    }
    fmt.Println(count)

}

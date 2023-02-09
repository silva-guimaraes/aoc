package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"os"
)

func main() {

    file, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    find_numbers, err := regexp.Compile("\\d+")
    if err != nil { panic(err) }

    scanner.Scan()
    ext := find_numbers.FindAllString(scanner.Text(), -1)

    var input []int
    for i := range ext {
        conv, _ := strconv.Atoi(ext[i])
        input = append(input, conv)
    }

    for i := 0; i < 80; i++ {
        var temp []int
        for i := range input {

            if input[i] == 0 {
                temp = append(temp, 8)
                input[i] = 6

            } else {
                input[i]--
            }
        }
        input = append(input, temp...)
    }

    fmt.Println(len(input))
}

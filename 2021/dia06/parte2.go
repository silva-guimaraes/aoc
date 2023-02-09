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

    input := make([]int, 9)
    for i := range ext {
        conv, _ := strconv.Atoi(ext[i])
        input[conv]++
    }

    for i := 0; i < 256; i++ {
        temp := input[0]
        input[0] = 0
        for i := 1; i < 9; i++ {
            input[i - 1] = input[i]
            input[i] = 0
        }
        input[6] += temp
        input[8] += temp
    }

    count := 0
    for i := range input {
        count += input[i]
    }
    fmt.Println(count)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strconv"
)


func main() {

	fmt.Println("hello world")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	horizontal, depth := 0, 0

	for i := range input {
		switch (input[i][0]) {
		case 'u':
			temp, _ := strconv.Atoi(string(input[i][3]))
			depth -= temp
		case 'f':
			temp, _ := strconv.Atoi(string(input[i][8]))
			horizontal += temp
		case 'd':
			temp, _ := strconv.Atoi(string(input[i][5]))
			depth += temp
		}
	}

	fmt.Println(horizontal * depth)
}

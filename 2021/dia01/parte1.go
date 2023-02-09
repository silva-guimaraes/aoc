package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {

	fmt.Println("hello world")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	last := 0
	count := 0

	for scanner.Scan() {
		temp, _ := strconv.Atoi(scanner.Text())
		if temp - last > 0 {
			count++
			// fmt.Println(last, temp)
		}
		last = temp
	}

	fmt.Println(count - 1)
}

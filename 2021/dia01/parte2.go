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

	var measures []int
	for scanner.Scan() {
		temp, _ := strconv.Atoi(scanner.Text())
		measures = append(measures, temp)
	}

	for i := 0; i + 2 < len(measures); i++ {
		temp := measures[i] + measures[i + 1] + measures[i + 2]

		if temp - last > 0 {
			count++
		}
		last = temp
	}

	fmt.Println(count - 1)
}

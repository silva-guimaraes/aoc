package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {

	// fmt.Println("hello world")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	out := make([]rune, len(input[0]))
	length := len(input)

	for i := 0; i < len(input[0]); i++ {
		count := 0
		for j := range input {
			if input[j][i] == '1' { count++ }
		}
		if count > length / 2 {
			out[i] = '1'

		} else {
			out[i] = '0'
		}
	}
	var binary int32 = 0x0
	for i := range out {
		binary = binary << 1
		if out[i] == '1' {
			binary = binary | 0x1
		}
	}

	fmt.Println(binary * (binary ^ 0xFFF /* 0x1F */))
}

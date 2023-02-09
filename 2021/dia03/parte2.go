package main

import (
	"bufio"
	"fmt"
	"os"
)

func foobar(input []string, a, b byte) int32 {

	out := make([]byte, len(input[0]))

	for i := 0; len(input) > 1; i++ {
		one, zero := 0, 0
		for j := range input {
			if input[j][i] == '1' {
				one++
			} else {
				zero++
			}
		}
		if one >= zero {
			out[i] = a

		} else {
			out[i] = b
		}

		var temp []string
		for j := range input  {
			if input[j][i] == out[i] {
				temp = append(temp, input[j])
			}
		}
		input = temp
		// fmt.Println("len", len(input), "onezero", one, zero)
		// fmt.Println(input, out)
	}

	var binary int32 = 0x0
	for i := range input[0] {
		binary = binary << 1
		if input[0][i] == '1' {
			binary |= 0x1
		}
	}
	return binary
}


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

	fmt.Println(foobar(input, '1', '0') * foobar(input, '0', '1'))

	// fmt.Println(binary * (binary ^ 0xFFF /* 0x1F */))
}

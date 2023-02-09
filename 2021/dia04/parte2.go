package main

import (
	"bufio"
	"regexp"
	"fmt"
	"strconv"
	"os"
)

const MARKED = -1

func main() {

	file, err := os.Open("input.txt")
	if err != nil { panic(err) }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	find_numbers, err := regexp.Compile("\\d+")
	if err != nil { panic(err) }

	scanner.Scan()
	raw_bingo := find_numbers.FindAllString(scanner.Text(), -1)
	bingo := make([]int, len(raw_bingo))
	for i := range raw_bingo {
		conv, err := strconv.Atoi(raw_bingo[i])
		if err != nil { panic(err) }
		bingo[i] = conv
	}

	var input []int
	for scanner.Scan() {
		if text := scanner.Text(); text != "" {
			raw_bingo = find_numbers.FindAllString(text, -1)
			for i := range raw_bingo {
				conv, _ := strconv.Atoi(raw_bingo[i])
				input = append(input, conv)
			}
		}
	}

	 var raw_boards [][]int
	 for i := 0; i < len(input); i += 5 {
	 	end := i + 5
	 	raw_boards = append(raw_boards, input[i:end])
	 }
	 var boards [][][]int
	 for i := 0; i < len(raw_boards); i += 5 {
	 	end := i + 5
	 	boards = append(boards, raw_boards[i:end])
	}

	// essa é a coisa mais feia que eu ja escrevi
	for current := range bingo {
		for board := 0; board < len(boards); board++ {
			count_unmarked := func () bool {
				if len(boards) > 1 { 
					boards = append(boards[:board], boards[board + 1:]...)
					board--
					return false
				}
				sum := 0
				for i := range boards[board] {
					for j := range boards[board][i] {
						if number := boards[board][i][j]; number != MARKED {
							sum += number
						}
					}
				}
				fmt.Println(sum * bingo[current])
				os.Exit(0)
				return true
			}
			for i := range boards[board] {
				horizontal_sum := 0
				for j := range boards[board][i] {
					number := boards[board][i][j]
					if  number == bingo[current] {
						boards[board][i][j] = MARKED
					}
					horizontal_sum += number
				}
				if horizontal_sum == MARKED * 5 {
					if !count_unmarked() {
						// meu deus do céu
						goto foobar
					}
				}
			}
			for i := range boards[board] {
				vertical_sum := 0
				for j := range boards[board][i] {
					vertical_sum += boards[board][j][i]
				}
				if vertical_sum == MARKED * 5 {
					if !count_unmarked() {
						goto foobar
					}
				}
			}
			foobar:
		}
	}
	fmt.Println("nada?")
}

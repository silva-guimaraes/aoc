package main

import (
    "fmt"
    "os"
    "bufio"
)

const (
    rock = 1
    paper = 2
    scissors = 3
)
const (
    lose = 1
    draw = 2
    win = 3
)

func guess(opp int, you int) int {
    if you == draw {
	return  3 + opp

    } else if you == win  {
	x := opp + 1
	if x == 4 { x = 1 }
	return 6 + x

    } else  { // derrota
	x := opp - 1
	if x == 0 { x = 3 }
	return 0 + x
    }
}

func main() {
    file, err := os.Open("./strategy_guide.txt")
    if err != nil {
	panic(err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    total_score := 0
    for scanner.Scan() {
	line := scanner.Text()
	opponent := int(byte(line[0]) - 'A') + 1
	you := int(byte(line[2]) - 'X') + 1

	score := guess(opponent, you)
	total_score += score
	// fmt.Println(opponent, you, score, total_score)
    }

    fmt.Println(total_score)
}

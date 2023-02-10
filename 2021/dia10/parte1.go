package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []byte

func (s *stack) push(x byte) {
    *s = append(*s, x)
}

func (s *stack) pop() byte {

    temp := (*s)[len(*s) - 1]

    *s = (*s)[:len(*s) - 1]

    return temp
}

func main() {

    file, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    if err != nil { panic(err) }

    var input []string
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }

    points := make(map[byte]int)

    points[')'] = 3
    points[']'] = 57
    points['}'] = 1197
    points['>'] = 25137

    parens := make(map[byte]byte)

    parens[')'] = '('
    parens[']'] = '['
    parens['}'] = '{'
    parens['>'] = '<'


    score := 0
    for i := range input {
        var checker stack
        for j := range input[i] {

            if input[i][j] == '(' || input[i][j] == '[' || input[i][j] == '{' || input[i][j] == '<' {
                checker.push(input[i][j])

            } else if last := checker.pop(); parens[input[i][j]] != last {
                score += points[input[i][j]]
                break
            }
        }
    }

    fmt.Println(score)
}

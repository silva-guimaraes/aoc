package main

import (
	"bufio"
	"fmt"
	"os"
    "sort"
)

type stack []byte

func (s *stack) push(x byte) {
    *s = append(*s, x)
}

func (s *stack) pop() (byte, bool) {

    if len(*s) == 0 {
        return 0, false
    }
    temp := (*s)[len(*s) - 1]

    *s = (*s)[:len(*s) - 1]

    return temp, true
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

    points[')'] = 1
    points[']'] = 2
    points['}'] = 3
    points['>'] = 4

    parens := make(map[byte]byte)

    parens[')'] = '('
    parens[']'] = '['
    parens['}'] = '{'
    parens['>'] = '<'

    parens['('] = ')'
    parens['['] = ']'
    parens['{'] = '}'
    parens['<'] = '>'

    var scores []int
    for i := range input {
        var checker stack
        for j := range input[i] {

            if input[i][j] == '(' || input[i][j] == '[' || input[i][j] == '{' || input[i][j] == '<' {
                checker.push(input[i][j])

            } else if last, _ := checker.pop(); parens[input[i][j]] != last {
                // score += points[input[i][j]]
                checker = stack{}
                break
            }

        }
        if len(checker) > 0 {
            // fmt.Println(checker)
            // for j := range checker { fmt.Print(points[parens[checker[j]]], " ") }
            // fmt.Println()
            sum := 0
            for {
                foo, exists := checker.pop()
                if !exists { 
                    break 
                }

                sum *= 5
                sum += points[parens[foo]]
            }
            scores = append(scores, sum)
        }
    }
    sort.Ints(scores)

    fmt.Println(scores[len(scores) / 2])
}

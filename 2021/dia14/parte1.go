package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type template struct {
    char, insert byte // insert guarda um caractere que precise se colocado logo a frente
}

// alguns dos nomes dessas variaveis não fazem sentido.

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

    input_output := input[0]
    input = input[2:]

    var output []template
    for i := range input_output {
        output = append(output, template{input_output[i], 0})
    }

    for step := 0; step < 10; step++ {

        for j := 0; j < len(output) - 1; j++ {
            test_pair := string(output[j].char) + string(output[j+1].char)

            for i := range input {
                pair := input[i][:2]
                letter := input[i][6]

                if test_pair == pair {
                    output[j].insert = letter
                    break
                }
            }

        }
        for i := 0; i < len(output); i++ {
            letter := output[i]
            if letter.char > 0 {
                output = append(output[:i+1], output[i:]...)
                output[i+1] = template{output[i].insert, 0}
                output[i].insert = 0
                i++
            }
            
        }
    }
    frequency := make(map[string]int)
    for i := range output {
        if output[i].char == 0 { continue }
        frequency[string(output[i].char)]++
    }
    var order []int
    for _, v := range frequency {
        order = append(order, v)
    }
    sort.Ints(order)

    fmt.Println(order[len(order) -1] - order[0])

}


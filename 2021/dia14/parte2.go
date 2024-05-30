package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

    raw_template := input[0]
    input = input[2:]

    template := make(map[string]int)
    middle := make(map[string]byte)

    for i := range input {
        raw_pair := input[i][:2]
        raw_letter := input[i][6]
        middle[raw_pair] = raw_letter
        template[raw_pair] = 0
    }
    
    for i := 0; i < len(raw_template) - 1; i++ {
        raw_pair := raw_template[i:i+2]
        template[raw_pair]++
    }

    buffer := make(map[string]int)

    for step := 0; step < 40; step++ {
        for pair, v := range template {

            if template[pair] > 0 {
                // basicamente conta todos os pares diferente do template em um hash map e trata 
                // um novo caractere como um par a menos e dois pares novos

                buffer[    string(pair[0])      + string(middle[pair])] += v
                buffer[    string(middle[pair]) + string(pair[1])] += v
                template[pair] -= v
            }
        }
        for k, v := range buffer {
            template[k] += v
            delete(buffer, k)
        }
    }
    count := make(map[byte]int)

    for k, v := range template {
        count[k[0]] += v
        count[k[1]] += v
    }

    count[raw_template[0]]++
    count[raw_template[len(raw_template) -1]]++

    var order []int
    for _, v := range count {
        order = append(order, v / 2)
    }

    sort.Ints(order)
    fmt.Println(order[len(order) -1] - order[0])
}


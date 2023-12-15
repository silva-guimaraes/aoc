package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    } 

    words := strings.Split(strings.TrimSpace(string(file)), ",")

    sum := 0
    for _, w := range words {
        current := 0
        for _, c:= range w {
            current += int(c)
            current *= 17
            current %= 256
        }
        sum += current
    }
    fmt.Println(sum)
}

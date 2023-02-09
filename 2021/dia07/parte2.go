package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

// tenho quase certeza que existe um termo na matematica pra um algoritmo desse tipo, só não sei qual.
func foobar(n int) int {
    if n == 0 { return 0 }
    return foobar(n - 1) + n
}

func main() {

    file, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    find_numbers, err := regexp.Compile("\\d+")
    if err != nil { panic(err) }

    scanner.Scan()
    ext := find_numbers.FindAllString(scanner.Text(), -1)

    var input []int
    max := 0
    for i := range ext {
        conv, _ := strconv.Atoi(ext[i])
        input = append(input, conv)
        if conv > max { max = conv }
    }
    needed_fuel := make([]int, max)


    for i := 0; i < max; i++ {
        sum := 0
        for j := range input {
            sum += foobar(int(math.Abs(float64(i - input[j]))))
        }
        needed_fuel[i] = sum
    }
    min := needed_fuel[0]
    for i := range needed_fuel {
        if needed_fuel[i] < min { min = needed_fuel[i]}
    }
    fmt.Println(min)
}

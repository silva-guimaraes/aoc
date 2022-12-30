package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
    file, err := os.Open("./calories.txt")
    if err != nil {
	panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var rank []int // lista com a soma de calorias
    sum := 0

    for scanner.Scan() {
	num, err := strconv.Atoi(scanner.Text())

	// salvar soma quando linha em branco for encontrada
	if err != nil {
	    rank = append(rank, sum)
	    sum = 0
	    continue
	} else {
	    sum += num
	}
    }
    // ultima soma
    rank = append(rank, sum)

    fmt.Println(rank)

    sort.Slice(rank, func(i, j int) bool {
	return rank[i] > rank[j]
    })

    fmt.Println(rank[0] + rank[1] + rank[2])
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// queria tentar fazer alguma coisa mais legivel dessa vez

func priority(c rune) int {
    if unicode.IsUpper(c) {
         return int(c) - 'A' + 27
    } else {
        return int(c) - 'a' + 1
    }
}

func filter_common_items(sack string, common_items []bool) []bool {

    occurrence := make([]bool, 52)

    for _, c := range sack {
	occurrence[priority(c) - 1] = true
    }

    for i := range occurrence {
	if occurrence[i] == false && common_items[i] == true {
	    common_items[i] = false
	}
    }
    return common_items
}

func main() {
    file, err := os.Open("./rucksacks.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    // ler todos as sacolas pra um slice
    var groups []string
    for scanner.Scan() {
	groups = append(groups, scanner.Text())
    }
    file.Close()

    // separar grupos de 3
    var trios [][]string
    for i := 0; i < len(groups); i += 3 {

	trio := make([]string, 3)
	trio[0] = groups[i + 0]
	trio[1] = groups[i + 1]
	trio[2] = groups[i + 2]

	trios = append(trios, trio)
    }

    priority_sum := 0

    // iterar cada grupo e encontrar item em comum
    for i := range trios {
	// lista com todos os items possiveis. iterar por cada elfo e retirar dessa lista
	// os items que aquele elfo não possuia. o elemento que sobrar depois de verificar
	// os 3 elfos é o elemento em comum entre os 3
	common_items := make([]bool, 52)
	for i := range common_items { common_items[i] = true } // surreal essa lingua

	for j := range trios[i] {
	    common_items = filter_common_items(trios[i][j], common_items)
	}

	for i := range common_items {
	    if common_items[i] == true {
		priority_sum += i + 1
		break
	    }
	}
    }

    fmt.Println(priority_sum)
}

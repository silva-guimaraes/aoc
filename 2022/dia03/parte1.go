package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
    file, err := os.Open("./rucksacks.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    priority_sum := 0

    // encontrar item repetido nos dois compartimentos enquanto lendo o arquivo
    for scanner.Scan() {
	lookup := make(map[byte]bool)
	line := scanner.Text()
	length := len(line)
	half := length / 2

	// salvar todos os itens diferentes do primeiro compartimento
	for  i := 0; i < half; i++ {
	    _, present := lookup[line[i]]
	    if !present {
		lookup[line[i]] = true
	    }
	}
	// iterar pelo segundo e parar quando item repetido for encontrado
	for i := half; i < length; i++ {

	    if _, present := lookup[line[i]]; present {
		// item encontrado. calcular prioridade e adicionar a soma total
		if unicode.IsUpper(rune(line[i])) {
		     priority_sum += int(line[i]) - 'A' + 27
		} else {
		    priority_sum += int(line[i]) - 'a' + 1
		}
		break
	    }
	}
    }
    fmt.Println(priority_sum)
}

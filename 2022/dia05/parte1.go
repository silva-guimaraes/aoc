package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	 "strings"
)

const (
    times = 0
    from = 1
    to = 2
)

type stack []string

func (s *stack) is_empty() bool {
    return len(*s) == 0
}

func (s *stack) push(str string){
    *s = append(*s, str)
}

func (s *stack) pop() (string, bool) {
    if s.is_empty() {
	return "", false
    } else {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
    }
}

func main() {
    // me diz que todo mundo fez o mesmo
    crates := []stack{	{"D", "L", "J", "R", "V", "G", "F"},
    			{"T", "P", "M", "B", "V", "H", "J", "S"},
    			{"V", "H", "M", "F", "D", "G", "P", "C"},
    			{"M", "D", "P", "N", "G", "Q"},
    			{"J", "L", "H", "N", "F"},
    			{"N", "F", "V", "Q", "D", "G", "T", "Z"},
    			{"F", "D", "B", "L"},
    			{"M", "J", "B", "S", "V", "D", "N"},
    			{"G", "L", "D"}}


    file, err := os.Open("./hanoi.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    // ler linhas
    var lines []string
    for scanner.Scan() {
	lines = append(lines, scanner.Text())
    }
    file.Close()

    r, _ := regexp.Compile("\\d+")

    // extrair instruções
    var steps [][]int
    for i := range lines {
	extract := r.FindAllString(lines[i], -1)

	indexes := make([]int, 3)

	for j := range extract {
	    indexes[j], _ = strconv.Atoi(extract[j])
	}

	steps = append(steps, indexes)
    }

    for _, step := range steps {
	// como dito, movendo uma caixa de uma vez, varias vezes
	for j := 0; j < step[times]; j++ {
	    crate, returned := crates[step[from] - 1].pop()
	    if !returned {
		break
	    }
	    crates[step[to] - 1].push(crate)
	}
    }

    var result strings.Builder

    for i := range crates {
	crate := crates[i]
	if !crate.is_empty() {
	    result.WriteString(crate[len(crate) - 1])
	}
    }
    fmt.Println(result.String())
}

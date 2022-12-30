package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	 "strings"
)

type stack []string

const (
    times = 0
    from = 1
    to = 2
)

func (s *stack) is_empty() bool {
    return len(*s) == 0
}

func (s *stack) push(str string){
    *s = append(*s, str)
}

func (s *stack) push_slice(slc []string){
    *s = append(*s, slc...)
}

func (s *stack) pop_slice(slice int) ([]string, bool) {
    if s.is_empty() || slice == 0 {
	return nil, false
    } else if len(*s) < slice {
	s.pop_slice(len(*s))
    } else {
	index := len(*s) - slice
	slice := (*s)[index:]
	*s = (*s)[:index]
	return slice, true
    }
    return nil, false
}

func (s *stack) pop() (string, bool) {
    a, err := s.pop_slice(1)
    return a[0], err
}



func main() {
    // crates := []stack{	{"Z", "N"},
    // 			{"M", "C", "D"},
    // 			{"P"} }

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

    var lines []string
    for scanner.Scan() {
	lines = append(lines, scanner.Text())
    }
    file.Close()

    r, _ := regexp.Compile("\\d+")

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
	crate, returned := crates[step[from] - 1].pop_slice(step[times])
	if !returned {
	    continue
	}
	crates[step[to] - 1].push_slice(crate)
    }

    var result strings.Builder

    for i := range crates {
	stack := crates[i]
	if !stack.is_empty() {
	    result.WriteString(stack[len(stack) - 1])
	}
    }
    fmt.Println(result.String())
}

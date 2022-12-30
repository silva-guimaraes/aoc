package main

import (
    "fmt"
    "bufio"
    "os"
)

const start = 242
const end = 228

type pos struct {
    i, j, counter int
}

type stack []pos

func (s *stack) is_empty() bool {
    return len(*s) == 0
}

func (s *stack) push(str pos){
    *s = append(*s, str)
}

func (s *stack) queue(p pos){
    *s = append(stack{p}, *s...)
}

func (s *stack) pop() (pos, bool) {
    if s.is_empty() {
	return pos{0, 0, 0}, false
    } else {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
    }
}

func remove_element(s []pos, index int) []pos {
    return append(s[:index], s[index + 1:]...)
}

func index_of(element int, data [][]int) (int, int) {
   for i := range data {
       for j := range data[i] {
	   if data[i][j] == element {
	       return i, j
	   }
       }
   }
   return 0, 0
}

func did_visit(pos pos, visited []pos) bool {
    for i := range visited {
	if visited[i].i == pos.i && visited[i].j == pos.j {
	    return true
	}
    }
    return false
}

func is_valid(current pos, target pos, heightmap [][]int) bool {

    if target.i >= 0 && target.i < len(heightmap) &&
    target.j >= 0 && target.j < len(heightmap[0]) {

	current_height := heightmap[current.i][current.j]
	target_height := heightmap[target.i][target.j]

	return target_height == current_height - 1 || target_height >= current_height ||
	target_height == 'E' - 'a'

    } else {
	return false
    }
}

func main() {
    file, err := os.Open("./heightmap.txt")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var lines []string
    for scanner.Scan() {
	lines = append(lines, scanner.Text())
    }

    heightmap := make([][]int, len(lines))
    for i := range heightmap {
	heightmap[i] = make([]int, len(lines[0]))
	for j := range heightmap[i] {
	    heightmap[i][j] = int(lines[i][j] - 'a')
	}
    }
    file.Close()


     starti, startj := index_of(start, heightmap)
    // start := pos{starti, startj, -1}
    endi, endj := index_of(end, heightmap)
    end := pos{endi, endj, 0}

    heightmap[endi][endj] = 'z' - 'a'
    heightmap[starti][startj] = 0

    var visited stack
    var queue stack
    queue.queue(end)

    fmt.Println(string(heightmap[endi][endj] + 'a'))

    for !queue.is_empty() {

	current, exists := queue.pop()
	if !exists {
	    panic(exists)
	}

	if did_visit(current, visited) {
	    continue
	}

	up := pos{	current.i - 1,	current.j,	current.counter + 1}
	down := pos{	current.i + 1,	current.j,	current.counter + 1}
	left := pos{	current.i,	current.j - 1,	current.counter + 1}
	right := pos{	current.i,	current.j + 1,	current.counter + 1}

	if is_valid(current, up, heightmap) && !did_visit(up, visited) {
	    queue.queue(up)
	}
	if is_valid(current, down, heightmap) && !did_visit(down, visited) {
	    queue.queue(down)
	}
	if is_valid(current, left, heightmap) && !did_visit(left, visited) {
	    queue.queue(left)
	}
	if is_valid(current, right, heightmap) && !did_visit(right, visited) {
	    queue.queue(right)
	}

	fmt.Print("\r[", current.i, " ", current.j,  "] len: ", len(queue), " visited: ", len(visited))

	visited.push(current)

	// tinha minhas duvidas se valeu mesmo a pena ter feito o algoritmo ao contrario 
	// na primeira parte mas essa segunda queria exatamento o que eu tinha feito =D
	if heightmap[current.i][current.j] == 'a' - 'a' {
	    fmt.Println()
	    fmt.Println(current.counter)
	    return

	}
    }
    fmt.Println("nada?")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type step struct {
    direction string
    times int
}

type pos struct {
    x, y int
}

const (
    x = 0
    y = 1
)


func is_unique(list [][]int, yy, xx int) bool {
    for i := range list {
	if list[i][y] == yy && list[i][x] == xx {
	    return false
	}
    }
    return true

}

func main() {
    file, err := os.Open("./rope.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    // ler linhas como strings
    var lines []string
    for scanner.Scan() {
	lines = append(lines, scanner.Text())
    }
    file.Close()


    // extrair instruções
    r, _ := regexp.Compile("\\d+")
    var steps []step
    for i := range lines {
	times, _ := strconv.Atoi(r.FindString(lines[i]))
	steps = append(steps, step{string(lines[i][0]), times})
    }

    tail := []int{0, 0}
    head := []int{0, 0}

    unique_visits := make(map[pos]int)
    var head_visited [][]int
    var tail_visited [][]int

    head_visited = append(head_visited, []int{head[y], head[x]})
    tail_visited = append(tail_visited, []int{tail[y], tail[x]})
    // unique_visits[pos{tail[x], tail[y]}]++

    // calcular todos os movimentos de head primeiro
    for i := range steps {
	for j := 0; j < steps[i].times; j++ {

	    // mover head de acordo com as direções
	    if steps[i].direction == "U" {
		head[y] -= 1

	    } else if steps[i].direction == "D" {
		head[y] += 1

	    } else if steps[i].direction == "L" {
		head[x] -= 1

	    } else if steps[i].direction == "R" {
		head[x] += 1
	    }

	    // salvar posições
	    head_visited = append(head_visited, []int{head[y], head[x]})
	}
    }

    for i := range head_visited {

	dx := head_visited[i][x] - tail[x]
	dy := head_visited[i][y] - tail[y]

	// se tail ficar a mais de uma unidade longe de head
	if (dx < -1 || dx > 1) || (dy < -1 || dy > 1) {
	    // mover tail pra ultima posição de head antes do movimento
	    tail[x] = head_visited[i - 1][x]
	    tail[y] = head_visited[i - 1][y]

	    tail_visited = append(tail_visited, []int{tail[y], tail[x]})
	}
    }

    for i := range tail_visited {
	unique_visits[pos{tail_visited[i][x], tail_visited[i][y]}] += 1
    }

    fmt.Println(len(unique_visits))

    // debug := make([][]int, 7)
    // for i := range debug {
    //     debug[i] = make([]int, 7)
    // }

    // for k := range tail_visited {
    //     debug[tail_visited[k][x]][tail_visited[k][y] * -1] += 1
    // }

    // for i := range debug {
    //     fmt.Println(debug[i])
    // }

}

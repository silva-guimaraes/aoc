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

func simulate_knot(head_visited [][]int) [][]int {

    tail := []int{0, 0}

    var knot_visited [][]int
    knot_visited = append(knot_visited, []int{tail[x], tail[y]})

    for i := range head_visited {

	dx := head_visited[i][x] - tail[x]
	dy := head_visited[i][y] - tail[y]

	// movimentos verticais/horizontais
	if (dx == 0 && (dy < -1 || dy > 1)) || (dy == 0 && (dx < -1 || dx > 1)) {
	    tail[x] = head_visited[i - 1][x]
	    tail[y] = head_visited[i - 1][y]


	} else if (dx < -1 || dx > 1) || (dy < -1 || dy > 1) {

	    // mover na diagonal quando o nó seguinte mover na diagonal. é isso o que o desafio quer.
	    nx := dx
	    if nx > 1 { nx-- }
	    if nx < -1 { nx++ }

	    ny := dy
	    if ny > 1 { ny-- }
	    if ny < -1 { ny++ }

	    tail[x] += nx
	    tail[y] += ny
	}

	knot_visited = append(knot_visited, []int{tail[x], tail[y]})
    }

    return knot_visited
}


func is_unique(list [][]int, yy, xx int) bool {
    for i := range list {
	if list[i][y] == yy && list[i][x] == xx {
	    return false
	}
    }
    return true

}

func main() {
    file, err := os.Open("./teste2.txt")
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

    head := []int{0, 0}

    unique_visits := make(map[pos]int)
    var head_visited [][]int

    head_visited = append(head_visited, []int{head[y], head[x]})
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

	    head_visited = append(head_visited, []int{head[y], head[x]})
	}
    }

    step := 9
    tail_visited := head_visited

    for i := 0; i < step; i++ {
	tail_visited = simulate_knot(tail_visited[1:])
	// fmt.Print(i + 1, " ")
	// for j := range tail_visited {
	//     fmt.Print(tail_visited[j], "\t\t")
	// }
	// fmt.Println()
    }

    // step := 10
    // tail_visited := make([][][]int, step)
    // tail_visited[0] = head_visited

    // for i := 0; i < step - 1; i++ {
    //     tail_visited[i + 1] = simulate_knot(tail_visited[i])
    // }

    for i := range tail_visited {
        unique_visits[pos{tail_visited[i][x], tail_visited[i][y]}] += 1
    }


    debug := make([][]int, 30)
    for i := range debug {
        debug[i] = make([]int, 50)
    }

    for i := range tail_visited {
        debug[tail_visited[i][x] + 18][tail_visited[i][y] + 28] = 9
    }

    for i := range debug {
	for j := 0; j < len(debug[0]); j++{

	    if debug[i][j] > 0 {
		fmt.Print("#")

	    } else {
		fmt.Print(" ")
	    }
	}
	fmt.Println()
    }

    fmt.Println(len(unique_visits))

}

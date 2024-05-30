package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type pos struct {
    i, j int
}

type path struct {
    pos pos
    count int
    prev pos
}

var directions []pos = []pos{
    {0, 1},
    {0, -1},
    {1, 0},
    {-1, 0},
}

type foobar struct { 
    current, prev pos
    visited []pos
}

type edge struct {
    weight float64
    node pos
}

func main() {

    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    grid := strings.Split(strings.TrimSpace(string(file)), "\n")

    start := pos{0, 1}
    end := pos{len(grid)-1, len(grid[0])-2}

    queue := []foobar {
        {start, start, []pos{start}},
    }

    graph := map[pos][]edge{}

    // peguei toda a ideia do código desse amigo
    // https://github.com/mgtezak/Advent_of_Code/blob/master/2023/Day_23.py
    // a ideia é comprimir todo o labirinto em um grafo bi direcional usando
    // as intercessões como nodes e logo após isso um bruteforce iterativo usando
    // DFS já que o problema de encontrar um caminho mais longo é considerado NP-hard.
    // uma coisa muito importante que eu aprendi hoje é que BFS comparado com DFS é
    // bastante intensivo em memória. de forma iterativa uma BFS se faz com uma queue,
    // caso queira fazer DFS uma stack seria necessária, que foi justamente o que o
    // nosso amigo fez. o único problema é que ele deu o nome dessa stack de queue
    // e eu não parei pra notar que o método .pop() removia o ultimo elemento da lista
    // como uma stack faz normalmente. quase congelei meu computador varias vezes 
    // com a memória (meus míseros 4gbs) indo alem dos 90%. esse erro do amigo me 
    // custou 3 dias no final das contas.
    for len(queue) > 0 {
        a := queue[0]
        queue = queue[1:]

        current := a.current

        if current == end {
            steps := len(a.visited)-1
            new := edge{float64(steps), current}

            graph[a.prev] = append(graph[a.prev], new)
            graph[current] = append(graph[a.prev], edge{float64(steps), a.prev})
            continue
        }

        neighbors := []pos{}

        for _, d := range directions {

            n := pos{d.i + current.i, d.j + current.j}

            if n.i < 0 || n.i >= len(grid) || n.j < 0 || n.j >= len(grid[0]) {
                continue
            }
            if  grid[n.i][n.j] ==  '#'{
                continue
            }
            if idx := slices.Index(a.visited, n); idx > -1 {
                continue
            }

            neighbors = append(neighbors, n)
        }

        if len(neighbors) == 1 {
            queue = append(
                queue, foobar{neighbors[0], a.prev, append(a.visited, current)},
            )
        } else if len(neighbors) > 1 {
            steps := len(a.visited)-1
            new := edge{float64(steps), current}

            if idx := slices.Index(graph[a.prev], new); idx > -1 {
                continue
            }

            graph[a.prev] = append(graph[a.prev], new)
            graph[current] = append(graph[current], edge{float64(steps), a.prev})
            for i := range neighbors {
                queue = append(queue, foobar{neighbors[i], current, []pos{
                    neighbors[i], current,
                }})
            }

        }


    }

    maxSteps := 0
    stack := []foobar{
        {current: start, prev: pos{0, 0}, visited: []pos{start}},
    }

    for len(stack) > 0 {
        a := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        steps := a.prev.i

        if a.current == end {
            maxSteps = max(steps, maxSteps)
            continue
        }
        for _, edge := range graph[a.current] {
            if idx := slices.Index(a.visited, edge.node); idx > -1 {
                continue
            }
            stack = append(stack, foobar{
                edge.node, 
                pos{a.prev.i + int(edge.weight), 0}, 
                append(a.visited, a.current)},
            )
        }
    }

    fmt.Println(maxSteps)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	// "unicode"
)

type graph map[string][]string

func (g graph) add_edge(from, to string) {
    g[from] = append(g[from], to)
    g[to] = append(g[to], from)
}

// func copy_slice[T any](slice []T) []T {
//     ret := make([]T, len(slice))
//     for i := range slice {
//         ret[i] = slice[i]
//     }
//     return ret
// }

func duplicates(visited []string) bool {
    count := make(map[string]int)
    for i := range visited {
        count[visited[i]]++
        if count[visited[i]] > 1 {
            return true
        }
    }
    return false
}

func did_visit(visited []string, current string) bool {
    for i := range visited {
        if visited[i] == current {
            return true
        }
    }
    return false
}

func (g graph) traverse(current string, visited []string) []string {

    if current == "end" {
        return []string{current} 
    }
    if unicode.IsLower(rune(current[0])) {
        if current == "start" && did_visit(visited, "start") {
            return []string{}
        }
        visited = append(visited, current)
    }

    var paths []string

    for _, node := range g[current] {

        if !duplicates(visited) || !did_visit(visited, node) {

            new_paths := g.traverse(node, visited)
            paths = append(paths, new_paths...)
        }
    }
    return paths
}

func main() {

    file, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var input []string
    for scanner.Scan() {
        input = append(input, strings.Split(scanner.Text(), "-")...)
    }
    
    paths := make(graph)

    for i := 0; i < len(input); i += 2 {
        paths.add_edge(input[i], input[i+1])
    }

    ret := paths.traverse("start", []string{})
    fmt.Println(len(ret))

}

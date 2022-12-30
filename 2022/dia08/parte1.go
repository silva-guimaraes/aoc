
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var i_len int
var j_len int

type grid [][]int
type visible [][]bool
// type horizontal func(grid, visible, int, int, int) int
// type corner func(grid, visible, int) int

// var left_right = func (grid grid, visible visible, tallest, i, j int) int {
//     if grid[i][j] > tallest {
// 	visible[i][j] = true
// 	return grid[i][j]
//     } else {
// 	return tallest
//     }
// }
// 
// var top_down = func (grid grid, visible visible, tallest, i, j int) int {
//     if j == 3 {
//         fmt.Println(grid[j][i] , tallest, grid[j][i] > tallest, i ,j)
//     }
//     if grid[j][i] > tallest {
// 	visible[j][i] = true
// 	return grid[j][i]
//     } else {
// 	return tallest
//     }
// }
// 
// 
// var right_left = func (grid grid, visible visible, tallest, i, j int) int {
//     return left_right(grid, visible, tallest, i_len - i, j_len - j)
// }
// var down_top = func (grid grid, visible visible, tallest, i, j int) int {
//     return top_down(grid, visible, tallest, i_len - i, j_len - j)
// }
// 
// 
// var left_corner = func (g grid, v visible, i int) int {
//     v[i][0] = true
//     return g[i][0]
// }
// var right_corner = func (g grid, v visible, i int) int {
//     v[i][j_len] = true
//     return g[i][j_len]
// }
// var top_corner = func (g grid, v visible, i int) int {
//     v[0][i] = true
//     return g[0][i]
// }
// var bottom_corner = func (g grid, v visible, i int) int {
//     // fmt.Println("bottom cornor", i, g[j_len][i])
//     v[j_len][i] = true
//     return g[j_len][i]
// }
// 
// func observe_from_edge(grid grid, visible visible, horizontal horizontal, corner corner)   {
// 
//     for i := 0 ; i < len(grid); i++{
// 
// 	tallest := corner(grid, visible, i)
// 
// 	for j := 0 ; j < len(grid[0]); j++ {
// 	    tallest = horizontal(grid, visible, tallest, i, j)
// 
// 	    if tallest == 9 {
// 		break
// 	    }
// 	}
//     }
// } 
// deixar isso aqui pra servir de exemplo pra eu nunca mais complicar problema nenhum.


func main() {
    file, err := os.Open("./grid.txt")
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

    // indice maximo das duas dimensões
    i_len = len(lines) - 1
    j_len = len(lines[0]) - 1

    // inicializar grids
    grid := make([][]int, i_len + 1) // Go não é pra preguiçoso bixo
    visible := make([][]bool, i_len + 1)
    for i := range grid {
	grid[i] = make([]int, j_len + 1)
	visible[i] = make([]bool, j_len + 1)
    }

    // transformar strings em numeros
    for i := range lines {
        for j := range lines[i] {
            grid[i][j], _ = strconv.Atoi(string(lines[i][j]))
        }
    }

    // iterar do lado de fora pro lado de dentro do grid em todos os quatro cantos, em linha reta.
    // marcar como arvore visivel cada numero que chegue a virar o numero maximo daquela reta,
    // ou seja, uma reta pode ter mais de uma arvore visivel
    // arvores são marcadas como visiveis em visible [][]bool pra evitar de marcar
    // a mesma arvore mais de uma vez.

    for i := 0 ; i < len(grid); i++{ // esquerda direita 

	// todas as arvores nos cantos são visivel automaticamente
	tallest := grid[i][0]
	visible[i][0] = true

	for j := 0 ; j < len(grid[0]); j++ {
	    if grid[i][j] > tallest {
		visible[i][j] = true
		tallest = grid[i][j]
	    }
	    if tallest == 9 {
		break
	    }
	}
    }
    for i := 0 ; i < len(grid); i++{ // direita esquerda

	tallest := grid[i][j_len]
	visible[i][j_len] = true

	for j := j_len ; j > 0; j-- {
	    if grid[i][j] > tallest {
		visible[i][j] = true
		tallest = grid[i][j]
	    }
	    if tallest == 9 {
		break
	    }
	}
    }
    for j := 0 ; j < len(grid); j++{ // cima baixo

	tallest := grid[0][j]
	visible[0][j] = true

	for i := 0 ; i < len(grid[0]); i++ {
	    if grid[i][j] > tallest {
		visible[i][j] = true
		tallest = grid[i][j]
	    }
	    if tallest == 9 {
		break
	    }
	}
    }
    for j := 0 ; j < len(grid); j++{ // baixo cima

	tallest := grid[i_len][j]
	visible[i_len][j] = true

	for i := i_len ; i > 0; i-- {
	    if grid[i][j] > tallest {
		visible[i][j] = true
		tallest = grid[i][j]
	    }
	    if tallest == 9 {
		break
	    }
	}
    }

    // observe_from_edge(grid, visible, left_right, left_corner)
    // observe_from_edge(grid, visible, right_left, right_corner)

    // observe_from_edge(grid, visible, top_down, top_corner)
    // observe_from_edge(grid, visible, down_top, bottom_corner) // inferno

    // contar arvores marcadas
    count := 0
    for i := range visible {
	for j := range visible[i] {
	    // fmt.Print("[", i, j, a[i][j], grid[i][j], "]\t") // debug
	    if visible[i][j] {
		count += 1
	    }
	}
	// fmt.Println()
    }

    fmt.Println(count)
}

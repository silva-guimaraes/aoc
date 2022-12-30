
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
    score := make([][]int, i_len + 1)

    for i := range grid { 
	grid[i] = make([]int, j_len + 1)
	score[i] = make([]int, j_len + 1) 
    }


    // transformar strings em numeros
    for i := range lines {
        for j := range lines[i] {
            grid[i][j], _ = strconv.Atoi(string(lines[i][j]))
        }
    }


    // em cada numero do grid partir nas quatro direções e contar o numero de arvores que são
    // menores que o ponto inicial

    for i := range grid {
	for j := range grid[i] {

	    spot := grid[i][j]
	    above_score, below_score, left_score, right_score := 0, 0, 0, 0

	    for above := i - 1; above >= 0; above-- { // cima

		above_score += 1
		if grid[above][j] >= spot {
		    break
		}
	    }
	    for below := i + 1; below <= i_len; below++ { // baixo

		below_score += 1
		if grid[below][j] >= spot {
		    break
		}
	    }
	    for left := j - 1; left >= 0; left-- { // esquerda

		left_score += 1
		if grid[i][left] >= spot {
		    break
		}
	    }
	    for right := j + 1; right <= j_len; right++ { // direita

		right_score += 1
		if grid[i][right] >= spot {
		    break
		}
	    }

	    score[i][j] = above_score * below_score * left_score * right_score
	}
    } 
    // procurar a arvore que tenha a maior pontuação calculada
    max := grid[0][0]
    for i := range grid {
	for j := range grid[i] {
	    if score[i][j] > max {
		max = score[i][j]
	    }
	    // fmt.Print("[", score[i][j], "]\t") debug
	}
    }

    fmt.Println(max)
}

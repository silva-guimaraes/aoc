package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"regexp"
)

const rock = '#'
const sand = '0'
const empty = 0

type pos struct {
    i, j int
}

func print_space(space [][]byte) {

    fmt.Println(len(space[0]))
    for i := range space {
	for j := range space[i] {
	    if space[i][j] > 0 {
		fmt.Print(string(space[i][j]))
	    } else {
		fmt.Print(" ")
	    }

	}
	fmt.Println()
    }
}

func check_void(unit pos, space [][]byte) bool {

    for i := unit.i; i < len(space); i++ {
	if space[i][unit.j] != empty {
	    return false
	}
    }
    return true
}

func draw_lines(a, b pos, space [][]byte) {

    // tendo a posição mais proxima de (0, 0), nós precisamos apenas iterar positivamente
    // enquanto o increment (inc, incj) diz qual sentido, horizontal ou vertical, seguir para 
    // posicionar o seguimento de pedras

    var i, j, end, start int
    inci, incj := 0, 0

    // decidir se as posições estão ou na vertical ou na horizontal e 
    // selecionar posição que mais estiver perto de (0, 0)

    if a.i == b.i && a.j != b.j { // posições na horizontal
	end = int(math.Max(float64(a.j), float64(b.j)))
	start = int(math.Min(float64(a.j), float64(b.j)))
	i = a.i
	j = start
	incj = 1

    } else { // posições na vertical
	end = int(math.Max(float64(a.i), float64(b.i)))
	start = int(math.Min(float64(a.i), float64(b.i)))
	i = start
	j = a.j
	inci = 1
    }

    for k := 0; k <= end - start; k++ {
	space[i + (k * inci)][j + (k * incj)] = rock
    }

}

func main() {
    file, err := os.Open("./rocks.txt")
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

    // lista com todos os seguimentos de pedras. 
    // seguimentos de pedras são pontos ligados em sequencia que forman uma parede de pedra
    traces := make([][]pos, len(lines))

    // valores minimos pra normalização
    min_x, max_x, max_y := 500, 0, 0

    // iterar por todas as linhas e extrair seguimentos
    for i := range lines {
	// extrair pares de numeros
	pairs := r.FindAllString(lines[i], -1)

	// salvar cada par como um ponto
	for j := 0; j < len(pairs); j += 2 {
	    x, _ := strconv.Atoi(pairs[j])
	    y, _ := strconv.Atoi(pairs[j + 1])

	    min_x = int(math.Min(float64(min_x), float64(x)))
	    max_x = int(math.Max(float64(max_x), float64(x)))
	    max_y = int(math.Max(float64(max_y), float64(y)))

	    traces[i] = append(traces[i], pos{y, x})
	}
    }
    // valores de X vem em numeros bem grandes e por isso normalização antes de
    // criar a array que vai conter toda a simulação
    min_x--
    for i := range traces {
        for j := range traces[i] {
            traces[i][j].j -= min_x
        }
    }
    // são duas da manhão e eu não sei por que isso não roda sem essas adições 
    space := make([][]byte, max_y + 3)
    for i := range space {
	space[i] = make([]byte, max_x - min_x + 2)
    }

    // desenhar cada seguimento de pedra
    for i := range traces {
	for j := 0; j < len(traces[i]) - 1; j++ {
	    draw_lines(traces[i][j], traces[i][j + 1], space)
	}
    }

    // numero de grãos de areia inseridos
    count := 0

    for {
	// controlar o grão de areia através dessa variavel e apenas inserir no espaço ao final
	// da simulação pra evitar por e retirar toda vez que for necessario alterarar a posição.
	current_unit := pos{0, 500 - min_x}

	for {
	    // fmt.Print("\r[", current_unit.j, ", ", current_unit.i, "] number: ", count)

	    // espaço abaixo?
	    if space[current_unit.i + 1][current_unit.j] == empty {

		// completo vazio abaixo?
		if check_void(current_unit, space) {
		    fmt.Println(count)
		    // print_space(space)
		    return

		} else { // não? então continuar caindo
		    current_unit.i++
		    continue
		}
	    // um passo pra esquerda e um pra baixo
	    } else if current_unit.j - 1 >= 0 &&
	    // space[current_unit.i][current_unit.j - 1] == empty &&
	    space[current_unit.i + 1][current_unit.j - 1] == empty {


		current_unit.i++
		current_unit.j--
		continue

	    // um passo pra direita e um pra baixo
	    } else if current_unit.j + 1 < len(space[0]) &&
	    // space[current_unit.i][current_unit.j + 1] == empty &&
	    space[current_unit.i + 1][current_unit.j + 1] == empty {


		current_unit.i++
		current_unit.j++
		continue

	    // sem espaço pra baixo nem diagonais então ficar aqui 
	    }  else {
		space[current_unit.i][current_unit.j] = sand
		break
	    }
	}
	count++
    }


}

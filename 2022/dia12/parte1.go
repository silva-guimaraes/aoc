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

// verficar se é possivel mover para essa casa
func is_valid(current pos, target pos, heightmap [][]int) bool {

    if target.i >= 0 && target.i < len(heightmap) && // evitar indices não permitidos
    target.j >= 0 && target.j < len(heightmap[0]) {

	current_height := heightmap[current.i][current.j]
	target_height := heightmap[target.i][target.j]

	// regras dizem que nós podemos apenas nos mover pra uma casa se ela ter ou a mesma altura,
	// ou um a mais de altura ou qualquer numero de alturas pra baixo.
	// fazemos isso só que ao contrario por que o programa começa do final (o topo) e faz caminho
	// até o inicio (a base)
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

    // ler arquivo
    // algumas linhas do mesmo tamanho que quando juntas formam um mapa sendo cada caractere uma casa
    // diferente
    var lines []string
    for scanner.Scan() {
	lines = append(lines, scanner.Text())
    }

    // transformar caracteres em numeros de 0 a 27
    heightmap := make([][]int, len(lines))
    for i := range heightmap {
	heightmap[i] = make([]int, len(lines[0]))
	for j := range heightmap[i] {
	    heightmap[i][j] = int(lines[i][j] - 'a')
	}
    }
    file.Close()


    // encontrar a posição do inicio e final do mapa
    starti, startj := index_of(start, heightmap)
    start := pos{starti, startj, -1}
    endi, endj := index_of(end, heightmap)
    end := pos{endi, endj, 0}

    // transformado em numero 'E' fica bem maior do que 'z'
    // tratar 'E' como 'z''
    heightmap[endi][endj] = 'z' - 'a'

    var visited stack // casas que ja vasculhamos e não queremos vasculhar novamente
    var queue stack // casas a serem vasculadas

    queue.queue(end) // começar do final e fazer caminho até o inicio

    // https://en.wikipedia.org/wiki/Pathfinding#Sample_algorithm
    for !queue.is_empty() {

	current, exists := queue.pop()
	if !exists {
	    panic(exists)
	}

	if did_visit(current, visited) {
	    // não deveria nem ser possivel adicionar uma casa ja visitada na queue mas o programa parece 
	    // não ligar pra isso e tenta visitar todas as casas do mesmo jeito. não sei como
	    continue
	}

	// procurar em todas as quatros direções casas que sigam as regras do desafio
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

	fmt.Print("\r[", current.i, " ", current.j,  "] queue: ", len(queue), " visited: ", len(visited))

	visited.push(current)

	if current.i == start.i && current.j == start.j { // parar quando casa inicial for encontrada
	    fmt.Println()
	    fmt.Println(current.counter)
	    return

	}
    }
    // for i := range visited { // debug
    //     heightmap[visited[i].i][visited[i].j] = 'V' - 'a'
    // }
    // for i := range heightmap {
    //     for j := range heightmap[i] {
    //         fmt.Print(string(heightmap[i][j] + 'a'))
    //     }
    //     fmt.Println()
    // }
    fmt.Println("nada?")
}

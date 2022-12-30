package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// inventei de usar recursão e não parei pra pensar naquela solução usando stacks que eu fiz no dia 5.
// se não fosse isso eu não teria feito essa parafernalha.

const first = 0

var r *regexp.Regexp

type folder struct {
    name string
    size int
    children []folder
}


// lê um ls até que algum comando seja encontrado
func read_ls(commands []string, index int) (int, int){

    sum := 0
    i := index

    for ; i < len(commands) && commands[i][first] != '$'; i++ {

	size, err := strconv.Atoi(r.FindString(commands[i]))

	if err != nil { // ignorar diretórios
	    continue
	} else { // somar tamanho dos arquivos
	    sum += size
	}
    }
    return sum, i // retornar o indice de onde essa função parou
}

// lê todos os comandos em serie e cria uma estrutura em arvore de todos os diretórios
func structure_directories(commands []string, index int, current *folder) (*folder, int) {

    if index >= len(commands){
	return current, -1
    }

    if commands[index][2:4] == "ls" { // $ ls

	total_sum, cont := read_ls(commands, index + 1)
	// cont de continuação pro proximo comando
	current.size = total_sum

	if cont == -1 {
	    return current, -1
	}

	return structure_directories(commands, cont, current)

    } else if commands[index][5:] == ".." { // $ cd ..  
	return current, index + 1 // retornar struct e seguir pro proximo comando

    } else { // $ cd diretório 
	name := commands[index][5:]
	new_folder := folder{name, 0, nil}

	child, cont := structure_directories(commands, index + 1, &new_folder)

	current.children = append(current.children, *child)

	if cont == -1 {
	    return current, -1
	}

	return structure_directories(commands, cont, current)
    }
}

func count_indirect_sizes(current *folder) int {

    for i := range current.children {
	current.size += count_indirect_sizes(&current.children[i])
    }

    return current.size
}

func sum_max_sizes(current *folder, max int, ret int64) int64 {

    for i := range current.children {
	if max > current.children[i].size {
	    ret += int64(current.children[i].size)
	}
	ret += sum_max_sizes(&current.children[i], max, 0)
    }

    return ret
}

// func list_directory_sizes(current *folder, ret []int) []int {
// 
//     for i := range current.children { 
// 	ret = append(ret, sum_max_sizes(&current.children[i], nil))
//     }
//     return []int{current.size}
// }

func main() {

    r, _ = regexp.Compile("^\\d+")

    file, err := os.Open("./dirs.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    // ler linhas
    var commands []string
    for scanner.Scan() {
	commands = append(commands, scanner.Text())
    }
    file.Close()

    root, _ := structure_directories(commands, 1, &folder{"root", 0, nil})

    count_indirect_sizes(root)

    fmt.Println(sum_max_sizes(root, 100000, 0))
}

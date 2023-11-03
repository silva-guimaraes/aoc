package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "sort"
    "strconv"
)

// eu tinha reclamado na primeira parte que seria melhor fazer a solução
// utilizando stacks, que recursão era muito dificil e que eu era burro
// demais pra por ter feito daquele jeito. fiz essa segunda parte utilizando 
// stacks como eu queria e acabou que não deu certo.
// voltei agora alguns meses depois pra tentar refazer esse problema usando
// erlang. não existe iteração em erlang então tive que fazer tudo com recursão do
// jeito que eu não queria e no fim consegui resolver as duas partes
// sem muitos problemas. estou descepcionado com Go.
//
// moral da história: recursão depende da lingua.


const first = 0

var r *regexp.Regexp

type folder struct {
    name string
    size int
    children []folder
}

type stack []folder

func (s *stack) is_empty() bool {
    return len(*s) == 0
}

func (s *stack) last() *folder {
    return &(*s)[len(*s) - 1]
}

func (s *stack) push(folder folder){
    *s = append(*s, folder)
}

func (s *stack) pop() (folder, bool) {
    if s.is_empty() {
        return folder{"nil", 0, nil}, false
    } else {
        index := len(*s) - 1
        element := (*s)[index]
        *s = (*s)[:index]
        return element, true
    }
}

func (s *stack) return_pop() {
    temp, _ := s.pop()
    s.last().children = append(s.last().children, temp)
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

func list_directory_sizes(current *folder, ret []int) []int {

    for i := range current.children {
        ret = append(ret, list_directory_sizes(&current.children[i], nil)...)
    }
    return append(ret, current.size)
}

func count_directories(current *folder, ret int) int {

    for i := range current.children {
        ret += count_directories(&current.children[i], 0)
    }

    return ret + 1
}

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

    root := folder{"root", 0, nil}
    var path stack
    path.push(root)

    commands = commands[1:]

    for i := 0; i < len(commands); i++ {

        if commands[i][2:4] == "ls" {
            for i += 1; i < len(commands) && commands[i][first] != '$'; i++ {

                size, err := strconv.Atoi(r.FindString(commands[i]))
                if err != nil {
                    continue
                } else {
                    path.last().size += size
                }
            }
        }
        if i >= len(commands) {
            break
        }
        if commands[i][5:] == ".." {
            path.return_pop()
        } else {
            name := commands[i][5:]
            new_folder := folder{name, 0, nil}
            path.push(new_folder)
        }
    }
    for len(path) > 1 {
        path.return_pop()
    }

    root, _ = path.pop()

    count_indirect_sizes(&root)

    fmt.Println(root)

    list := list_directory_sizes(&root, nil)

    sort.Slice(list, func (i, j int) bool {
        return list[i] > list[j]
    })

    max_index := 0
    for i := range list {
        if list[i] < 8381165 {
            break
        } else if list[max_index] > list[i] {
            max_index = i
        }

    }

    fmt.Println(list)
    fmt.Println(list[max_index]) // ta certa a resposta aaaaaaaaa
}

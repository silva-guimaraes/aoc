package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "math"
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


func simulate_knot(head_visited []pos) []pos {
    var tail_visited []pos
    tail := pos{0, 0}
    for i := range head_visited {

        dx := head_visited[i].x - tail.x
        dy := head_visited[i].y - tail.y
        d := int(math.Abs(float64(dx)) + math.Abs(float64(dy)))

        if d == 2 {
            if dx >  1 { tail.x++; }
            if dy >  1 { tail.y++; }
            if dx < -1 { tail.x--; }
            if dy < -1 { tail.y--; }
        } else if d > 2 {
            if dx > 0 { tail.x++; }
            if dy > 0 { tail.y++; }
            if dx < 0 { tail.x--; }
            if dy < 0 { tail.y--; }
        } 

        tail_visited = append(tail_visited, tail)
    }
    return tail_visited
}


func is_unique(list []pos, yy, xx int) bool {
    for i := range list {
        if list[i].y == yy && list[i].x == xx {
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

    tail := pos{0, 0}
    head := pos{0, 0}

    unique_visits := make(map[pos]int)
    var head_visited []pos
    var tail_visited []pos

    head_visited = append(head_visited, head)
    tail_visited = append(tail_visited, tail)
    // unique_visits[pos{tail.x, tail.y}]++

    // calcular todos os movimentos de head primeiro
    for i := range steps {
        for j := 0; j < steps[i].times; j++ {

            // mover head de acordo com as direções
            if steps[i].direction == "U" {
                head.y -= 1

            } else if steps[i].direction == "D" {
                head.y += 1

            } else if steps[i].direction == "L" {
                head.x -= 1

            } else if steps[i].direction == "R" {
                head.x += 1
            }

            // salvar posições
            head_visited = append(head_visited, head)
        }
    }

    // trata cada knot como um head e calcular o proximo tail.
    // o tail logo em seguida vira um head e assim por diante
    for i := 0; i < 9; i++ {
        head_visited = simulate_knot(head_visited)
    }

    tail_visited = head_visited


    for i := range tail_visited {
        unique_visits[tail_visited[i]] += 1
    }

    fmt.Println(len(unique_visits))

}

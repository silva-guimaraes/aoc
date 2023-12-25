package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
    "slices"
)

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    cubes := [][2]int{
        {0, 0},
    }

    sum := 0

    for _, line := range strings.Split(strings.TrimSpace(string(file)), "\n") {

        split := strings.Split(line, " ")

        times, err := strconv.ParseInt(split[2][2:len(split[2])-2], 16, 32)
        if err != nil {
            panic(err)
        }

        last := cubes[len(cubes)-1]

        sum += int(times)

        switch split[2][len(split[2])-2:len(split[2])-1] {
        case "0":
            cubes = append(cubes, [2]int{last[0] + int(times), last[1]})
        case "1":
            cubes = append(cubes, [2]int{last[0], last[1] - int(times)})
        case "2":
            cubes = append(cubes, [2]int{last[0] - int(times), last[1]})
        case "3":
            cubes = append(cubes, [2]int{last[0], last[1] + int(times)})
        }

    }

    slices.Reverse(cubes)

    // shoelace formula
    for i := 0; i < len(cubes)-1; i++ {
        sum += cubes[i][0] * cubes[i+1][1] - cubes[i+1][0] * cubes[i][1]
    }

    // não faço ideia o do porque desse +1. isso da a resposta correta.
    fmt.Println(sum / 2 + 1)

    // não tinha percebido isso antes e até aogra não sei o porque a resposta foi 
    // a correta mas a soma do perimetro também ta sendo dividia por 2 do jeito que
    // eu fiz.
    // normalmente não faria sentido isso na minha mente, a formula de shoelace calcula
    // a area de um polígono e o problema pedia a area + perimetro. mesmo assim o que
    // eu acabei fazendo foi (area + perimetro) / 2.
    // se eu sem querer tivesse percebido isso eu teria ficado mais uma hora coçando
    // a cabeça procurando algum bug.
}

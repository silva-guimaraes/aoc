package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mustAtoi64(s string) int64 {
    x, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return int64(x)
}

type pos struct {
    x, y int64
}

type machine struct{
    buttonA, buttonB, prize pos
}

// Não ouça as mentiras de Eric.
// Existe apenas uma ou nenhuma solução pra cada problema, nunca mais de uma.
// Era fútil procurar pela resposta minima. Mesmo assim essa foi a solução
// que eu, desatento, fiz:
func part1(machines []machine) {
    sum := int64(0)
    for _, machine := range machines {
        bA := machine.buttonA
        bB := machine.buttonB
        prize := machine.prize
        minTokens := math.Inf(1)
        for a := range int64(100) {
            bX := float64(prize.x - bA.x * a) / float64(bB.x)
            bY := float64(prize.y - bA.y * a) / float64(bB.y)
            if bX != bY {
                continue
            }
            minTokens = min(float64(a) * 3 + bX, minTokens)
        }
        if minTokens != math.Inf(1) {
            sum += int64(minTokens)
        }
    }
    fmt.Println(sum)
}

// Isso usa regra de Cramer. Claramente a resposta correta.
func part2(machines []machine) {
    extra := int64(10000000000000)
    sum := int64(0)

    for _, machine := range machines {
        bA := machine.buttonA
        bB := machine.buttonB
        prize := pos {
            x: machine.prize.x + extra,
            y: machine.prize.y + extra,
        }
        BX := bB.x
        BY := bB.y
        AX := bA.x
        AY := bA.y
        PY := prize.y
        PX := prize.x

        a := (PX*BY - BX*PY) / (AX*BY - BX*AY)
        b := (AX*PY - PX*AY) / (AX*BY - BX*AY)
        if AX*a + BX*b != PX || AY*a + BY*b != PY {
            continue
        }
        sum += a * 3 + b
    }
    fmt.Println(sum)
}

func main() {
    bytes, err := os.ReadFile("13.txt")
    if err != nil {
        panic(err)
    }
    _machines := strings.Split(strings.TrimSpace(string(bytes)), "\n\n")
    re := regexp.MustCompile(`\d+`)

    var machines []machine

    for _, m := range _machines {
        l := strings.Split(m, "\n")
        buttonA := re.FindAllString(l[0], -1)
        buttonB := re.FindAllString(l[1], -1)
        prize := re.FindAllString(l[2], -1)
        machines = append(machines, machine{
            buttonA: pos{
                mustAtoi64(buttonA[0]), mustAtoi64(buttonA[1]),
            },
            buttonB: pos{
                mustAtoi64(buttonB[0]), mustAtoi64(buttonB[1]),
            },
            prize: pos{
                mustAtoi64(prize[0]), mustAtoi64(prize[1]),
            },
        })
    }
    part1(machines)
    part2(machines)
}

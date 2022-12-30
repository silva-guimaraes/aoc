package main

import (
    "fmt"
    "os"
    "bufio"
)

const (
    rock = 1
    paper = 2
    scissors = 3
)

func play(opp int, you int) int {
    if opp == you { // empate
	return  3 + you
    }
    // se oponente tiver -1 do que a nossa mão a vitória é nossa
    opp += 1
    if opp == 4 { opp = 1 } // voltar para o inicio depois de tesoura
    if opp == you {
	return 6 + you
    } else { // derrota
	return  0 + you
    }
}


func main() {

    file, err := os.Open("./strategy_guide.txt")
    if err != nil {
	panic(err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    total_score := 0
    for scanner.Scan() {
	line := scanner.Text()
	// transformar inputs em numeros de 1 a 3
	opponent := int(byte(line[0]) - 'A') + 1
	you := int(byte(line[2]) - 'X') + 1

	score := play(opponent, you)
	total_score += score
	// fmt.Println(opponent, you, score, total_score)
    }

    fmt.Println(total_score)
}

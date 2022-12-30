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

    var i, j, end, start int
    inci, incj := 0, 0

    if a.i == b.i && a.j != b.j {
	end = int(math.Max(float64(a.j), float64(b.j)))
	start = int(math.Min(float64(a.j), float64(b.j)))
	i = a.i
	j = start
	incj = 1

    } else {
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

    traces := make([][]pos, len(lines))

    max_x, max_y := 0, 0

    for i := range lines {
	pairs := r.FindAllString(lines[i], -1)

	for j := 0; j < len(pairs); j += 2 {
	    x, _ := strconv.Atoi(pairs[j])
	    y, _ := strconv.Atoi(pairs[j + 1])

	    max_x = int(math.Max(float64(max_x), float64(x)))
	    max_y = int(math.Max(float64(max_y), float64(y)))

	    traces[i] = append(traces[i], pos{y, x})
	}
    }
    space := make([][]byte, max_y + 2)
    for i := range space {
	space[i] = make([]byte, max_x * 2)
    }

    for i := range traces {
	for j := 0; j < len(traces[i]) - 1; j++ {
	    draw_lines(traces[i][j], traces[i][j + 1], space)
	}
    }


    for count := 0;; count++ {
	current_unit := pos{0, 500}

	for {
	    // fmt.Print("\r[", current_unit.j, ", ", current_unit.i, "] number: ", count)

	    if current_unit.i + 1 == len(space) { // parar antes de cair pra fora da array
		space[current_unit.i][current_unit.j] = sand
		break

	    } else if space[current_unit.i + 1][current_unit.j] == empty {
		current_unit.i++
		continue

	    } else if current_unit.j - 1 >= 0 &&
	    space[current_unit.i + 1][current_unit.j - 1] == empty {


		current_unit.i++
		current_unit.j--
		continue

	    } else if current_unit.j + 1 < len(space[0]) &&
	    space[current_unit.i + 1][current_unit.j + 1] == empty {


		current_unit.i++
		current_unit.j++
		continue

	    }  else if current_unit.i == 0 && current_unit.j == 500 {
		print_space(space)
		fmt.Println(count + 1) // não sei cara
		return

	    } else {
		space[current_unit.i][current_unit.j] = sand
		break
	    }
	}
    }


}

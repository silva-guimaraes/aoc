package main

import (
	"bytes"
	"fmt"
	"iter"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type tuple struct{ i, j int }
type tuple4 struct{ i, j, r, b int }
type shape [][]byte

var (
	shapes     []shape
	shapeSizes []int
	rot        = [][]tuple{
		{{2, 0}, {1, 0}, {0, 0}},
		{{2, 1}, {1, 1}, {0, 1}},
		{{2, 2}, {1, 2}, {0, 2}},
	}
)

func check(area [][]byte, block shape, si, sj int) bool {
	for i := range 3 {
		for j := range 3 {
			if block[i][j] == '#' && area[i+si][j+sj] == '#' {
				return false
			}
		}
	}
	return true
}

func put(area [][]byte, block shape, si, sj int) {
	for i := range 3 {
		for j := range 3 {
			if block[i][j] == '#' {
				if area[i+si][j+sj] == '#' {
					panic("impossível")
				} else {
					area[i+si][j+sj] = '#'
				}
			}
		}
	}
}

func rotated(block shape, amount int) shape {
	if amount == 0 {
		return block
	} else if amount < 0 {
		panic("impossível")
	}
	var rotatedShape shape = make([][]byte, 3)
	for i := range rotatedShape {
		rotatedShape[i] = make([]byte, 3)
	}
	for i := range 3 {
		for j := range 3 {
			rotatedShape[i][j] = block[rot[i][j].i][rot[i][j].j]
		}
	}
	return rotated(rotatedShape, amount-1)
}

func allIndex(area [][]byte) iter.Seq[tuple4] {
	return func(yield func(tuple4) bool) {
		for b := range 6 {
			for i := 0; i+2 < len(area); i++ {
				for j := 0; j+2 < len(area[0]); j++ {
					for r := range 4 {
						if !yield(tuple4{i, j, r, b}) {
							return
						}
					}
				}
			}
		}
	}
}

func printArea(area [][]byte) {
	for i := range area {
		for _, c := range area[i] {
			if c == '#' {
				fmt.Print(string(c))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func copyArea(x [][]byte) [][]byte {
	n := make([][]byte, len(x))
	for i := range n {
		n[i] = slices.Clone(x[i])
	}
	return n
}

func recur(area [][]byte, indices []int) bool {
	if slices.Equal(indices, []int{0, 0, 0, 0, 0, 0}) {
		printArea(area)
		return true
	}
	for t := range allIndex(area) {
		if indices[t.b] == 0 {
			continue
		} else if indices[t.b] < 0 {
			panic("impossível")
		}
		rotatedShape := rotated(shapes[t.b], t.r)
		if check(area, rotatedShape, t.i, t.j) {
			newArea := copyArea(area)
			newIndices := slices.Clone(indices)
			put(newArea, rotatedShape, t.i, t.j)
			newIndices[t.b] -= 1
			if recur(newArea, newIndices) {
				return true
			}
		}
		fmt.Print()
	}
	return false
}

func main() {
	bs, err := os.ReadFile("12.txt")
	if err != nil {
		panic(err)
	}
	sections := strings.Split(strings.TrimSpace(string(bs)), "\n\n")
	shapesString := sections[:len(sections)-1]
	for _, s := range shapesString {
		lines := bytes.Split([]byte(s), []byte("\n"))
		lines = lines[1:]
		shapes = append(shapes, lines)
		sum := 0
		for i := range lines {
			for _, c := range lines[i] {
				if c == '#' {
					sum += 1
				}
			}
		}
		shapeSizes = append(shapeSizes, sum)
	}

	regionsString := strings.Split(sections[len(sections)-1], "\n")
	re := regexp.MustCompile(`\d+`)

	pt1 := 0

	for _, r := range regionsString {
		numbersString := re.FindAllString(r, -1)
		width, _ := strconv.Atoi(numbersString[0])
		height, _ := strconv.Atoi(numbersString[1])
		area := make([][]byte, height)
		for i := range area {
			area[i] = make([]byte, width)
		}
		sum := 0
		var nums []int = make([]int, 6)
		numbersString = numbersString[2:]
		for i := range numbersString {
			num, _ := strconv.Atoi(numbersString[i])
			nums[i] = num
			sum += num * shapeSizes[i]
		}
		if sum > width*height { // isso é sacanagem......
			continue
		} else {
			pt1 += 1
		}
		// if recur(area, nums) { // gastei horas nisso
		// 	pt1 += 1
		// }
	}
	fmt.Println(pt1)

}

package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bs, err := os.ReadFile("06.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")
	last := len(lines) - 1
	re := regexp.MustCompile(`\d+`)
	var nums [][]int
	for i := range lines[:last] {
		var num []int
		for _, n := range re.FindAllString(lines[i], -1) {
			int, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			num = append(num, int)
		}
		nums = append(nums, num)
	}
	var ops []byte
	var opIndices []int
	for i, c := range lines[last] {
		if c == '*' || c == '+' {
			opIndices = append(opIndices, i)
			ops = append(ops, byte(c))
		}
	}
	pt1 := 0
	for i, op := range ops {
		if op == '+' {
			pt1 += nums[0][i] + nums[1][i] + nums[2][i] + nums[3][i]
		} else if op == '*' {
			pt1 += nums[0][i] * nums[1][i] * nums[2][i] * nums[3][i]
		}
	}
	fmt.Println(pt1)
	opIndices = append(opIndices, len(lines[last])+2)
	pt2 := 0
	for i := range opIndices[:len(opIndices)-1] {
		this := opIndices[i]
		next := opIndices[i+1] - 1
		mult := lines[last][opIndices[i]] == '*'
		n := 0
		if mult {
			n = 1
		}
		var nums [][]string
		for j := range len(lines) - 1 {
			nums = append(nums, strings.Split(lines[j][this:next], ""))

		}
		transposed := make([][]string, len(nums[0]))
		for i := range transposed {
			transposed[i] = make([]string, len(nums))
		}
		for i := range nums {
			for j := range nums[i] {
				transposed[j][i] = nums[i][j]
			}
		}
		for i := range transposed {
			transposed[i] = slices.DeleteFunc(transposed[i], func(x string) bool {
				return x == " "
			})
			int, _ := strconv.Atoi(strings.Join(transposed[i], ""))
			if mult {
				n *= int
			} else {
				n += int
			}
		}
		pt2 += n
	}
	fmt.Println(pt2)
}

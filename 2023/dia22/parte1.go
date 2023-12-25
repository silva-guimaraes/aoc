package main

import (
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type pos []int

type cube []pos


func simulate(cubes []cube) ([][]bool, [][]bool) {

    // cubes := copyCubes(old)
    supports := make([][]bool, len(cubes))
    supported := make([][]bool, len(cubes))
    for i := range cubes {
        supports[i] = make([]bool, len(cubes))
        supported[i] = make([]bool, len(cubes))
    }

    for i := range cubes {
        for {
            if cubes[i][0][2] == 1 {
                break
            }
            rest := false
            for _, c := range cubes[i] {

                for j := i-1; j >= 0; j-- {
                    for _, c2 := range cubes[j] {
                        if c2[0] == c[0] && c2[1] == c[1] && c2[2] == c[2]-1{
                            supports[j][i] = true
                            supported[i][j] = true
                            rest = true
                        }
                    }
                }
            }

            if rest {
                break
            }

            for j := range cubes[i] {
                cubes[i][j][2] -= 1
            }
        }
    }

    return supports, supported
}

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    getNums := regexp.MustCompile(`\d+`)

    cubes := []cube{}

    for _, line := range  strings.Split(strings.TrimSpace(string(file)), "\n") {

        nums := []int{}
        for _, n := range  getNums.FindAllString(line, -1) {
            num, err := strconv.Atoi(n)
            if err != nil {
                panic(err)
            }
            nums = append(nums, num)
        }
        a := nums[:3]
        b := nums[3:]
        idx := 0

        // if b[0] < a[0] || b[1] < a[1] || b[2] < a[2] {
        //     panic(i)
        // }

        switch {
        case b[0] > a[0]:
            idx = 0
        case b[1] > a[1]:
            idx = 1
        case b[2] > a[2]:
            idx = 2
        }

        c := cube{}
        for i := a[idx]; i <= b[idx]; i++ {

            d := pos{}
            d = append(d, a...)

            d[idx] = i

            c = append(c, d)
        }

        cubes = append(cubes, c)
    }


    slices.SortFunc(cubes, func(a, b cube) int {
        return cmp.Compare(a[0][2], b[0][2])
    })


    count := 0
    supports, supported := simulate(cubes)
    fmt.Println(cubes)
    for i := range supports {
        c := 0
        for j := range supports[i] {
            if supports[i][j] {
                l := len(slices.DeleteFunc(
                    slices.Clone(supported[j]), func(a bool) bool {
                        return !a
                    }),
                )
                if l == 1 {
                    c++
                    break
                }
            }
        }
        if c == 0 {
            count++
        }
    }
    fmt.Println(count)
    
}

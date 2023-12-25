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


func copyCubes(a []cube) []cube {
    ret := []cube{}
    for _, c := range a {
        cube := []pos{}
        for _, p := range c {
            p1 := pos{}
            p1 = append(p1, p...)
            cube = append(cube, p1)
        }
        ret = append(ret, cube)
    }
    return ret
}

func simulate(cubes []cube) []cube {

    cubes = copyCubes(cubes)

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
                            rest = true
                            goto end
                        }
                    }
                }
            }
            end:

            if rest {
                break
            }

            for j := range cubes[i] {
                cubes[i][j][2] -= 1
            }
        }
    }

    return cubes
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


    cubes = simulate(cubes)
    count := 0
    for i := range cubes {
        test := slices.Delete(copyCubes(cubes), i, i+1)
        compare := simulate(test)

        for j := range compare {
            for k := range compare[j] {
                if compare[j][k][2] != test[j][k][2] {
                    count++
                    break
                } 
            }
        }
    }
    fmt.Println(count)
    
}

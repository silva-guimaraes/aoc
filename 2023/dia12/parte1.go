package main

import (
	// "bufio"
	"fmt"
	"math"
	"os"

	// "slices"
	"strconv"
	"strings"
	// "strings"
)

// ESSE FOI O BRUTEFORCE MAIS COMPLICADO QUE EU JA ESCREVI

type springs struct {
    onMask, offMask int
    springs string
    nums []int
}

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    var input []springs
    for _, l := range strings.Split(string(file), "\n") {
        if l == "" { continue }
        seq := strings.Split(l, " ")

        var current springs

        current.springs = seq[0]

        numbers := strings.Split(seq[1], ",")

        for _, n := range numbers {
            num, err := strconv.Atoi(n)
            if err != nil {
                panic(err)
            }
            current.nums = append(current.nums, num)
        }

        for _, s := range current.springs {
            current.onMask <<= 1
            current.offMask <<= 1
            if s != '.' {
                current.offMask |= 1
            }
            if s == '#' {
                current.onMask |= 1
            } 
        }
        input = append(input, current)
    }

    // fmt.Println(input)

    count := 0
    // input = input[1:2]

    for _, s := range input {
        lim := int(math.Pow(2, float64(len(s.springs))))
        // lim = 1
        // last := 0
        last := make(map[int]bool)
        for i := 0; i < lim; i++ {
            test := (i | s.onMask) & s.offMask

            _, exists := last[test]
            if exists { 
                continue 
            }
            last[test] = true

            // if test != 180030 { continue }

            // fmt.Println(test, i, s.onMask, s.offMask)

            inside := false
            var nums []int
            nums = append(nums, s.nums...)
            k := len(nums)-1
            valid := true

            // 1 dentro 0 0 -> fora
            // 2 dentro 0 >1 -> nao valido
            // 3 dentro 1 0 -> nao valido
            // 4 dentro 1 >1 -> continua
            // 5 fora 0 0 -> panic
            // 6 fora 0 >1 -> continua
            // 7 fora 1 0 -> panic
            // 8 fora 1 >1 -> dentro
            // 9 terminar >1 -> nao valido

            for j := range s.springs {
                a := (test >> j) & 1
                b := (test >> j)
                if k == -1 {
                    break
                }

                // fmt.Println(j, k, nums, inside, a)
                if b == 0 {
                    break

                } else if inside && a == 0 && nums[k] == 0 {
                    inside = false
                    k--

                } else if inside && a == 0 && nums[k] > 0 {
                    valid = false
                    break

                } else if inside && a == 1 && nums[k] == 0 {
                    valid = false
                    break

                } else if inside && a == 1 && nums[k] > 0 {
                    //

                } else if !inside && a == 0 && nums[k] > 0 {
                    //

                } else if !inside && a == 1 && nums[k] == 0 {
                    panic(7)

                } else if !inside && a == 1 && nums[k] > 0 {
                    inside = true
                }

                if inside {
                    nums[k]--
                }


            }

            // fmt.Println()
            if k < 0 {
                continue
            }
            if k == 0 && nums[k] == 0 && valid {
                count++
            }
        }
    }
    fmt.Println(count)
}

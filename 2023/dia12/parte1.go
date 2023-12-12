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

type springs struct {
    onMask, offMask int
    springs string
    nums []int
}

func main() {
    file, err := os.ReadFile("teste.txt")
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
            // s := current.springs[i]
            // 1xx10xx000x11
            // 0111011000111 and
            // 1001000000011 or 
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

    fmt.Println(input)

    count := 0
    input = input[:1]

    for _, s := range input {
        lim := int(math.Pow(2, float64(len(s.springs))))
        // lim = 1
        last := 0
        for i := 0; i < lim; i++ {
            test := (i | s.onMask) & s.offMask
            if test == last { continue }
            last = test
            fmt.Println(test, i, s.onMask, s.offMask)

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

                if inside && a == 0 && nums[k] == 0 {
                    inside = false

                } else if inside && a == 0 && nums[k] > 0 {
                    valid = false
                    break

                } else if inside && a == 1 && nums[k] == 0 {
                    valid = false
                    break

                } else if inside && a == 1 && nums[k] > 0 {
                    nums[k]--

                } else if !inside && a == 0 && nums[k] > 0 {
                    //

                } else if !inside && a == 1 && nums[k] == 0 {
                    panic(7)

                } else if !inside && a == 1 && nums[k] > 1 {
                    inside = true
                }


                // if nums[k] == 0 && a == 1 {
                //     break
                // }
                //
                // if inside && nums[k] > 0 && a == 0 {
                //     break
                //
                // } else if inside && nums[k] == 0 && a == 0 {
                //     k--
                //     inside = false
                //
                // } else if a == 1 && !inside {
                //     inside = true
                // }
                //
                // if inside {
                //     nums[k]--
                // }

            }
            if test == 87 {
                fmt.Println(nums, k)
            }
            if k == 0 && nums[k] == 0 && valid {
                fmt.Println(test)
                count++
            }
        }
    }
    fmt.Println(count)
}

// ???.### 1,1,3
// xxx0111
// .??..??...?##. 1,1,3

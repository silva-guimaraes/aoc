package main

import (
	"os"
    "fmt"
	"strconv"
	"strings"
    "slices"
)

type state struct {
    springs string
    nums string
}

func recur(springs []byte, nums []byte, cache map[state]int) int {
    // fmt.Println(string(springs))

    if len(springs) == 0 {
        if len(nums) == 0 { return 1 } else { return 0 }
    }

    if len(nums) == 0 {
        if idx := slices.Index(springs, '#'); idx > -1 { return 0 } else { return 1 } 
    }

    key := state{string(springs), string(nums)}
    if restult, ok := cache[key]; ok {
        return restult
    }

    result := 0
    if springs[0] == '.' || springs[0] == '?' {
        result += recur(springs[1:], nums, cache)
    }

    if springs[0] == '#' || springs[0] == '?' {
        if int(nums[0]) <= len(springs)  {
            idx := slices.Index(springs[:nums[0]], '.')
            if idx == -1 && (int(nums[0]) == len(springs) || springs[nums[0]] != '#') {
                if int(nums[0]) + 1 > len(springs) {
                    result += recur(springs[nums[0]:], nums[1:], cache)
                } else {
                    result += recur(springs[nums[0]+1:], nums[1:], cache)
                }
            }    
        }
    }

    cache[key] = result

    return result
}

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    cache := make(map[state]int)
    
    sum := 0 
    for _, l := range strings.Split(string(file), "\n") {
        if l == "" { continue }
        seq := strings.Split(l, " ")

        // springs := seq[0]
        springs := seq[0] + "?" + seq[0] + "?" + seq[0] + "?" + 
        seq[0] + "?" + seq[0]
        numbers := strings.Split(seq[1], ",")

        var temp []byte
        for _, n := range numbers {
            num, err := strconv.Atoi(n)
            if err != nil {
                panic(err)
            }
            if num > 255 {panic(255)} 
            temp = append(temp, byte(num))
        }

        var nums []byte
        nums = append(nums, temp...)
        nums = append(nums, temp...)
        nums = append(nums, temp...)
        nums = append(nums, temp...)
        nums = append(nums, temp...)

        sum += recur([]byte(springs), nums, cache)
    }

    // fmt.Println(cache)

    fmt.Println(sum)
}

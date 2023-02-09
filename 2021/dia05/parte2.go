
package main

import (
    "bufio"
    "regexp"
    "fmt"
    // "math"
    "strconv"
    "os"
)

const MARKED = -1

type point struct {
    x, y int
}

func main() {

    file, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    find_numbers, err := regexp.Compile("\\d+")
    if err != nil { panic(err) }


    var input []point
    for scanner.Scan() {

        ext := find_numbers.FindAllString(scanner.Text(), -1)

        conv1, _ := strconv.Atoi(ext[0])
        conv2, _ := strconv.Atoi(ext[1])
        conv3, _ := strconv.Atoi(ext[2])
        conv4, _ := strconv.Atoi(ext[3])

        input = append(input, point{conv1, conv2})
        input = append(input, point{conv3, conv4})
    }

    field := make(map[point]int)

    for i := 0; i < len(input); i += 2 {
        first, second := input[i], input[i + 1]

        if first.y == second.y {
            if first.x > second.x {
                first, second = second, first
            }
            for k := first.x; k <= second.x; k++ {
                field[point{k, first.y}]++
            }

        } else if first.x == second.x {
            if first.y > second.y {
                first, second = second, first
            }
            for k := first.y; k <= second.y; k++ {
                field[point{first.x, k}]++
            }
        } else {
            if first.x < second.x && first.y < second.y {
                for i := 0; i <= second.x - first.x; i++ {
                    field[point{first.x + i, first.y + i}]++
                }

            } else if first.x > second.x && first.y < second.y {
                for i := 0; i <= first.x - second.x; i++ {
                    field[point{first.x - i, first.y + i}]++
                }

            } else if first.x < second.x && first.y > second.y {
                for i := 0; i <= second.x - first.x; i++ {
                    field[point{first.x + i, first.y - i}]++
                }

            } else if first.x > second.x && first.y > second.y {
                for i := 0; i <= first.x - second.x; i++ {
                    field[point{first.x - i, first.y - i}]++
                }
            } 
        }
    }
    count := 0
    for _, v := range field {
        if v > 1 { count++ }
    }
    fmt.Println(count)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func find_by_len(x []string, n int) string {
    for i := range x {
        if len(x[i]) == n {
            temp := x[i]
            x[i] = ""
            return temp
        }
    }
    return ""
}

func to_binary(x string) uint8 {
    var ret uint8 = 0x0
    for i := range x {
        ret |= 0x1 << int(x[i] - 'a')
    }
    if ret == 0x0 { panic("to binary 0x0") }
    return ret
}

func count_bits(x uint8) int {
    count := 0
    for x != 0 {
        count += int(x & 1)
        x >>= 1
    }
    return count
}

// func print_binary(n uint8) {
//     for i := 7; i >= 0; i-- {
//         if n&(1<<uint(i)) != 0 {
//             fmt.Print("1")
//         } else {
//             fmt.Print("0")
//         }
//     }
//     fmt.Println()
// }

func main() {

    file, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    find_chars, err := regexp.Compile("\\w+")
    if err != nil { panic(err) }

    var input [][]string
    for scanner.Scan() {
        ext := find_chars.FindAllString(scanner.Text(), -1)
        input = append(input, ext)
    }

    // lookup := make(map[int]rune)

    count := 0
    for i := range input {
        // sabendo 1, 7 e 4, nós sabemos os segmentos abcdf
        one := to_binary(find_by_len(input[i], 2))
        seven := to_binary(find_by_len(input[i], 3))
        four := to_binary(find_by_len(input[i], 4))
        eight := to_binary(find_by_len(input[i], 7))

        var nine uint8 = 0x0
        for j := range input[i] {
            if len(input[i][j]) != 6 { continue }

            // 9 é o unico numero de 6 segmentos com os segmentos abcdf
            if foo := to_binary(input[i][j]); count_bits((four | seven) ^ foo) == 1 {
                nine = foo
                input[i][j] = ""
                break
            }
        }
        // com isso agora temos o segmentos e, sendo o unico segmento que o nove não tem

        var zero uint8
        for j := range input[i] {
            if len(input[i][j]) != 6 { continue }

            // "x ^ 0x80" remove o ultimo bit
            not_nine := (^nine ^ 0x80) 

            // 0 é o unico numero de 6 segmentos com os segmentos abcefg
            if foo := to_binary(input[i][j]); count_bits(foo ^ (not_nine | one)) == 3 {
                zero = foo
                input[i][j] = ""
                break
            }
        }

        var three uint8
        for j := range input[i] {
            if len(input[i][j]) != 5 { continue }

            // posição do unico segmento que o zero não tem
            not_three := (^zero ^ 0x80)

            // 3 é o unico numero de 5 segmentos com os segmentos cdf
            if foo := to_binary(input[i][j]); count_bits(foo & (not_three | one)) == 3 {
                three = foo
                input[i][j] = ""
                break
            }
        }


        six := to_binary(find_by_len(input[i], 6))

        var two uint8
        for j := range input[i] {
            if len(input[i][j]) != 5 { continue }

            not_nine := (^nine ^ 0x80) 

            if foo := to_binary(input[i][j]); foo & not_nine > 0 {
                two = foo
                input[i][j] = ""
                break
            }
        }

        var five uint8
        for j := range input[i] {
            if len(input[i][j]) == 5 { 
                five = to_binary(input[i][j])
                input[i][j] = ""
                break
            }
        }

        decoded := []uint8{zero, one, two, three, four, five, six, seven, eight, nine}

        var output []uint8
        for j := range input[i][10:] {
            output = append(output, to_binary(input[i][j + 10]))
        }

        new_string := "" // gambiarra
        for k := range output {
            for m := range decoded {
                if decoded[m] == output[k]  {
                    new_string += string(int('0') + m)
                    break
                }
            }
        }

        conv, err := strconv.Atoi(new_string)
        if err != nil { panic(err) }
        count += conv
    }

    fmt.Println(count)
}

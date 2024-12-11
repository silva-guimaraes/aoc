package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)



func part1(diskMap []byte) {
    sum := 0

    for pos, head, tail := 0, 0, len(diskMap)-1; head <= tail; head++ {
        headId := head/2
        for range diskMap[head] {
            sum += headId * pos
            pos++
        }

        head += 1
        n := diskMap[head]
        tailId := tail/2
        for range n {
            if diskMap[tail] == 0 {
                tail -= 2
                tailId = tail/2
            }
            if diskMap[head] == 0 {
                break
            }
            if head > tail {
                break
            }
            diskMap[head]--
            diskMap[tail]--
            sum += tailId * pos
            pos++
        }
    }
    fmt.Println(sum)
}


func part2(diskMap []byte) {


    // file não é um ótimo nome
    type file struct {
        id int
        size byte
    }
    files := make([]*file, 0, len(diskMap)/2)
    for i := 0; i < len(diskMap); i++ {
        files = append(files, &file{
            id: i/2,
            size: diskMap[i],
        })

        if i == len(diskMap)-1 {
            break
        }
        i += 1

        files = append(files, &file{
            id: -1,
            size: diskMap[i],
        })

    }
    for tail := len(files)-1; tail >= 0; tail-- {
        f := files[tail]
        if f.id == -1 {
            continue
        }
        for i := 1; i <= tail; i++ {
            space := files[i]
            if space.id != -1 {
                continue
            }
            if f.size > space.size {
                continue
            }
            space.size -= f.size
            files = slices.Insert(files, i, f)
            files[tail+1] = &file{
                id: -1,
                size: f.size,
            }
            break
        }
    }
    sum := 0
    pos := 0
    for _, file := range files {
        if file.id == -1 {
            pos += int(file.size)
            continue
        }
        for range file.size {
            sum += file.id * pos
            pos++
        }
    }
    fmt.Println(sum)
}


func main() {
    bytes, err := os.ReadFile("09.txt")
    if err != nil {
        panic(err)
    }
    line := strings.TrimSpace(string(bytes))
    diskMap1 := make([]byte, len(line))
    diskMap2 := make([]byte, len(line))
    for i := range line {
        diskMap1[i] = line[i] - '0'
        diskMap2[i] = line[i] - '0'
    }

    part1(diskMap1)
    part2(diskMap2)
}

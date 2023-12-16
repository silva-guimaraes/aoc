package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type lens struct {
    label string
    focalLength int
}

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    } 

    words := strings.Split(strings.TrimSpace(string(file)), ",")

    boxes := make([][]lens, 256)

    for _, w := range words {

        var label string
        focalLength := 0

        if w[len(w)-1] == '-' {
            label = w[:len(w)-1]
        } else {
            split := strings.Split(w, "=")
            label = split[0]
            focalLength, _ = strconv.Atoi(split[1])
        }

        hash := 0
        for _, c:= range label {
            hash += int(c)
            hash *= 17
            hash %= 256
        }

        idx := slices.IndexFunc(boxes[hash], func(l lens) bool {
            return label == l.label
        })
        lens := lens{
            label: label,
            focalLength: focalLength,
        }
        if idx > -1 && focalLength == 0 {
            boxes[hash] = slices.Delete(boxes[hash], idx, idx+1)

        } else if idx > -1 {
            slices.Replace(boxes[hash], idx, idx+1, lens)

        } else if idx == -1 && focalLength > 0 {
            boxes[hash] = append(boxes[hash], lens)
        }
    }

    sum := 0
    for k, v := range boxes {
        for i, lens := range v {
            sum += (k+1) * (i+1) * lens.focalLength
        }
    }
    fmt.Println(sum)

}

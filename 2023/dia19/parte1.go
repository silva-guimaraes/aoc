package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
    category int
    compare string
    num int
    destination string
}

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    s := strings.Split(strings.TrimSpace(string(file)), "\n\n")

    workflows := map[string][]rule{}

    // regex deve ser mais facíl de entender do que isso
    for _, x := range strings.Split(s[0], "\n") {
        a := strings.Split(x, "{")
        label := a[0]
        r := strings.Split(a[1][:len(a[1])-1], ",")
        def := r[len(r)-1]
        rules := []rule{}
        for _, y := range r[:len(r)-1] {
            n := strings.Split(y[2:], ":")
            num, err := strconv.Atoi(n[0])
            if err != nil { panic(err) }
            var category int
            switch y[0] {
            case 'x':
                category = 0
            case 'm':
                category = 1
            case 'a':
                category = 2
            case 's':
                category = 3
            }
            rules = append(rules, rule{
                category: category,
                compare: string(y[1]),
                num: num,
                destination: n[1],
            })
        }
        rules = append(rules, rule{
            category: 0,
            compare: ">",
            num: 0,
            destination: def,
        })
        workflows[label] = rules
    }

    parts := [][4]int{}

    for _, x := range strings.Split(s[1], "\n") {
        part := [4]int{}
        for i, p := range strings.Split(x[1:len(x)-1], ",") {
            num, err := strconv.Atoi(p[2:])
            if err != nil { panic(err) }
            part[i] = num
        }
        parts = append(parts, part)
    }

    sum := 0
    for _, p := range parts {
        label := "in"
        for label != "A" && label != "R" {
        // fmt.Println(p, label)

            rules := workflows[label]
            for _, r := range rules {

                // fmt.Println(p, p[r.category], r)

                if r.compare == ">" && p[r.category] > r.num {
                    label = r.destination
                    break

                } else if r.compare == "<" && p[r.category] < r.num {
                    label = r.destination
                    break

                } else {
                    continue
                }
            }
        }
        // fmt.Println(label)
        if label == "A" {
            sum += p[0] + p[1] + p[2] + p[3]
        }
    }
    fmt.Println(sum)
}

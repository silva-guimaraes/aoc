package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
    // "slices"
)

type rule struct {
    category int
    compare string
    num int
    destination string
}

type Range struct {
    start, end int
}


// isso funcionou de primeira kkk mal creio
// fiquei traumatizado com o dia 5 e ja esperava que isso seria mais complicado
func recur(range1 []Range, current string, workflows map[string][]rule) int64 {

    var sum int64 = 0

    if current == "A" {
        sum = 1
        for _, r := range range1 {
            sum *= int64(r.end - r.start) + 1
        }
        return sum
    } else if current == "R" {
        return 0
    }

    for _, r := range workflows[current] {
        if r.compare == ">" && 
        range1[r.category].start > r.num && range1[r.category].end > r.num{
            return recur(range1, r.destination, workflows) + sum

        } else if r.compare == ">" && 
        range1[r.category].start < r.num && range1[r.category].end > r.num{
            newRange := make([]Range, 4)
            copy(newRange, range1)
            newRange[r.category].start = r.num + 1
            sum += recur(newRange, r.destination, workflows)
            range1[r.category].end = r.num
        }

        if r.compare == "<" && 
        range1[r.category].start < r.num && range1[r.category].end < r.num{
            return recur(range1, r.destination, workflows) + sum

        } else if r.compare == "<" && 
        range1[r.category].start < r.num && range1[r.category].end > r.num{
            newRange := make([]Range, 4)
            copy(newRange, range1)
            newRange[r.category].end = r.num - 1
            sum += recur(newRange, r.destination, workflows)
            range1[r.category].start = r.num
        }
    }

    return sum
}

func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    s := strings.Split(strings.TrimSpace(string(file)), "\n\n")

    workflows := map[string][]rule{}

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

    fmt.Println(recur([]Range{
        {1, 4000},
        {1, 4000},
        {1, 4000},
        {1, 4000},
    }, "in", workflows))

}

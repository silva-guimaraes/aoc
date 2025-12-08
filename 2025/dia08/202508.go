package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type box struct {
	x, y, z int
	nodes   []int16
}

type pair struct {
	distance float32
	a, b     int16
}

var boxes []box

func euclidean(a, b box) float32 {
	return float32(math.Sqrt(
		math.Pow(float64(a.x)-float64(b.x), 2) +
			math.Pow(float64(a.y)-float64(b.y), 2) +
			math.Pow(float64(a.z)-float64(b.z), 2),
	))
}

func main() {
	bs, err := os.ReadFile("teste.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")
	for i := range lines {
		coords := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		boxes = append(boxes, box{x, y, z, nil})
	}
	var pairs []pair
	for i := range boxes {
		for j := range boxes {
			if i == j {
				continue
			}
			pairs = append(pairs, pair{euclidean(boxes[i], boxes[j]), int16(i), int16(j)})
		}
	}
	slices.SortFunc(pairs, func(a, b pair) int {
		return cmp.Compare(a.distance, b.distance)
	})
	for _, pair := range pairs {
		boxes[pair.a].nodes = append(boxes[pair.a].nodes, pair.b)
		boxes[pair.b].nodes = append(boxes[pair.b].nodes, pair.a)
	}
	var visited = make(map[int]bool)
	var circuits = []int{}
	for i := range boxes {
		if _, ok := visited[i]; ok || len(boxes[i].nodes) == 0 {
			continue
		}
		size := 0
		queue := []int{i}
		for len(queue) > 0 {
			pop := queue[0]
			queue = queue[1:]
			visited[pop] = true
			size += 1
			for _, j := range boxes[pop].nodes {
				if _, ok := visited[int(j)]; ok {
					continue
				} else {
					queue = append(queue, int(j))
				}
			}
		}
		circuits = append(circuits, size)
	}
	slices.Sort(circuits)
	slices.Reverse(circuits)
	fmt.Println(circuits)
	pt1 := circuits[0] * circuits[1] * circuits[2]
	fmt.Println(circuits)
	fmt.Println(pt1)
}

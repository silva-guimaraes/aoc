package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const not_visited = -1
const INFINITY = 999999

type valve struct {
	label string
	tunnels []*valve
	flow_rate, minute int
}

type queue []*valve

func (s *queue) is_empty() bool {
	return len(*s) == 0
}

func (s *queue) push(str *valve){
	*s = append(*s, str)
}

func (s *queue) queue(p *valve){
	*s = append(queue{p}, *s...)
}

func (s *queue) pop() (*valve, bool) {
	if s.is_empty() {
		return &valve{"", nil, 0, 0}, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func is_visited(valves queue, valve *valve) bool {

	for i := range valves {
		if valves[i] == valve {
			return true
		}
	}
	return false
}


// calcula a distancia em minutos se movendo de uma valvula para a outra
func minute_distance(valves queue, current, target int) *valve {

	if current == target {
		return nil
	}

	var q queue
	var visited queue

	for i := range valves {
		valves[i].minute = INFINITY
	}
	valves[current].minute = 0
	q.queue(valves[current])

	for !q.is_empty() {
		valve, _ := q.pop()

		for i := range valve.tunnels {
			if is_visited(visited, valve.tunnels[i]) {
				continue
			}
			q.queue(valve.tunnels[i])

			if valve.minute + 1 < valve.tunnels[i].minute {
				valve.tunnels[i].minute = valve.minute + 1;
			}
		}

		visited.queue(valve)
	}

	return valves[target]
}

func find_valve(valves []*valve, label string) *valve {

	for i := range valves {
		if valves[i].label == label {
			return valves[i]
		}
	}
	panic(label)
}

func find_valve_index(valves []*valve, target *valve) int {

	for i := range valves {
		if valves[i] == target {
			return i
		}
	}
	panic(target.label)
}

func all_valves_opened(valves queue, open map[*valve]bool) bool {

	for i := range valves {
		if open[valves[i]] == false {
			return false
		}
	}
	return true
}

/// aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa

func foobar(maps queue, open map[string]bool, current *valve, time, pressure int) int  {

	if all_valves_opened(maps, open) || time > 30 {
		return pressure
	}

	fmt.Println(current.label, time, pressure)

	max := 0

	for i := range maps {
		for j := range maps[i].tunnels {
	}
	return max
}


// func findPaths(graph map[int][]int, start int, end int) [][]int {
// 	var paths [][]int
// 	if start == end {
// 		return [][]int{{start}}
// 	}

// 	for _, nextNode := range graph[start] {
// 		for _, path := range findPaths(graph, nextNode, end) {
// 			paths = append(paths, append([]int{start}, path...))
// 		}
// 	}

// 	return paths
// }

func max_flow_rate(tunnels []*valve) *valve {

	max := tunnels[0]
	for i := range tunnels {
		if tunnels[i].flow_rate > max.flow_rate {
			max = tunnels[i]
		}
	}
	return max
}
func main() {
	file, err := os.Open("./teste.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	r, _ := regexp.Compile("\\d+")

	var valves queue
	for i := range lines {
		flow_rate, _ := strconv.Atoi(r.FindString(lines[i]))
		label := lines[i][6:8]
		valves = append(valves, &valve{label, nil, flow_rate, 0})
	}

	find_tunnels := regexp.MustCompile(`([A-Z][A-Z]),?`)

	for i := range lines {
		tunnels := find_tunnels.FindAllStringSubmatch(lines[i], -1)[1:]
		for j := range tunnels {
			valves[i].tunnels = append(valves[i].tunnels, find_valve(valves, tunnels[j][1]))
		}
	}

	graph := map[int][]int{
		1: {2, 3},
		2: {3, 4},
		3: {4},
		4: {},
	}
	paths := findPaths(graph, 1, 4)
	fmt.Println(paths)

}

package main

import (
	"fmt"
	"os"
	"strings"
)

type pos struct{ i, j int }

var directions = []pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func part1(grid []string, guardPos pos) {
	current_direction := 0
	visited := map[pos]bool{}

	for {
        if grid[guardPos.i][guardPos.j] == '#' {
            d := directions[current_direction]
            guardPos = pos{guardPos.i-d.i, guardPos.j-d.j}
        }

		visited[guardPos] = true
		next := directions[current_direction]

		nextPos := pos{
			guardPos.i + next.i,
			guardPos.j + next.j,
		}
		if nextPos.i < 0 || nextPos.i >= len(grid) || nextPos.j < 0 || nextPos.j >= len(grid[0]) {
			break
		} else if grid[nextPos.i][nextPos.j] == '#' {
			current_direction += 1
			current_direction %= 4

			next = directions[current_direction]

			nextPos = pos{
				guardPos.i + next.i,
				guardPos.j + next.j,
			}
		}
		guardPos = nextPos
	}
	fmt.Println(len(visited))
}

func simulate(grid []string, guardPos pos, obstacle pos) bool {
	current_direction := 0
	visited := map[pos]int{}

	if grid[obstacle.i][obstacle.j] == '#' || guardPos == obstacle {
		return false
	}

	for {

		if guardPos.i < 0 || guardPos.i >= len(grid) || guardPos.j < 0 || guardPos.j >= len(grid[0]) {
			return false

		} else if grid[guardPos.i][guardPos.j] == '#' || guardPos == obstacle {

            d := directions[current_direction]
            guardPos = pos{guardPos.i-d.i, guardPos.j-d.j}

			current_direction += 1
			current_direction %= 4

            continue
		}

		if v, ok := visited[guardPos]; ok && v&(0x1<<current_direction) > 0 {
			return true
		}

		visited[guardPos] |= 0x1 << current_direction
		next := directions[current_direction]

		nextPos := pos{
			guardPos.i + next.i,
			guardPos.j + next.j,
		}


		guardPos = nextPos
	}
}

func part2(grid []string, guardPos pos) {
    sum := 0
    for i, line := range grid {
        for j := range line {
            if simulate(grid, guardPos, pos{i,j}) {
                sum += 1
            }
        }
    }
    fmt.Println(sum)
}

func main() {
	bytes, err := os.ReadFile("./06.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	var guardPos pos
	for i, l := range lines {
		for j := range l {
			if l[j] == '^' {
				guardPos = pos{i, j}
			}
		}
	}


	part1(lines, guardPos)
	part2(lines, guardPos)
}

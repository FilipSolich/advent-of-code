package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func findGuard(lines []string) (int, int) {
	for i, line := range lines {
		for j, ch := range line {
			if ch == '^' {
				return j, i
			}
		}
	}
	return 0, 0
}

func loop(lines []string, x, y int) []string {
	direction := 0
	directionDiffs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for {
		nextX := x + directionDiffs[direction][0]
		nextY := y + directionDiffs[direction][1]
		if nextX < 0 || nextX > len(lines[0])-1 || nextY < 0 || nextY > len(lines)-1 {
			lines[y] = lines[y][:x] + "X" + lines[y][x+1:]
			break
		}
		if lines[nextY][nextX] == '#' {
			direction++
			direction = direction % 4
		} else {
			lines[y] = lines[y][:x] + "X" + lines[y][x+1:]
			x = nextX
			y = nextY
		}
	}
	return lines
}

func part1(input string) (int, error) {
	lines := strings.Fields(input)
	x, y := findGuard(lines)
	lines = loop(lines, x, y)

	fields := 0
	for _, line := range lines {
		for _, ch := range line {
			if ch == 'X' {
				fields++
			}
		}
	}

	return fields, nil
}

func plantObstacle(lines []string, x, y int) []string {
	if lines[y][x] == '#' || lines[y][x] == '^' {
		return nil
	}
	lines[y] = lines[y][:x] + "#" + lines[y][x+1:]
	return lines
}

func doCycle(lines []string, x, y int) bool {
	direction := 0
	directionDiffs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	lookup := make([][]int, len(lines)*len(lines[0]))

	for {
		nextX := x + directionDiffs[direction][0]
		nextY := y + directionDiffs[direction][1]
		if nextX < 0 || nextX > len(lines[0])-1 || nextY < 0 || nextY > len(lines)-1 {
			return false
		}
		if lines[nextY][nextX] == '#' {
			direction++
			direction = direction % 4
			idx := y*len(lines[0]) + x
			if slices.Contains(lookup[idx], direction) {
				return true
			}
			lookup[idx] = append(lookup[idx], direction)
		} else {
			x = nextX
			y = nextY
		}
	}
}

func part2(input string) (int, error) {
	lines := strings.Fields(input)
	x, y := findGuard(lines)

	count := 0
	for i, line := range lines {
		for j := range line {
			maze := plantObstacle(slices.Clone(lines), j, i)
			if maze == nil {
				continue
			}
			if doCycle(maze, x, y) {
				count++
			}
		}
	}

	return count, nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2024 Day 06", input, config, part1, part2)
}

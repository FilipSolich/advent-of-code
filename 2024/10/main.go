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

func compareCoord(x, y int) func([]int) bool {
	return func(coord []int) bool {
		return coord[0] == x && coord[1] == y
	}
}

func findNext(lines []string, x int, y int, next int, unique *[][]int) int {
	if x < 0 || x > len(lines[0])-1 || y < 0 || y > len(lines)-1 {
		return 0
	}

	if string(lines[y][x]) == "." {
		return 0
	}

	ch := utils.MustAtoi(string(lines[y][x]))
	if ch == next {
		if ch == 9 {
			if slices.ContainsFunc(*unique, compareCoord(x, y)) {
				return 0
			}
			*unique = append(*unique, []int{x, y})
			return 1
		}
		sum := 0
		for _, dir := range [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			sum += findNext(lines, x+dir[0], y+dir[1], next+1, unique)
		}
		return sum
	}
	return 0
}

func findNext2(lines []string, x int, y int, next int) int {
	if x < 0 || x > len(lines[0])-1 || y < 0 || y > len(lines)-1 {
		return 0
	}

	if string(lines[y][x]) == "." {
		return 0
	}

	ch := utils.MustAtoi(string(lines[y][x]))
	if ch == next {
		if ch == 9 {
			return 1
		}
		sum := 0
		for _, dir := range [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			sum += findNext2(lines, x+dir[0], y+dir[1], next+1)
		}
		return sum
	}
	return 0
}

func part1(input string) (int, error) {
	lines := strings.Fields(input)
	sum := 0
	for y, line := range lines {
		for x := range line {
			inc := findNext(lines, x, y, 0, &[][]int{})
			sum += inc
		}
	}

	return sum, nil
}

func part2(input string) (int, error) {
	lines := strings.Fields(input)
	sum := 0
	for y, line := range lines {
		for x := range line {
			inc := findNext2(lines, x, y, 0)
			sum += inc
		}
	}

	return sum, nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2024 Day 10", input, config, part1, part2)
}

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

func part1(input string) (int, error) {
	letters := map[rune][][]int{
		'X': {},
		'M': {},
		'A': {},
		'S': {},
	}
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	xLastIdx := len(lines[0]) - 1
	yLastIdx := len(lines) - 1
	for i, line := range lines {
		for j, ch := range line {
			if strings.Contains("XMAS", string(ch)) {
				letters[ch] = append(letters[ch], []int{i, j})
			}
		}
	}

	sum := 0
	for _, coord := range letters['X'] {
	direction:
		for _, inc := range [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}} {
			start := slices.Clone(coord)
			for _, nextLetter := range []rune{'M', 'A', 'S'} {
				start = []int{start[0] + inc[0], start[1] + inc[1]}
				if start[0] < 0 || start[0] > xLastIdx || start[1] < 0 || start[1] > yLastIdx {
					continue direction
				}
				if !slices.ContainsFunc(letters[nextLetter], func(n []int) bool {
					return n[0] == start[0] && n[1] == start[1]
				}) {
					continue direction
				}
			}
			sum++
		}
	}

	return sum, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	xLastIdx := len(lines[0]) - 1
	yLastIdx := len(lines) - 1
	sum := 0
	for i, line := range lines {
	l:
		for j, ch := range line {
			if ch == 'A' {
				var letter rune = 0
				for _, inc := range [][]int{{1, 1}, {-1, -1}} {
					x := j + inc[0]
					y := i + inc[1]
					if x < 0 || x > xLastIdx || y < 0 || y > yLastIdx {
						continue l
					}
					if lines[y][x] == 'M' || lines[y][x] == 'S' {
						if letter == 0 {
							letter = rune(lines[y][x])
						} else if letter == rune(lines[y][x]) {
							continue l
						}
					} else {
						continue l
					}
				}
				letter = 0
				for _, inc := range [][]int{{1, -1}, {-1, 1}} {
					x := j + inc[0]
					y := i + inc[1]
					if x < 0 || x > xLastIdx || y < 0 || y > yLastIdx {
						continue l
					}
					if lines[y][x] == 'M' || lines[y][x] == 'S' {
						if letter == 0 {
							letter = rune(lines[y][x])
						} else if letter == rune(lines[y][x]) {
							continue l
						}
					} else {
						continue l
					}
				}
				sum++
			}
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

	core.RunSolutions("AoC 2024 Day 04", input, config, part1, part2)
}

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

func parse(input string) ([]string, map[rune][][]int) {
	lines := strings.Fields(input)
	out := map[rune][][]int{}
	for y, line := range lines {
		for x, antena := range line {
			if antena != '.' {
				out[antena] = append(out[antena], []int{x, y})
			}
		}
	}

	return lines, out
}

func countAntinodes(lines []string) int {
	count := 0
	for _, line := range lines {
		for _, antinode := range line {
			if antinode == '#' {
				count++
			}
		}
	}
	return count
}

func antinodeCoord(x []int, y []int) []int {
	diff := []int{(y[0] - x[0]) * (-1), (y[1] - x[1]) * (-1)}
	return []int{x[0] + diff[0], x[1] + diff[1]}
}

func antinodeCoordDiff(x []int, y []int) []int {
	return []int{(y[0] - x[0]) * (-1), (y[1] - x[1]) * (-1)}
}

func part1(input string) (int, error) {
	lines, in := parse(input)
	for _, coords := range in {
		for idx, coord := range coords {
			rest := slices.Delete(slices.Clone(coords), idx, idx+1)
			for _, other := range rest {
				antinode := antinodeCoord(coord, other)
				if antinode[0] < 0 || antinode[0] > len(lines[0])-1 || antinode[1] < 0 || antinode[1] > len(lines)-1 {
					continue
				}
				lines[antinode[1]] = lines[antinode[1]][:antinode[0]] + "#" + lines[antinode[1]][antinode[0]+1:]
			}
		}
	}

	count := countAntinodes(lines)
	return count, nil
}

func part2(input string) (int, error) {
	lines, in := parse(input)
	antinodes := slices.Clone(lines)
	for _, coords := range in {
		for idx, coord := range coords {
			rest := slices.Delete(slices.Clone(coords), idx, idx+1)
			for _, other := range rest {
				antinodes[coord[1]] = antinodes[coord[1]][:coord[0]] + "#" + antinodes[coord[1]][coord[0]+1:]
				diff := antinodeCoordDiff(coord, other)
				x, y := coord[0], coord[1]
				for {
					x += diff[0]
					y += diff[1]
					if x < 0 || x > len(lines[0])-1 || y < 0 || y > len(lines)-1 {
						break
					}
					antinodes[y] = antinodes[y][:x] + "#" + antinodes[y][x+1:]
				}
			}
		}
	}

	count := countAntinodes(antinodes)
	return count, nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2024 Day 08", input, config, part1, part2)
}

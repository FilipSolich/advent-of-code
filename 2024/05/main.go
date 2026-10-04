package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func parse(input string) (map[int][]int, []string, error) {
	data := map[int][]int{}

	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[1], "\n")
	lines = lines[:len(lines)-1]

	sc := bufio.NewScanner(strings.NewReader(parts[0]))
	for sc.Scan() {
		nums := sc.Text()
		var x, y int
		_, err := fmt.Sscanf(nums, "%d|%d", &x, &y)
		if err != nil {
			return nil, nil, err
		}
		_, ok := data[x]
		if !ok {
			data[x] = make([]int, 0)
		}
		data[x] = append(data[x], y)
	}

	return data, lines, nil
}

func part1(input string) (int, error) {
	data, lines, err := parse(input)
	if err != nil {
		return 0, err
	}

	sum := 0

nextL:
	for _, line := range lines {
		nums := strings.Split(line, ",")
		seen := map[int]struct{}{}
		for _, ch := range nums {
			num := utils.MustAtoi(ch)
			seen[num] = struct{}{}
			yy, ok := data[num]
			if !ok {
				continue
			}
			for _, y := range yy {
				_, ok := seen[y]
				if ok {
					continue nextL
				}
			}
		}
		inc := utils.MustAtoi(nums[len(nums)/2])
		sum += inc
	}

	return sum, nil
}

func sortLine(data map[int][]int, line []int) []int {
	out := []int{line[0]}
next:
	for _, num := range line[1:] {
		for i, inLine := range out {
			if slices.Contains(data[num], inLine) {
				out = slices.Insert(out, i, num)
				continue next
			}
			if i == len(out)-1 {
				out = append(out, num)
			}
		}
	}
	return out
}

func part2(input string) (int, error) {
	data, lines, err := parse(input)
	if err != nil {
		return 0, err
	}

	badLines := []int{}

nextL:
	for i, line := range lines {
		nums := strings.Split(line, ",")
		seen := map[int]struct{}{}
		for _, ch := range nums {
			num := utils.MustAtoi(ch)
			seen[num] = struct{}{}
			yy, ok := data[num]
			if !ok {
				continue
			}
			for _, y := range yy {
				_, ok := seen[y]
				if ok {
					badLines = append(badLines, i)
					continue nextL
				}
			}
		}
	}

	sum := 0
	for _, idx := range badLines {
		var line []int
		for _, num := range strings.Split(lines[idx], ",") {
			line = append(line, utils.MustAtoi(num))
		}
		line = sortLine(data, line)

		sum += line[len(line)/2]
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

	core.RunSolutions("AoC 2024 Day 05", input, config, part1, part2)
}

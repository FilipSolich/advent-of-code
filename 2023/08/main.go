package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func parseLines(lines []string) map[string][2]string {
	m := map[string][2]string{}
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		dest := strings.Split(parts[1][1:][:len(parts[1])-2], ", ")
		m[parts[0]] = [2]string{dest[0], dest[1]}
	}
	return m
}

func part1(input string) (int, error) {
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]

	in := lines[0]
	lines = lines[2:]
	nodes := parseLines(lines)
	start := "AAA"

	steps := 0
outer:
	for {
		for _, r := range in {
			s := string(r)
			if start == "ZZZ" {
				break outer
			}
			if s == "L" {
				start = nodes[start][0]
			} else {
				start = nodes[start][1]
			}
			steps++
		}
	}

	return steps, nil
}

func part2(input string) (int, error) {
	return 0, errors.New("not implemented")
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2023 Day 08", input, config, part1, part2)
}

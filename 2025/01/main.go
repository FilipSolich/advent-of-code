package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func part1(input string) (int, error) {
	var (
		res = 0
		cur = 50
		num int
		dir string
	)
	for line := range strings.Lines(input) {
		dir = line[:1]
		numStr := line[1 : len(line)-1]
		num = utils.MustAtoi(numStr)
		switch dir {
		case "L":
			cur -= num
		case "R":
			cur += num
		}
		cur = cur % 100
		if cur == 0 {
			res++
		}
	}

	return res, nil
}

func part2(input string) (int, error) {
	var (
		res = 0
		cur = 50
		num int
		dir string
	)
	for line := range strings.Lines(input) {
		dir = line[:1]
		numStr := line[1 : len(line)-1]
		num = utils.MustAtoi(numStr)
		for range num {
			inc := 1
			if dir == "L" {
				inc = -1
			}
			cur += inc
			switch cur {
			case 0:
				res++
			case -1:
				cur = 99
			case 100:
				res++
				cur = 0
			default:
			}
		}
	}

	return res, nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2025 Day 01", input, config, part1, part2)
}

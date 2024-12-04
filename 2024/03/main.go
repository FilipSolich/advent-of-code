package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func part1(input string) (int, error) {
	out := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`).FindAll([]byte(input), -1)
	sum := 0
	for _, n := range out {
		var x, y int
		_, err := fmt.Sscanf(string(n), "mul(%d,%d)", &x, &y)
		if err != nil {
			return 0, err
		}
		sum += x * y
	}
	return sum, nil
}

func part2(input string) (int, error) {
	input = "do()" + input
	doSplit := strings.Split(input, "do()")
	sum := 0
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	for _, part := range doSplit {
		dontSplit := strings.Split(part, "don't()")
		muls := re.FindAll([]byte(dontSplit[0]), -1)
		for _, n := range muls {
			var x, y int
			_, err := fmt.Sscanf(string(n), "mul(%d,%d)", &x, &y)
			if err != nil {
				return 0, err
			}
			sum += x * y
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

	core.RunSolutions("AoC 2024 Day 03", input, config, part1, part2)
}

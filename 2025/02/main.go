package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func isInvalid(x int) bool {
	str := strconv.Itoa(x)
	if len(str)%2 != 0 {
		return false
	}

	part1 := str[:len(str)/2]
	part2 := str[len(str)/2:]

	return part1 == part2
}

func isInvalid2(x int) bool {
	str := strconv.Itoa(x)
outter:
	for i := 1; i <= len(str)/2; i++ {
		if len(str)%i != 0 {
			continue
		}

		seq := str[:i]
		l := 0
		h := i
		for h < len(str)+1 {
			if str[l:h] != seq {
				continue outter
			}
			l = h
			h += i
		}
		return true
	}

	return false
}

func part1(input string) (int, error) {
	res := 0
	input = strings.TrimRight(input, "\n")
	for r := range strings.SplitSeq(input, ",") {
		var low, high int
		fmt.Sscanf(r, "%d-%d", &low, &high) // nolint:errcheck
		for i := low; i < high; i++ {
			if isInvalid(i) {
				res += i
			}
		}
	}
	return res, nil
}

func part2(input string) (int, error) {
	res := 0
	input = strings.TrimRight(input, "\n")
	for r := range strings.SplitSeq(input, ",") {
		var low, high int
		fmt.Sscanf(r, "%d-%d", &low, &high) // nolint:errcheck
		for i := low; i <= high; i++ {
			if isInvalid2(i) {
				res += i
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

	core.RunSolutions("AoC 2025 Day 02", input, config, part1, part2)
}

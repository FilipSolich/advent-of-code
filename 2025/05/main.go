package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func part1(input string) (int, error) {
	var fresh int
	var ranges [][]int
	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		ids := strings.Split(line, "-")
		switch len(ids) {
		case 1:
			if ids[0] == "" {
				continue
			}
			id, err := strconv.Atoi(ids[0])
			if err != nil {
				return 0, err
			}
			for _, r := range ranges {
				if id >= r[0] && id <= r[1] {
					fresh++
					break
				}
			}
		case 2:
			low, err := strconv.Atoi(ids[0])
			if err != nil {
				return 0, err
			}
			high, err := strconv.Atoi(ids[1])
			if err != nil {
				return 0, err
			}
			ranges = append(ranges, []int{low, high})
		default:
			continue
		}
	}
	return fresh, nil
}

func part2(input string) (int, error) {
	var ranges [][]int
	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		ids := strings.Split(line, "-")
		switch len(ids) {
		case 2:
			low, err := strconv.Atoi(ids[0])
			if err != nil {
				return 0, err
			}
			high, err := strconv.Atoi(ids[1])
			if err != nil {
				return 0, err
			}
			ranges = append(ranges, []int{low, high})
		default:
			continue
		}
	}
	for {
		change := false
		for i, r := range ranges {
			for j, rr := range ranges {
				popRange := false
				if r[0] < rr[0] && r[1] >= rr[0] {
					ranges[j][0] = r[0]
					popRange = true
				}
				if r[1] > rr[1] && r[0] <= rr[1] {
					ranges[j][1] = r[1]
				}
				if popRange {
					ranges = slices.Delete(ranges, i, i)
					change = true
				}
			}
		}
		if !change {
			break
		}
	}
	lowR := make(map[int]struct{})
	var fresh int
	for _, cr := range ranges {
		_, ok := lowR[cr[0]]
		if ok {
			continue
		}
		fresh += cr[1] - cr[0] + 1
		lowR[cr[0]] = struct{}{}
	}
	return fresh, nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2025 Day 05", input, config, part1, part2)
}

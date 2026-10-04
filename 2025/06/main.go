package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func part1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	var nums [][]int
	var op []string
	for i := 0; i < len(lines)-1; i++ {
		numStrs := strings.Fields(lines[i])
		var row []int
		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %v", err)
			}
			row = append(row, num)
		}
		nums = append(nums, row)
	}
	op = strings.Fields(lines[len(lines)-1])

	var total int
	for i := range len(nums[0]) {
		var sum int
		switch op[i] {
		case "+":
			sum = 0
			for j := range len(nums) {
				sum += nums[j][i]
			}
		case "*":
			sum = 1
			for j := range len(nums) {
				sum *= nums[j][i]
			}
		}
		total += sum
	}

	return total, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	col := 0
	for {
		for row := len(lines) {
			if

		}

		col++
	}

	return 0, errors.New("not implemented")
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2025 Day 06", input, config, part1, part2)
}

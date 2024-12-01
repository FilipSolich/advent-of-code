package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/math"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func part1(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))

	var left, right []int
	for sc.Scan() {
		nums := strings.Split(sc.Text(), "   ")
		l, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		r, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, err
		}
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := range len(left) {
		sum += math.Abs(left[i] - right[i])
	}

	return sum, nil
}

func part2(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))

	left := make(map[int]int)
	right := make(map[int]int)
	for sc.Scan() {
		nums := strings.Split(sc.Text(), "   ")
		l, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		r, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, err
		}
		left[l]++
		right[r]++
	}

	sum := 0
	for k, v := range left {
		sum += v * (k * right[k])
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

	core.RunSolutions("AoC 2024 Day 01", input, config, part1, part2)
}

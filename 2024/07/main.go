package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func part1(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))

	count := 0
l:
	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		test := utils.MustAtoi(parts[0][:len(parts[0])-1])
		nums := []int{}
		for _, n := range parts[1:] {
			nums = append(nums, utils.MustAtoi(n))
		}

		for i := 0; i < int(math.Pow(2, float64(len(nums)-1))); i++ {
			sum := nums[0]
			for idx := range len(nums) - 1 {
				if i&(0b1<<idx) == 0 {
					sum += nums[idx+1]
				} else {
					sum *= nums[idx+1]
				}
			}
			if test == sum {
				count += test
				continue l
			}
		}
	}
	return count, nil
}

func apply(x, y int) []int {
	return []int{x + y, x * y, utils.MustAtoi(fmt.Sprintf("%d%d", x, y))}
}

func part2(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))

	count := 0

	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		test := utils.MustAtoi(parts[0][:len(parts[0])-1])
		nums := []int{}
		for _, n := range parts[1:] {
			nums = append(nums, utils.MustAtoi(n))
		}

		acc := []int{nums[0]}
		newAcc := []int{}
		for i := range len(nums) - 1 {
			for _, n := range acc {
				newAcc = append(newAcc, apply(n, nums[i+1])...)
			}
			acc = slices.Clone(newAcc)
			newAcc = []int{}

			if i == len(nums)-2 && slices.Contains(acc, test) {
				count += test
			}
		}
	}
	return count, nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2024 Day 07", input, config, part1, part2)
}

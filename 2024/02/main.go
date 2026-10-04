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
	valid := 0
	for sc.Scan() {
		line := sc.Text()
		numbersStr := strings.Split(line, " ")
		var nums []int
		for _, n := range numbersStr {
			num, err := strconv.Atoi(n)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
		}
		inc := nums[0] < nums[1]
		validLine := true
		for i := range len(nums) - 1 {
			if !validate(nums[i], nums[i+1], inc) {
				validLine = false
				break
			}
		}
		if validLine {
			valid++
		}
	}

	return valid, nil
}

func validate(a int, b int, inc bool) bool {
	sub := a - b
	abs := math.Abs(a - b)
	return (abs >= 1 && abs <= 3) && ((inc && sub < 0) || (!inc && sub > 0))
}

func part2(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))
	valid := 0
	for sc.Scan() {
		line := sc.Text()
		numbersStr := strings.Split(line, " ")
		var nums []int
		for _, n := range numbersStr {
			num, err := strconv.Atoi(n)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
		}
	outter:
		for i := -1; i < len(nums); i++ {
			slice := slices.Clone(nums)
			if i >= 0 {
				slice = slices.Delete(slice, i, i+1)
			}

			inc := slice[0] < slice[1]
			for j := range len(slice) - 1 {
				if !validate(slice[j], slice[j+1], inc) {
					continue outter
				}
			}
			valid++
			break
		}
	}
	return valid, nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2024 Day 02", input, config, part1, part2)
}

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

var cache map[int][]int

func parse(input string) []int {
	strNums := strings.Fields(input)

	var nums []int
	for _, num := range strNums {
		nums = append(nums, utils.MustAtoi(num))
	}
	return nums
}

func strToInt(s string) int {
	if s == strings.Repeat("0", len(s)) {
		return 0
	}
	return utils.MustAtoi(s)
}

func blink(stones []int) []int {
	newStones := make([]int, 0, len(stones))
	for _, stone := range stones {
		str := fmt.Sprint(stone)
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if len(str)%2 == 0 {
			first, second := str[:len(str)/2], str[len(str)/2:]
			newStones = append(newStones, strToInt(first), strToInt(second))
		} else {
			newStones = append(newStones, stone*2024)
		}
	}
	return newStones
}

func nextStone(stone int) []int {
	cached, ok := cache[stone]
	if ok {
		return cached
	}

	var next []int
	str := fmt.Sprint(stone)
	if stone == 0 {
		next = []int{1}
	} else if len(str)%2 == 0 {
		first, second := str[:len(str)/2], str[len(str)/2:]
		next = []int{strToInt(first), strToInt(second)}
	} else {
		next = []int{stone * 2024}
	}
	cache[stone] = next
	return next
}

func blink2(stones map[int]int) map[int]int {
	result := map[int]int{}
	for stone, count := range stones {
		values := nextStone(stone)
		for _, value := range values {
			result[value] += count
		}
	}

	return result
}

func part1(input string) (int, error) {
	stones := parse(input)
	cache = map[int][]int{}
	for range 25 {
		stones = blink(stones)
	}
	return len(stones), nil
}

func part2(input string) (int, error) {
	in := parse(input)
	stones := map[int]int{}
	for _, stone := range in {
		stones[stone]++
	}

	cache = map[int][]int{}
	for range 75 {
		stones = blink2(stones)
	}

	sum := 0
	for _, count := range stones {
		sum += count
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

	core.RunSolutions("AoC 2024 Day 11", input, config, part1, part2)
}

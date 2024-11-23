package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func getNumStr(str string) string {
	for _, n := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		if str == n {
			return str
		}
	}
	switch str {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}

func part1(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))

	sum := 0
	for sc.Scan() {
		first, last := "", ""
		for _, char := range sc.Text() {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}

			if first == "" {
				first = string(char)
			}
			last = string(char)
		}

		num, err := strconv.Atoi(first + last)
		if err != nil {
			return 0, err
		}

		sum += num
	}
	if err := sc.Err(); err != nil {
		return 0, err
	}

	return sum, nil
}

func part2(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))
	sum := 0
	for sc.Scan() {
		line := sc.Text()
		var first, last string
		for i := 0; i < len(line); i++ {
			for j := i; j < len(line); j++ {
				num := getNumStr(line[i : j+1])
				if num != "" {
					if first == "" {
						first = num
					}
					last = num
				}
			}
		}

		num, err := strconv.Atoi(first + last)
		if err != nil {
			return 0, err
		}

		sum += num
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

	core.RunSolutions("AoC 2023 Day 01", input, config, part1, part2)
}

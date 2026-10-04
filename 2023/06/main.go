package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func strToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func parseInput1(line string) []int {
	result := []int{}
	nums := strings.Fields(line)[1:]
	for _, num := range nums {
		result = append(result, strToInt(num))
	}
	return result
}

func parseInput2(line string) []int {
	num := strings.Join(strings.Fields(line)[1:], "")
	return []int{strToInt(num)}
}

func part1(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))
	in := [][]int{}
	for sc.Scan() {
		in = append(in, parseInput1(sc.Text()))
	}

	result := 1
	for i := 0; i < len(in[0]); i++ {
		raceOpts := 0
		for j := 0; j <= in[0][i]; j++ {
			dist := j * (in[0][i] - j)
			if dist > in[1][i] {
				raceOpts++
			}
		}

		result *= raceOpts
	}
	return result, nil
}

func part2(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))
	in := [][]int{}
	for sc.Scan() {
		in = append(in, parseInput2(sc.Text()))
	}

	result := 1
	for i := 0; i < len(in[0]); i++ {
		raceOpts := 0
		for j := 0; j <= in[0][i]; j++ {
			dist := j * (in[0][i] - j)
			if dist > in[1][i] {
				raceOpts++
			}
		}

		result *= raceOpts
	}
	return result, nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2023 Day 06", input, config, part1, part2)
}

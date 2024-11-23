package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func part1(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))
	sum := 0
	for sc.Scan() {
		cardSplit := strings.Split(sc.Text(), ": ")
		wins := []int{}
		gets := []int{}
		winPart := true
		for _, numStr := range strings.Fields(cardSplit[1]) {
			if numStr == "|" {
				winPart = false
				continue
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			if winPart {
				wins = append(wins, num)
			} else {
				gets = append(gets, num)
			}
		}

		winCount := 0
		for _, n := range gets {
			for _, w := range wins {
				if n == w {
					winCount++
				}
			}
		}

		if winCount > 0 {
			sum += int(math.Pow(2, float64(winCount-1)))
		}
	}
	return sum, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	sum := 0
	copies := make([]int, len(lines))
	for i := range copies {
		copies[i] = 1
	}
	for i, line := range lines {
		cardSplit := strings.Split(line, ": ")
		wins := []int{}
		gets := []int{}
		winPart := true
		for _, numStr := range strings.Fields(cardSplit[1]) {
			if numStr == "|" {
				winPart = false
				continue
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			if winPart {
				wins = append(wins, num)
			} else {
				gets = append(gets, num)
			}
		}

		winCount := 0
		for _, n := range gets {
			for _, w := range wins {
				if n == w {
					winCount++
				}
			}
		}

		for c := copies[i]; c > 0; c-- {
			for j := 1; j <= winCount; j++ {
				copies[i+j]++
			}
		}
	}

	for _, c := range copies {
		sum += c
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

	core.RunSolutions("AoC 2023 Day 04", input, config, part1, part2)
}

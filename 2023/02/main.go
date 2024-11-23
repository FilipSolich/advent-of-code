package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

func part1(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))

	sum := 0
	for sc.Scan() {
		gameLine := strings.Split(sc.Text(), ": ")
		var gameID int
		fmt.Sscanf(gameLine[0], "Game %d", &gameID)

		gameFitsIn := true
		sets := strings.Split(gameLine[1], "; ")
	outer:
		for _, set := range sets {
			colors := strings.Split(set, ", ")
			for _, color := range colors {
				var num int
				var colorStr string
				fmt.Sscanf(color, "%d %s", &num, &colorStr)
				if (colorStr == "red" && num > MaxRed) || (colorStr == "green" && num > MaxGreen) || (colorStr == "blue" && num > MaxBlue) {
					gameFitsIn = false
					break outer
				}
			}
		}
		if gameFitsIn {
			sum += gameID
		}
	}

	return sum, nil
}

func part2(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))

	sum := 0
	for sc.Scan() {
		gameLine := strings.Split(sc.Text(), ": ")
		var gameID int
		fmt.Sscanf(gameLine[0], "Game %d", &gameID)

		sets := strings.Split(gameLine[1], "; ")
		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, set := range sets {
			colors := strings.Split(set, ", ")
			for _, color := range colors {
				var num int
				var colorStr string
				fmt.Sscanf(color, "%d %s", &num, &colorStr)
				switch colorStr {
				case "red":
					if num > maxRed {
						maxRed = num
					}
				case "green":
					if num > maxGreen {
						maxGreen = num
					}
				case "blue":
					if num > maxBlue {
						maxBlue = num
					}
				}
			}
		}
		gamePower := maxRed * maxGreen * maxBlue
		sum += gamePower
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

	core.RunSolutions("AoC 2023 Day 02", input, config, part1, part2)
}

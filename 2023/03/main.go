package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func isSpecial(char string) bool {
	return strings.Contains("*+&=@/%$#-", char)
}

func isGear(char string) bool {
	return strings.Contains("*", char)
}

func hasNeihborSymbol1(lines []string, lineN int, i int) bool {
	if i > 0 {
		if isSpecial(string(lines[lineN][i-1])) {
			return true
		}
	}
	if i > 0 && lineN > 0 {
		if isSpecial(string(lines[lineN-1][i-1])) {
			return true
		}
	}
	if lineN > 0 {
		if isSpecial(string(lines[lineN-1][i])) {
			return true
		}
	}
	if i < len(lines[lineN])-1 && lineN > 0 {
		if isSpecial(string(lines[lineN-1][i+1])) {
			return true
		}
	}
	if i < len(lines[lineN])-1 {
		if isSpecial(string(lines[lineN][i+1])) {
			return true
		}
	}
	if i < len(lines[lineN])-1 && lineN < len(lines)-1 {
		if isSpecial(string(lines[lineN+1][i+1])) {
			return true
		}
	}
	if lineN < len(lines)-1 {
		if isSpecial(string(lines[lineN+1][i])) {
			return true
		}
	}
	if i > 0 && lineN < len(lines)-1 {
		if isSpecial(string(lines[lineN+1][i-1])) {
			return true
		}
	}

	return false
}

func hasNeihborSymbol2(lines []string, lineN int, i int) (bool, int, int) {
	if i > 0 {
		if isGear(string(lines[lineN][i-1])) {
			return true, lineN, i - 1
		}
	}
	if i > 0 && lineN > 0 {
		if isGear(string(lines[lineN-1][i-1])) {
			return true, lineN - 1, i - 1
		}
	}
	if lineN > 0 {
		if isGear(string(lines[lineN-1][i])) {
			return true, lineN - 1, i
		}
	}
	if i < len(lines[lineN])-1 && lineN > 0 {
		if isGear(string(lines[lineN-1][i+1])) {
			return true, lineN - 1, i + 1
		}
	}
	if i < len(lines[lineN])-1 {
		if isGear(string(lines[lineN][i+1])) {
			return true, lineN, i + 1
		}
	}
	if i < len(lines[lineN])-1 && lineN < len(lines)-1 {
		if isGear(string(lines[lineN+1][i+1])) {
			return true, lineN + 1, i + 1
		}
	}
	if lineN < len(lines)-1 {
		if isGear(string(lines[lineN+1][i])) {
			return true, lineN + 1, i
		}
	}
	if i > 0 && lineN < len(lines)-1 {
		if isGear(string(lines[lineN+1][i-1])) {
			return true, lineN + 1, i - 1
		}
	}

	return false, 0, 0
}

func part1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	sum := 0
	for lineN, line := range lines {
		num := ""
		neighborSymbol := false
		for i, c := range line {
			char := string(c)
			if strings.Contains("0123456789", char) {
				num += char
				if !neighborSymbol {
					if hasNeihborSymbol1(lines, lineN, i) {
						neighborSymbol = true
					}
				}
			}
			if i == len(line)-1 || !strings.Contains("0123456789", char) {
				if neighborSymbol {
					intNum, err := strconv.Atoi(num)
					if err != nil {
						return 0, err
					}
					sum += intNum
				}
				num = ""
				neighborSymbol = false
			}
		}
	}

	return sum, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	sum := 0
	gearsMap := map[string][]int{}
	for lineN, line := range lines {
		num := ""
		key := ""
		neighborGear := false
		for i, c := range line {
			char := string(c)
			if strings.Contains("0123456789", char) {
				num += char
				if !neighborGear {
					neighbor, x, y := hasNeihborSymbol2(lines, lineN, i)
					if neighbor {
						neighborGear = true
						key = strconv.Itoa(x) + ":" + strconv.Itoa(y)
					}
				}
			}
			if i == len(line)-1 || !strings.Contains("0123456789", char) {
				if neighborGear {
					intNum, err := strconv.Atoi(num)
					if err != nil {
						log.Fatal(err)
					}

					if gearsMap[key] == nil {
						gearsMap[key] = []int{}
					}
					gearsMap[key] = append(gearsMap[key], intNum)
				}
				num = ""
				neighborGear = false
			}
		}
	}

	for _, v := range gearsMap {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
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

	core.RunSolutions("AoC 2023 Day 03", input, config, part1, part2)
}

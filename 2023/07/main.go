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
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

var cardValues1 = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var cardValues2 = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

var kindValues = map[string]int{
	"five":      7,
	"four":      6,
	"fullhouse": 5,
	"three":     4,
	"twopair":   3,
	"onepair":   2,
	"highcard":  1,
}

type Hand struct {
	Cards string
	Bid   int
}

func (h Hand) Kind1() int {
	var top int
	m := map[rune]int{}
	for _, card := range h.Cards {
		m[card]++
	}
	double := false
	triple := false
	for _, v := range m {
		switch v {
		case 5:
			top = kindValues["five"]
		case 4:
			top = max(top, kindValues["four"])
		case 3:
			if double {
				top = max(top, kindValues["fullhouse"])
			} else {
				top = max(top, kindValues["three"])
			}
			triple = true
		case 2:
			if double {
				top = max(top, kindValues["twopair"])
			} else if triple {
				top = max(top, kindValues["fullhouse"])
			} else {
				top = max(top, kindValues["onepair"])
			}
			double = true
		case 1:
			top = max(top, kindValues["highcard"])
		}
	}
	return top
}

func (h Hand) Kind2() int {
	var top int
	jCards := 0
	m := map[rune]int{}
	for _, card := range h.Cards {
		if card == 'J' {
			jCards++
		}
		m[card]++
	}
	if jCards > 0 {
		if len(m) == 2 || len(m) == 1 {
			return kindValues["five"]
		} else if len(m) == 3 {
			if jCards != 1 {
				return kindValues["four"]
			}
			for k, s := range m {
				if k != 'J' {
					if s == 3 || s == 1 {
						return kindValues["four"]
					} else {
						return kindValues["fullhouse"]
					}
				}
			}
		} else if len(m) == 4 {
			return kindValues["three"]
		} else if len(m) == 5 {
			return kindValues["onepair"]
		}
	}
	double := false
	triple := false
	for _, v := range m {
		switch v {
		case 5:
			top = kindValues["five"]
		case 4:
			top = max(top, kindValues["four"])
		case 3:
			if double {
				top = max(top, kindValues["fullhouse"])
			} else {
				top = max(top, kindValues["three"])
			}
			triple = true
		case 2:
			if double {
				top = max(top, kindValues["twopair"])
			} else if triple {
				top = max(top, kindValues["fullhouse"])
			} else {
				top = max(top, kindValues["onepair"])
			}
			double = true
		case 1:
			top = max(top, kindValues["highcard"])
		}
	}
	return top
}

func parseHand(line string) Hand {
	parts := strings.Fields(line)
	bid, _ := strconv.Atoi(parts[1])
	return Hand{Cards: parts[0], Bid: bid}
}

func sortFunc1(a, b Hand) int {
	if a.Kind1() != b.Kind1() {
		return a.Kind1() - b.Kind1()
	}
	for i := 0; i < len(a.Cards); i++ {
		x, y := cardValues1[string(a.Cards[i])], cardValues1[string(b.Cards[i])]
		if x != y {
			return x - y
		}
	}
	return 0
}

func sortFunc2(a, b Hand) int {
	if a.Kind2() != b.Kind2() {
		return a.Kind2() - b.Kind2()
	}
	for i := 0; i < len(a.Cards); i++ {
		x, y := cardValues2[string(a.Cards[i])], cardValues2[string(b.Cards[i])]
		if x != y {
			return x - y
		}
	}
	return 0
}

func part1(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))
	hands := []Hand{}
	for sc.Scan() {
		hands = append(hands, parseHand(sc.Text()))
	}

	slices.SortFunc(hands, sortFunc1)

	result := 0
	for i, hand := range hands {
		result += (i + 1) * hand.Bid
	}
	return result, nil
}

func part2(input string) (int, error) {
	sc := bufio.NewScanner(strings.NewReader(input))
	hands := []Hand{}
	for sc.Scan() {
		hands = append(hands, parseHand(sc.Text()))
	}

	slices.SortFunc(hands, sortFunc2)

	result := 0
	for i, hand := range hands {
		result += (i + 1) * hand.Bid
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

	core.RunSolutions("AoC 2023 Day 07", input, config, part1, part2)
}

package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

type machine struct {
	ax, ay float64
	bx, by float64
	px, py float64
}

func parse(input string) []machine {
	var machines []machine
	mstr := strings.Split(input, "\n\n")
	for _, in := range mstr {
		var m machine
		fmt.Sscanf(in, "Button A: X+%f, Y+%f\nButton B: X+%f, Y+%f\nPrize: X=%f, Y=%f", &m.ax, &m.ay, &m.bx, &m.by, &m.px, &m.py)
		machines = append(machines, m)
	}
	return machines
}

func det(m machine) float64 {
	return m.ax*m.by - m.ay*m.bx
}

func solveLinear(m machine) (int, int) {
	d := det(m)
	a := (m.by/d)*m.px + (-1*m.bx/d)*m.py
	b := (-1*m.ay/d)*m.px + (m.ax/d)*m.py
	return int(math.Round(a)), int(math.Round(b))
}

func validate(m machine, a, b int) bool {
	return math.Round(m.ax*float64(a)+m.bx*float64(b)) == m.px && math.Round(m.ay*float64(a)+m.by*float64(b)) == m.py
}

func part1(input string) (int, error) {
	machines := parse(input)
	res := 0
	for _, m := range machines {
		a, b := solveLinear(m)
		if validate(m, a, b) {
			res += int(a*3 + b)
		}
	}
	return res, nil
}

func part2(input string) (int, error) {
	machines := parse(input)
	res := 0
	for _, m := range machines {
		m.px += 10000000000000
		m.py += 10000000000000
		a, b := solveLinear(m)
		if validate(m, a, b) {
			res += int(a*3 + b)
		}
	}
	return res, nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2024 Day 13", input, config, part1, part2)
}

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

var (
	maxX = 101
	maxY = 103
)

type robot struct {
	x, y   int
	vx, vy int
}

func parse(input string) []robot {
	sc := bufio.NewScanner(strings.NewReader(input))
	var robots []robot
	for sc.Scan() {
		r := robot{}
		fmt.Sscanf(sc.Text(), "p=%d,%d v=%d,%d", &r.x, &r.y, &r.vx, &r.vy)
		robots = append(robots, r)
	}
	return robots
}

func move(r robot) robot {
	r.x, r.y = r.x+r.vx, r.y+r.vy
	if r.x > maxX-1 {
		r.x = r.x % maxX
	} else if r.x < 0 {
		r.x = maxX + r.x
	}
	if r.y > maxY-1 {
		r.y = r.y % maxY
	} else if r.y < 0 {
		r.y = maxY + r.y
	}
	return r
}

func calculateSecurity(robots []robot) int {
	quadrants := make([]int, 4)
	for _, r := range robots {
		if r.x < maxX/2 && r.y < maxY/2 {
			quadrants[0]++
		} else if r.x > maxX/2 && r.y < maxY/2 {
			quadrants[1]++
		} else if r.x < maxX/2 && r.y > maxY/2 {
			quadrants[2]++
		} else if r.x > maxX/2 && r.y > maxY/2 {
			quadrants[3]++
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func printRobots(robots []robot, s int) {
	fmt.Printf("After %d seconds.\n", s)
	grid := make([][]int, maxY)
	for i := range maxY {
		grid[i] = make([]int, maxX)
	}
	for _, r := range robots {
		grid[r.y][r.x]++
	}
	for y := range maxY {
		for x := range maxX {
			num := strconv.Itoa(grid[y][x])
			if num == "0" {
				num = "."
			}
			fmt.Printf("%s", num)
		}
		fmt.Println()
	}
}

func part1(input string) (int, error) {
	robots := parse(input)
	for i, r := range robots {
		for range 100 {
			r = move(r)
		}
		robots[i] = r
	}

	return calculateSecurity(robots), nil
}

func part2(input string) (int, error) {
	robots := parse(input)
	for s := range 10000 {
		printRobots(robots, s)
		for i, r := range robots {
			robots[i] = move(r)
		}
	}

	return calculateSecurity(robots), nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2024 Day 14", input, config, part1, part2)
}

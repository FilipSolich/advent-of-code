package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

var (
	xMax    int
	yMax    int
	visited map[string]struct{}
	fences  int
)

func key(x, y int) string {
	return fmt.Sprintf("X%dY%d", x, y)
}

func process(lines []string, x, y int, plant byte) (int, int) {
	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	fence := 4
	neigbors := 0
	neigborsFence := 0
	visited[key(x, y)] = struct{}{}
	for _, dir := range dirs {
		xx, yy := x+dir[0], y+dir[1]
		if xx < 0 || xx > xMax-1 || yy < 0 || yy > yMax-1 {
			continue
		}
		if lines[yy][xx] != plant {
			continue
		}
		if _, ok := visited[key(xx, yy)]; ok {
			fence--
			continue
		}
		n, f := process(lines, xx, yy, plant)
		fence--
		neigbors += n
		neigborsFence += f
	}
	return 1 + neigbors, fence + neigborsFence
}

//func rememberFence(x, y, xx, yy int) {
//	if y < yy {
//		fences["^"+strconv.Itoa(y)] = struct{}{}
//	} else if y > yy {
//		fences["v"+strconv.Itoa(y)] = struct{}{}
//	} else if x < xx {
//		fences["<"+strconv.Itoa(x)] = struct{}{}
//	} else {
//		fences[">"+strconv.Itoa(x)] = struct{}{}
//	}
//}

func rememberFence(toDir, dir int) {
	if toDir != dir+1%4 {
		fences++
	}
}

func process2(lines []string, x, y int, plant byte, toDir int) int {
	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	neigbors := 0
	visited[key(x, y)] = struct{}{}
	for i, dir := range dirs {
		xx, yy := x+dir[0], y+dir[1]
		if xx < 0 || xx > xMax-1 || yy < 0 || yy > yMax-1 {
			//rememberFence(x, y, xx, yy)
			rememberFence(toDir, i)
			continue
		}
		if lines[yy][xx] != plant {
			//rememberFence(x, y, xx, yy)
			rememberFence(toDir, i)
			continue
		}
		if _, ok := visited[key(xx, yy)]; ok {
			continue
		}
		n := process2(lines, xx, yy, plant, i)
		neigbors += n
	}
	return 1 + neigbors
}

func part1(input string) (int, error) {
	lines := strings.Fields(input)
	xMax, yMax = len(lines[0]), len(lines)

	visited = map[string]struct{}{}
	sum := 0
	for y := range yMax {
		for x := range xMax {
			if _, ok := visited[key(x, y)]; ok {
				continue
			}
			plant := lines[y][x]
			area, perimetr := process(lines, x, y, plant)
			sum += area * perimetr
		}
	}

	return sum, nil
}

func part2(input string) (int, error) {
	lines := strings.Fields(input)
	xMax, yMax = len(lines[0]), len(lines)

	visited = map[string]struct{}{}
	sum := 0
	for y := range yMax {
		for x := range xMax {
			if _, ok := visited[key(x, y)]; ok {
				continue
			}
			plant := lines[y][x]
			area := process2(lines, x, y, plant, 0)
			sum += area * fences
			fences = 0
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

	core.RunSolutions("AoC 2024 Day 12", input, config, part1, part2)
}

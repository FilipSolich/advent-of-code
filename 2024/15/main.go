package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func parse(input string) ([]string, string) {
	parts := strings.Split(input, "\n\n")
	warehouse := strings.Fields(parts[0])
	moves := strings.Join(strings.Fields(parts[1]), "")
	return warehouse, moves
}

func findRobot(warehouse []string) (int, int) {
	for y, line := range warehouse {
		for x, symbol := range line {
			if symbol == '@' {
				return x, y
			}
		}
	}
	return 0, 0
}

var moveMap = map[rune][]int{
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
}

func swap(warehouse *[]string, x, y, xx, yy int) {
	symbol := (*warehouse)[yy][xx]
	line := []rune((*warehouse)[yy])
	line[xx] = rune((*warehouse)[y][x])
	(*warehouse)[yy] = string(line)
	line = []rune((*warehouse)[y])
	line[x] = rune(symbol)
	(*warehouse)[y] = string(line)
}

func move(warehouse *[]string, x, y int, dir rune) (bool, int, int) {
	diff := moveMap[dir]
	xx, yy := x+diff[0], y+diff[1]
	switch (*warehouse)[yy][xx] {
	case '#':
		return false, x, y
	case '.':
		swap(warehouse, x, y, xx, yy)
		return true, xx, yy
	case 'O':
		ok, _, _ := move(warehouse, xx, yy, dir)
		if ok {
			swap(warehouse, x, y, xx, yy)
			return true, xx, yy
		}
	}

	return false, x, y
}

func move2(warehouse *[]string, x, y int, dir rune) (bool, int, int) {
	diff := moveMap[dir]
	xx, yy := x+diff[0], y+diff[1]
	switch (*warehouse)[yy][xx] {
	case '#':
		return false, x, y
	case '.':
		swap(warehouse, x, y, xx, yy)
		return true, xx, yy
	case '[', ']':
		var ok bool
		if dir == '<' || dir == '>' {
			ok, _, _ = move2(warehouse, xx, yy, dir)
		} else {
			if (*warehouse)[yy][xx] == '[' {
				ok = wideMove(warehouse, xx, xx+1, yy, dir)
			} else {
				ok = wideMove(warehouse, xx-1, xx, yy, dir)
			}
		}
		if ok {
			swap(warehouse, x, y, xx, yy)
			return true, xx, yy
		}
	}

	return false, x, y
}

func canMove(warehouse *[]string, x, y int, dir rune) bool {
	diff := moveMap[dir]
	xx, yy := x+diff[0], y+diff[1]
	nextSymbol := (*warehouse)[yy][xx]
	switch nextSymbol {
	case '#':
		return false
	case '.':
		return true
	case '[', ']':
		if dir == '<' || dir == '>' {
			return canMove(warehouse, xx, yy, dir)
		}
		if nextSymbol == '[' {
			return canMove(warehouse, xx, yy, dir) && canMove(warehouse, xx+1, yy, dir)
		} else {
			return canMove(warehouse, xx, yy, dir) && canMove(warehouse, xx-1, yy, dir)
		}
	}

	return false
}

func move3(warehouse *[]string, x, y int, dir rune) (bool, int, int) {
	diff := moveMap[dir]
	xx, yy := x+diff[0], y+diff[1]
	nextSymbol := (*warehouse)[yy][xx]
	switch nextSymbol {
	case '.':
		swap(warehouse, x, y, xx, yy)
		return true, xx, yy
	case '[', ']':
		if dir == '<' || dir == '>' {
			if ok, _, _ := move3(warehouse, xx, yy, dir); ok {
				swap(warehouse, x, y, xx, yy)
				return true, xx, yy
			}
		}
		if nextSymbol == '[' {
			move3(warehouse, xx, yy, dir)
			swap(warehouse, x, y, xx, yy)
			if (*warehouse)[yy+diff[1]][xx] == ']' {
				move3(warehouse, xx+1, yy, dir)
			}
			if (*warehouse)[y][x] != '@' {
				swap(warehouse, x+1, y, xx+1, yy)
			}
			return true, xx, yy
		} else {
			move3(warehouse, xx, yy, dir)
			swap(warehouse, x, y, xx, yy)
			if (*warehouse)[yy+diff[1]][xx] == '[' {
				move3(warehouse, xx-1, yy, dir)
			}
			if (*warehouse)[y][x] != '@' {
				swap(warehouse, x-1, y, xx-1, yy)
			}
			return true, xx, yy
		}
	}
	return false, x, y
}

func wideMove(warehouse *[]string, x1, x2, y int, dir rune) bool {
	diff := moveMap[dir]
	xx1, xx2, yy := x1+diff[0], x2+diff[0], y+diff[1]
	sym1, sym2 := (*warehouse)[yy][xx1], (*warehouse)[yy][xx2]
	if sym1 == '#' || sym2 == '#' {
		return false
	} else if sym1 == '.' && sym2 == '.' {
		swap(warehouse, x1, y, xx1, yy)
		swap(warehouse, x2, y, xx2, yy)
		return true
	} else if sym1 == '[' && sym2 == ']' {
		if wideMove(warehouse, xx1, xx2, yy, dir) {
			swap(warehouse, x1, y, xx1, yy)
			swap(warehouse, x2, y, xx2, yy)
			return true
		}
	} else {
		ok := true
		if sym1 == ']' {
			ok = wideMove(warehouse, xx1-1, xx1, yy, dir)
		}
		if sym2 == '[' && ok {
			ok = wideMove(warehouse, xx2, xx2+1, yy, dir)
			if !ok {
				// TODO: Undo first move
				inc := 1
				if dir == '^' {
					inc = -1
				}
				swap(warehouse, xx1-1, yy, xx1-1, yy+inc)
				swap(warehouse, xx1, yy, xx1, yy+inc)
			}
		}
		if ok {
			swap(warehouse, x1, y, xx1, yy)
			swap(warehouse, x2, y, xx2, yy)
			return true
		}
	}

	return false
}

func countBoxes(warehouse []string) int {
	var sum int
	for y, line := range warehouse {
		for x, symbol := range line {
			if symbol == 'O' {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func countBoxes2(warehouse []string) int {
	var sum int
	for y, line := range warehouse {
		for x, symbol := range line {
			if symbol == '[' {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func widenWarehouse(warehouse []string) []string {
	var new []string
	for _, line := range warehouse {
		var newLine []string
		for _, symbol := range line {
			switch symbol {
			case '#':
				newLine = append(newLine, "##")
			case 'O':
				newLine = append(newLine, "[]")
			case '.':
				newLine = append(newLine, "..")
			case '@':
				newLine = append(newLine, "@.")
			}
		}
		new = append(new, strings.Join(newLine, ""))
	}
	return new
}

func part1(input string) (int, error) {
	warehouse, moves := parse(input)
	x, y := findRobot(warehouse)
	for _, m := range moves {
		_, x, y = move(&warehouse, x, y, m)
	}
	return countBoxes(warehouse), nil
}

//func part2(input string) (int, error) {
//	warehouse, moves := parse(input)
//	warehouse = widenWarehouse(warehouse)
//	x, y := findRobot(warehouse)
//	for i, m := range moves {
//		_, x, y = move2(&warehouse, x, y, m)
//		fmt.Println(i)
//	}
//	return countBoxes2(warehouse), nil
//}

func part2(input string) (int, error) {
	warehouse, moves := parse(input)
	warehouse = widenWarehouse(warehouse)
	x, y := findRobot(warehouse)
	for i, m := range moves {
		if canMove(&warehouse, x, y, m) {
			_, x, y = move3(&warehouse, x, y, m)
		}
		fmt.Println(i)
	}
	return countBoxes2(warehouse), nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2024 Day 15", input, config, part1, part2)
}

package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func genList(elem int, count int) []int {
	out := make([]int, count)
	for i := range out {
		out[i] = elem
	}
	return out
}

func split(input string) ([]int, []int) {
	var files, spaces []int
	for i, ch := range input {
		if i%2 == 0 {
			files = append(files, utils.MustAtoi(string(ch)))
		} else {
			spaces = append(spaces, utils.MustAtoi(string(ch)))
		}
	}
	return files, spaces
}

func merge(files, spaces []int) []int {
	var fs []int
	fIdx := 0
	fLast := len(files) - 1
	sIdx := 0
	file := true
	for {
		if file {
			fs = append(fs, genList(fIdx, files[fIdx])...)
			files[fIdx] = 0
			fIdx++
			if fIdx == fLast {
				fs = append(fs, genList(fIdx, files[fIdx])...)
				break
			}
			file = false
		} else {
			fill := min(spaces[sIdx], files[fLast])
			fs = append(fs, genList(fLast, fill)...)
			files[fLast] -= fill
			if files[fLast] == 0 {
				fLast--
				if fLast < fIdx {
					break
				}
			}
			spaces[sIdx] -= fill
			if spaces[sIdx] == 0 {
				sIdx++
				file = true
			}
		}
	}
	return fs
}

func swap(fs []int, fromIdx int, toIdx int, count int) []int {
	for i := range count {
		fs[toIdx+i] = fs[fromIdx+i]
		fs[fromIdx+i] = 0
	}

	return fs
}

func merge2(files, spaces []int) []int {
	var fs []int
	var filesPos [][]int
	var spacePos [][]int

	for i := range len(files) {
		filesPos = append(filesPos, []int{len(fs), files[i]})
		fs = append(fs, genList(i, files[i])...)
		if i != len(files)-1 && spaces[i] != 0 {
			spacePos = append(spacePos, []int{len(fs), spaces[i]})
			fs = append(fs, genList(0, spaces[i])...)
		}
	}

	for fileIdx := len(files) - 1; fileIdx >= 0; fileIdx-- {
		if filesPos[fileIdx][0] <= spacePos[0][0] {
			break
		}
		for i := 0; ; i++ {
			if i > len(spacePos)-1 || spacePos[i][0] >= filesPos[fileIdx][0] {
				break
			}
			if files[fileIdx] <= spacePos[i][1] {
				fs = swap(fs, filesPos[fileIdx][0], spacePos[i][0], files[fileIdx])
				spacePos[i][1] -= files[fileIdx]
				spacePos[i][0] += files[fileIdx]
				if spacePos[i][1] == 0 {
					spacePos = slices.Delete(spacePos, i, i+1)
				}
				break
			}
		}
	}
	return fs
}

func checksum(fs []int) int {
	sum := 0
	for i, num := range fs {
		sum += i * num
	}
	return sum
}

func part1(input string) (int, error) {
	input = strings.TrimSpace(input)
	files, spaces := split(input)
	fs := merge(files, spaces)
	checksum := checksum(fs)
	return checksum, nil
}

func part2(input string) (int, error) {
	input = strings.TrimSpace(input)
	files, spaces := split(input)
	fs := merge2(files, spaces)
	checksum := checksum(fs)
	return checksum, nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2024 Day 09", input, config, part1, part2)
}

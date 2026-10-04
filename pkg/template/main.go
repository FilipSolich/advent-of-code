// This is template file. When used replace YYYY with year, DD with day and implement part1 and part2 functions.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func part1(input string) (int, error) {
	return 0, errors.New("not implemented")
}

func part2(input string) (int, error) {
	return 0, errors.New("not implemented")
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC YYYY Day DD", input, config, part1, part2)
}

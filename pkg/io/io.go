package io

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

func ReadInputFile(filename string) (string, error) {
	var err error
	reader := os.Stdin
	if filename != "" {
		reader, err = os.Open(filename)
		if err != nil {
			return "", fmt.Errorf("error opening input file: %w", err)
		}
		defer reader.Close()
	}

	content, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("error reading input file: %w", err)
	}
	input := string(content)

	return input, nil
}

const lineFormat = "%-13s%v\n"

func printGuards(title string) (string, string) {
	guardLineLength := 50
	titlePart := fmt.Sprintf("=== %s ", title)
	titleLine := titlePart + strings.Repeat("=", guardLineLength-len(titlePart))
	endLine := strings.Repeat("-", guardLineLength)
	return titleLine, endLine
}

func printStats(duration time.Duration, config utils.Config) {
	fmt.Printf(lineFormat, "Input file:", config.InputFileName)
	fmt.Printf(lineFormat, "Total time:", duration)
	fmt.Printf(lineFormat, "Time limit:", config.Timeout)
}

func PrintAnswer(title string, result string, duration time.Duration, config utils.Config) {
	titleLine, endLine := printGuards(title)
	fmt.Println(titleLine)
	printStats(duration, config)
	fmt.Printf(lineFormat, "Result:", fmt.Sprintf("\033[36m%s\033[0m", result))
	fmt.Println(endLine)
}

func PrintError(title string, err error, duration time.Duration, config utils.Config) {
	titleLine, endLine := printGuards(title)
	fmt.Println(titleLine)
	printStats(duration, config)
	fmt.Printf(lineFormat, "Error:", fmt.Sprintf("\033[1m\033[31m%v\033[0m", err))
	fmt.Println(endLine)
}

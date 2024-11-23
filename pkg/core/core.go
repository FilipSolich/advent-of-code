package core

import (
	"context"
	"fmt"
	"time"

	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

type SolutionFunc func(input string) (int, error)

type Solution struct {
	Output int
	Err    error
}

func RunSolution(title string, input string, fnc SolutionFunc, config utils.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	start := time.Now()
	solutionCh := make(chan Solution)
	go func() {
		result, err := fnc(input)
		solutionCh <- Solution{Output: result, Err: err}
	}()

	select {
	case result := <-solutionCh:
		if result.Err != nil {
			io.PrintError(title, result.Err, time.Since(start), config)
			return
		}
		io.PrintAnswer(title, result.Output, time.Since(start), config)
	case <-ctx.Done():
		io.PrintError(title, ctx.Err(), time.Since(start), config)
	}
}

func RunSolutions(title string, input string, config utils.Config, solutions ...SolutionFunc) {
	if config.PartIdx >= 0 {
		RunSolution(fmt.Sprintf("%s - Part %d", title, config.PartIdx+1), input, solutions[config.PartIdx], config)
		return
	}

	for i, fnc := range solutions {
		RunSolution(fmt.Sprintf("%s - Part %d", title, i+1), input, fnc, config)
	}
}

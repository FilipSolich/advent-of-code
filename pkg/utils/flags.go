package utils

import (
	"flag"
	"time"
)

type Config struct {
	InputFileName string
	PartIdx       int
	Timeout       time.Duration
}

func ParseFlags() Config {
	inputFileName := flag.String("input", "", "Input file. If empty read from stdin.")
	part := flag.Int("part", 0, "Solution part. If empty run both parts.")
	timeout := flag.Duration("timeout", 1*time.Minute, "Solution timeout.")
	flag.Parse()

	return Config{
		InputFileName: *inputFileName,
		PartIdx:       *part - 1,
		Timeout:       *timeout,
	}
}

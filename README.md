# Advent of Code

Solutions for Advent of Code written in Go.

## State of solutions

| Year | 01 | 02 | 03 | 04 | 05 | 06 | 07 | 08 | 09 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 |
|------|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|
| 2023 | ** | ** | ** | ** | ** | ** | ** | *  |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |
| 2024 | ** | ** |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |

## Usage

### Generate template for day and download input

```sh
task generate YEAR=2024 DAY=1
task aoc:input YEAR=2024 DAY=1
```

### Run solution

```sh
task run YEAR=2024 DAY=1
```

To pass flags to program append `-- <flags>`. Possible options:

```sh
-input string
      Input file. If empty read from stdin.
-part int
      Solution part. If empty run both parts.
-timeout duration
      Solution timeout. (default 1m0s)
```

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/FilipSolich/advent-of-code/pkg/core"
	"github.com/FilipSolich/advent-of-code/pkg/io"
	"github.com/FilipSolich/advent-of-code/pkg/utils"
)

var cache map[int][]int

func parse(input string) []int {
	strNums := strings.Fields(input)

	var nums []int
	for _, num := range strNums {
		nums = append(nums, utils.MustAtoi(num))
	}
	return nums
}

func strToInt(s string) int {
	if s == strings.Repeat("0", len(s)) {
		return 0
	}
	return utils.MustAtoi(s)
}

func blink(stones []int) []int {
	newStones := make([]int, 0, len(stones))
	for _, stone := range stones {
		str := fmt.Sprint(stone)
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if len(str)%2 == 0 {
			first, second := str[:len(str)/2], str[len(str)/2:]
			newStones = append(newStones, strToInt(first), strToInt(second))
		} else {
			newStones = append(newStones, stone*2024)
		}
	}
	return newStones
}

func insert(stones []int, idx int, in []int) []int {
	stones[idx] = in[0]
	if len(in) == 1 {
		return stones
	}
	if idx+1 == len(stones) {
		return append(stones, in[1])
	}
	return slices.Insert(stones, idx+1, in[1])
}

func blink2(stones []int) []int {
	newStones := make([]int, 0, len(stones))
	for _, stone := range stones {
		cached, ok := cache[stone]
		if ok {
			newStones = append(newStones, cached...)
			continue
		}

		var new []int
		str := fmt.Sprint(stone)
		if stone == 0 {
			new = []int{1}
		} else if len(str)%2 == 0 {
			first, second := str[:len(str)/2], str[len(str)/2:]
			new = []int{strToInt(first), strToInt(second)}
		} else {
			new = []int{stone * 2024}
		}
		cache[stone] = new
		newStones = append(newStones, new...)
	}
	return newStones
}

func insert2(mem []int, mu *sync.Mutex, in []int, idx int) {
	mu.Lock()
	defer mu.Unlock()
	mem[idx] = in[0]
	if len(in) == 2 {
		if idx+1 > len(mem)-1 {
			mem = append(mem, in[1])
		} else {
			mem[idx+1] = in[1]
		}
	}
}

func process(mem []int, mu *sync.Mutex, stones []int, idx int) {
	offset := 0
	for i, stone := range stones {
		cached, ok := cache[stone]
		if ok {
			insert2(mem, mu, cached, idx+i+offset)
			if len(cached) == 2 {
				offset++
			}
			continue
		}

		var new []int
		str := fmt.Sprint(stone)
		if stone == 0 {
			new = []int{1}
		} else if len(str)%2 == 0 {
			first, second := str[:len(str)/2], str[len(str)/2:]
			new = []int{strToInt(first), strToInt(second)}
		} else {
			new = []int{stone * 2024}
		}
		cache[stone] = new
		insert2(mem, mu, new, idx+i+offset)
		if len(new) == 2 {
			offset++
		}

	}
}

func blink3(stones []int) []int {
	split := 0
	for _, s := range stones {
		if len(fmt.Sprint(s))%2 == 0 {
			split++
		}
	}
	stones = slices.Grow(stones, split)
	process(stones, &sync.Mutex{}, slices.Clone(stones), 0)
	return stones
}

func blink4(read string, write string) error {
	rf, err := os.OpenFile(read, os.O_CREATE|os.O_RDWR, 0o664)
	if err != nil {
		return err
	}
	defer rf.Close()
	wf, err := os.OpenFile(write, os.O_CREATE|os.O_RDWR, 0o664)
	if err != nil {
		return err
	}
	defer wf.Close()
	rs := bufio.NewScanner(rf)
	for rs.Scan() {
		stones := parse(rs.Text())
		stones = blink2(stones)

		for {
			count := min(len(stones), 4096)
			write := stones[:count]
			var str []string
			for _, stone := range write {
				str = append(str, strconv.Itoa(stone))
			}
			wf.WriteString(strings.Join(str, " "))
			wf.WriteString("\n")
			stones = stones[count:]
			if len(stones) == 0 {
				break
			}
		}
	}

	if rs.Err() != nil {
		return rs.Err()
	}

	return nil
}

func part1(input string) (int, error) {
	stones := parse(input)
	for range 25 {
		stones = blink(stones)
	}
	return len(stones), nil
}

func part2(input string) (int, error) {
	os.Truncate("./1.txt", 0)
	f, err := os.OpenFile("./1.txt", os.O_CREATE|os.O_RDWR, 0o664)
	if err != nil {
		return 0, err
	}
	f.WriteString(input)
	f.Close()

	cache = map[int][]int{}
	read, write := "./1.txt", "./2.txt"
	for i := range 75 {
		fmt.Println(i)
		if err := blink4(read, write); err != nil {
			return 0, err
		}
		read, write = write, read
		if err := os.Truncate(write, 0); err != nil {
			return 0, err
		}
	}

	content, err := os.ReadFile(read)
	if err != nil {
		return 0, err
	}
	stones := parse(string(content))
	return len(stones), nil
}

func main() {
	config := utils.ParseFlags()

	input, err := io.ReadInputFile(config.InputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v\n", err)
		os.Exit(1)
	}

	core.RunSolutions("AoC 2024 Day 11", input, config, part1, part2)
}

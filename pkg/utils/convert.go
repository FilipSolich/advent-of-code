package utils

import "strconv"

func MustAtoi(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return n
}

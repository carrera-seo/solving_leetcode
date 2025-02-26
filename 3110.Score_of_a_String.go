package main

import (
	"math"
)

func scoreOfString(s string) int {

	if len(s) < 2 {
		return 0
	}

	total := 0

	runes := []rune(s)

	for i := 0; i < len(runes)-1; i++ {
		sum := float64(runes[i] - runes[i+1])
		total += int(math.Abs(sum))
	}
	return total
}

package main

import (
	"strconv"
)

func isPalindrome(x int) bool {

	if x > -1 && x < 10 {
		return true
	}

	input_value := strconv.Itoa(x)
	len := len(input_value)
	half_index := len / 2

	for i := 0; i <= half_index; i++ {
		if input_value[i] != input_value[len-1-i] {
			return false
		}
	}
	return true
}

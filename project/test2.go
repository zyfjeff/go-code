package main

import (
	"fmt"
)

func Min(a ...int) int {
	min := int(^uint(0) >> 1) //largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}

	return min
}

func main() {
	minvalue := Min(10, 203, 11)
	fmt.Println(minvalue)
}

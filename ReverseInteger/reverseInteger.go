package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	neg := false
	var answer int

	if x < 0 {
		x = x * -1
		neg = true
	}

	for x != 0 {
		n := x % 10
		answer = n + answer*10
		if answer > math.MaxInt32 || answer < math.MinInt32 {
			return 0
		}
		x = x / 10
	}

	if neg == true {
		answer *= -1
		neg = false
	}

	return answer
}

func main() {
	x := 123
	fmt.Printf("%d -> %d\n", x, reverse(x))
	x = -123
	fmt.Printf("%d -> %d\n", x, reverse(x))
	x = 1534236469
	fmt.Printf("%d -> %d\n", x, reverse(x))
}

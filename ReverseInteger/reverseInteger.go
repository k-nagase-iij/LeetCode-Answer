package ReverseInteger

import (
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

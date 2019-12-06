package StringtoInteger

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}
	var sign int
	unsignedStr := str
	switch str[0] {
	case '-':
		sign = -1
		unsignedStr = str[1:]
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign = 1
	default:
		return 0
	}

	for i, c := range unsignedStr {
		if c < '0' || '9' < c {
			unsignedStr = unsignedStr[:i]
			break
		}
	}

	return convertInt(unsignedStr, sign)
}

func convertInt(str string, sign int) int {
	ans := 0
	for _, c := range str {
		ans = ans*10 + int(c-'0')
		if sign == 1 && math.MaxInt32 < ans {
			return math.MaxInt32
		} else if sign == -1 && ans*sign < math.MinInt32 {
			return math.MinInt32
		}
	}
	return ans * sign
}

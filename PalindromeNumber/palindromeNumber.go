package PalindromeNumber

import "strconv"

func isPalindrome(x int) bool {
	runes := []rune(strconv.Itoa(x))
	for i := 0;i < len(runes) / 2;i++ {
		if runes[i] != runes[len(runes) - 1 - i] {
			return false
		}
	}
	return true
}
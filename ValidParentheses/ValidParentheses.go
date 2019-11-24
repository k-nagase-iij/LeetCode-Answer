package ValidParentheses

func isValid(s string) bool {
	index, bracket := 0, make([]rune, len(s))
	bracketComb := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, b := range s {
		if b == '(' || b == '[' || b == '{' {
			bracket[index] = b
			index++
		} else if b == ')' || b == ']' || b == '}' {
			if index == 0 || bracketComb[b] != bracket[index-1] {
				return false
			}
			index--
		}
	}
	return index == 0
}

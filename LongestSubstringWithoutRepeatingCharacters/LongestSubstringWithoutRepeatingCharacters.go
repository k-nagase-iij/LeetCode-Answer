package LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	m, max, start := make(map[rune]int), 0, 0
	for i, c := range s {
		_, ok := m[c]
		if ok == true && start <= m[c] {
			if i-start > max {
				max = i - start
			}
			start = m[c] + 1
		}
		m[c] = i
	}
	if len(s)-start > max {
		max = len(s) - start
	}
	return max
}

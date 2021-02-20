package LongestCommonPrefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		input []string
		want string
	}{
		{"2 lenght", []string{"flower", "flow", "flight"}, "fl"},
		{"no common str", []string{"dog", "racecar", "car"}, ""},
		{"same strings", []string{"dogs", "dogs", "dogs"}, "dogs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.input); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

package PalindromeNumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		input int
		want bool
	}{
		{"121", 121, true},
		{"1", 1, true},
		{"12", 12, false},
		{"123", 123, false},
		{"-12345", -12345, false},
		{"-11", -11, false},
		{"123321", 123321, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.input); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

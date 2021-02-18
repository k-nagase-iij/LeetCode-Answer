package RomanToInteger

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		input string
		want int
	}{
		{"1", "I", 1},
		{"2", "II", 2},
		{"3", "III", 3},
		{"4", "IV", 4},
		{"5", "V", 5},
		{"6", "VI", 6},
		{"7", "VII", 7},
		{"8", "VIII", 8},
		{"9", "IX", 9},
		{"10", "X", 10},
		{"39", "XXXIX", 39},
		{"40", "XL", 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.input); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
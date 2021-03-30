package RemoveDuplicatesSortedArray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"no.1", []int{1, 1, 2}, 2},
		{"no.2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{"no.3", []int{0, 1, 2, 3, 3}, 4},
		{"no.4", []int{3, 4, 5, 6}, 4},
		{"no.5", []int{3}, 1},
		{"no.6", []int{3, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.input); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

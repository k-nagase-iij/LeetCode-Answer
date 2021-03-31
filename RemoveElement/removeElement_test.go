package RemoveElement

import "testing"

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"no.1", args{[]int{3, 2, 2, 3}, 3}, 2},
		{"no.2", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
		{"no.3", args{[]int{2, 2}, 2}, 0},
		{"no.4", args{[]int{0}, 2}, 1},
		{"no.5", args{[]int{4, 5}, 1}, 2},
		{"no.6", args{[]int{}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

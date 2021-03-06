package MergeTwoSortedLists

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1,1,2,3,4,4", args{l1: []int{1, 2, 4}, l2: []int{1, 3, 4}}, []int{1, 1, 2, 3, 4, 4}},
		{"0", args{l1: []int{}, l2: []int{0}}, []int{0}},
		{"nil pointer", args{l1: []int{}, l2: []int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(CreateLinkedList(tt.args.l1), CreateLinkedList(tt.args.l2))

			assertLinkedList(t, got, CreateLinkedList(tt.want))
		})
	}
}

func assertLinkedList(t *testing.T, got, want *ListNode) {
	for got != nil || want != nil {
		if got.Val != want.Val {
			t.Errorf("got: %d, want: %d", got.Val, want.Val)
		}

		got = got.Next
		want = want.Next
	}
}

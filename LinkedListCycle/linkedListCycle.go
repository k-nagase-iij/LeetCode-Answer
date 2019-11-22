package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	a, b := head, head.Next
	for b != nil && b.Next != nil && a != b {
		a, b = b, b.Next.Next
	}
	return a == b
}

func main() {
	head1 := ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}
	head2 := ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	fmt.Println(hasCycle(&head1))
	fmt.Println(hasCycle(&head2))
}

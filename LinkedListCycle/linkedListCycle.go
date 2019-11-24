package LinkedListCycle

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

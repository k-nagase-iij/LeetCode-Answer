package LinkedListCycle2

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

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	a, b := head.Next, head.Next.Next
	for b != nil && b.Next != nil && a != b {
		a, b = a.Next, b.Next.Next
	}
	if a != b {
		return nil
	}

	a = head
	for a != b {
		a, b = a.Next, b.Next
	}
	return a
}

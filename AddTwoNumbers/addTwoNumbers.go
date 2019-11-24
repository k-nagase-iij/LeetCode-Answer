package AddTwoNumbers

/*
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	answer := list
	curry := 0

	for l1 != nil || l2 != nil || curry != 0 {
		sum := curry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curry = sum / 10
		answer.Next = &ListNode{Val: sum % 10}
		answer = answer.Next
	}
	return list.Next
}

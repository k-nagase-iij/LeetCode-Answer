package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	ans := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}
	return ans
}

func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	ans := deleteDuplicates(&head)
	fmt.Println(ans)
}

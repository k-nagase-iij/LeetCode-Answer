package MergeTwoSortedLists

import (
	"sort"
)

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var mergedInt []int

	for l1 != nil {
		mergedInt = append(mergedInt, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		mergedInt = append(mergedInt, l2.Val)
		l2 = l2.Next
	}

	sort.Ints(mergedInt)

	ret := CreateLinkedList(mergedInt)

	return ret
}

func CreateLinkedList(intList []int) *ListNode {
	ret := &ListNode{}
	answer := ret

	for _, v := range intList {
		answer.Next = &ListNode{}
		answer = answer.Next

		answer.Val = v
	}

	return ret.Next
}

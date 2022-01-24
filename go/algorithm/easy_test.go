package algorithm

import "testing"

func TestAlgorithm(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, &ListNode{3, nil}}
	traverse(head)
	reverse := reverseList(head)
	traverse(reverse)
	reverse2 := reverseList2(head)
	traverse(reverse2)
}

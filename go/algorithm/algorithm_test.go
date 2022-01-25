package algorithm

import (
	"fmt"
	"testing"
)

func TestAlgorithm(t *testing.T) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	traverse(list1)
	reverse := reverseList(list1)
	traverse(reverse)
	reverse2 := reverseList2(list1)
	traverse(reverse2)
}

func traverse(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Printf("%d ", head.Val)
	}
	fmt.Println()
}

func foo() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	curr1 := list1
	curr2 := list2

	fmt.Println(curr1)
	fmt.Println(curr2)
}

func TestList(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	var head *ListNode
	cur := list
	for cur != nil {
		node := cur
		cur = cur.Next
		node.Next = head
		head = node
	}
	traverse(head)
}

func TestMergeTwoLists1(t *testing.T) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	list3 := mergeTwoLists1(list1, list2)
	traverse(list3)
}

func TestMergeTwoLists2(t *testing.T) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	list3 := mergeTwoLists2(list1, list2)
	traverse(list3)
}

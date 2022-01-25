package algorithm

import (
	"fmt"
	"testing"
)

func populateList(nums []int) *ListNode {
	dummy := &ListNode{0, nil}
	curr := dummy
	for _, num := range nums {
		node := &ListNode{num, nil}
		curr.Next = node
		curr = curr.Next
	}
	return dummy.Next
}

func TestReverseList1(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	reverse := reverseList1(list)
	traverse(reverse)

	node := *list
	list.Val = 10
	node.Val = 20
	fmt.Printf("node:%p, list:%p %p\n", &node, list, &list)
	traverse(list)

	a := 1
	b := a
	fmt.Println(a, b)
	fmt.Printf("%p %p\n", &a, &b)
}
func TestReverseList3(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	reverse := reverseList3(list)
	traverse(reverse)
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

func traverse(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Printf("%d ", head.Val)
	}
	fmt.Println()
}

func TestReverseList(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	traverse(reverseList(list))
}

func TestHalfList(t *testing.T) {
	list := populateList([]int{1, 2, 3, 4, 5})
	traverse(halfList(list))
}

func TestIsPalindrome(t *testing.T) {
	list := populateList([]int{1, 2, 3, 2, 1})
	t.Log(isPalindrome(list))
	traverse(list)
}

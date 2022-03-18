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

func TestHasCycle(t *testing.T) {
	list := populateList([]int{1, 2, 3, 4, 5})
	curr := list
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = list

	t.Log(hasCycle1(list))
	// list2 := populateList([]int{1, 1, 1, 1})
	// t.Log(hasCycle(list2))
	t.Log(hasCycle2(list))
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	rotate(nums, 3)

	tmp := make([][]int, 3)
	for i := range tmp {
		tmp[i] = make([]int, 3)
	}
	fmt.Println(tmp, len(tmp), len(tmp[0]))
}

func TestFirst(t *testing.T) {
	var a interface{}
	var i int = 5
	s := "Hello World"
	a = i
	fmt.Println(a)
	a = s
	fmt.Println(a)

	s1 := s[:]
	fmt.Println(s1)
	fmt.Printf("%p %p\n", &s, &s1)

	in := [3]string{"a", "b", "c"}
	var out []*string

	for _, v := range in {
		v := v
		out = append(out, &v)
	}

	fmt.Println("Values:", *out[0], *out[1], *out[2])
}

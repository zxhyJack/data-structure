package algorithm

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// ! 值传递只是修改数据的副本
// ! 指针传递是修改原始数据
// ! 函数传递和变量传递均是如此
// 链表反转1
// 头插法建立新的链表，实现反转
func reverseList1(head *ListNode) *ListNode {
	current := head
	var reverse *ListNode

	for ; current != nil; current = current.Next {
		// 取下原链表的第一个节点
		node := *current // 拿到的是节点的**副本**而不是指针，否则对node指针操作会影响current
		// fmt.Printf("node:%+v curr:%+v\n", &node.Val, &current.Val)
		node.Next = reverse // 将取下的节点放入第二个链表（头插法）
		reverse = &node
	}

	return reverse
}

// 链表反转2
// 对同一个链表操作，只改变指针，将当前节点的Next指针指向前一个节点
func reverseList2(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

// 链表反转2优化版
func reverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	var next *ListNode

	for ; current != nil; current = next {
		next = current.Next
		current.Next = prev
		prev = current

	}
	return prev
}

// 链表反转3
// 将原链表的节点拆下来，使用头插法放到新链表中
// 原链表被修改了
func reverseList3(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	for current != nil {
		node := current
		current = current.Next
		node.Next = prev
		prev = node
	}
	return prev
}

// 两有序链表重组1
// 不修改原链表，新建节点放到新链表中
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy = &ListNode{0, nil} // 虚拟头，解决空指针的问题
	new := dummy
	for list1 != nil && list2 != nil {
		if list1.Val == list2.Val {
			new.Next = &ListNode{list1.Val, nil}
			new = new.Next

			new.Next = &ListNode{list2.Val, nil}
			new = new.Next

			list1 = list1.Next
			list2 = list2.Next
		} else if list1.Val < list2.Val {
			new.Next = &ListNode{list1.Val, nil}
			new = new.Next
			list1 = list1.Next
		} else {
			new.Next = &ListNode{list2.Val, nil}
			new = new.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		new.Next = &ListNode{list1.Val, nil}
		new = new.Next
		list1 = list1.Next
	}
	for list2 != nil {
		new.Next = &ListNode{list2.Val, nil}
		new = new.Next
		list2 = list2.Next
	}
	return dummy.Next
}

// 两有序链表重组2
// 无需新建节点，直接在原链表上操作
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	p := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			p = p.Next // 这一行可以放在if判断后面，两种情况都需要移动指针
			list1 = list1.Next
		} else {
			p.Next = list2
			p = p.Next //
			list2 = list2.Next
		}
	}

	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}
	return dummy.Next
}

// 判断回文链表
func isPalindrome(head *ListNode) bool {
	// 寻找链表的中间节点
	secondHalf := halfList(head)
	// 反转链表的后半部分
	secondHalfReverse := reverseList(secondHalf)

	// 判断是否回文
	p1 := head
	p2 := secondHalfReverse
	res := true // 保存判断的结果，避免不是回文直接返回，为了后面还原链表
	for res && p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			res = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表
	reverseList(secondHalfReverse)

	return res
}

// 寻找链表的中点
func halfList(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 环形链表
// 快慢指针 比较的是指针里面存储的地址
func hasCycle1(head *ListNode) bool {
	fast := head // 快指针
	slow := head // 慢指针
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next // 快指针每次走两步
		slow = slow.Next
		fmt.Printf("%v %v %p %p\n", fast, slow, fast, slow)
		if fast == slow {
			return true
		}
	}
	return false
}

// 环形链表
// hashmap
func hasCycle2(head *ListNode) bool {
	hashmap := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := hashmap[head]; ok {
			return true
		}
		hashmap[head] = struct{}{}
		head = head.Next
	}
	return false
}

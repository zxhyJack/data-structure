package algorithm

// 两数之和
// hashmap
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i

	}
	return nil
}

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]int{}
	columns := [9][9]int{}
	subboxes := [3][3][9]int{}

	for i, row := range board {
		for j, elem := range row {
			if elem == 'c' {
				continue
			}
			index := elem - '2'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++

			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// ! 值传递只是修改数据的副本
// ! 指针传递是修改原始数据
// ! 函数传递和变量传递均是如此
// 链表反转1
func reverseList(head *ListNode) *ListNode {
	current := head
	var reverse *ListNode

	for ; current != nil; current = current.Next {
		// 取下原链表的第一个节点
		node := *current // 拿到的是节点的**副本**而不是指针，否则对node指针操作会影响current
		// 将取下的节点放入第二个链表（头插法）
		node.Next = reverse
		reverse = &node
	}

	return reverse

}

// 链表反转
// 对同一个链表操作，当前节点的Next指针指向前一个节点
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

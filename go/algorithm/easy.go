package algorithm

import "fmt"

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

// 链表反转1
func reverseList(head *ListNode) *ListNode {
	current := head
	head = nil

	for ; current != nil; current = current.Next {
		// 取下原链表的第一个节点
		node := *current // 拿到的是节点而不是指针，否则对node指针操作会影响current
		// 将取下的节点放入第二个链表（头插法）
		node.Next = head
		head = &node
	}

	return head

}

// 链表反转
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

func traverse(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Printf("%d ", head.Val)
	}
	fmt.Println()
}

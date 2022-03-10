package algorithm

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最大深度1
// 递归
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth1(root.Left), maxDepth1(root.Right))
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// 二叉树的最大深度2
// 层次遍历
// 构建两个队列，队列2存放队列1的子节点
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue1 := []*TreeNode{}
	queue1 = append(queue1, root)
	ans := 0
	for len(queue1) > 0 {
		queue2 := []*TreeNode{}
		for len(queue1) > 0 {
			node := queue1[0]
			queue1 = queue1[1:]
			if node.Left != nil {
				queue2 = append(queue2, node.Left)
			}
			if node.Right != nil {
				queue2 = append(queue2, node.Right)
			}
		}
		queue1 = queue2
		ans++
	}
	return ans
}

// 二叉树的最大深度3
// 方法2的优化版本，无需新建队列，利用size存储上一层节点的数量，
// 迭代寻找子节点并存放在队列中，当上一层节点访问完毕，队列中存放的是子节点
func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		ans++
	}
	return ans
}

// 验证二叉搜索树1
// 递归 验证节点值在所允许的范围之间
func isValidBST1(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}

// 验证二叉搜索树2
// 非递归中序遍历，记录当前节点的前一个节点值，
// 由前置节点值是否小于当前节点值验证二叉搜索树
func isValidBST2(root *TreeNode) bool {
	inorder := math.MinInt64
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.Val <= inorder {
				return false
			}
			inorder = root.Val
			root = root.Right
		}
	}
	return true
}

// 验证二叉搜索树3
// 非递归中序遍历
func isValidBST3(root *TreeNode) bool {
	inorder := math.MinInt64
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		for root != nil { // 左节点全部放入栈中
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

// 对称二叉树1
// 迭代
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{}
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil { // 如果都为空，则取消往队列中放节点
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		// 对称放入队列
		queue = append(queue, left.Left)
		queue = append(queue, right.Right)
		queue = append(queue, left.Right)
		queue = append(queue, right.Left)
	}
	return true
}

// 对称二叉树2
// 递归
func isSymmetric2(root *TreeNode) bool {
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left) // 此处放入的节点的顺序是对称的
}

// 二叉树的层序遍历
// 广度优先搜索
// 只使用一个队列，使用size控制该层节点的遍历，size为0时，该层节点遍历完毕
// 也可以使用两个队列，队列1存储该层节点，队列2存储下一层节点
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		innerList := []int{}
		for size > 0 {
			// 出队
			node := queue[0]
			queue = queue[1:]
			// visit(node)
			innerList = append(innerList, node.Val)
			// 非空子节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		res = append(res, innerList)
	}
	return res
}

// 将有序数组转化为二叉搜索树
// 递归，中序遍历 逆向
func sortedArrayToBST(nums []int) *TreeNode {
	return ArrayToBSTHelper(nums, 0, len(nums)-1)
}

func ArrayToBSTHelper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = ArrayToBSTHelper(nums, left, mid-1)
	root.Right = ArrayToBSTHelper(nums, mid+1, right)
	return root
}

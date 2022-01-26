package algorithm

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

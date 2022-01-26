package algorithm

import "testing"

// func populateTree(nums []interface{}) *TreeNode {
// 	for _, num := range nums {
// 		if num != nil
// 	}
// }

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{9, nil, nil},
		Right: &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}},
	}
	depth := maxDepth1(root)
	t.Log("maxDepth1:", depth)

	queue := []*TreeNode{nil}
	t.Log(len(queue))

	depth = maxDepth2(root)
	t.Log("maxDepth2:", depth)

	depth = maxDepth3(root)
	t.Log("maxDepth3:", depth)
}

package algorithm

import (
	"testing"
)

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

func TestIsValidBST(t *testing.T) {
	root := &TreeNode{
		Val:   10,
		Left:  &TreeNode{5, nil, &TreeNode{11, nil, nil}},
		Right: &TreeNode{15, &TreeNode{6, nil, nil}, &TreeNode{20, nil, nil}},
	}
	t.Log(isValidBST1(root))
}

func TestIsValidBST2(t *testing.T) {
	root1 := &TreeNode{
		Val:   10,
		Left:  &TreeNode{5, nil, &TreeNode{11, nil, nil}},
		Right: &TreeNode{15, &TreeNode{6, nil, nil}, &TreeNode{20, nil, nil}},
	}
	root2 := &TreeNode{
		Val:   10,
		Left:  &TreeNode{5, nil, &TreeNode{9, nil, nil}},
		Right: &TreeNode{15, &TreeNode{12, nil, nil}, &TreeNode{20, nil, nil}},
	}
	t.Log(isValidBST2(root1))
	t.Log(isValidBST2(root2))
	t.Log(isValidBST3(root2))
}

func TestIsSymmetric(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, &TreeNode{3, nil, nil}},
		Right: &TreeNode{2, &TreeNode{3, nil, nil}, nil},
	}
	t.Log(isSymmetric1(root))
	t.Log(isSymmetric2(root))
}

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, &TreeNode{3, nil, nil}},
		Right: &TreeNode{2, &TreeNode{3, nil, nil}, nil},
	}
	t.Log(levelOrder(root))
}

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4}
	root := sortedArrayToBST(nums)
	t.Log(levelOrder(root))
}

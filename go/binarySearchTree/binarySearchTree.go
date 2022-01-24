package binarySearchTree

import (
	"errors"
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(elem int) {
	if bst.root == nil {
		bst.root = &Node{data: elem}
		return
	}

	insertRecursive(bst.root, elem)

}

func insertRecursive(root *Node, elem int) *Node {
	if root == nil {
		return &Node{data: elem}

	}

	if elem <= root.data {
		root.left = insertRecursive(root.left, elem)
	}

	if elem > root.data {
		root.right = insertRecursive(root.right, elem)
	}
	return root
}

func (bst *BST) PreOrder() {
	preOrderRecursive(bst.root)
	fmt.Println()
}

func preOrderRecursive(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.data)
		preOrderRecursive(root.left)
		preOrderRecursive(root.right)
	}
}

func (bst *BST) InOrder() {
	inOrderRecursive(bst.root)
	fmt.Println()
}

func inOrderRecursive(root *Node) {
	if root != nil {
		inOrderRecursive(root.left)
		fmt.Printf("%d ", root.data)
		inOrderRecursive(root.right)
	}
}

func (bst *BST) PostOrder() {
	postOrderRecursive(bst.root)
	fmt.Println()
}

func postOrderRecursive(root *Node) {
	if root != nil {
		postOrderRecursive(root.left)
		postOrderRecursive(root.right)
		fmt.Printf("%d ", root.data)
	}
}

func (bst *BST) Min() (int, error) {
	root := bst.root

	if root == nil {
		return 0, errors.New("the tree is empty")
	}
	for root.left != nil {
		root = root.left
	}
	return root.data, nil
}

func (bst *BST) Max() (int, error) {
	root := bst.root
	if root == nil {
		return 0, errors.New("tree is empty")
	}

	for root.right != nil {
		root = root.right
	}

	return root.data, nil
}

func (bst *BST) Search(elem int) bool {
	return searchRecursive(bst.root, elem)
}

func searchRecursive(root *Node, elem int) bool {
	if root == nil {
		return false
	}

	if elem < root.data {
		return searchRecursive(root.left, elem)
	} else if elem > root.data {
		return searchRecursive(root.right, elem)
	} else {
		return true
	}
}

// important: in order(not recursive)
// 重点：左链入栈
func (bst *BST) InOrder2() []int {
	res := []int{}
	root := bst.root
	stack := make([]*Node, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.left
		} else {
			length := len(stack)
			root = stack[length-1]
			stack = stack[:length-1]
			res = append(res, root.data)
			root = root.right
		}
	}
	return res
}

// level order
// 重点：每一层的节点先入队再出队访问
func (bst *BST) LevelOrder() []int {
	res := []int{}
	p := bst.root
	queue := []*Node{}
	queue = append(queue, p)
	for len(queue) != 0 {
		p = queue[0]
		queue = queue[1:]
		res = append(res, p.data)

		if p.left != nil {
			queue = append(queue, p.left)
		}
		if p.right != nil {
			queue = append(queue, p.right)
		}
	}
	return res
}

// 使用BFS层次遍历二叉树
func levelOrder(root *Node) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*Node{root} // 队列1
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*Node{} // 队列2，存放队列1中节点的子节点
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.data)
			if node.left != nil {
				p = append(p, node.left)
			}
			if node.right != nil {
				p = append(p, node.right)
			}
		}
		q = p
	}
	return ret
}

// todo: remove a element form the tree
// func (bst *BST) Remove(elem int) bool {
// 	return removeRecursive(bst.root, elem)
// }

// func removeRecursive(root *Node, elem int) bool {

// 	return false
// }

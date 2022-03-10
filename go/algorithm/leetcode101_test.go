package algorithm

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	ratings := []int{1, 2, 2}
	t.Log(candy(ratings))
}

func TestSearchRange(t *testing.T) {
	nums := []int{1}
	target := 1
	ret := searchRange(nums, target)
	fmt.Println(ret)
}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 1, 2, 5, 4, 6}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestTopK(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	topKFrequent(nums, k)
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, [2]int{1, 3})
	heap.Push(h, [2]int{2, 2})
	heap.Push(h, [2]int{3, 1})
	heap.Push(h, [2]int{4, 1})
	heap.Push(h, [2]int{5, 2})
	heap.Push(h, [2]int{6, 3})
	h.Push([2]int{7, 2})
	h.Push([2]int{8, 3})

	ret := topKFrequent2([]int{5, 3, 1, 1, 1, 3, 73, 1}, 2)
	fmt.Println(ret)
}

func TestGcd(t *testing.T) {
	ret := gcd(4, 10)
	fmt.Println(ret)
	ret = lcm(4, 10)
	fmt.Println(ret)
}

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	n := lengthOfLIS(nums)
	fmt.Println(n)
}

func TestIsAnagram(t *testing.T) {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("ab", "a"))
}

func TestIsIsomorphic(t *testing.T) {
	fmt.Println(isIsomorphic("title", "paper"))
	fmt.Println(isIsomorphic("foo", "bar"))
}

func TestLengthOfLast(t *testing.T) {
	str := "hello newcoder"
	fmt.Println(lengthOfLast(str))
	var a string
	fmt.Println(a)
	b := make([]byte, 10)
	fmt.Println(b[0])
	c := "hello"
	fmt.Println(c[0])
	
}

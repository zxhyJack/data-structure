package algorithm

import (
	"testing"
)

func TestMerge1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge1(nums1, 3, nums2, 3)
	t.Log(nums1)
}

func TestMerge2(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge2(nums1, 3, nums2, 3)
	t.Log(nums1)
}

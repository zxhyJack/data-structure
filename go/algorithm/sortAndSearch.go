package algorithm

import "sort"

// 合并两个有序数组1
// 逆向双指针（直接在原数组上修改）
// 双指针，从后往前移动，比较所指元素的大小
// 再有一个指针指向nums1的末尾，存放双指针比较得到的较大的数
func merge1(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// 合并两个有序数组2
// 双指针（需要额外的数组）
// 双指针，从前往后移动，比较所指元素的大小
func merge2(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)
	p1, p2, k := 0, 0, 0
	for p1 < m && p2 < n {
		if nums1[p1] <= nums2[p2] {
			nums3[k] = nums1[p1]
			k++
			p1++
		} else {
			nums3[k] = nums2[p2]
			k++
			p2++
		}
	}
	if p1 < m {
		copy(nums3[k:], nums1[p1:])
	}
	if p2 < n {
		copy(nums3[k:], nums2[p2:])
	}
	copy(nums1, nums3)
}

// 第一个错误的版本
func firstBadVersion1(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func firstBadVersion2(n int) int {
	return sort.Search(n, isBadVersion)
}

func firstBadVersion3(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func isBadVersion(n int) bool {
	return n == 3
}

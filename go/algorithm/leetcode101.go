package algorithm

import (
	"container/heap"
	"sort"
)

// 455 分发饼干
// 贪心
// 给剩下的孩子中最小的孩子以能饱腹的最小的饼干
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans := 0
	m, n := len(g), len(s)
	for i, j := 0, 0; i < m && j < n; i++ {
		for j < n && g[i] > s[j] {
			j++
		}
		if j < n {
			ans++
			j++
		}
	}
	return ans
}

// 135 分发糖果
// 困难
// 两次遍历
func candy(ratings []int) int {
	ans := []int{}
	for i := 0; i < len(ratings); i++ {
		ans = append(ans, 1)
	}

	// 正向遍历
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			ans[i] = ans[i-1] + 1
		}
	}

	// 反向遍历
	for j := len(ratings) - 2; j >= 0; j-- {
		// if ratings[j] > ratings[j+1] && ans[j] <= ans[j+1] {
		// 	ans[j] = ans[j+1] + 1 // j比j+1大，并且j的糖果不如j+1多
		// }
		if ratings[j] > ratings[j+1] {
			ans[j] = max(ans[j+1]+1, ans[j])
		}
	}

	sum := 0
	for _, num := range ans {
		sum += num
	}

	return sum
}

// 435 无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	ans, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			ans++
			right = p[1]
		}
	}
	return n - ans
}

// 两数之和
func twoSum1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	hashmap := map[int]int{}
	for i, v := range nums {
		if j, ok := hashmap[target-v]; ok {
			return []int{j, i}
		} else {
			hashmap[v] = i
		}
	}
	return []int{}
}

// 167. 两数之和II-输入有序数组
// 双指针法
func twoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

// 88. 归并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	q := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[q] = nums1[p1]
			p1--
		} else {
			nums1[q] = nums2[p2]
			p2--
		}
		q--
	}
	for p1 >= 0 {
		nums1[q] = nums1[p1]
		p1--
		q--
	}

	for p2 >= 0 {
		nums1[q] = nums2[p2]
		p2--
		q--
	}
}

// 142. 环形链表II
// 给定一个链表，如果有环路，找出环路的开始点
// 快慢指针
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next

		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// 76. 最小覆盖字串
// 给定两个字符串 S 和 T，求 S 中包含 T 所有字符的最短连续子字符串的长度，同时要求时间复杂度不得超过 O(n)
func minWindow(s string, t string) string {
	return ""
}

// 69. x的平方根
func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	lower := leftBound(nums, target)
	upper := rightBound(nums, target) - 1
	if lower == len(nums) || nums[lower] != target {
		return []int{-1, -1}
	}
	return []int{lower, upper}
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 快排
func quickSort(arr []int, left int, right int) []int {
	if left < right {
		partitionpos := partition2(arr, left, right)
		quickSort(arr, left, partitionpos-1)
		quickSort(arr, partitionpos+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}

func partition2(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}

// 数组中第k个最大元素
func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	target := len(nums) - k
	for left < right {
		index := quickSelect(nums, left, right)
		if index == target {
			return nums[index]
		}
		if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
	return nums[left]
}

func quickSelect(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]

		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

// 347. 前k个高频元素
// 给定一个数组，求前 k 个最频繁的数字
// 用堆排序，小顶堆
func topKFrequent(nums []int, k int) []int {
	// 用哈希存储数组中数字及出现的次数，数字做键，出现的次数做值
	hashmap := map[int]int{}
	for _, v := range nums {
		if _, ok := hashmap[v]; ok {
			hashmap[v]++
		} else {
			hashmap[v] = 1
		}
	}
	// for _, num := range nums {
	// 	hashmap[num]++
	// }
	h := &IntHeap{}
	heap.Init(h)
	for key, value := range hashmap {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-1-i] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 桶排序
func topKFrequent2(nums []int, k int) []int {
	// hash统计数字出现的次数
	hashmap := map[int]int{}
	max_count := 0 // 记录数字出现的最大次数
	for _, num := range nums {
		hashmap[num]++
		if max_count < hashmap[num] {
			max_count = hashmap[num]
		}
	}

	arr := make([][]int, max_count+1)
	for key, value := range hashmap {
		arr[value] = append(arr[value], key)
	}
	ret := []int{}
	for i := 0; i <= k && len(ret) <= k; i++ {
		ret = append(ret, arr[max_count-i]...)
		if len(ret) == k {
			break
		}
	}
	return ret
}

// 最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 最小公倍数
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 300. 最长递增子序列
// 给定一个未排序的整数数组，求最长的递增子序列
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	max_length := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1) // 状态转移方程
			}
		}
		max_length = max(max_length, dp[i])
	}
	return max_length
}

// 1143. 最长公共子序列
//给定两个字符串，求它们最长的公共子序列长度
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 //
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) //
			}
		}
	}
	return dp[m][n]
}

// 242. 有效的字母异位词
// 判断两个字符串包含的字符是否完全相同
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := [26]int{}
	for _, c := range s {
		index := c - 'a'
		arr[index]++
	}
	for _, c := range t {
		arr[c-'a']--
		if arr[c-'a'] < 0 {
			return false
		}
	}
	return true
}

// 205. 同构字符串
func isIsomorphic(s string, t string) bool {
	st1 := [256]int{}
	st2 := [256]int{}
	for i := 0; i < len(s); i++ {
		if st1[s[i]] != st2[t[i]] {
			return false
		}
		st1[s[i]] = i + 1
		st2[t[i]] = i + 1
	}
	return true
}

// 647. 回文子字符串
// 给定一个字符，求其有多少个回文子字符串。回文的定义是左右对称

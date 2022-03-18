package algorithm

// 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target
// 在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解法：hashmap
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

// 有效的数独
// 规则：
// 1. 数字 1-9 在每一行只能出现一次。
// 2. 数字 1-9 在每一列只能出现一次。
// 3. 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次
// 输入：board =
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// 输出：true
func isValidSudoku(board [][]byte) bool {
	rows := [9][9]int{}
	columns := [9][9]int{}
	subboxes := [3][3][9]int{}

	for i, row := range board {
		for j, elem := range row {
			if elem == '.' {
				continue
			}
			index := elem - '1'
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

// 移动零
// 将给定一个数组 nums 的所有 0 移动到数组的末尾，同时保持非零元素的相对顺序
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 双指针法
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)

	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// 加一
// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一
// 输入：digits = [1,2,3]
// 输出：[1,2,4]
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- { // 倒序寻找第一个不是9的数字
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ { // 该数字后面的数字全置为0
				digits[j] = 0
			}
			return digits
		}
	}

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

// 350 两个数组的交集II
// 以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2,2]
// hashtable
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int) // hash 统计 nums1 数字出现的次数

	for _, v := range nums1 {
		m[v]++
	}

	intersection := make([]int, 0) // 存放两数组的交集
	for _, v := range nums2 {
		if m[v] > 0 {
			intersection = append(intersection, v)
			m[v]--
		}
	}
	return intersection
}

// 136. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
// 要求时间复杂度O(n), 空间复杂度O(1)
// 位运算，异或
// 利用异或的性质
// 1. 任何数和0异或，结果仍然是原来的数
// 2. 任何数和自身异或， 结果为0
// 3. 异或运算满足交换律和结合率
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}

	return single
}

// 217. 存在重复元素
// hash 存储元素的个数，在存储时，如果该元素已经存在，说明该元素是重复元素，返回true，
// 如果存储完所有元素都没有发现元素重复，则返回false
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, num := range nums {
		if m[num] > 0 {
			return true
		}
		m[num]++
	}
	return false
}

// 189. 轮转数组1
// 每次移动一个数字，数组头存储最后一个数字， 其他数字顺次后移
// 性能差，超时
func rotate(nums []int, k int) {
	n := len(nums)
	for i := 0; i < k; i++ {
		tmp := nums[n-1]
		for j := n - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
}

// 189. 轮转数组1
// 数组反转
// 先整体反转，再单独反转
func rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

// 122. 买卖股票的最佳时机 II
func maxProfit(prices []int) int {
	n := len(prices)

	if n <= 1 {
		return 0
	}

	ret := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			ret += prices[i] - prices[i-1] // 交易次数不受限，可以把所有的上坡全部收集到
		}
	}
	return ret
}

// 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 双指针
	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] { // 当快指针和慢指针指向的数字不同时，慢指针的下一位存储快指针指向的值
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

//
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 双指针
	slow := 1

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] { // 使用快指针判断是否重复
			nums[slow] = nums[fast] // 不重复时赋值给慢指针
			slow++
		}
	}

	return slow
}

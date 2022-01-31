package algorithm

import "math"

// 爬楼梯1
// 递归 fibonacci
func climbStairs1(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}

// 爬楼梯2
// 滚动数组法，只有三个变量，用来存放f(x),f(x-1),f(x-2)，优化空间复杂度
func climbStairs2(n int) int {
	p := 0
	q := 1
	cur := 1
	for i := 2; i <= n; i++ { // 从n=2开始计算fibonacci
		p = q
		q = cur
		cur = p + q
	}
	return cur
}

// 使用数组来存储每一级的结果
func climbStairs3(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1) // n从0计数，总共有n+1级台阶
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs4(n int) int {
	if n < 2 {
		return 1
	}
	dp := []int{1, 1}
	for i := 2; i < n+1; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}

// 买卖股票的最佳时机
// 暴力解法
func maxProfit1(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

// 买卖股票的最佳时机
// 一次遍历
func maxProfit2(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

// 最大子序和
// 动态规划
// 转移公式：dp[i] = max(dp[i-1], 0) + dp[i]
// 或 dp[i] = max(dp[i-1]+nums[i], nums[i])
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	maxVal := nums[0]
	dp[0] = nums[0] // 边界条件
	for i := 1; i < len(nums); i++ {
		// if dp[i-1] < 0 {
		// 	dp[i] = nums[i]
		// } else {
		// 	dp[i] = dp[i-1] + nums[i]
		// }
		dp[i] = max(dp[i-1], 0) + nums[i] // 转移公式
		maxVal = max(maxVal, dp[i])       // 记录最大值
	}
	return maxVal
}

// 最大子序和
// 动态规划优化
// 无需使用数组存储每一项的前最大子序和
// 仅使用变量存储前一个变量的最大子序和
func maxSubArray2(nums []int) int {
	cur := nums[0]
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		cur = max(cur, 0) + nums[i]
		maxVal = max(maxVal, cur)
	}
	return maxVal
}

// 打家劫舍
// 动态规划 使用额外数组
// dp[i] = max(dp[i-2]+nums[i], dp[i-1])
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[length-1]
}

// 打家劫舍
// 动态规划优化 使用滚动数组简化空间复杂度
// dp[i] = max(dp[i-2]+nums[i], dp[i-1])
func rob2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	pre1 := nums[0]
	pre2 := max(nums[0], nums[1])
	cur := max(nums[0], nums[1]) // 第一个指向nums[2]
	for i := 2; i < length; i++ {
		cur = max(pre1+nums[i], pre2)
		pre1 = pre2
		pre2 = cur
	}
	return cur
}

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

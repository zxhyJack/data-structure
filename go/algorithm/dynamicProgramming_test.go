package algorithm

import "testing"

func TestClimbStairs1(t *testing.T) {
	t.Log(climbStairs1(5))
}

func TestClimbStairs2(t *testing.T) {
	t.Log(climbStairs2(2))
}

func TestClimbStairs3(t *testing.T) {
	t.Log(climbStairs3(4))
}

func TestClimbStairs4(t *testing.T) {
	t.Log(climbStairs4(4))
}

func TestMaxProfit1(t *testing.T) {
	t.Log(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
}

func TestMaxProfit2(t *testing.T) {
	t.Log(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, -1}
	t.Log(maxSubArray(nums))
}

func TestMaxSubArray2(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArray2(nums))
}

func TestRob(t *testing.T) {
	nums := []int{1, 2}
	t.Log(rob(nums))
	t.Log(max(1, 2))
}

func TestRob2(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	t.Log(rob2(nums))
}

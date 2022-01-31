package algorithm

import (
	"fmt"
	"math/rand"
)

// 打乱数组
// 主要是数组随机打乱的算法
type Solution struct {
	data []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Reset() []int {
	return this.data
}

func (this *Solution) Shuffle() []int {
	// this.count++
	tmp := append([]int{}, this.data...)
	shuffle := []int{}
	length := len(tmp)
	// rand.Seed(this.count)
	for i := 0; i < length; i++ {
		j := rand.Intn(len(tmp))
		fmt.Printf("%d ", j)
		shuffle = append(shuffle, tmp[j])
		tmp = append(tmp[:j], tmp[j+1:]...)
	}
	return shuffle
}

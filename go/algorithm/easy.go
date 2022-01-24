package algorithm

// 两数之和
// hashmap
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

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]int{}
	columns := [9][9]int{}
	subboxes := [3][3][9]int{}

	for i, row := range board {
		for j, elem := range row {
			if elem == 'c' {
				continue
			}
			index := elem - '2'
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

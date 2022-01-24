package binarySearchTree

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestBST(t *testing.T) {
	bst := BST{}

	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)

	// t.Log(bst.root.data)
	bst.PreOrder()
	bst.InOrder()
	bst.PostOrder()

	t.Log(bst.Min())
	t.Log(bst.Max())

	t.Log(bst.Search(4))

	t.Log(bst.InOrder2())

	t.Log(bst.LevelOrder())

	t.Log(levelOrder(bst.root))

	q := []*Node{nil} // 需要提前判断放入队列中的节点是否未nil
	t.Log(q)

	t.Log(levelOrder(bst.root))

	t.Log(string("8"))

	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	// for i := 0; i < len(board); i++ {
	// 	for j := 0; j < len(board[i]); j++ {
	// 		fmt.Printf("%c ", board[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("---------------------")
	// for j := 0; j < len(board[0]); j++ {
	// 	for i := 0; i < len(board); i++ {
	// 		fmt.Printf("%c ", board[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("---------------------")
	// for i := 0; i < len(board); i++ {
	// 	for j := 0; j < len(board); j++ {
	// 		fmt.Printf("%c ", board[j][i])
	// 	}
	// 	fmt.Println()
	// }

	t.Log(isValidSudoku(board))
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	t.Log(s)
	reverseString(s)
	t.Log(s)

	t.Log(reverse(1563847412), math.MaxInt32)
	t.Log(-13 % 10)
	t.Log(-1<<31, math.MinInt32)
	t.Log(1<<31-1, math.MaxInt32)

	hashmap := map[rune]int{}
	cnt := [26]int{}
	for i, v := range "leetcode" {
		t.Log(i, v-'a')
		cnt[v-'a']++
		hashmap[v-'a']++
	}

	t.Log(isPalindrome("race a car"))

	t.Log(strings.ToLower(" ;"))
	// str := " 123"
	r := int('1')
	t.Log(r)
	t.Log(myAtoi("-2147483647"))

	strs := []string{"flower", "flow", "flight"}
	for j := 0; j < len(strs); j++ {
		for i := 0; i < 3; i++ {
			fmt.Printf("%c ", strs[i][j])
		}
	}

}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	// 双指针
	for i := 0; i <= len(haystack)-len(needle); i++ {
		a := i
		b := 0
		for b < len(needle) && haystack[a] == needle[b] {
			a++
			b++
		}
		if b == len(needle) {
			return i
		}
	}
	return -1
}

func myAtoi(s string) int {
	res := 0
	// 处理空格
	index := 0
	for index < len(s) && s[index] == ' ' {
		index++
	}
	// 处理全是空格的情况
	if index == len(s) {
		return 0
	}

	// 处理符号位
	flag := 1
	if s[index] == '-' {
		flag = -1
		index++
	} else if s[index] == '+' {
		index++
	}
	// 处理数字和字符
	for ; index < len(s); index++ {
		if s[index] > '9' || s[index] < '0' {
			break
		}

		if res < -1<<31/10 || (res == -1<<31/10 && int(s[index]-'0') > 1<<31%10) {
			return -1 << 31
		}
		if res > (1<<31-1)/10 || (res == (1<<31-1)/10 && int(s[index]-'0') > (1<<31-1)%10) {
			return 1<<31 - 1
		}

		res = res*10 + int(s[index]-'0')*flag
	}
	return res
}

func isPalindrome(s string) bool {
	str := ""
	for i, ch := range s {
		if isNumAndLetter(s[i]) {
			str += string(ch)
		}
	}

	str = strings.ToLower(str)
	fmt.Println(str)

	for left, right := 0, len(str)-1; left < right; {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isNumAndLetter(ch byte) bool {
	if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
		return true
	}
	return false
}

func reverse(x int) int {
	ret := 0
	digit := x % 10
	x /= 10
	ret = ret*10 + digit
	for x != 0 {
		if ret < -1<<31/10 || ret > (1<<31-1)/10 {
			return 0
		}
		digit = x % 10
		x /= 10
		ret = ret*10 + digit
	}
	return ret
}

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

func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

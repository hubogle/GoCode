// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
// 示例 1：
//
//
// 输入：n = 4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
// 输入：n = 1
// 输出：[["Q"]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
//
// Related Topics 数组 回溯 👍 1214 👎 0

package main

import "fmt"

func main() {
	result := solveNQueens(4)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	var ans [][]string
	var backtracking func(index int)
	used := make([][]byte, n)
	for k, _ := range used {
		var v []byte
		for i := 0; i < n; i++ {
			v = append(v, '.')
		}
		used[k] = v
	}
	backtracking = func(index int) {
		if index == n {
			var s []string
			for _, v := range used {
				s = append(s, string(v))
			}
			ans = append(ans, s)
			return
		}
		for i := 0; i < n; i++ {
			if isUsed(used, index, i) {
				used[index][i] = 'Q'
				backtracking(index + 1) // 下一层，和上一层所在的位置
				used[index][i] = '.'
			}
		}
	}
	backtracking(0)
	return ans
}

func isUsed(used [][]byte, row, col int) bool {
	for i := row - 1; i >= 0; i-- {
		if used[i][col] == 'Q' {
			return false
		}
	} // 同一列不能有
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if used[i][j] == 'Q' {
			return false
		}
	} // 左上不能有
	for i, j := row-1, col+1; i >= 0 && j < len(used); i, j = i-1, j+1 {
		if used[i][j] == 'Q' {
			return false
		}
	} // 右上不能有
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目解析：主要是每层回溯递归时。
皇后所在位置：
1. 同一行不能有；
2. 同一列不能有；
3. 当前皇后所在的对角线不能有；

递归是行遍历，循环是列遍历
*/

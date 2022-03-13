// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
//
//
//
//
//
// 示例 1：
//
//
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
// "ABCCED"
// 输出：true
//
//
// 示例 2：
//
//
// 输入：board = [["a","b"],["c","d"]], word = "abcd"
// 输出：false
//
//
//
//
// 提示：
//
//
// 1 <= board.length <= 200
// 1 <= board[i].length <= 200
// board 和 word 仅由大小写英文字母组成
//
//
//
//
// 注意：本题与主站 79 题相同：https://leetcode-cn.com/problems/word-search/
// Related Topics 数组 回溯 矩阵 👍 541 👎 0

package main

import "fmt"

func main() {
	// result := exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB")
	result := exist([][]byte{{'c', 'a', 'a'}, {'a', 'a', 'a'}, {'b', 'c', 'd'}}, "aab")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	var dfs func(row, col, index int) bool // 上一个匹配的位置，以及该匹配第几个字符串
	wordLen := len(word)
	rowLen := len(board)
	colLen := len(board[0])
	dfs = func(row, col, index int) bool {
		if row < 0 || row >= rowLen || col < 0 || col >= colLen || board[row][col] != word[index] {
			return false
		}
		if index == wordLen-1 {
			return true
		} // 长度满足之间返回
		board[row][col] = ' '
		if dfs(row+1, col, index+1) || dfs(row-1, col, index+1) || dfs(row, col-1, index+1) || dfs(row, col+1,
			index+1) {
			return true // dfs 搜索
		} else {
			board[row][col] = word[index] // 恢复原来的字符
			return false
		}
	}
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)

func exist1(board [][]byte, word string) bool {
	var backtracking func(row, col, index int) bool // 上一个匹配的位置，以及该匹配第几个字符串
	wordLen := len(word)
	rowLen := len(board)
	colLen := len(board[0])
	backtracking = func(row, col, index int) bool {
		if index == wordLen {
			return true
		}
		if index == 0 { // 第一个位置随意
			for i := row; i < rowLen; i++ {
				for j := col; j < colLen; j++ {
					val := board[i][j]
					if val != '0' && val == word[index] { // 未被用过
						board[i][j] = '0'
						if backtracking(i, j, index+1) {
							return true
						}
						board[i][j] = val
					}
				}
			}
		} else {
			if row+1 < rowLen && board[row+1][col] == word[index] {
				val := board[row+1][col]
				board[row+1][col] = '0'
				if backtracking(row+1, col, index+1) {
					return true
				}
				board[row+1][col] = val
			}
			if row > 0 && board[row-1][col] == word[index] {
				val := board[row-1][col]
				board[row-1][col] = '0'
				if backtracking(row-1, col, index+1) {
					return true
				}
				board[row-1][col] = val
			}
			if col > 0 && board[row][col-1] == word[index] {
				val := board[row][col-1]
				board[row][col-1] = '0'
				if backtracking(row, col-1, index+1) {
					return true
				}
				board[row][col-1] = val
			}
			if col+1 < colLen && board[row][col+1] == word[index] {
				val := board[row][col+1]
				board[row][col+1] = '0'
				if backtracking(row, col+1, index+1) {
					return true
				}
				board[row][col+1] = val
			}
		}
		return false
	}
	return backtracking(0, 0, 0)
}

/*
解题思路：
一开始：
1. 第一位可以全图任意位置。
2. 后面的必须在前一位附近。
3. 不能用到已使用的字符。 `0` 标识已使用
后续：
dfs 思路，两层 for 循环遍历每个位置。
1. 不符合的条件 row < 0 || row >= rowLen || col < 0 || col >= colLen || board[row][col] != word[index]
2. dfs 遍历 上下左右四个位置
3. 恢复原来的字符
*/

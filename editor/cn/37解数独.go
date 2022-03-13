// 编写一个程序，通过填充空格来解决数独问题。
//
// 数独的解法需 遵循如下规则：
//
//
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//
//
// 数独部分空格内已填入了数字，空白格用 '.' 表示。
//
//
//
//
//
//
// 示例：
//
//
// 输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".
// ",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".
// ","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6
// "],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[
// ".",".",".",".","8",".",".","7","9"]]
// 输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8
// "],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],[
// "4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9",
// "6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4",
// "5","2","8","6","1","7","9"]]
// 解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
//
//
//
//
//
//
// 提示：
//
//
// board.length == 9
// board[i].length == 9
// board[i][j] 是一位数字或者 '.'
// 题目数据 保证 输入数独仅有一个解
//
//
//
//
// Related Topics 数组 回溯 矩阵 👍 1158 👎 0

package main

func main() {
	solveSudoku([][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}})
}

// leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	var backtracking func() bool
	n := len(board)
	backtracking = func() bool {
		for row := 0; row < n; row++ {
			for col := 0; col < n; col++ {
				if board[row][col] != '.' {
					continue
				}
				for i := '1'; i <= '9'; i++ {
					if isUsed(board, row, col, byte(i)) {
						board[row][col] = byte(i)
						if backtracking() {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtracking()
}

func isUsed(used [][]byte, row, col int, v byte) bool {
	n := len(used)
	for i := 0; i < n; i++ {
		if used[i][col] == v {
			return false
		}
	}
	for j := 0; j < n; j++ {
		if used[row][j] == v {
			return false
		}
	}
	row = row / 3
	col = col / 3
	for i := row * 3; i < row*3+3; i++ {
		for j := col * 3; j < col*3+3; j++ {
			if used[i][j] == v {
				return false
			}
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

// 错误写法，一层for循环
func solveSudoku1(board [][]byte) {
	var backtracking func(row int) bool
	n := len(board)
	backtracking = func(row int) bool {
		for i := 0; i < n && row < n; i++ {
			if board[row][i] != '.' {
				continue
			}
			for v := '1'; v <= '9'; v++ {
				if isUsed(board, row, i, byte(v)) {
					board[row][i] = byte(v)
					if backtracking(row + 1) {
						return true
					}
					board[row][i] = '.'
				}
			}
			return false
		}
		return true
	}
	backtracking(0)
}

/*
解题思路：
一层 for 循环根本不能把当前层每个元素值的情况考虑到。
只有两层 for 循环。
与 N 皇后问题不同，N皇后当前层只需要考虑一个位置的存放。
而数独问题需要考虑当前层所有位置的存放。
*/

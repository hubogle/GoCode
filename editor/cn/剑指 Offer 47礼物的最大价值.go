// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直
// 到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
//
//
//
// 示例 1:
//
// 输入:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// 输出: 12
// 解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
//
//
//
// 提示：
//
//
// 0 < grid.length <= 200
// 0 < grid[0].length <= 200
//
// Related Topics 数组 动态规划 矩阵 👍 250 👎 0

package main

import "fmt"

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	result := maxValue(grid)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxValue(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	for i := 1; i < row; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < col; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			grid[i][j] += max(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[row-1][col-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func maxValue1(grid [][]int) int {
	var dp = make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for row := 1; row < len(grid); row++ {
		dp[row][0] = grid[row][0] + grid[row-1][0]
	} // dp 数组第一行累加
	for col := 1; col < len(grid[0]); col++ {
		dp[0][col] = grid[0][col] + grid[0][col-1]
	} // dp 数组第一列累加
	for row := 1; row < len(grid); row++ {
		for col := 1; col < len(grid[0]); col++ {
			dp[row][col] = max(dp[row-1][col], dp[row][col-1]) + grid[row][col]
		} // 从当前行的上面，左边获取最大值然后加上当前的值
	}
	return dp[len(grid)][len(grid[0])-1]
}

func maxValue2(grid [][]int) int {
	var dp = make([][]int, len(grid)+1)
	for i := range dp {
		dp[i] = make([]int, len(grid[0])+1)
	}
	for row := 1; row <= len(grid); row++ {
		for col := 1; col <= len(grid[0]); col++ {
			dp[row][col] = max(dp[row-1][col], dp[row][col-1]) + grid[row-1][col-1]
		}
	}
	return dp[len(grid)][len(grid[0])]
}

/*
题目思路：动态规划 dp 定义好 dp 数组

dp 数组两种方式：
方式1：定义 dp 数组大小与原数组大小一致 dp[row][col] = max(dp[row-1][col], dp[row][col-1]) + grid[row][col]
方式2：定义 dp 数组大小比原数组增加一列和一行 dp[row][col] = max(dp[row-1][col], dp[row][col-1]) + grid[row-1][col-1]
方式3：不需要 dp 数组在原数组上进行累加目标值

方式2相对方式1好处在于，少些部分代码，将累加第一行，第一列的代码都放在 for 循环中去做

0000
0000
0000
0000

131
151
421

*/

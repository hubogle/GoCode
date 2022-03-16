// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
// 问总共有多少条不同的路径？
//
//
//
// 示例 1：
//
//
// 输入：m = 3, n = 7
// 输出：28
//
// 示例 2：
//
//
// 输入：m = 3, n = 2
// 输出：3
// 解释：
// 从左上角开始，总共有 3 条路径可以到达右下角。
// 1. 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右
// 3. 向下 -> 向右 -> 向下
//
//
// 示例 3：
//
//
// 输入：m = 7, n = 3
// 输出：28
//
//
// 示例 4：
//
//
// 输入：m = 3, n = 3
// 输出：6
//
//
//
// 提示：
//
//
// 1 <= m, n <= 100
// 题目数据保证答案小于等于 2 * 10⁹
//
// Related Topics 数学 动态规划 组合数学 👍 1287 👎 0

package main

import "fmt"

func main() {
	result := uniquePaths(7, 3)
	fmt.Println(result)
}

/*
3 * 3
100
000
001

*/

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	} // dp 数组从二维压缩到一维
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[col] += dp[col-1]
		}
	}
	return dp[n-1]
}

// leetcode submit region end(Prohibit modification and deletion)

// dp 初始化数组
func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	} // 从 0,0 到 m,0
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[row][col] = dp[row-1][col] + dp[row][col-1]
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

// 也是 dp，只不过将 dp 边的初始化放在里面
func uniquePaths2(m int, n int) int {
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				dp[i][j] += dp[i][j-1]
			} else if i > 0 && j == 0 {
				dp[i][j] += dp[i-1][j]
			} else if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// dfs 深度优先遍历，时间复杂度 O(2^(m+n+1) -1)
func uniquePaths3(m int, n int) int {
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row > m || col > n {
			return 0
		}
		if row == m && col == n {
			return 1
		}
		return dfs(row+1, col) + dfs(row, col+1)
	}
	return dfs(1, 1)
}

/*
解题思路：动态规划 dp[i][j] = dp[i-1][j] + dp[i][j-1]
解题思路1：申请二维数据，两层 for 循环分别遍历 m、n
解体思路2：申请一维数组，两层 for 循环分别遍历 m、n，类似于滚动数组

滚动数组：将行遍历的结果和列遍历的结果想成 dp[i] = dp[i] + dp[i-1] 从二维输出抽出来
*/

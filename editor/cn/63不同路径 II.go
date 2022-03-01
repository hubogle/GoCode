// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
//
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//
//
// 示例 1：
//
//
// 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// 输出：2
// 解释：3x3 网格的正中间有一个障碍物。
// 从左上角到右下角一共有 2 条不同的路径：
// 1. 向右 -> 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右 -> 向右
//
//
// 示例 2：
//
//
// 输入：obstacleGrid = [[0,1],[0,0]]
// 输出：1
//
//
//
//
// 提示：
//
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] 为 0 或 1
//
// Related Topics 数组 动态规划 矩阵 👍 736 👎 0

package main

import "fmt"

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	result := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row && obstacleGrid[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < col && obstacleGrid[0][j] != 1; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[row-1][col-1]
}

// leetcode submit region end(Prohibit modification and deletion)
/*
解体思路：相对于62不同路径，增加了路障的情况。

步骤1：在初始化二维数组边缘时，遇到路障后后面都不可达。
步骤2：在遍历二维数组内部时，遇到路障就不可达进行特殊处理。
*/

// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
//
//
// 示例 1：
//
//
// 输入：n = 3
// 输出：5
//
//
// 示例 2：
//
//
// 输入：n = 1
// 输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 19
//
// Related Topics 树 二叉搜索树 数学 动态规划 二叉树 👍 1618 👎 0

package main

import "fmt"

func main() {
	result := numTrees(3)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
			// dp[j-1] 以 j 为根节点的左子树的数量
			// dp[i-j] 以 j 为根节点的右子树的数量
		}
	}
	return dp[n]
}

// leetcode submit region end(Prohibit modification and deletion)
/*
贪心思路：
发现重叠子问题.
dp[i] += dp[j-1] * dp[i-j] 左子树节点数量 * 右子树节点数量

dp[3] += dp[0] * dp[2]
dp[3] += dp[1] * dp[1]
dp[3] += dp[2] * dp[0]
*/

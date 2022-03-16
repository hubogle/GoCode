// 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
//
// 返回 你可以获得的最大乘积 。
//
//
//
// 示例 1:
//
//
// 输入: n = 2
// 输出: 1
// 解释: 2 = 1 + 1, 1 × 1 = 1。
//
// 示例 2:
//
//
// 输入: n = 10
// 输出: 36
// 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
//
//
//
// 提示:
//
//
// 2 <= n <= 58
//
// Related Topics 数学 动态规划 👍 746 👎 0

package main

import "fmt"

func main() {
	result := integerBreak(10)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ { // 遍历拆分数字
			ans := max(j*(i-j), dp[i-j]*j) // 寻找是两个数乘机积，还是拆分数乘积大
			dp[i] = max(ans, dp[i])
		}
	}
	return dp[n]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// leetcode submit region end(Prohibit modification and deletion)
/*
递推公式：
从 1 遍历 i 遍历过程数值为 j
dp[i] = j * (i - j) 整数拆分两个相乘
dp[i] = j * dp[i-j] 拆分两个及两个以上的数
dp[i] = max(dp[i], i*(i-j), j * dp[i-j])
遍历顺序先有 dp[i-j] 再有 dp[i]
*/

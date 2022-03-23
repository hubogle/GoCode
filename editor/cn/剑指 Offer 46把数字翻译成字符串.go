// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可
// 能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
//
//
// 示例 1:
//
// 输入: 12258
// 输出: 5
// 解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
//
//
//
// 提示：
//
//
// 0 <= num < 2³¹
//
// Related Topics 字符串 动态规划 👍 397 👎 0

package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := translateNum(12258)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func translateNum(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	ans := 1 // 首个字符必有一种
	cur, next := 0, 1
	for i := 1; i < n; i++ {
		cur, next, ans = next, ans, 0 // 滚动数组
		ans += next
		if s[i-1:i+1] <= "25" && s[i-1:i+1] >= "10" {
			ans += cur
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func translateNum1(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1 // 第一个字符只有一种情况
	for i := 2; i <= n; i++ {
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '5') {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}

/*
解题思路：动态规划
第一个字符只有一种情况，第二个字符的情况要看前一位字符是否为 1/2
dp[i] = dp[i-1] + dp[i-2] 的情况是 s[i-2] 字符等于 1 or s[i-2]="2" s[i-1]<="5" 否则就是 dp[i] = dp[i-1]

进一步：将 dp 数组优化为滚动数组，空间复杂度为 O(1)
*/

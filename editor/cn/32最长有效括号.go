// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
// 输入：s = ""
// 输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
//
// Related Topics 栈 字符串 动态规划 👍 1713 👎 0

package main

import "fmt"

func main() {
	result := longestValidParentheses(")()())")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	ans := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' { // 与上一位匹配 ()
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			}
			if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' { // (()) 这种匹配
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			if dp[i] > ans {
				ans = dp[i]
			}
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

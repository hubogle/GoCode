// 给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
// 输入：s = "cbbd"
// 输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
// Related Topics 字符串 动态规划 👍 5170 👎 0

package main

import "fmt"

func main() {
	result := longestPalindrome("babad")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, n := 0, 0, len(s)
	ansl, ansr := 0, 0
	for left < n {
		for right+1 < n && s[right+1] == s[left] { // 有相同字母的情况
			right++
		}
		for left-1 >= 0 && right+1 < n && s[left-1] == s[right+1] { // 寻找边界
			left--
			right++
		}
		if right-left > ansr-ansl {
			ansr, ansl = right, left
		}
		left = (left+right)/2 + 1 // 重置到下一次寻找回文的中心
		right = left
	}
	return s[ansl : ansr+1]
}

// leetcode submit region end(Prohibit modification and deletion)

/*
5.最长回文子串

解题思路：滑动窗口
有相同字母的情况下，右指针向右边移动下。
当遇到相同字母时，左右指针同时向两边扩展，判断最大的左右指针边界。
左指针是上次中心节点的右边。
*/

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 字符串 滑动窗口 👍 7143 👎 0

package main

import "fmt"

func main() {
	result := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	cache := map[byte]int{}
	left, right, ans := 0, 0, 0
	for left < n {
		if right < n && cache[s[right]] == 0 { // 右指针符合条件的话，一直右移
			cache[s[right]]++
			right++
		} else { // 不符合条件，则左指针左移动，且从 cache 值减 1
			cache[s[left]]--
			left++
		}
		if right-left > ans {
			ans = right - left
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
func lengthOfLongestSubstring1(s string) int {
	n := len(s)
	cache := map[byte]int{}
	right := 0
	ans := 0
	for left := 0; left < len(s); left++ {
		if left > 0 {
			delete(cache, s[left-1])
		}
		for right < n && cache[s[right]] == 0 {
			cache[s[right]]++
			right++
		}
		if right-left > ans {
			ans = right - left
		}
	}
	return ans
}

/*
解题思路：滑动窗口
左边窗口一直遍历 +1
右边窗口一直遍历，若遇到重复字符串则结束遍历，当前左右窗口长度就是符合的条件。
理解为：维持一个队列队列里的值不能有重复的，遇到重复的就将左边的值弹出。
*/

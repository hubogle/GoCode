// 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
//
//
//
// 示例 1:
//
//
// 输入: s = "abab"
// 输出: true
// 解释: 可由子串 "ab" 重复两次构成。
//
//
// 示例 2:
//
//
// 输入: s = "aba"
// 输出: false
//
//
// 示例 3:
//
//
// 输入: s = "abcabcabcabc"
// 输出: true
// 解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
//
//
//
//
// 提示：
//
//
//
//
// 1 <= s.length <= 10⁴
// s 由小写英文字母组成
//
// Related Topics 字符串 字符串匹配 👍 620 👎 0

package main

import "fmt"

func main() {
	s := "abcabcabcabc"
	result := repeatedSubstringPattern(s)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	next := make([]int, n)
	j := 0
	next[0] = j
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)

/*
next[n-1] 最长相同前后缀的长度

一个循环周期：数组长度减去最长相同前后缀的长度相当于是第一个周期的长度

如果这个周期可以被整除，就说明整个数组就是这个周期的循环。
*/

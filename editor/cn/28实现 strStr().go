// 实现 strStr() 函数。
//
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如
// 果不存在，则返回 -1 。
//
//
//
// 说明：
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。
//
//
//
// 示例 1：
//
//
// 输入：haystack = "hello", needle = "ll"
// 输出：2
//
//
// 示例 2：
//
//
// 输入：haystack = "aaaaa", needle = "bba"
// 输出：-1
//
//
// 示例 3：
//
//
// 输入：haystack = "", needle = ""
// 输出：0
//
//
//
//
// 提示：
//
//
// 0 <= haystack.length, needle.length <= 5 * 10⁴
// haystack 和 needle 仅由小写英文字符组成
//
// Related Topics 双指针 字符串 字符串匹配 👍 1267 👎 0

package main

import "fmt"

func main() {
	haystack := "mississippi"
	needle := "issip"
	result := strStr(haystack, needle)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	next := make([]int, n)
	nextArray(next, needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 前 k 相同位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

func nextArray(next []int, s string) {
	j := 0
	next[0] = 0
	for i := 1; i < len(next); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		} // 一直寻找相同的位置
		if s[i] == s[j] {
			j++
		} // 如果相同的话
		next[i] = j
	}
}

// leetcode submit region end(Prohibit modification and deletion)

/*
问题思路：KMP 方法
前缀表存储的是相同前后缀的长度。
前缀表是用来回退的，它记录了模式串与主串(文本串)不匹配的时候，模式串应该从哪里开始重新匹配。
时间复杂度：O(m+n) 匹配字符串长度 m 前缀数组 n
空间复杂度：O(n)
*/

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//
//
//
// 示例 1:
//
//
// 输入: s = "anagram", t = "nagaram"
// 输出: true
//
//
// 示例 2:
//
//
// 输入: s = "rat", t = "car"
// 输出: false
//
//
//
// 提示:
//
//
// 1 <= s.length, t.length <= 5 * 10⁴
// s 和 t 仅包含小写字母
//
//
//
//
// 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// Related Topics 哈希表 字符串 排序 👍 505 👎 0

package main

import (
	"fmt"
)

func main() {
	s := "anagram"
	t := "nagaram"
	result := isAnagram(s, t)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	nums := make([]int, 26, 26)
	for _, v := range s {
		nums[v-'a']++
	}
	for _, v := range t {
		nums[v-'a']--
	}
	for _, v := range nums {
		if v != 0 {
			return false
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func isAnagram1(s string, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c1[ch-'a']++
	}
	return c1 == c2
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} // 首先确保字符串长度一定一样
	cnt := make([]rune, 26, 26) // 若字符是 Unicode 类型的情况
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 { // 前提字符串长度一致
			return false
		}
	}
	return true
}

/*
题目：判断两个字符串包含的字符是否相同
思路：通过哈希确定重复值，因为字符种类共 26 个，所以可以通过申请 26 个 int 大小的数组。
如果种类过多的话，需要用 `map`，这种情况种类少并且字符重复率高占用不了太多空间。
空间复杂度：O(S) S 为 26 个空间
时间复杂度：O(n) 需要遍历两次字符串 + 一次数组
*/

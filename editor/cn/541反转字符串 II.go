// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
//
//
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
//
//
//
//
// 示例 1：
//
//
// 输入：s = "abcdefg", k = 2
// 输出："bacdfeg"
//
//
// 示例 2：
//
//
// 输入：s = "abcd", k = 2
// 输出："bacd"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由小写英文组成
// 1 <= k <= 10⁴
//
// Related Topics 双指针 字符串 👍 253 👎 0

package main

import "fmt"

func main() {
	s := "abcdefg"
	result := reverseStr(s, 8)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	chars := []byte(s)
	for i := 0; i < len(chars); i += 2 * k { // 每次区间增加 2k 长度，再区分处理
		if i+k < len(chars) { // 剩余字符长度 >= k
			reverse(chars[i : i+k])
		} else { // 剩余字符长度 < k
			reverse(chars[i:])
		}
	}
	return string(chars)
}
func reverse(b []byte) {
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：反转字符串，每间隔 2k 个字符，子区间字符串长度 >= k 则反转 k 个字符串，子区间字符串 < k 反转全部字符串
时间复杂度：O(n)
空间复杂度：O(1)

解体思路：需要固定规律一段一段去处理字符串的时候，在 for 循环内部进行区别处理每段字符串。
*/

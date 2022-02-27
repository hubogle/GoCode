// 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数
// 将返回左旋转两位得到的结果"cdefgab"。
//
//
//
// 示例 1：
//
// 输入: s = "abcdefg", k = 2
// 输出: "cdefgab"
//
//
// 示例 2：
//
// 输入: s = "lrloseumgh", k = 6
// 输出: "umghlrlose"
//
//
//
//
// 限制：
//
//
// 1 <= k < s.length <= 10000
//
// Related Topics 数学 双指针 字符串 👍 197 👎 0

package main

import "fmt"

func main() {
	s := "abcdefg"
	result := reverseLeftWords(s, 2)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func reverseLeftWords(s string, n int) string {
	chars := []byte(s)
	l, r := 0, n-1
	for l < r {
		chars[l], chars[r] = chars[r], chars[l]
		l++
		r--
	}
	l, r = n, len(chars)-1
	for l < r {
		chars[l], chars[r] = chars[r], chars[l]
		l++
		r--
	}
	l, r = 0, len(chars)-1
	for l < r {
		chars[l], chars[r] = chars[r], chars[l]
		l++
		r--
	}
	return string(chars)
}

// leetcode submit region end(Prohibit modification and deletion)

/*
字符串反转前 k 位字符串，abcdefg 变为 cdefqab
整体反转的前一步：baqfedc，abcdefg 部分反转字符串得到的结果
解体思路：局部反转+整体反转

*/

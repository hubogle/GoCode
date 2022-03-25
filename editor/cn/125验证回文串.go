// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
//
//
// 示例 1:
//
//
// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 解释："amanaplanacanalpanama" 是回文串
//
//
// 示例 2:
//
//
// 输入: "race a car"
// 输出: false
// 解释："raceacar" 不是回文串
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 10⁵
// 字符串 s 由 ASCII 字符组成
//
// Related Topics 双指针 字符串 👍 497 👎 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	result := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	str := ""
	for i := 0; i < len(s); i++ {
		if isStr(s[i]) {
			str += string(s[i])
		}
	}
	str = strings.ToLower(str)
	l := 0
	r := len(str) - 1
	for l < r {
		if str[l] == str[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isStr(s byte) bool {
	if (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9') || (s >= 'A' && s <= 'Z') {
		return true
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)

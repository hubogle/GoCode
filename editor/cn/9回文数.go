// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//
// 例如，121 是回文，而 123 不是。
//
//
//
//
// 示例 1：
//
//
// 输入：x = 121
// 输出：true
//
//
// 示例 2：
//
//
// 输入：x = -121
// 输出：false
// 解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
//
// 示例 3：
//
//
// 输入：x = 10
// 输出：false
// 解释：从右向左读, 为 01 。因此它不是一个回文数。
//
//
//
//
// 提示：
//
//
// -2³¹ <= x <= 2³¹ - 1
//
//
//
//
// 进阶：你能不将整数转为字符串来解决这个问题吗？
// Related Topics 数学 👍 1894 👎 0

package main

import "fmt"

func main() {
	result := isPalindrome(121)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	y := 0
	for ; x > y; x /= 10 {
		v := x % 10
		y = y*10 + v
	}
	return x == y || x == y/10
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：
负数返回 false
正数如果是偶数位则判断中间切分后两边是否相等。
正数如果是奇数位则排除中间位判断两边是否相等。
特殊情况：如果数字的最后一位是 0，为了使该数字为回文，只能为0
*/

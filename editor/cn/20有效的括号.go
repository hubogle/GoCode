// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
//
//
// 示例 1：
//
//
// 输入：s = "()"
// 输出：true
//
//
// 示例 2：
//
//
// 输入：s = "()[]{}"
// 输出：true
//
//
// 示例 3：
//
//
// 输入：s = "(]"
// 输出：false
//
//
// 示例 4：
//
//
// 输入：s = "([)]"
// 输出：false
//
//
// 示例 5：
//
//
// 输入：s = "{[]}"
// 输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
// Related Topics 栈 字符串 👍 3028 👎 0

package main

import "fmt"

func main() {
	s := "{}}"
	result := isValid1(s)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	var stack []rune
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else if len(stack) == 0 {
			return false
		} else {
			if v == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else if v == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else if v == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// leetcode submit region end(Prohibit modification and deletion)

func isValid1(s string) bool {
	n := len(s)
	if n&0x01 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if v, ok := pairs[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

/*
解题思路：利用栈的特性。
遍历字符串时遇到左括号则进栈，遇到右括号时，取出栈顶的左括号判断是否能闭合，
若能闭合则取出，若不能闭合则返回 false
当便利完成后若栈中还有数据则返回 false
时间复杂度：O(n)
空间复杂度：O(n) + O(m) 哈希表存储额外空间本题中 m = 6
*/

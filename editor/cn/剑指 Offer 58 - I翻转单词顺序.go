// 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，
// 则输出"student. a am I"。
//
//
//
// 示例 1：
//
// 输入: "the sky is blue"
// 输出: "blue is sky the"
//
//
// 示例 2：
//
// 输入: "  hello world!  "
// 输出: "world! hello"
// 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//
//
// 示例 3：
//
// 输入: "a good   example"
// 输出: "example good a"
// 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//
//
//
//
// 说明：
//
//
// 无空格字符构成一个单词。
// 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//
//
// 注意：本题与主站 151 题相同：https://leetcode-cn.com/problems/reverse-words-in-a-string/
//
//
// 注意：此题对比原题有改动
// Related Topics 双指针 字符串 👍 176 👎 0

package main

import "fmt"

func main() {
	s := "hello world"
	result := reverseWords(s)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	chars := []byte(s)
	l, r := 0, len(chars)-1
	for l < len(chars) && chars[l] == ' ' {
		l++
	}
	for r >= 0 && chars[r] == ' ' {
		r--
	}
	if r < l {
		return ""
	}
	chars = chars[l : r+1]
	l, r = 0, 0
	for ; r < len(chars); l, r = l+1, r+1 {
		chars[l] = chars[r]
		for r < len(chars) && chars[r] == ' ' && chars[r+1] == ' ' {
			r++
		}
	}
	chars = chars[:l]
	reverse(&chars, 0, len(chars)-1)
	n := len(chars) - 1
	for l, r = 0, 0; r <= n; {
		for r < n && chars[r] != ' ' && chars[r+1] != ' ' {
			r++
		}
		reverse(&chars, l, r)
		l = r + 1
		r++
	}
	return string(chars)
}

func reverse(chars *[]byte, l, r int) {
	for l < r {
		(*chars)[l], (*chars)[r] = (*chars)[r], (*chars)[l]
		l++
		r--
	}
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：
注意点：
1. 考虑空字符串，以及输入字符的临界值
2. 清空周围空格，双指针将多个空格合并为一个空格
3. 将字符串整体反转，将单词部分反转

时间复杂度：O(n)
空间复杂度：O(1)
如果借用临时数组，那么空间复杂度为 O(n)
*/

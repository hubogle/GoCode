// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//
//
// 示例 1：
//
// 输入：s = "We are happy."
// 输出："We%20are%20happy."
//
//
//
// 限制：
//
// 0 <= s 的长度 <= 10000
// Related Topics 字符串 👍 216 👎 0

package main

import (
	"fmt"
)

func main() {
	s := "We are happy."
	result := replaceSpace(s)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
	num := 0
	for _, v := range s {
		// if strconv.Itoa(v) == " "
		if v == 32 {
			num++
		}
	}
	n := len(s) + num*2
	newByte := make([]byte, n, n)
	for i, j := len(s)-1, n-1; i >= 0 && j >= 0; i-- {
		if s[i] != 32 {
			newByte[j] = s[i]
			j--
		} else {
			newByte[j] = '0'
			newByte[j-1] = '2'
			newByte[j-2] = '%'
			j -= 3
		}
	}
	return string(newByte)
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：申请替换空格后的字符串长度的容量，从后往前替换掉原字符串里面的空格
空间复杂度：额外申请的字符串容量
时间复杂度：O(n) 原字符串长度
*/

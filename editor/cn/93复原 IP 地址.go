// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
// 和 "192.168@1.1" 是 无效 IP 地址。
//
//
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新
// 排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
//
//
// 示例 1：
//
//
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
//
//
// 示例 2：
//
//
// 输入：s = "0000"
// 输出：["0.0.0.0"]
//
//
// 示例 3：
//
//
// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 20
// s 仅由数字组成
//
// Related Topics 字符串 回溯 👍 825 👎 0

package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := restoreIpAddresses("25525511135")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	var ans []string
	var val []string
	var n = len(s)
	var backtracking func(index int)
	backtracking = func(index int) {
		if index == n && len(val) == 4 {
			res := ""
			for _, v := range val {
				res += v
				res += "."
			}
			ans = append(ans, res[:len(res)-1])
			return
		}
		for i := index; i < n; i++ {
			if i > index && s[index] == '0' {
				return
			}
			v, _ := strconv.Atoi(s[index : i+1])
			if v >= 0 && v <= 255 && len(val) < 4 {
				val = append(val, s[index:i+1])
				backtracking(i + 1)
				val = val[:len(val)-1]
			}
		}
	}
	backtracking(0)
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：需要判断切分字符串不符合条件的情况
1. 数字不能 "0" 开头，可以为 0
2. 数字在 [0,255]
3.
*/

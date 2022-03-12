// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。
//
//
//
// 示例 1：
//
//
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
//
//
// 示例 2：
//
//
// 输入：s = "a"
// 输出：[["a"]]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成
//
// Related Topics 字符串 动态规划 回溯 👍 1028 👎 0

package main

import "fmt"

func main() {
	result := partition("aab")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	var ans [][]string
	var str []string
	var backtracking func(index int)
	n := len(s)
	backtracking = func(index int) {
		if index == n { // 只要遍历到 n
			ans = append(ans, append([]string{}, str...))
			return
		}
		for i := index; i < n; i++ {
			if isPar(s, index, i) {
				str = append(str, s[index:i+1])
				backtracking(i + 1)
				str = str[:len(str)-1]
			}
		}
	}
	backtracking(0)
	return ans
}
func isPar(str string, l, r int) bool {
	for ; l < r; l, r = l+1, r-1 {
		if str[l] != str[r] {
			return false
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：当遇到切分字符串时，判断字符串是否为回文串。
终止条件，只要纵深度遍历完就终止。
`for` 循环中只要切分的串不符合回就不加入到数据中，不继续进行深度遍历。
*/

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
//
//
// 示例 2：
//
//
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
// Related Topics 字符串 👍 2125 👎 0

package main

import "fmt"

func main() {
	result := longestCommonPrefix()
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// leetcode submit region end(Prohibit modification and deletion)
/*
解题思路：纵向遍历每个字符串的第 i 位
以第一个字符串为基础，依次遍历后面每个字符串的第 i 个字符
*/

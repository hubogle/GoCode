// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
// 输入：digits = ""
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：digits = "2"
// 输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
// Related Topics 哈希表 字符串 回溯 👍 1766 👎 0

package main

import (
	"fmt"
)

func main() {
	result := letterCombinations("23")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var letterMap = []string{"", "",
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	} // 数组存储数字对应的数值
	var ans []string
	var val string
	var backtracking func(index int)
	n := len(digits)
	backtracking = func(index int) {
		if index == n {
			ans = append(ans, val)
			return
		}
		num := digits[index] - '0' // 字符串转换为 int
		for i := 0; i < len(letterMap[num]); i++ {
			val += string(letterMap[num][i])
			backtracking(index + 1)
			val = val[:len(val)-1]
		}
	}
	backtracking(0)
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
注意：数字字符串变化为 int 的操作
二维数组存储数字对应索引位置的字符串
每次循环遍历当前位置能出现的所有字符情况
*/

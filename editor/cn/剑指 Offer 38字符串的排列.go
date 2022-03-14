// 输入一个字符串，打印出该字符串中字符的所有排列。
//
//
//
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
//
//
//
// 示例:
//
// 输入：s = "abc"
// 输出：["abc","acb","bac","bca","cab","cba"]
//
//
//
//
// 限制：
//
// 1 <= s 的长度 <= 8
// Related Topics 字符串 回溯 👍 515 👎 0

package main

import "fmt"

func main() {
	result := permutation("aabc")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func permutation(s string) []string {
	var ans []string
	var val string
	n := len(s)
	use := make([]bool, n) // 递归使用
	var backtracking func(index int)
	backtracking = func(index int) {
		if index == n {
			ans = append(ans, val)
			return
		}
		used := make(map[byte]int, 0) // 遍历使用，这种方式是针对不能排列的，若可以排列可用 for 循环
		for i := 0; i < n; i++ {
			_, ok := used[s[i]]
			if ok || use[i] { // 已经使用过
				continue
			}
			used[s[i]]++
			use[i] = true
			val += string(s[i])
			backtracking(index + 1)
			use[i] = false
			val = val[:len(val)-1]
		}
	}
	backtracking(0)
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
/*
题目：不能包含重复串
方式一：进行排序，通过 for i>0 && nums[i] == num[i-1] && used[i] == false{}
方式二：用一个 map 在 for 循环时，标记每个值是否使用。
最外层要用一个 used 标记整个深度下不能使用同一个位置的值，但是可以使用相同的值
*/

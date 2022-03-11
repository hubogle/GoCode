// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
//
//
//
// 示例 1：
//
//
// 输入：n = 4, k = 2
// 输出：
// [
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
// ]
//
// 示例 2：
//
//
// 输入：n = 1, k = 1
// 输出：[[1]]
//
//
//
// 提示：
//
//
// 1 <= n <= 20
// 1 <= k <= n
//
// Related Topics 数组 回溯 👍 892 👎 0

package main

import "fmt"

func main() {
	result := combine(4, 2)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	var path []int
	var ans [][]int
	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		if len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := startIndex; i <= n-(k-len(path))+1; i++ { // 后续数量不满足，注意减枝操作
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：数组的组合情况
递归终止条件：临时数组里数量为两个值
递归 `for` 循环：遍历集合所有的情况，递归条件改变，进行下次递归。
递归终止条件优化：当后续数量已经不满足时可以进行减枝操作
*/

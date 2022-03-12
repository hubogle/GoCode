// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。
//
// 注意：解集不能包含重复的组合。
//
//
//
// 示例 1:
//
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
//
// 示例 2:
//
//
// 输入: candidates = [2,5,2,1,2], target = 5,
// 输出:
// [
// [1,2,2],
// [5]
// ]
//
//
//
// 提示:
//
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
// Related Topics 数组 回溯 👍 871 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	result := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int
	var sum int
	n := len(candidates)
	sort.Ints(candidates)
	used := make([]bool, len(candidates))
	var backtracking func(index int)
	backtracking = func(index int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
		}
		for i := index; i < n && sum < target; i++ {
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false { // 这里标记 false
				// if i > index && candidates[i] == candidates[i-1] { // 只有第一个树枝可用，后面的树枝不能用。不使用 used 的写法
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			used[i] = true // 同一个树标记已经用过了（可以用）
			backtracking(i + 1)
			used[i] = false // 同一层树标记用过了
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目解析：要理解树层去重，还是数枝去重
树层重复：can[i] == can[i-1] & used[i-1] == false 同一层相同数据已用过，前一个树枝已用过，不能重复选取
树枝重构：can[i] == can[i-1] & used[i-1] == true 同一个树已用过，可以重复选取
谨记：for 循环是树层遍历，递归是树深遍历
*/

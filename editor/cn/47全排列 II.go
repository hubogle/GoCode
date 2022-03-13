// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// Related Topics 数组 回溯 👍 970 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	result := permuteUnique([]int{1, 1, 2})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var val []int
	var backtracking func(index int)
	sort.Ints(nums)
	n := len(nums)
	used := make([]bool, n)
	backtracking = func(index int) {
		if index == n {
			ans = append(ans, append([]int{}, val...))
			return
		}
		for i := 0; i < n; i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false { // 同一层跳过相同值
				continue
			}
			if used[i] == false { // 一棵树，只有没用过的才可以添加
				val = append(val, nums[i])
				used[i] = true
				backtracking(index + 1)
				used[i] = false
				val = val[:len(val)-1]
			}
		}
	}
	backtracking(0)
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：同一层不能有重复的值

去重复一定要对元素进行排列，这样才能通过相邻元素判断是否使用过。
层去重效率高于，树去重复，同一层去重效率高，
*/

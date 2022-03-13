// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
// 输入：nums = [1]
// 输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
// Related Topics 数组 回溯 👍 1834 👎 0

package main

import "fmt"

func main() {
	result := permute([]int{1, 2, 3})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	var ans [][]int
	var val []int
	n := len(nums)
	used := make([]bool, len(nums))
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(val) == n {
			ans = append(ans, append([]int{}, val...))
			return
		}
		for i := 0; i < n; i++ {
			if used[i] == true {
				continue
			}
			val = append(val, nums[i])
			used[i] = true
			backtracking(index + 1)
			val = val[:len(val)-1]
			used[i] = false
		}
	}
	backtracking(0)
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：同一棵树不能用用过的值，也就是 for 循环每次都从 0 开始，要跳过已经使用过的值；
排列和组合问题最大的不同是，排列每层都是从 0 开始搜索，需要用 used 记录使用的元素
组合是从 index 开始，不能有本身。
*/

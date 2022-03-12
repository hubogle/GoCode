// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
//
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//
//
// 示例 2：
//
//
// 输入：nums = [0]
// 输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
//
//
//
// Related Topics 位运算 数组 回溯 👍 759 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	result := subsetsWithDup([]int{1, 2, 3})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	var val []int
	n := len(nums)
	sort.Ints(nums)
	var backtracking func(index int)
	backtracking = func(index int) {
		ans = append(ans, append([]int{}, val...))
		for i := index; i < n; i++ {
			if i > index && nums[i] == nums[i-1] { // 从大于起始索引位置开始
				continue
			}
			val = append(val, nums[i])
			backtracking(i + 1)
			val = val[:len(val)-1]
		}
	}
	backtracking(0)
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func subsetsWithDup1(nums []int) [][]int {
	var ans [][]int
	var val []int
	n := len(nums)
	used := make([]bool, n)
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(val) <= index {
			ans = append(ans, append([]int{}, val...))
		}
		for i := index; i < n; i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}
			val = append(val, nums[i])
			used[i] = true
			backtracking(i + 1)
			used[i] = false
			val = val[:len(val)-1]
		}
	}
	backtracking(0)
	return ans
}

/*
题目解析：
遇到求不能有相同重复值的子集，优先考虑是否要排序。
1. 有两种写法，第一种借助额外数组，进行标识。
2. 第二种方法：判断本身与前一个值相同，则跳过。
*/

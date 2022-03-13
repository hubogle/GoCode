// 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
//
// 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
//
//
//
// 示例 1：
//
//
// 输入：nums = [4,6,7,7]
// 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
//
//
// 示例 2：
//
//
// 输入：nums = [4,4,3,2,1]
// 输出：[[4,4]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 15
// -100 <= nums[i] <= 100
//
// Related Topics 位运算 数组 哈希表 回溯 👍 402 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	result := findSubsequences([]int{4, 6, 7, 7})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findSubsequences(nums []int) [][]int {
	var ans [][]int
	var val []int
	n := len(nums)
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(val) > 1 {
			ans = append(ans, append([]int{}, val...))
		}
		used := make([]bool, 210)
		for i := index; i < n; i++ {
			if index > 0 && nums[index-1] > nums[i] || used[nums[i]+100] == true { // 这里必须两个条件判断
				continue
			} // 当前数组最后一个
			used[nums[i]+100] = true
			val = append(val, nums[i])
			backtracking(i + 1)
			val = val[:len(val)-1]
		}
	}
	backtracking(0)
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
// 不能排序直接求解
func findSubsequences1(nums []int) [][]int {
	var ans [][]int
	var val []int
	sort.Ints(nums)
	n := len(nums)
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(val) <= index && len(val) >= 2 {
			ans = append(ans, append([]int{}, val...))
		}
		for i := index; i < n; i++ {
			if i > index && nums[i] == nums[i-1] {
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

/*
解题思路：不能对数组进行排序，这样得到的结果不对，得到的都是有序 6,5,4 这种没有递增子序列的
40组合总和II
两个 used 标记的区别在于，该题 used 用于当前层判断是否有值相同。
40 题 的 used 用于整棵树，每层进行判断是否有重复值。
40 题是排序之后的，有序了，使用 used 标记整棵树，也要先排序
491 题是要求递增的，没法直接排序。
*/

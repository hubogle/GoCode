// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
// 序列。
//
//
// 示例 1：
//
//
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
//
// 示例 2：
//
//
// 输入：nums = [0,1,0,3,2,3]
// 输出：4
//
//
// 示例 3：
//
//
// 输入：nums = [7,7,7,7,7,7,7]
// 输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
//
// Related Topics 数组 二分查找 动态规划 👍 2479 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	result := lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// leetcode submit region end(Prohibit modification and deletion)
func lengthOfLIS1(nums []int) int {
	dp := make([]int, len(nums)+1)
	result := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		result = Max(result, dp[i])
	}
	return result
}

/*
最长递增子序列：
动态规划 DP，时间复杂度为 O(n^2)
初始化 DP 数组默认值为 1，状态转移方程为 dp[i] = max( 1 + dp[j]) ，其中 j < i && nums[j] > nums[i]

二分查找方法：时间复杂度 O(nlogn)
记录上升子序列的最末尾的一个数字，维护一个 dp 数组，值代表递增子序列的最大值。
*/

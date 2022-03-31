// 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
//
// 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那
// 么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,3,5,4,7]
// 输出：3
// 解释：最长连续递增序列是 [1,3,5], 长度为3。
// 尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
//
//
// 示例 2：
//
//
// 输入：nums = [2,2,2,2,2]
// 输出：1
// 解释：最长连续递增序列是 [2], 长度为1。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
// Related Topics 数组 👍 278 👎 0

package main

import "fmt"

func main() {
	result := findLengthOfLCIS([]int{1, 2, 3})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findLengthOfLCIS(nums []int) int {
	if len(nums) < 0 {
		return 0
	}
	ans, count := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count > ans {
			ans = count
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
func findLengthOfLCIS1(nums []int) int {
	if len(nums) < 2 {
		return 1
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

/*
动态规划：O(n)
申请 dp 数组，当当前值大于前一个值时 dp[i] = dp[i-1] + 1
贪心算法：O(n)
遇到 nums[i] > nums[i-1] 情况 count ++ 其它情况 count = 1
*/

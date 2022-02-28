// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。
//
//
//
// 示例 1：
//
//
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//
//
// 示例 2：
//
//
// 输入：nums = [1]
// 输出：1
//
//
// 示例 3：
//
//
// 输入：nums = [5,4,-1,7,8]
// 输出：23
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
// Related Topics 数组 分治 动态规划 👍 4424 👎 0

package main

import "fmt"

func main() {
	nums := []int{5, 4, -1, 7, 8}
	result := maxSubArray(nums)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	// ans := nums[0] // 默认最大为第一个
	// for i := 1; i < len(nums); i++ {
	// 	nums[i] = max(nums[i], nums[i]+nums[i-1])
	// 	ans = max(ans, nums[i])
	// }
	// return ans
	pre, ans := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		pre = max(nums[i], pre+nums[i])
		ans = max(pre, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func maxSubArray1(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := nums[0] // 默认最大为第一个
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 { // 如果前面的值大于 0
			dp[i] = dp[i-1] + nums[i]
		} else { // 前面的值小于 0 就不用累加
			dp[i] = nums[i]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
func maxSubArray2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := nums[0] // 默认最大为第一个
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i]
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
func maxSubArray3(nums []int) int {
	ans := nums[0] // 默认最大为第一个
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}

/*
题目思路：dp 数组，存储每个索引位置 i 当前的最大值。
情况1：dp[i-1] > 0 则 dp[i] = dp[i-1] + nums[i]
情况2：dp[i-1] < 0 则 dp[i] = nums[i]
分治法就是分别获取左右最大值，然后两边最大值进行相加

*/

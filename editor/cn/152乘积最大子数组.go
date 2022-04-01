// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
// 测试用例的答案是一个 32-位 整数。
//
// 子数组 是数组的连续子序列。
//
//
//
// 示例 1:
//
//
// 输入: nums = [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
//
//
// 示例 2:
//
//
// 输入: nums = [-2,0,-1]
// 输出: 0
// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 2 * 10⁴
// -10 <= nums[i] <= 10
// nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
//
// Related Topics 数组 动态规划 👍 1645 👎 0

package main

import "fmt"

func main() {
	result := maxProduct([]int{-4, -3, -2})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxAns, minAns, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			minAns, maxAns = maxAns, minAns // 交换最小和最大值
		}
		maxAns = Max(nums[i], maxAns*nums[i])
		minAns = Min(nums[i], minAns*nums[i])
		ans = Max(ans, maxAns)
	}
	return ans
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// leetcode submit region end(Prohibit modification and deletion)

/*
动态规划，记录上次一的最大值和最小值。
当本次值小于 0 时，最大值和最小值进行交换。
状态转移方程是：最大值是 Max(f(n)) = Max( Max(f(n-1)) * n, Min(f(n-1)) * n)；最小值是 Min(f(n)) = Min( Max(f(n-1)) * n, Min(f(n-1)) * n)
*/

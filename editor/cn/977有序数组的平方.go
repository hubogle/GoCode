// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//
//
//
//
//
// 示例 1：
//
//
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]
// 排序后，数组变为 [0,1,9,16,100]
//
// 示例 2：
//
//
// 输入：nums = [-7,-3,2,3,11]
// 输出：[4,9,9,49,121]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 已按 非递减顺序 排序
//
//
//
//
// 进阶：
//
//
// 请你设计时间复杂度为 O(n) 的算法解决本问题
//
// Related Topics 数组 双指针 排序 👍 439 👎 0

package main

import "fmt"

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	result := sortedSquares(nums)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums), len(nums))
	for l, r, n := 0, len(nums)-1, len(nums)-1; l <= r && n >= 0; n-- {
		if -nums[l] < nums[r] { // 因为是升序排列，当左边为负取反小于的话证明平方也是小于
			ans[n] = nums[r] * nums[r]
			r--
		} else {
			ans[n] = nums[l] * nums[l]
			l++
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：将原升序数组，平方后还按升序排列
时间复杂度：O(n)
空间复杂度：O(1)
思路从左右两指针，当左边取反后大于右边值，则左指针移动，否则右指针移动。
*/

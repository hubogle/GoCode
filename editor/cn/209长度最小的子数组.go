// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长
// 度。如果不存在符合条件的子数组，返回 0 。
//
//
//
// 示例 1：
//
//
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//
//
// 示例 2：
//
//
// 输入：target = 4, nums = [1,4,4]
// 输出：1
//
//
// 示例 3：
//
//
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0
//
//
//
//
// 提示：
//
//
// 1 <= target <= 10⁹
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
//
//
//
//
// 进阶：
//
//
// 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
//
// Related Topics 数组 二分查找 前缀和 滑动窗口 👍 949 👎 0

package main

import (
	"math"
)

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
// 时间组复杂度为 n^2, 空间复杂度 O(1)
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans, sum := math.MaxInt32, 0
	for i := 0; i < n; i++ {
		sum = 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= target {
				ans = mix(ans, j-i+1)
				break // 后面符合条件都不是最小值
			}
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 时间复杂度 n * log(n) 前缀和 + 二分查找
func minSubArrayLen1(target int, nums []int) int {
	n := len(nums)
	ans := math.MaxInt32
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		j := binarySearch(sums, target+sums[i-1])
		if j != -1 {
			ans = mix(ans, j-i+1)
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// 二分查找, 大于目标的最小值
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r { // 左闭右开
		mid := l + (r-l)>>2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[l] >= target {
		return l
	} else {
		return -1
	}
}

// 滑动窗口
func minSubArrayLen2(target int, nums []int) int {
	ans := math.MaxInt32
	l, r, sum := 0, 0, 0
	for ; r < len(nums); r++ { // 右指针一直滑动
		sum += nums[r]
		for sum >= target { // 当满足的时候，再移动左指针
			ans = mix(ans, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// 最小值
func mix(x, y int) int {
	if x < y {
		return x
	}
	return y
}

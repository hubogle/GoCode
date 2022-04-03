// 统计一个数字在排序数组中出现的次数。
//
//
//
// 示例 1:
//
//
// 输入: nums = [5,7,7,8,8,10], target = 8
// 输出: 2
//
// 示例 2:
//
//
// 输入: nums = [5,7,7,8,8,10], target = 6
// 输出: 0
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// nums 是一个非递减数组
// -10⁹ <= target <= 10⁹
//
//
//
//
// 注意：本题与主站 34 题相同（仅返回值不同）：https://leetcode-cn.com/problems/find-first-and-last-
// position-of-element-in-sorted-array/
// Related Topics 数组 二分查找 👍 310 👎 0

package main

import "fmt"

func main() {
	result := search([]int{}, 1)
	// result := search([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	var l, r int
	l, r = 0, len(nums)-1
	for l <= r {
		mid := l + ((r - l) >> 2)
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	right := l
	if r >= 0 && nums[r] != target { // 数组中无 target ，则提前返回
		return 0
	}
	l, r = 0, len(nums)-1
	for l <= r {
		mid := l + ((r - l) >> 2)
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1 // 跳出时指向左
		}
	}
	left := r
	return right - left - 1
}

// leetcode submit region end(Prohibit modification and deletion)

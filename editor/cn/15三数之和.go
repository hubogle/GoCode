// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
// 复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
// 输入：nums = []
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：nums = [0]
// 输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// Related Topics 数组 双指针 排序 👍 4343 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	var ans [][]int
	for i, _ := range nums {
		if nums[i] > 0 {
			return ans // 排序后如果第一个大于 0 则总和不可能等于 0
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} // 重复元素跳过，防止有重复解答
		j, k := i+1, n-1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] { // 相等时进行左右移动
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
/*
时间复杂度：O(n^2) 数组排序 n*log(n) + 遍历数组 O(n) * 双指针遍历 O(n)
空间复杂度：log(n) 排序占用

解题思路：排序 + 双指针

1. 长度小于 3 直接返回 []
2. 遍历排序数组
	2.1 若 nums[i] > 0 排序后的第一个值不能大于 0
	2.2 对于重复元素进行跳过，防止重复解出现
	2.3 左、右指针分别移动，寻找判断符合条件值，移动中若后一位和当前值相同则跳过。
	2.3.1 和大于 0 则 num[r] 太大，左移
	2.3.2 和小于 0 则 num[l] 太小，右移
*/

// 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
// b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
//
//
// 你可以按 任意顺序 返回答案 。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,0,-1,0,-2,2], target = 0
// 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// 示例 2：
//
//
// 输入：nums = [2,2,2,2,2], target = 8
// 输出：[[2,2,2,2]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
//
// Related Topics 数组 双指针 排序 👍 1130 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	// nums := []int{-2, -1, -1, 1, 1, 2, 2}
	// nums := []int{2, 2, 2, 2, 2}
	target := 0
	result := fourSum(nums, target)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}
	sort.Ints(nums)
	var ans [][]int
	for i, _ := range nums {
		// if nums[i] > target {
		// 	continue
		// } // 不能减树因为 target 是任意值，存在负数的情况

		// if i < n-1 && nums[i] == nums[i+1] {
		// 	continue
		// } // 不能这样往后移动，不然 -1, -1, 1, 1 这种会跳过

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			} // 注意这里，不是 j > 1 的判断。
			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}

		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：四个数之和是目标 target
与三数之和为 0 不同的是，target 可以是任意数。

解题思路：排序后，两层 for 循环，最后两层用双指针减少循环。
时间复杂度：O(n^3)
空间复杂度：O(log(n)) 快排占用

注意 target 是任意值，不需要进行减支操作
注意 前两层循环时，不能提前判断后面的值是否相等，只有之前值相等时才跳过。不然会跳过 -1,-1,1,1 的答案。
*/

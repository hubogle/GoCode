// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
//
//
// 示例 1：
//
//
// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
// 示例 2：
//
//
// 输入：nums = [0,0,0], target = 1
// 输出：0
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
// Related Topics 数组 双指针 排序 👍 1138 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	result := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := math.MaxInt
	res := nums[0]
	for k, _ := range nums {
		l, r := k+1, len(nums)-1
		for l < r {
			val := nums[k] + nums[l] + nums[r]
			abs := math.Abs(float64(val - target)) // 相差最小值
			if int(abs) < ans {
				ans = int(abs)
				res = val
			}
			if val >= target {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)

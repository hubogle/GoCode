// 给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
// 假设你总是可以到达数组的最后一个位置。
//
//
//
// 示例 1:
//
//
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
//
// 示例 2:
//
//
// 输入: nums = [2,3,0,1,4]
// 输出: 2
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// 0 <= nums[i] <= 1000
//
// Related Topics 贪心 数组 动态规划 👍 1460 👎 0

package main

import "fmt"

func main() {
	result := jump([]int{2, 3, 1, 1, 4})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	ans := 0
	cur := 0  // 当前跳跃位置
	next := 0 // 下一个跳跃位置
	n := len(nums)
	for i := 0; i < n-1; i++ { // 下标范围在 [0, n-2]
		if nums[i]+i > next {
			next = nums[i] + i
		}
		if i == cur {
			cur = next
			ans++
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
func jump1(nums []int) int {
	ans := 0
	cur := 0  // 当前跳跃位置
	next := 0 // 下一个跳跃位置
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i]+i > next {
			next = nums[i] + i
		} // 获取下一步能最大跳跃的位置
		if i == cur { // 走到当前最大的距离
			if cur != n-1 { // 未走到终点
				ans++
				cur = next
				if next >= n-1 {
					break
				}
			} else {
				break
			}
		}
	}
	return ans
}

/*
贪心思想：
每次并不一定是走最大的步，当走到最大步时未到达终点，走遍历过程中的最大值。
1. 在遍历走步的时，也要保存遍历过程中，下一步能达到的最大距离；
2. 当走到最大步时，若是终点则返回，不是终点则取下一步并且步数加1；
局部最优：当前可移动的情况下，尽可能多走，没到达终点则步数 + 1；
全局最优：最小的步数到达终点；

统一处理方法：下标遇到当前覆盖范围的最大下标，则步数 +1
*/

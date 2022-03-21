// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1：
//
//
// 输入：[3,2,3]
// 输出：3
//
// 示例 2：
//
//
// 输入：[2,2,1,1,1,2,2]
// 输出：2
//
//
//
//
// 进阶：
//
//
// 尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
//
// Related Topics 数组 哈希表 分治 计数 排序 👍 1366 👎 0

package main

import "fmt"

func main() {
	result := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	ans := 0
	level := 0
	for _, v := range nums {
		if level == 0 {
			ans = v
		}
		if ans == v {
			level++
		} else {
			level--
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：投票选举
当票数为0时，设当前值为众数，且投票数为 1
若下一个与当前临时众数相同，则票加1，不同则票减1
*/

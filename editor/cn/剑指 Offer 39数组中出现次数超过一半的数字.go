// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//
//
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1:
//
// 输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
// 输出: 2
//
//
//
// 限制：
//
// 1 <= 数组长度 <= 50000
//
//
//
// 注意：本题与主站 169 题相同：https://leetcode-cn.com/problems/majority-element/
//
//
// Related Topics 数组 哈希表 分治 计数 排序 👍 246 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	result := majorityElement(nums)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	votes := 0 // 投票统计
	x := 0     // 临时众数
	for _, num := range nums {
		if votes == 0 { // 当前票数为 0
			x = num // 新选临时众数据
		}
		if x == num {
			votes++
		} else {
			votes--
		}
	}
	count := 0 // 检查是否是众数
	for _, num := range nums {
		if num == x {
			count++
		}
	}
	if count > len(nums)/2 {
		return x
	} else {
		return 0
	}
}

// leetcode submit region end(Prohibit modification and deletion)

/*
哈希表法：遍历数组 nums 用 hash 表统计各数组的数量，找出众数据。时间和空间复杂度均为 O(n)
数组排序法：将数组 nums 排序，数组中向下取整的中间值一定为众数；时间复杂度 O(nlog(n)) 空间复杂度 `O(log(n))`
摩尔投票法：核心思想 票数正负抵消，此方法时间复杂度 O(log(n)) 空间复杂度 O(1)

设输入数组 nums 的众数为 x ，数组长度为 n 。

推论一： 若记 众数 的票数为 +1 ，非众数 的票数为 −1 ，则一定有所有数字的 票数和 >0 。
推论二： 若数组的前 a 个数字的 票数和 =0 ，则 数组剩余 (n−a) 个数字的 票数和一定仍 >0 ，即后 (n−a) 个数字的 众数仍为 x 。

*/

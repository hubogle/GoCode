// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出
// 这个数字。
//
//
//
// 示例 1:
//
// 输入: [0,1,3]
// 输出: 2
//
//
// 示例 2:
//
// 输入: [0,1,2,3,4,5,6,7,9]
// 输出: 8
//
//
//
// 限制：
//
// 1 <= 数组长度 <= 10000
// Related Topics 位运算 数组 哈希表 数学 二分查找 👍 226 👎 0

package main

import "fmt"

func main() {
	nums := []int{0, 1, 3}
	result := missingNumber(nums)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func missingNumber(nums []int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l <= r {
		mid = l + ((r - l) >> 1)
		if nums[mid] == mid {
			l = mid + 1
		} else { // 只要中值不等于，都往左子数查找
			r = mid - 1
		}
	}
	return l // 考虑临界值情况 [0]
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：有序序列 n 中，包含的值范围为 0~n-1 但是缺少一个。
题目相关：268 丢失的数字

循环二分 l <= r 时循环，当闭区间 [l,r] 为空时跳出，计算中点向下取整法 m = (l+r)//2
若 nums[m] = m 则缺少的数一定在右边
若 nums[m] != m 缺少的数一定在左边
跳出循环时遍历 l 和 r 分别指向右子数的首位元素，左子数的末位元素。因此返回 l

*/

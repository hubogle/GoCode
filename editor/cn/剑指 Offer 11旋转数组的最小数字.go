// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
//
// 给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。例如，数组 [3
// ,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为1。
//
// 示例 1：
//
//
// 输入：[3,4,5,1,2]
// 输出：1
//
//
// 示例 2：
//
//
// 输入：[2,2,2,0,1]
// 输出：0
//
//
// 注意：本题与主站 154 题相同：https://leetcode-cn.com/problems/find-minimum-in-rotated-
// sorted-array-ii/
// Related Topics 数组 二分查找 👍 523 👎 0

package main

import "fmt"

func main() {
	// nums := []int{3, 4, 5, 1, 2}
	// nums := []int{3, 1, 2}
	nums := []int{2, 3, 1}
	result := minArray(nums)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		mid := l + ((r - l) >> 1)
		if numbers[mid] < numbers[r] {
			r = mid
		} else if numbers[mid] > numbers[r] {
			l = mid + 1
		} else {
			r--
		} // 存在重复值
	}
	return numbers[l]
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解体思路：因为是升序旋转的数组并且还存在重复值，在其中寻找最小的数值。
要考虑正常情况下，如何判断
1. 为了防止死循环，l < r 不能等于，否则 nums[mid] > nums[r] 时 会死循环。
2. 因为存在重复值，所以当 nums[mid] == nums[r] 时 r-- 不然没办法考虑最小值在那边。

只要枚举下面三种情况，考虑进去就可以。
1,2,3
3,1,2
2,3,1
*/

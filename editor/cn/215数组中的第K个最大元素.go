// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
// 示例 1:
//
//
// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
//
//
// 示例 2:
//
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 1540 👎 0

package main

import "fmt"

func main() {
	result := findKthLargest([]int{1, 2, 3}, 2)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	return Search(nums, 0, len(nums)-1, len(nums)-k)
}

func Search(nums []int, l, r, k int) int {
	q := SearchPartition(nums, l, r)
	if q == k {
		return nums[q]
	}
	if q < k {
		return Search(nums, q+1, r, k)
	} else {
		return Search(nums, l, q-1, k)
	}
}

func SearchPartition(nums []int, l, r int) int {
	base := l
	index := l + 1
	for i := index; i <= r; i++ {
		if nums[i] < nums[base] {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[base], nums[index-1] = nums[index-1], nums[base]
	return index - 1
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：
快速排序的思想，取最左边一个值，将比改值小的放在左边，大的放在右边，返回中间值。
如果需要的最大值，等于中间值，则直接返回。
否则就在左右区间里面去寻找。
*/

// 给你一个整数数组 nums，请你将该数组升序排列。
//
//
//
//
//
//
// 示例 1：
//
//
// 输入：nums = [5,2,3,1]
// 输出：[1,2,3,5]
//
//
// 示例 2：
//
//
// 输入：nums = [5,1,1,2,0,0]
// 输出：[0,0,1,1,2,5]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁴
// -5 * 10⁴ <= nums[i] <= 5 * 10⁴
//
// Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 559 👎 0

package main

import "fmt"

func main() {
	result := sortArray([]int{5, 3, 2, 1})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	return quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) []int {
	if left < right {
		mid := searchMid(nums, left, right)
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
	return nums
}
func searchMid(nums []int, left, right int) int {
	base := left
	index := left + 1
	for i := index; i <= right; i++ {
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
快速排序
选取中间值，将比中间值小的放在左边，大的放在右边。
下次递归处理左半部，和右半部
直到 left >= right
*/

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
// 示例 1:
//
//
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
//
//
// 示例 2:
//
//
// 输入: nums = [0]
// 输出: [0]
//
//
//
// 提示:
//
//
//
// 1 <= nums.length <= 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
//
// 进阶：你能尽量减少完成的操作次数吗？
// Related Topics 数组 双指针 👍 1502 👎 0

package main

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	l, r, n := 0, 0, len(nums)
	for r < n {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func moveZeroes1(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		for j < i && nums[j] != 0 {
			j++
			if j == n {
				return
			}
		}
		if nums[j] == 0 && nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
}

/*
解题思路：双指针
快指针找到非 0 的，与慢指针交换，慢指针 ++
*/

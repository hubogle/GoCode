// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
//
//
//
// 示例：
//
//
// 输入：nums = [1,2,3,4]
// 输出：[1,3,2,4]
// 注：[3,1,2,4] 也是正确的答案之一。
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 50000
// 0 <= nums[i] <= 10000
//
// Related Topics 数组 双指针 排序 👍 198 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 3, 5}
	result := exchange(nums)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func exchange(nums []int) []int {
	for low, fast := 0, 0; fast < len(nums); fast++ {
		if nums[fast]&0x1 == 1 {
			nums[fast], nums[low] = nums[low], nums[fast]
			low++
		} // 双指针移动，fast 指向奇数，low 指向下一个交换位置
		if low < fast && nums[low]&0x1 == 1 {
			low++
		}
	}
	return nums
}

// leetcode submit region end(Prohibit modification and deletion)

func exchange0(nums []int) []int {
	n := len(nums) - 1
	for l, r := 0, n; l < r; l, r = l+1, r-1 {
		for l <= n && nums[l]&0x1 == 1 {
			l++
		}
		for r >= 0 && nums[r]&0x1 == 0 {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	return nums
}

func exchange1(nums []int) []int {
	for l, r := 0, len(nums)-1; l < r; {
		if nums[l]&0x1 == 1 {
			l++
			continue
		}
		if nums[r]&0x1 == 0 {
			r--
			continue
		}
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return nums
}

func exchange2(nums []int) []int {
	for low, fast := 0, 0; fast < len(nums); fast++ {
		if nums[fast]&0x1 == 1 {
			nums[fast], nums[low] = nums[low], nums[fast]
			low++
		} // 双指针移动，fast 指向奇数，low 指向下一个交换位置
		if low < fast && nums[low]&0x1 == 1 {
			low++
		}
	}
	return nums
}

// 11121

/*
题目：将一个数组中，奇数移动到前面，偶数移动到后面
时间复杂度：O(n)
空间复杂度：O(1)

解体思路：双指针，左指针只将偶数，右指针只将奇数进行交换。
注意左、右指针移动过程中必须符合切片索引获取条件。
解体思路2:快慢指针，快指针指向奇数位置，慢指针指向下一个奇数存放的位置。
fast 搜索到奇数时与 low 交换，此时 low 向前移动
*/

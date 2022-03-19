// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,0]
// 输出：3
//
//
// 示例 2：
//
//
// 输入：nums = [3,4,-1,1]
// 输出：2
//
//
// 示例 3：
//
//
// 输入：nums = [7,8,9,11,12]
// 输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
//
// Related Topics 数组 哈希表 👍 1397 👎 0

package main

import "fmt"

func main() {
	result := firstMissingPositive([]int{3, 4, -1, 1})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] < n && nums[i] != nums[nums[i]-1] { // nums[i] = i + 1
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return n + 1
}

// leetcode submit region end(Prohibit modification and deletion)
// 将数组下标当作哈希表使用
func firstMissingPositive1(nums []int) int {
	n := len(nums) // 数组值范围在 [1,n]
	for k, v := range nums {
		if v <= 0 {
			nums[k] = n + 1
		}
	} // 先将负数变为不符合条件的整数
	for i := 0; i < n; i++ {
		num := nums[i]
		if num < 0 {
			num = -num
		}
		if num <= n {
			v := nums[num-1]
			if v > 0 {
				nums[num-1] = -v
			}
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

/*
解题思路：
利用数组作为哈希表进行标识。
1. 首先将 <= 0 的数变为 n+1 的值
2. 循环遍历 val 将下标为 |val| 对应的数变为负数。
3. 找到第一个不为负数的下标，改下标 + 1 就是最小值。

解题思路：将值与对应下标的值进行交换 nums[i] != nums[nums[i]-1] 注意 [1,1] 这种情况
i + 1 = nums[i]
*/

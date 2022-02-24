// 给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
//
//
//
//
//
//
// 示例 1：
//
//
// 输入：nums = [3,0,1]
// 输出：2
// 解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
//
// 示例 2：
//
//
// 输入：nums = [0,1]
// 输出：2
// 解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。
//
// 示例 3：
//
//
// 输入：nums = [9,6,4,2,3,5,7,0,1]
// 输出：8
// 解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。
//
// 示例 4：
//
//
// 输入：nums = [0]
// 输出：1
// 解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁴
// 0 <= nums[i] <= n
// nums 中的所有数字都 独一无二
//
//
//
//
// 进阶：你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?
// Related Topics 位运算 数组 哈希表 数学 排序 👍 568 👎 0

package main

import "fmt"

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	result := missingNumber(nums)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func missingNumber(nums []int) int {
	target := 0
	for _, v := range nums {
		target ^= v
	} // 获取所有元素值
	for i := 0; i <= len(nums); i++ {
		target ^= i
	} // ^ 成对儿的数字都会消为 0
	return target
}

// leetcode submit region end(Prohibit modification and deletion)

func missingNumber1(nums []int) int {
	n := len(nums)
	target := 0
	target ^= n // 因为遍历索引 k 是从 0 开始的，缺少最大值 n
	for k, v := range nums {
		target ^= k
		target ^= v
	} // ^ 成对的数都会抵消掉
	return target
}

/*
因为数组范围数是无序的，不然可用二分查找
那么用 ^ 异或运算，可以将成对的数消除掉。

第二种方式区别第一种就是，for 循环的索引缺少一个最大值
*/

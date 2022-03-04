// 从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，
// 可以看成任意数字。A 不能视为 14。
//
//
//
// 示例 1:
//
//
// 输入: [1,2,3,4,5]
// 输出: True
//
//
//
// 示例 2:
//
//
// 输入: [0,0,1,2,5]
// 输出: True
//
//
//
// 限制：
//
// 数组长度为 5
//
// 数组的数取值为 [0, 13] .
// Related Topics 数组 排序 👍 206 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{9, 10, 4, 0, 9}
	result := isStraight(nums)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isStraight(nums []int) bool {
	flag := int(^uint(0) >> 1) // 存储最小值
	for _, v := range nums {
		if v < flag && v != 0 {
			flag = v
		}
	}
	for k := range nums {
		// k + flag = num[k]
		if nums[k]-flag >= 5 {
			return false
		} // 减去最小值小于 5
		for k+flag != nums[k] && nums[k] != 0 { // 2 + 2 != nums[2] != 4
			if nums[k] == nums[nums[k]-flag] {
				return false
			} // 重复数字
			nums[k], nums[nums[k]-flag] = nums[nums[k]-flag], nums[k]
		}
	}
	for k, v := range nums {
		if k+flag != v && v != 0 {
			return false
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func isStraight1(nums []int) bool {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			ans += 1
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[ans] < 5
}

/*
题目：判断五个数字是否为顺子
解题思路：排序 + 遍历
时间复杂度：O(n*log(n))
空间复杂度：O(log(n))

注意：不要单纯的将问题想复杂度，不要一味追求解答方案的极致

五个数字为顺子的条件：
1. 除大小王外不能有重复的
2. 牌最大值-牌最小值 < 5

一开始的思路，将里面的牌进行遍历，将最小值放在 nums[0] 的位置，将最大值放在右边，满足最大值减去最小值 < 5，切空位有 0 填充


解题思路：遍历数组
时间复杂度：O(n)
空间复杂度：O(1)

1. 第一次遍历出最小值
2. 第二次遍历
	2.1 将数组中的值减去最小值 >= 5 不符合条件
	2.2 将 num[i] != i + min 则进行交换两者 num[k], num[num[k]-min] 如果两者重复则不符合条件
3. 第三次遍历，判断 num[i] + min != val && val != 0 不符合条件
*/

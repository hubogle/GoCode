// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//
// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
//
//
//
// 示例 1：
//
// 输入：target = 9
// 输出：[[2,3,4],[4,5]]
//
//
// 示例 2：
//
// 输入：target = 15
// 输出：[[1,2,3,4,5],[4,5,6],[7,8]]
//
//
//
//
// 限制：
//
//
// 1 <= target <= 10^5
//
//
//
// Related Topics 数学 双指针 枚举 👍 384 👎 0

package main

import "fmt"

func main() {
	target := 15
	result := findContinuousSequence(target)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findContinuousSequence(target int) [][]int {
	var ans [][]int
	if target <= 2 {
		return ans
	}
	l, r := 1, 2
	sum := l
	for l < r && r <= target {
		if sum < target {
			sum += r
			r++
		} else {
			if sum == target {
				var val []int
				for i := l; i < r; i++ {
					val = append(val, i)
				}
				ans = append(ans, val)
			}
			sum -= l
			l++
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func findContinuousSequence1(target int) [][]int {
	var ans [][]int
	for l, r := 1, 2; l < r; {
		sum := (l + r) * (r - l + 1) / 2 // 直接 O(1) 求出
		if sum == target {
			val := make([]int, r-l+1)
			for i := l; i <= r; i++ {
				val[i-l] = i
			}
			ans = append(ans, val)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return ans
}

/*
题目：输入一个正整数 target ，输出所有和为 target 的连续正整数序列
解题思路：双指针，左右指针区间的和累加等于 target 时
	情况1：累加和大于 target 左指针前移动
	情况2：累加和小于 target 右指针前移动
思路2：求连续递增数组的总和公式 (l + r)(r - l + 1) / 2
*/

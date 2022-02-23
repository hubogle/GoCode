// 给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
//
//
//
// 示例 1：
//
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2]
//
//
// 示例 2：
//
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[9,4]
// 解释：[4,9] 也是可通过的
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
//
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 490 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	result := intersection1(nums1, nums2)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	s1 := map[int]struct{}{}
	s2 := map[int]struct{}{}
	var res []int
	for _, v := range nums1 {
		s1[v] = struct{}{}
	}
	for _, v := range nums2 {
		s2[v] = struct{}{}
	}
	if len(s1) > len(s2) { // 遍历长度最小的可以节省时间
		s1, s2 = s2, s1
	}
	for v := range s2 {
		if _, has := s1[v]; has {
			res = append(res, v)
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func intersection1(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int // ⚠️ 插入值的时候判断是否已经存在
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] { // 要插入的数值只会大于，等于的话就是重复了
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return res
}

/*
题目：返回两个数组的交集
思路1：两个数组用 `map` 存储，然后遍历其中一个数组，查看是否在另一个数组中存在
时间复杂度：O(n+m)
空间复杂度：O(n+m)

思路2：将两个数组先排序，然后双指针在两个数组中移动查找重复值。
时间复杂度：O(n * log(n) + m * log(m))
空间复杂度：O(log(m) + log(n)) 快排需要的额外空间
*/

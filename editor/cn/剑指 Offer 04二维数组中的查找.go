// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个
// 整数，判断数组中是否含有该整数。
//
//
//
// 示例:
//
// 现有矩阵 matrix 如下：
//
//
// [
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
// ]
//
//
// 给定 target = 5，返回 true。
//
// 给定 target = 20，返回 false。
//
//
//
// 限制：
//
// 0 <= n <= 1000
//
// 0 <= m <= 1000
//
//
//
// 注意：本题与主站 240 题相同：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
// Related Topics 数组 二分查找 分治 矩阵 👍 570 👎 0

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 0
	result := findNumberIn2DArray(matrix, target)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	col := len(matrix[0]) - 1 // 列
	line := len(matrix) - 1   // 行
	i, j := 0, line           // 没办法从0,0 开始因为是最小值，需要从中值开始判断，左下或者右上
	for i <= col && j >= 0 {
		if target == matrix[j][i] {
			return true
		} else if target > matrix[j][i] {
			i++
		} else {
			j--
		}
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)

/*
时间复杂度：O(n + m) 行和列数量综合
空间复杂度：O(1)

题目特点：二维数组每行从左到右，从上到下递增特点，从对角线开始搜索可以成半过滤。
思路：从左下\右上过滤元素查找(中间值)，无法从左上\右下查找(极值)

* 数组为空返回 false
* 通过下标访问元素位置，直到下标超过边界
	1. 获取下标位置元素
	2. 如果元素相等返回 true
	3. 如果 num > target 行减1
	4. 如果 num < target 列加1

*/

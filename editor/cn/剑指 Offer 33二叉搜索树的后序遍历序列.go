// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
//
//
//
// 参考以下这颗二叉搜索树：
//
//      5
//    / \
//   2   6
//  / \
// 1   3
//
// 示例 1：
//
// 输入: [1,6,3,2,5]
// 输出: false
//
// 示例 2：
//
// 输入: [1,3,2,6,5]
// 输出: true
//
//
//
// 提示：
//
//
// 数组长度 <= 1000
//
// Related Topics 栈 树 二叉搜索树 递归 二叉树 单调栈 👍 476 👎 0

package main

import "fmt"

func main() {
	result := verifyPostorder([]int{})
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func verifyPostorder(postorder []int) bool {
	var seach func(l, r int) bool
	seach = func(l, r int) bool {
		if l >= r {
			return true
		}
		v := postorder[r]
		i := l // 每次的左边
		for postorder[i] < v {
			i++
		}
		mid := i
		for postorder[i] > v {
			i++
		}
		return i == r && seach(l, mid-1) && seach(mid, r-1) // (l, mid-1) (mid, r-1)
	}
	return seach(0, len(postorder)-1)
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：二分法
每次取最后一个值为 mid，从数组里面找出来第一个大于mid的位置，迭代查找(l, mid-1) 为左边 (mid, r-1) 为右边
判断条件 每次查找遍历到最右边，i == r && search(l, mid-1) && search(mid, r-1)
*/

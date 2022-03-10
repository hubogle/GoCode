// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 差值是一个正数，其数值等于两值之差的绝对值。
//
//
//
// 示例 1：
//
//
// 输入：root = [4,2,6,1,3]
// 输出：1
//
//
// 示例 2：
//
//
// 输入：root = [1,0,48,null,null,12,49]
// 输出：1
//
//
//
//
// 提示：
//
//
// 树中节点的数目范围是 [2, 10⁴]
// 0 <= Node.val <= 10⁵
//
//
//
//
// 注意：本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-
// nodes/ 相同
// Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 二叉树 👍 313 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	result := getMinimumDifference(new(TreeNode))
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	var pre *TreeNode
	var stack []*TreeNode
	var ans int
	ans = int(^uint(0) >> 1)
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil {
				ans = min(ans, int(math.Abs(float64(root.Val)-float64(pre.Val))))
			}
			pre = root
			root = root.Right
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// leetcode submit region end(Prohibit modification and deletion)

// 递归写法
func getMinimumDifference1(root *TreeNode) int {
	var pre *TreeNode
	var ans int
	ans = int(^uint(0) >> 1)
	var getMin func(node *TreeNode)
	getMin = func(node *TreeNode) {
		if node == nil {
			return
		}
		getMin(node.Left)
		if pre != nil {
			ans = min(ans, int(math.Abs(float64(pre.Val)-float64(node.Val))))
		}
		pre = node
		getMin(node.Right)
	}
	getMin(root)
	return ans
}

/*
题目：求二叉搜索树两个节点之间的最小相对值

思路：暴力上一个节点，利用中序遍历当前节点的值与上一个节点的值相减获得最小绝对值
*/

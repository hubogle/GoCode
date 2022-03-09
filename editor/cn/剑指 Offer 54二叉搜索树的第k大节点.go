// 给定一棵二叉搜索树，请找出其中第 k 大的节点的值。
//
//
//
// 示例 1:
//
//
// 输入: root = [3,1,4,null,2], k = 1
//   3
//  / \
// 1   4
//  \
//    2
// 输出: 4
//
// 示例 2:
//
//
// 输入: root = [5,3,6,2,4,null,null,1], k = 3
//       5
//      / \
//     3   6
//    / \
//   2   4
//  /
// 1
// 输出: 4
//
//
//
// 限制：
//
//
// 1 ≤ k ≤ 二叉搜索树元素个数
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 261 👎 0

package main

import "fmt"

func main() {
	result := kthLargest(new(TreeNode), 1)
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
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var stack []*TreeNode
	var ans []int
	stack = append(stack, root)
	for n := len(stack); n > 0; n = len(stack) {
		node := stack[n-1]
		stack = stack[:n-1]
		if node == nil {
			node = stack[n-2]
			stack = stack[:n-2]
			ans = append(ans, node.Val)
		} else {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}
	return ans[len(ans)-k]
}

// leetcode submit region end(Prohibit modification and deletion)

func kthLargest1(root *TreeNode, k int) int {
	var stack []*TreeNode
	var ans []int
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, root.Val)
			root = root.Right
		}
	}
	return ans[len(ans)-k]
}

/*
求最大值的问题以及 top k 问题，都可以中序遍历获取数组来得到结果。
中序迭代遍历的两种写法
*/

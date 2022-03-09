// 输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
//
// 例如：
//
// 给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
//
//
//
// 提示：
//
//
// 节点总数 <= 10000
//
//
// 注意：本题与主站 104 题相同：https://leetcode-cn.com/problems/maximum-depth-of-binary-
// tree/
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 167 👎 0

package main

import (
	"fmt"
)

func main() {
	result := maxDepth(new(TreeNode))
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth func(node *TreeNode, n int) int
	// depth = func(node *TreeNode, n int) int {
	// 	if node == nil {
	// 		return n
	// 	}
	// 	return max(depth(node.Left, n+1), depth(node.Right, n+1))
	// }
	// return depth(root, 0)
	depth = func(node *TreeNode, n int) int {
		if node.Left == nil && node.Right == nil {
			return n
		}
		left := n
		right := n
		if node.Left != nil {
			left = depth(node.Left, n+1)
		}
		if node.Right != nil {
			right = depth(node.Right, n+1)
		}
		return max(right, left)
	}
	return depth(root, 1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// leetcode submit region end(Prohibit modification and deletion)

// 迭代层级遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	var ans int
	queue = append(queue, root)
	for n := len(queue); n > 0; {
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
		ans++
		n = len(queue)
	}
	return ans
}

/*
题目：二叉树最大深度
在递归判断的时候，要注意 return 情况是判断为 nil 还是左右节点为 nil
*/

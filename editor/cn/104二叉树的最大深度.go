// 给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1135 👎 0

package main

import "fmt"

func main() {
	result := maxDepth(nil)
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
	var depth func(node *TreeNode, n int)
	var ans int
	if root == nil {
		return ans
	}
	depth = func(node *TreeNode, n int) {
		ans = max(ans, n)
		if node == nil {
			return
		}
		if node.Left != nil {
			depth(node.Left, n+1)
		}
		if node.Right != nil {
			depth(node.Right, n+1)
		}
	}
	depth(root, 1)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

// 递归解答
// 后序遍历
func maxDepth1(root *TreeNode) int {
	var depth func(node *TreeNode, n int) int
	depth = func(node *TreeNode, n int) int {
		if node == nil {
			return n
		}
		// left := depth(node.Left, n +1) // 左
		// right := depth(node.Right, n+1) // 右
		// ans := max(left, right) // 中
		return max(depth(node.Left, n+1), depth(node.Right, n+1))
	}
	return depth(root, 0)
}

// 递归解答
// 前序遍历
func maxDepth2(root *TreeNode) int {
	var depth func(node *TreeNode, n int)
	var ans int
	if root == nil {
		return ans
	} // 考虑节点为 nil 的情况
	depth = func(node *TreeNode, n int) {
		ans = max(ans, n) // 前序遍历
		if node == nil {
			return
		}
		if node.Left != nil {
			depth(node.Left, n+1)
		}
		if node.Right != nil {
			depth(node.Right, n+1)
		}
	}
	depth(root, 1)
	return ans
}

// 迭代解答
// 层序遍历
func maxDepth3(root *TreeNode) int {
	var stack []*TreeNode
	var node *TreeNode
	var ans int
	if root != nil {
		stack = append(stack, root)
	} // 注意 root 为 nil 时不能添加

	for n := len(stack); n > 0; {
		m := n
		for i := 0; i < m; i++ {
			node = stack[i]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		stack = stack[m:]
		ans++
		n = len(stack)
	}
	return ans
}

/*
题目：二叉树的最大深度
解题思路：深度指从上到下开始查找，所以用前序遍历进行。

递归解题：前序（中左右）或后序遍历（左右中），使用前序求的就是深度，使用后序求的是高度。
		而根节点的高度就是二叉树的最大深度
		前序遍历：要考虑 nil 节点的情况，因为是先算当前节点的深度，然后遍历左右
		后序遍历：不用考虑 nil 节点，因为先左右后当前节点

迭代解题：层序遍历法
		用前序遍历的迭代方法不好写
*/

// 给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。
//
//
//
// 示例 1：
//
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2
//
//
// 示例 2：
//
//
// 输入：root = [2,null,3,null,4,null,5,null,6]
// 输出：5
//
//
//
//
// 提示：
//
//
// 树中节点数的范围在 [0, 10⁵] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 681 👎 0

package main

import (
	"fmt"
)

func main() {
	result := minDepth(new(TreeNode))
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
func minDepth(root *TreeNode) int {
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right != nil {
			return 1 + depth(root.Right)
		}
		if root.Left != nil && root.Right == nil {
			return 1 + depth(root.Left)
		}
		return min(depth(root.Left), depth(root.Right)) + 1
	}
	return depth(root)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

// 迭代 前序遍历
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root.Left == nil && root.Right == nil {
			return 1
		}
		ans := int(^uint(0) >> 1)
		if root.Left != nil {
			ans = min(depth(root.Left), ans)
		}
		if root.Right != nil {
			ans = min(depth(root.Right), ans)
		}
		// ans = mix(depth(root.Left), depth(root.Right)) + 1 错误写法
		return ans + 1
	}
	return depth(root)
}

// 迭代 后序遍历
func minDepth2(root *TreeNode) int {
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right != nil {
			return 1 + depth(root.Right)
		}
		if root.Left != nil && root.Right == nil {
			return 1 + depth(root.Left)
		}
		return min(depth(root.Left), depth(root.Right)) + 1
	}
	return depth(root)
}

// 层序遍历
func minDepth3(root *TreeNode) int {
	ans := int(^uint(0) >> 1)
	depth := 0
	var queue []*TreeNode
	if root == nil {
		return 0
	}
	queue = append(queue, root)
	for n := len(queue); n > 0; {
		m := n
		depth++
		for i := 0; i < m; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				ans = min(ans, depth)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[m:]
		n = len(queue)
	}
	return ans
}

/*
题目：二叉树最小深度

误区：要注意只有没有左右子树的节点，才叫做叶子节点。

解题思路：
递归法：
	后序遍历：递归终止情况
		node == nil return 0
		node.Left == nil && node.Right != nil 则右节点 + 1
		node.Right == nil && node.Left != nil 则左节点 + 1
		min(node.Left, node.Right) + 1
    前序遍历：
		左右节点为 nil 时返回 1
		不然的话取出左右节点的最小值 + 1
	// ans = mix(depth(root.Left), depth(root.Right)) + 1 前序遍历错误写法

后序遍历的情况遇到 nil 返回 0，前序遍历遇到 nil 返回 1 因为前序遍历首先访问根节点，后序遍历最后访问根节点为 nil
后序遍历相对前序遍历，处理 nil 节点的情况要注重判断

迭代法：层序遍历，只有遇到真正的叶子节点，才与目标返回值做比较。

*/

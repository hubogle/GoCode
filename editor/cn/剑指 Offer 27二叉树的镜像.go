// 请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//
// 例如输入：
//
// 4
// / \
// 2 7
// / \ / \
// 1 3 6 9
// 镜像输出：
//
// 4
// / \
// 7 2
// / \ / \
// 9 6 3 1
//
//
//
// 示例 1：
//
// 输入：root = [4,2,7,1,3,6,9]
// 输出：[4,7,2,9,6,3,1]
//
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 1000
//
// 注意：本题与主站 226 题相同：https://leetcode-cn.com/problems/invert-binary-tree/
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 219 👎 0

package main

import "fmt"

func main() {
	result := mirrorTree(new(TreeNode))
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
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for n := len(stack); n > 0; {
		node := stack[n-1]
		stack = stack[:n-1]
		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
		} else {
			node = stack[n-2]
			stack = stack[:n-2]
			node.Left, node.Right = node.Right, node.Left
		}
		n = len(stack)
	}
	return root
}

// leetcode submit region end(Prohibit modification and deletion)

// 前序遍历
func mirrorTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for n := len(stack); n > 0; {
		node := stack[n-1]
		stack = stack[:n-1]
		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		n = len(stack)
	}
	return root
}

/*
二叉树的前序列遍历、后序遍历，层级遍历。不能中序遍历会有节点交换两次
*/

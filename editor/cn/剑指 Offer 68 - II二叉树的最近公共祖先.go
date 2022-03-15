// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
// 一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]
//
//
//
//
//
// 示例 1:
//
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出: 3
// 解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
//
//
// 示例 2:
//
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出: 5
// 解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
//
//
//
//
// 说明:
//
//
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉树中。
//
//
// 注意：本题与主站 236 题相同：https://leetcode-cn.com/problems/lowest-common-ancestor-of-
// a-binary-tree/
// Related Topics 树 深度优先搜索 二叉树 👍 397 👎 0

package main

import "fmt"

func main() {
	result := lowestCommonAncestor(new(TreeNode), new(TreeNode), new(TreeNode))
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var search func(node *TreeNode) *TreeNode
	search = func(node *TreeNode) *TreeNode {
		if node == p || node == q || node == nil {
			return node
		}
		left := search(node.Left)
		right := search(node.Right)
		if left == nil {
			return right
		}
		if right == nil {
			return left
		}
		return node
	}
	return search(root)
}

// leetcode submit region end(Prohibit modification and deletion)
/*
题目：寻找两个节点的公共节点
解题思路：后序遍历整个二叉树。
当左右节点返回值都不为 nil 时，当前节点就是要返回的节点。
迭代法不适合回溯模拟的过程。
*/

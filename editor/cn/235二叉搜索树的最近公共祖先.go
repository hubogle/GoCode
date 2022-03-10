// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
// 一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5]
//
//
//
//
//
// 示例 1:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// 输出: 6
// 解释: 节点 2 和节点 8 的最近公共祖先是 6。
//
//
// 示例 2:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
// 输出: 2
// 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
//
//
//
// 说明:
//
//
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉搜索树中。
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 788 👎 0

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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var search func(node *TreeNode) *TreeNode
	search = func(node *TreeNode) *TreeNode {
		if node == nil {
			return node
		} // 前序遍历 中左右
		if node.Val > p.Val && node.Val > q.Val {
			left := search(node.Left)
			if left != nil {
				return left
			}
		}
		if node.Val < p.Val && node.Val < q.Val {
			right := search(node.Right)
			if right != nil {
				return right
			}
		}
		return node // 最终返回的还是 root
	}
	return search(root)
}

// leetcode submit region end(Prohibit modification and deletion)

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return root
}

/*
题目：二叉搜索树中查找公共节点
与普通二叉树查找公共节点，简单。
二叉搜索树是有序的，可以利用前序遍历，当该前节点处于两个节点之间时，该节点就是要找节点。
二叉搜索树，可以一条边遍历，而二叉树就不可以。
二叉树要拿到左右节点进行判断是否符合条件（这里表示的就是回溯，处理中间节点），因为二叉搜索树有序可以直接返回。
所以二叉树搜索公共节点只能后序遍历，而二叉搜索树的情况只要自上而下遍历二是区区是叉树就可以
*/

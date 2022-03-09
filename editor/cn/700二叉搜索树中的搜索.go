// 给定二叉搜索树（BST）的根节点 root 和一个整数值 val。
//
// 你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 null 。
//
//
//
// 示例 1:
//
//
//
//
// 输入：root = [4,2,7,1,3], val = 2
// 输出：[2,1,3]
//
//
// Example 2:
//
//
// 输入：root = [4,2,7,1,3], val = 5
// 输出：[]
//
//
//
//
// 提示：
//
//
// 数中节点数在 [1, 5000] 范围内
// 1 <= Node.val <= 10⁷
// root 是二叉搜索树
// 1 <= val <= 10⁷
//
// Related Topics 树 二叉搜索树 二叉树 👍 247 👎 0

package main

import "fmt"

func main() {
	result := searchBST(new(TreeNode), 0)
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
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}

// leetcode submit region end(Prohibit modification and deletion)

func searchBST1(root *TreeNode, val int) *TreeNode {
	var search func(node *TreeNode) *TreeNode
	search = func(node *TreeNode) *TreeNode {
		if node == nil || node.Val == val {
			return node
		}
		if node.Val > val {
			return search(node.Left)
		} else {
			return search(node.Right)
		}
	}
	return search(root)
}

func searchBST2(root *TreeNode, val int) *TreeNode {
	var stack []*TreeNode
	stack = append(stack, root)
	for n := len(stack); n > 0; n = len(stack) {
		node := stack[n-1]
		stack = stack[:n-1]
		if node == nil || node.Val == val {
			return node
		}
		if node.Val > val {
			stack = append(stack, node.Left)
		} else {
			stack = append(stack, node.Right)
		}
	}
	return nil
}

/*
二叉搜索树中搜索：
利用二叉搜索树的特性，比当前节点大的话则去右节点找。可以用递归和迭代解答
*/

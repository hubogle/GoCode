// 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,null,2,3]
// 输出：[3,2,1]
//
//
// 示例 2：
//
//
// 输入：root = []
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：root = [1]
// 输出：[1]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 767 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	node3 := TreeNode{Val: 2}
	node2 := TreeNode{Val: 1}
	node1 := TreeNode{Val: 0, Left: &node2, Right: &node3}
	result := postorderTraversal(&node1)
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	stack := list.New() // 迭代遍历
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		ans = append(ans, node.Val) // 中右左遍历
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	// 中右左 倒序后 左中右
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func postorderTraversal1(root *TreeNode) []int {
	var ans []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		ans = append(ans, node.Val)
	}
	traverse(root)
	return ans
}

/*
题目：二叉树后序遍历
后序迭代遍历与前序迭代遍历的区别在于，将中右左的结果反转成 左中右的结果。
*/

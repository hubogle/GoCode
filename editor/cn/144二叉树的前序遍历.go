// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,null,2,3]
// 输出：[1,2,3]
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
// 示例 4：
//
//
// 输入：root = [1,2]
// 输出：[1,2]
//
//
// 示例 5：
//
//
// 输入：root = [1,null,2]
// 输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 742 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	node2 := TreeNode{Val: 2}
	node1 := TreeNode{Val: 1}
	node := TreeNode{Val: 0, Left: &node1, Right: &node2}
	result := preorderTraversal(&node)
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
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var node []*TreeNode
	if root == nil {
		return nil
	}
	node = append(node, root)
	for len(node) != 0 {
		v := node[len(node)-1]
		node = node[:len(node)-1]
		ans = append(ans, v.Val)
		if v.Right != nil {
			node = append(node, v.Right)
		}
		if v.Left != nil {
			node = append(node, v.Left)
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func preorderTraversal1(root *TreeNode) []int {
	var ans []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return ans
}

func preorderTraversal2(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return nil
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}

/*
二叉树前序遍历
方法1：递归遍历，写递归函数。
	注意递归函数停止结束条件判断，以及递归函数定义方式，从根节点递归，插入值后递归左节点，然后递归右节点
方式2: 迭代遍历，利用栈的特性
	将根节点入栈，每次从栈顶弹出一个数据，然后入右节点，然后左节点，依次循环遍历节点。
*/

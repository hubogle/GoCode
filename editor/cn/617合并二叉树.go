// 给你两棵二叉树： root1 和 root2 。
//
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠
// ，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
//
// 返回合并后的二叉树。
//
// 注意: 合并过程必须从两个树的根节点开始。
//
//
//
// 示例 1：
//
//
// 输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// 输出：[3,4,5,5,4,null,7]
//
//
// 示例 2：
//
//
// 输入：root1 = [1], root2 = [1,2]
// 输出：[2,2]
//
//
//
//
// 提示：
//
//
// 两棵树中的节点数目在范围 [0, 2000] 内
// -10⁴ <= Node.val <= 10⁴
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 897 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	result := mergeTrees()
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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue := list.New()
	queue.PushBack(root1)
	queue.PushBack(root2)
	for queue.Len() > 0 {
		node1 := queue.Remove(queue.Front()).(*TreeNode)
		node2 := queue.Remove(queue.Front()).(*TreeNode)
		node1.Val += node2.Val
		if node1.Left != nil && node2.Left != nil {
			queue.PushBack(node1.Left)
			queue.PushBack(node2.Left)
		}
		if node1.Right != nil && node2.Right != nil {
			queue.PushBack(node1.Right)
			queue.PushBack(node2.Right)
		}
		if node1.Left == nil && node2.Left != nil {
			node1.Left = node2.Left
		}
		if node1.Right == nil && node2.Right != nil {
			node1.Right = node2.Right
		}
	}
	return root1
}

// leetcode submit region end(Prohibit modification and deletion)

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var merge func(node1 *TreeNode, node2 *TreeNode) *TreeNode
	merge = func(node1 *TreeNode, node2 *TreeNode) *TreeNode {
		if node1 == nil {
			return node2
		}
		if node2 == nil {
			return node1
		}
		node1.Val += node2.Val
		node1.Left = merge(node1.Left, node2.Left)
		node1.Right = merge(node1.Right, node2.Right)
		return node1
	}
	return merge(root1, root2)
}

/*
题目：合并两个二叉树
递归写法：考虑两个节点什么情况下节点相加，当其中一个节点为 nil 的情况
迭代写法：用队列模拟同时控制两棵树节点的进入，只有两棵树节点同时不为空时才进入栈，否则就赋值给另一棵树的节点。
*/

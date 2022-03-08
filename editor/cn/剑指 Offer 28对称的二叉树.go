// 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
// 1
// / \
// 2 2
// / \ / \
// 3 4 4 3
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
// 1
// / \
// 2 2
// \ \
// 3 3
//
//
//
// 示例 1：
//
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
//
//
// 示例 2：
//
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 1000
//
// 注意：本题与主站 101 题相同：https://leetcode-cn.com/problems/symmetric-tree/
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 294 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	result := isSymmetric(new(TreeNode))
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var symmetric func(left *TreeNode, right *TreeNode) bool
	symmetric = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || right.Val != left.Val {
			return false
		}
		return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
	}
	return symmetric(root.Left, root.Right)
}

// leetcode submit region end(Prohibit modification and deletion)

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue *list.List
	queue = list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() > 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue.PushBack(left.Left)
		queue.PushBack(right.Right)
		queue.PushBack(left.Right)
		queue.PushBack(right.Left)
	}
	return true
}

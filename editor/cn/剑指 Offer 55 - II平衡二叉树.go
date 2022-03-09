// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
//
//
//
// 示例 1:
//
// 给定二叉树 [3,9,20,null,null,15,7]
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回 true 。
//
// 示例 2:
//
// 给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
//
//
// 返回 false 。
//
//
//
// 限制：
//
//
// 0 <= 树的结点个数 <= 10000
//
//
// 注意：本题与主站 110 题相同：https://leetcode-cn.com/problems/balanced-binary-tree/
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 237 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	result := isBalanced(new(TreeNode))
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
func isBalanced(root *TreeNode) bool {
	var balanced func(node *TreeNode) int // 返回树当前高度，若为 -1 则不平衡
	balanced = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := balanced(node.Left)
		right := balanced(node.Right)
		if left == -1 || right == -1 {
			return -1
		}
		if math.Abs(float64(left-right)) > 1 {
			return -1
		}
		return int(math.Max(float64(left), float64(right))) + 1
	}
	return balanced(root) != -1
}

// leetcode submit region end(Prohibit modification and deletion)

/*
判断平衡二叉树的时候，用后序遍历最好。
先求左右节点，若左右节点为 -1 直接返回 -1，否则返回左右节点其中的最大一个高度。
*/

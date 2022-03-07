// 给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
//
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
//
//
//
//
// 示例 1：
//
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
//
//
// 示例 2：
//
//
// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
//
//
// 示例 3：
//
//
// 输入：root = []
// 输出：true
//
//
//
//
// 提示：
//
//
// 树中的节点数在范围 [0, 5000] 内
// -10⁴ <= Node.val <= 10⁴
//
// Related Topics 树 深度优先搜索 二叉树 👍 912 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	result := isBalanced(new(TreeNode))
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
func isBalanced(root *TreeNode) bool {
	var balanced func(node *TreeNode) int
	balanced = func(node *TreeNode) int {
		if node == nil {
			return 0
		} // 后序遍历不用判断左右树是否为空
		left := balanced(node.Left)
		// 左节点
		right := balanced(node.Right) // 右节点
		// 右节点
		if left == -1 || right == -1 {
			return -1
		}
		if math.Abs(float64(left-right)) > 1 {
			return -1
		}
		return int(math.Max(float64(left), float64(right))) + 1 // 中节点
	}
	return balanced(root) != -1
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：平衡二叉树判断

# 递归解题
后序遍历的解法，先遍历左右节点，然后判断左右树高度差是否大于 1
1. 递归函数参数和返回值：参数为节点，返回值为节点的高度，如果不是平衡树则返回 -1
2. 终止条件：遇到空节点返回 0
3. 单层递归逻辑：判断左右树高度差是否大于1，或取左右最高树高度 + 1

# 迭代解题
1. 需要定义一个单独求树高度的方法（前序遍历）
2. 每层遍历节点左右树高度差
该方式有很多重复计算
*/

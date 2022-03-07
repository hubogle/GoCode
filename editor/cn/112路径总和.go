// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和
// targetSum 。如果存在，返回 true ；否则，返回 false 。
//
// 叶子节点 是指没有子节点的节点。
//
//
//
// 示例 1：
//
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// 输出：true
// 解释：等于目标和的根节点到叶节点路径如上图所示。
//
//
// 示例 2：
//
//
// 输入：root = [1,2,3], targetSum = 5
// 输出：false
// 解释：树中存在两条根节点到叶子节点的路径：
// (1 --> 2): 和为 3
// (1 --> 3): 和为 4
// 不存在 sum = 5 的根节点到叶子节点的路径。
//
// 示例 3：
//
//
// 输入：root = [], targetSum = 0
// 输出：false
// 解释：由于树是空的，所以不存在根节点到叶子节点的路径。
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 802 👎 0

package main

import "fmt"

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node := &TreeNode{Val: 3, Left: node1, Right: node2}
	result := hasPathSum(node, 5)
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var stack []*TreeNode
	var sums []int
	stack = append(stack, root)
	sums = append(sums, root.Val)
	for n := len(stack); n > 0; {
		node := stack[n-1]
		sum := sums[n-1]
		stack = stack[:n-1]
		sums = sums[:n-1]
		if node.Left == nil && node.Right == nil && sum == targetSum {
			return true
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			sums = append(sums, sum+node.Val)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			sums = append(sums, sum+node.Val)
		}
		n = len(stack)
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)

func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val // 将targetSum在遍历每层的时候都减去本层节点的值
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var pathSum func(node *TreeNode, target int)
	var flag bool
	pathSum = func(node *TreeNode, target int) {
		if node.Left == nil && node.Right == nil {
			if target+node.Val == targetSum {
				flag = true
			}
			return
		}
		if node.Left != nil {
			pathSum(node.Left, target+node.Val)
		}
		if node.Right != nil {
			pathSum(node.Right, target+node.Val)
		}
	}
	pathSum(root, 0)
	return flag
}

/*
题目：二叉树搜索节点累加值是否有目标数

递归解题：
1. 正常的遍历，只不过递归参数传递累加值。
2. 判断叶子节点加值是否为目标值

迭代解题目：
1. 两个栈一个栈存储 node，另一个栈存储当前 node 的 sum
*/

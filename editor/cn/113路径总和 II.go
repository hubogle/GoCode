// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
//
//
//
// 示例 1：
//
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：[[5,4,11,2],[5,8,4,5]]
//
//
// 示例 2：
//
//
// 输入：root = [1,2,3], targetSum = 5
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：root = [1,2], targetSum = 0
// 输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点总数在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
//
//
//
// Related Topics 树 深度优先搜索 回溯 二叉树 👍 686 👎 0

package main

import "fmt"

func main() {
	result := pathSum(new(TreeNode), 0)
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var path func(node *TreeNode, sum int, path []int)
	path = func(node *TreeNode, sum int, nodePath []int) {
		sum += node.Val
		if node.Left == nil && node.Right == nil && targetSum == sum {
			nodePath = append(nodePath, node.Val)
			ans = append(ans, nodePath)
			return
		}
		if node.Left == nil && node.Right == nil {
			return
		}
		if node.Left != nil {
			nodeP := make([]int, len(nodePath))
			copy(nodeP, nodePath)
			nodeP = append(nodeP, node.Val)
			path(node.Left, sum, nodeP)
		}
		if node.Right != nil {
			nodeP := make([]int, len(nodePath))
			copy(nodeP, nodePath)
			nodeP = append(nodeP, node.Val)
			path(node.Right, sum, nodeP)
		}
	}
	path(root, 0, []int{})
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
//
// 示例 1：
//
//
//
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：[[5,4,11,2],[5,8,4,5]]
//
//
// 示例 2：
//
//
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
// 注意：本题与主站 113 题相同：https://leetcode-cn.com/problems/path-sum-ii/
// Related Topics 树 深度优先搜索 回溯 二叉树 👍 301 👎 0

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
func pathSum(root *TreeNode, target int) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var searchPath func(node *TreeNode, num int, path []int)
	searchPath = func(node *TreeNode, num int, path []int) {
		if node.Left == nil && node.Right == nil {
			if num == target {
				ans = append(ans, path)
			}
			return
		}
		if node.Left != nil {
			val := make([]int, len(path))
			copy(val, path)
			val = append(val, node.Left.Val)
			searchPath(node.Left, num+node.Left.Val, val)
		}
		if node.Right != nil {
			val := make([]int, len(path))
			copy(val, path)
			val = append(val, node.Right.Val)
			searchPath(node.Right, num+node.Right.Val, val)
		}
	}
	searchPath(root, root.Val, []int{root.Val})
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

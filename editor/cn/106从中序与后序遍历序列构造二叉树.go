// 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并
// 返回这颗 二叉树 。
//
//
//
// 示例 1:
//
//
// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 输出：[3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
// 输入：inorder = [-1], postorder = [-1]
// 输出：[-1]
//
//
//
//
// 提示:
//
//
// 1 <= inorder.length <= 3000
// postorder.length == inorder.length
// -3000 <= inorder[i], postorder[i] <= 3000
// inorder 和 postorder 都由 不同 的值组成
// postorder 中每一个值都在 inorder 中
// inorder 保证是树的中序遍历
// postorder 保证是树的后序遍历
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 744 👎 0

package main

import "fmt"

func main() {
	result := buildTree([]int{1, 2, 3}, []int{2, 3, 4})
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	postorderLen := len(postorder)
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[postorderLen-1]}
	postorder = postorder[:postorderLen-1]
	for k, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(inorder[:k], postorder[:len(inorder[:k])])
			root.Right = buildTree(inorder[k+1:], postorder[len(inorder[:k]):])
		}
	}
	return root
}

// leetcode submit region end(Prohibit modification and deletion)
/*
解题思路：通过中序遍历确定左右区间，没有中序无法处理。
首先通过中序遍历的结果从后往前处理，
从前序遍历找到中序遍历的位置，然后递归依次处理。
*/

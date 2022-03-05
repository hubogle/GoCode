// 给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
//
//
// 示例 1：
//
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[15,7],[9,20],[3]]
//
//
// 示例 2：
//
//
// 输入：root = [1]
// 输出：[[1]]
//
//
// 示例 3：
//
//
// 输入：root = []
// 输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 广度优先搜索 二叉树 👍 540 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	result := levelOrderBottom(new(TreeNode))
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	var val []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		n := stack.Len()
		val = []int{}
		for n > 0 {
			node := stack.Remove(stack.Front()).(*TreeNode)
			val = append(val, node.Val)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			n--
		}
		ans = append(ans, val)
	}
	n := len(ans)
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
正常层序遍历，将最后的结果从中间进行反转。
*/

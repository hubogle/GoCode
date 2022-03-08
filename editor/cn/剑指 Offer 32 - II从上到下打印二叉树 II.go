// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
//
//
//
// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层次遍历结果：
//
// [
//  [3],
//  [9,20],
//  [15,7]
// ]
//
//
//
//
// 提示：
//
//
// 节点总数 <= 1000
//
//
// 注意：本题与主站 102 题相同：https://leetcode-cn.com/problems/binary-tree-level-order-
// traversal/
// Related Topics 树 广度优先搜索 二叉树 👍 187 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	result := levelOrder(new(TreeNode))
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	var queue *list.List
	queue = list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		n := queue.Len()
		val := make([]int, 0, n)
		for ; n > 0; n-- {
			node := queue.Remove(queue.Front()).(*TreeNode)
			val = append(val, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans = append(ans, val)
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

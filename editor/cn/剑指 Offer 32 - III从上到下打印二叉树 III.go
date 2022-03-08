// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
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
//  [20,9],
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
// Related Topics 树 广度优先搜索 二叉树 👍 188 👎 0

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
	var queue *list.List
	var node *TreeNode
	var ans [][]int
	queue = list.New()
	queue.PushBack(root)
	for n := queue.Len(); n > 0; n = queue.Len() {
		val := make([]int, 0, n*2)
		for ; n > 0; n-- { // 反向存储而不是读取
			node = queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			val = append(val, node.Val)
		}
		if len(ans)&0x01 == 1 { // 奇数反转
			reverse(val)
		}
		ans = append(ans, val)
	}
	return ans
}
func reverse(val []int) {
	for i, j := 0, len(val)-1; i < j; i, j = i+1, j-1 {
		val[i], val[j] = val[j], val[i]
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue *list.List
	var node *TreeNode
	var flag int // 偶数往右
	var ans [][]int
	queue = list.New()
	queue.PushBack(root)
	for n := queue.Len(); n > 0; {
		val := make([]int, 0, n*2)
		stack := make([]*TreeNode, 0, n*2)
		for n > 0 { // 反向存储而不是读取
			if flag%2 == 1 {
				node = queue.Remove(queue.Back()).(*TreeNode)
				if node.Right != nil {
					stack = append(stack, node.Right)
				}
				if node.Left != nil {
					stack = append(stack, node.Left)
				}
			} else {
				node = queue.Remove(queue.Front()).(*TreeNode)
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			}
			val = append(val, node.Val)
			n--
		}
		flag++
		for m := len(stack) - 1; m >= 0; m-- {
			queue.PushBack(stack[m])
		}
		ans = append(ans, val)
		n = queue.Len()
	}
	return ans
}

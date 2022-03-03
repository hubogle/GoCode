// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
//
//
// 示例 2：
//
//
// 输入：root = []
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：root = [1]
// 输出：[1]
//
//
// 示例 4：
//
//
// 输入：root = [1,2]
// 输出：[2,1]
//
//
// 示例 5：
//
//
// 输入：root = [1,null,2]
// 输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1288 👎 0

package main

import (
	"container/list"
)

func main() {
	// result := inorderTraversal()
	// fmt.Println(result)
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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int // 结果集
	var node *TreeNode
	stack := list.New() // 栈
	stack.PushBack(root)
	for stack.Len() > 0 {
		e := stack.Back() // 出栈顺序：左中右
		stack.Remove(e)
		if e.Value == nil { // 如果为空，则表明是需要处理中间节点
			node = stack.Remove(stack.Back()).(*TreeNode) // 删除元素，并加入到结果集合
			ans = append(ans, node.Val)                   // 将中间节点加入到结果集中
		} else {
			node = e.Value.(*TreeNode)
			// 压栈顺序：右中左
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			stack.PushBack(node) // 中间节点压栈后再压入nil作为中间节点的标志符
			stack.PushBack(nil)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func inorderTraversal1(root *TreeNode) []int {
	var ans []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		ans = append(ans, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return ans
}

func inorderTraversal2(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return nil
	}
	stack := list.New()
	cur := root
	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return ans
}

/*
1. 迭代法
	迭代到最左元素，然后将元素从栈中弹出
2. 统一迭代法
	无法同时解决访问节点（遍历节点）和处理节点（将元素放进结果集）不一致的情况。
	那我们就将访问的节点放入栈中，把要处理的节点也放入栈中但是要做标记。
	如何标记呢，就是要处理的节点放入栈之后，紧接着放入一个空指针作为标记。 这种方法也可以叫做标记法。
*/

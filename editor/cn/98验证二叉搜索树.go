// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
// 有效 二叉搜索树定义如下：
//
//
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
//
//
// 示例 1：
//
//
// 输入：root = [2,1,3]
// 输出：true
//
//
// 示例 2：
//
//
// 输入：root = [5,1,4,null,null,3,6]
// 输出：false
// 解释：根节点的值是 5 ，但是右子节点的值是 4 。
//
//
//
//
// 提示：
//
//
// 树中节点数目范围在[1, 10⁴] 内
// -2³¹ <= Node.val <= 2³¹ - 1
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 1451 👎 0

package main

import (
	"fmt"
)

func main() {
	result := isValidBST(new(TreeNode))
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var pre *TreeNode
	var stack []*TreeNode
	stack = append(stack, root)
	for n := len(stack); n > 0; n = len(stack) {
		node := stack[n-1]
		if node != nil {
			stack = stack[:n-1]
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			node = stack[n-2]
			stack = stack[:n-2]
			if pre != nil && pre.Val >= node.Val {
				return false
			}
			pre = node
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

// 迭代，中序遍历。判断数组有序
func isValidBST1(root *TreeNode) bool {
	var stack []*TreeNode
	var ans []int
	stack = append(stack, root)
	for n := len(stack); n > 0; n = len(stack) {
		node := stack[n-1]
		stack = stack[:n-1]
		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			node = stack[n-2]
			stack = stack[:n-2]
			ans = append(ans, node.Val)
		}
	}
	for i := 0; i < len(ans)-1; i++ {
		if ans[i+1] < ans[i] {
			return false
		}
	}
	return true
}

// 递归写法
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var valid func(node *TreeNode, min, max int64) bool
	valid = func(node *TreeNode, min, max int64) bool {
		if node == nil {
			return true
		}
		if int64(node.Val) < min || int64(node.Val) > max {
			return false
		}
		return valid(node.Left, min, int64(node.Val)) && valid(node.Right, int64(node.Val), max)
	}
	return valid(root, int64(^uint64(0)>>1), int64(^uint64(0)>>1))
}

// 递归写法，使用临时节点存储上一个值
func isValidBST3(root *TreeNode) bool {
	var pre *TreeNode
	var valid func(node *TreeNode) bool
	valid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		left := valid(node.Left)
		if pre != nil && pre.Val >= node.Val {
			return false
		} // 存储上一个节点，因为是前序遍历，所以只存在 >= 的情况
		pre = node
		right := valid(node.Right)
		return left && right
	}
	return valid(root)
}

// 迭代遍历
func isValidBST4(root *TreeNode) bool {
	var pre *TreeNode
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && pre.Val >= root.Val {
				return false
			}
			pre = root
			root = root.Right
		}
	}
	return true
}

/*
题目：判断是否为二叉查找树
解题思路：二叉树中序遍历是有序数组的特性，将二叉树有序排序后。
解题思路2：利用前序遍历的特性，先遍历的最小值，将一个节点存储上一个节点的值，永远是节点的最小值

递归易错误区：
	1. 不能单纯的比较左节点小于中间节点，右节点大于中间节点就完事了。
	2. 最小节点 可能是int的最小值，如果这样使用最小的int来比较也是不行的。取 64 位的最大值

递归写法：
	1. 存储上一个节点，上一个节点为最小节点。
	2. 这样的话就是中序遍历，上一个就是左边就是最小的值。
	3. 正常的中序遍历，只不过存储上一个节点， pre 不能为 nil
*/

// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
// 示例 1：
//
//
// 输入：root = [1,2,3,null,5]
// 输出：["1->2->5","1->3"]
//
//
// 示例 2：
//
//
// 输入：root = [1]
// 输出：["1"]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 100] 内
// -100 <= Node.val <= 100
//
// Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 670 👎 0

package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := binaryTreePaths(new(TreeNode))
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
func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	var stack []*TreeNode
	var paths []string
	stack = append(stack, root)
	paths = append(paths, "")
	for n := len(stack); n > 0; {
		node := stack[n-1]
		path := paths[n-1]
		stack = stack[:n-1]
		paths = paths[:n-1]
		if node.Left == nil && node.Right == nil {
			ans = append(ans, path+strconv.Itoa(node.Val))
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
		n = len(stack)
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 递归写法：前序遍历
func binaryTreePaths1(root *TreeNode) []string {
	var ans []string
	var treePath func(node *TreeNode, path string)
	treePath = func(node *TreeNode, path string) {
		if node.Left == nil && node.Right == nil {
			ans = append(ans, path+strconv.Itoa(node.Val))
			return
		}
		if node.Left != nil {
			treePath(node.Left, path+strconv.Itoa(node.Val)+"->") // 隐藏回溯的写法
		}
		if node.Right != nil {
			treePath(node.Right, path+strconv.Itoa(node.Val)+"->")
		}
	}
	treePath(root, "")
	return ans
}

/*
题目：遍历二叉树所有到达叶子节点的路径
# 递归解题
前序遍历写法
1. 递归函数参数和返回值：参数为节点，path 为路径
2. 结束条件：左右节点都为 nil 时
3. 单层递归逻辑：依次遍历左右节点

# 迭代解题
1. 存储 node 的 stack 和 存储 path 的 stack
2. 正常的前序遍历 Tree 树的顺序，携带着添加 path
*/

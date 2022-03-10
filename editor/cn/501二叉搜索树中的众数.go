// 给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
//
// 如果树中有不止一个众数，可以按 任意顺序 返回。
//
// 假定 BST 满足如下定义：
//
//
// 结点左子树中所含节点的值 小于等于 当前节点的值
// 结点右子树中所含节点的值 大于等于 当前节点的值
// 左子树和右子树都是二叉搜索树
//
//
//
//
// 示例 1：
//
//
// 输入：root = [1,null,2,2]
// 输出：[2]
//
//
// 示例 2：
//
//
// 输入：root = [0]
// 输出：[0]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
//
//
//
//
// 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 418 👎 0

package main

import "fmt"

func main() {
	result := findMode(new(TreeNode))
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
func findMode(root *TreeNode) []int {
	var ans []int
	var flag int      // 累计当前相同数量
	var maxNum int    // 最大众数数量
	var pre *TreeNode // 上一个指针
	var node *TreeNode
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre == nil || pre.Val != node.Val {
				flag = 1
			} else {
				flag++
			}
			if flag == maxNum {
				ans = append(ans, node.Val)
			} else if flag > maxNum {
				maxNum = flag
				ans = []int{node.Val}
			}
			pre = node
			root = node.Right
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：二叉搜索树中的众数
解题思路：中序遍历二叉树，用一个指针保存上一个节点。
第一次遇到就是 `1` 然后中序遍历，如果相同则数量累加。
上一个指针有三种情况：
1. nil flag = 1
2. pre.Val = node.Val flag++
3. pre.Val != node.Val flag = 1

因为众数不止有一个所以，每当遇到数量更多的众数，之前的统计就要清空。
*/

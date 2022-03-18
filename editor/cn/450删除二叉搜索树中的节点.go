// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的
// 根节点的引用。
//
// 一般来说，删除节点可分为两个步骤：
//
//
// 首先找到需要删除的节点；
// 如果找到了，删除它。
//
//
//
//
// 示例 1:
//
//
//
//
// 输入：root = [5,3,6,2,4,null,7], key = 3
// 输出：[5,4,6,2,null,null,7]
// 解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
// 一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
// 另一个正确答案是 [5,2,6,null,4,null,7]。
//
//
//
//
// 示例 2:
//
//
// 输入: root = [5,3,6,2,4,null,7], key = 0
// 输出: [5,3,6,2,4,null,7]
// 解释: 二叉树不包含值为 0 的节点
//
//
// 示例 3:
//
//
// 输入: root = [], key = 0
// 输出: []
//
//
//
// 提示:
//
//
// 节点数的范围 [0, 10⁴].
// -10⁵ <= Node.val <= 10⁵
// 节点值唯一
// root 是合法的二叉搜索树
// -10⁵ <= key <= 10⁵
//
//
//
//
// 进阶： 要求算法时间复杂度为 O(h)，h 为树的高度。
// Related Topics 树 二叉搜索树 二叉树 👍 666 👎 0

package main

import "fmt"

func main() {
	result := deleteNode(new(TreeNode), 0)
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	var pre *TreeNode // 删除节点的父节点
	var cur *TreeNode
	cur = root
	for cur != nil {
		if cur.Val == key {
			break
		}
		pre = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if pre == nil { // 头节点为目标删除
		return deleteN(root)
	}
	if pre.Left != nil && pre.Left.Val == key {
		pre.Left = deleteN(pre.Left)
	}
	if pre.Right != nil && pre.Right.Val == key {
		pre.Right = deleteN(pre.Right)
	}
	return root
}

func deleteN(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	if node.Right == nil {
		return node.Left
	} // 右节点为空直接返回左节点
	cur := node.Right
	for cur.Left != nil {
		cur = cur.Left
	} // 右节点不为空，一直寻找到右节点的最左节点。
	cur.Left = node.Left // 将删除节点左孩子，插入到右孩子的最左侧
	return node.Right
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：删除二搜索树叉树的目标节点
1. 首先找到二叉搜索树要删除的目标节点（额外考虑删除根节点），和删除节点的父节点。
2. 判断删除删除节点在父节点的左侧还是右侧；
3. 移除节点时的判断
	3.1 若右孩子为空，则直接返回左孩子
	3.2 右孩子不为空，则寻找右孩子的最左孩子，将删除节点的左孩子放在这里。
*/

// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
//
// B是A的子结构， 即 A中有出现和B相同的结构和节点值。
//
// 例如:
// 给定的树 A:
//
// 3
// / \
// 4 5
// / \
// 1 2
// 给定的树 B：
//
// 4
// /
// 1
// 返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
//
// 示例 1：
//
// 输入：A = [1,2,3], B = [3,1]
// 输出：false
//
//
// 示例 2：
//
// 输入：A = [3,4,5,1,2], B = [4,1]
// 输出：true
//
// 限制：
//
// 0 <= 节点个数 <= 10000
// Related Topics 树 深度优先搜索 二叉树 👍 520 👎 0

package main

import "fmt"

func main() {
	result := isSubStructure(new(TreeNode), new(TreeNode))
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
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	var exist func(node1 *TreeNode, node2 *TreeNode) bool
	exist = func(node1 *TreeNode, node2 *TreeNode) bool {
		if node2 == nil {
			return true
		}
		if node1 == nil {
			return false
		}
		if node1.Val == node2.Val {
			return exist(node1.Left, node2.Left) && exist(node1.Right, node2.Right)
		} else {
			return false
		}
	}

	var search func(node *TreeNode) bool
	search = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if node.Val == B.Val && exist(node, B) {
			return true
		}
		return search(node.Left) || search(node.Right)
	}
	return search(A)
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：深度优先遍历
没办法同时遍历 A、B 来获取是否为子树
只有先遍历 A 遇到 B 的 root 节点，判断 A 和 B 子树是否一致
需要两个遍历，一个遍历 A 所有节点，另一个遍历 A子节点是否等于 B树
*/

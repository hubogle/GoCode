// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//
//
// 示例 1：
//
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
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
// -100 <= Node.val <= 100
//
// Related Topics 树 广度优先搜索 二叉树 👍 606 👎 0

package main

import "fmt"

func main() {
	result := zigzagLevelOrder(new(TreeNode))
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	var ans [][]int
	if root == nil {
		return ans
	}
	queue = append(queue, root)
	for n, level := len(queue), 0; n > 0; n, level = len(queue), level+1 {
		var val []int
		for i := 0; i < n; i++ {
			if queue[i] != nil {
				val = append(val, queue[i].Val)
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
			}
		}
		if len(val) == 0 {
			break
		}
		if level&0x1 == 1 {
			for l, r := 0, len(val)-1; l < r; l, r = l+1, r-1 {
				val[l], val[r] = val[r], val[l]
			}
		}
		ans = append(ans, val)
		queue = queue[n:]
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

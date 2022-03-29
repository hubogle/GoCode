// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
//
//
// 示例 3：
//
//
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
// Related Topics 递归 链表 数学 👍 8016 👎 0

package main

import "fmt"

func main() {
	result := addTwoNumbers(new(ListNode), new(ListNode))
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	v1, v2, val, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || val != 0 {
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (v1 + v2 + val) % 10}
		current = current.Next
		val = (v1 + v2 + val) / 10
	}
	return head.Next
}

// leetcode submit region end(Prohibit modification and deletion)
/*
2.两数相加
新建立一个虚拟节点，这样虚拟节点的 Next 指向真的 head 最后返回 head.Next 即可。
主要是建立虚拟节点不用单独判断，循环终止条件是 p != nil
主要是两个链表进行相加操作，依次从低位相加到高位置，每次预留高位的进一位值 val = (v1+v2+val) / 10
*/

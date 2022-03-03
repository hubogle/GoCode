// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
// 示例1：
//
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
//
// 限制：
//
// 0 <= 链表长度 <= 1000
//
// 注意：本题与主站 21 题相同：https://leetcode-cn.com/problems/merge-two-sorted-lists/
// Related Topics 递归 链表 👍 217 👎 0

package main

import (
	"fmt"
)

func main() {
	result := mergeTwoLists()
	fmt.Println(result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	empty := new(ListNode)
	head := empty
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			empty.Next = l1
			l1 = l1.Next
		} else {
			empty.Next = l2
			l2 = l2.Next
		}
		empty = empty.Next
	}
	if l1 != nil {
		empty.Next = l1
	} else {
		empty.Next = l2
	}
	return head.Next
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解体思路：引入伪头节点，方便后续判断
双指针 l,r 遍历两链表，根据大小关系添加顺序，两节点交替前进，直到遍历完成
*/

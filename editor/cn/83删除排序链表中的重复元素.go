// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
//
//
//
// 示例 1：
//
//
// 输入：head = [1,1,2]
// 输出：[1,2]
//
//
// 示例 2：
//
//
// 输入：head = [1,1,2,3,3]
// 输出：[1,2,3]
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序 排列
//
// Related Topics 链表 👍 751 👎 0

package main

import "fmt"

func main() {
	result := deleteDuplicates(new(ListNode))
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	cur := head.Next
	for pre != nil {
		if cur == nil {
			pre.Next = cur
			break
		}
		if cur.Val == pre.Val {
			cur = cur.Next
			continue
		}
		pre.Next = cur
		pre = cur
		cur = cur.Next
	}
	return head
}

// leetcode submit region end(Prohibit modification and deletion)
/*
解题思路：
注意不能 for 循环 cur 指针，这样遇到最后是连续重复的话，无法去重复。
*/

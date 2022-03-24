// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//
//
//
// 示例 1：
//
//
// 输入：head = [1,2,3,3,4,4,5]
// 输出：[1,2,5]
//
//
// 示例 2：
//
//
// 输入：head = [1,1,1,2,3]
// 输出：[2,3]
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
// Related Topics 链表 双指针 👍 846 👎 0

package main

import "fmt"

func main() {
	// result := deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, nil}}}, )
	result := deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, nil}}})
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := &ListNode{0, head}
	cur := tmp
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val { // 判断下一个与下下一个是否相等
			x := cur.Next.Val
			if cur.Next != nil && cur.Next.Val == x { // 记录重复值，一直跳过无重复值
				cur.Next = cur.Next.Next
			}
		} else { // 不相等就是无重复
			cur = cur.Next
		}
	}
	return tmp.Next
}

// leetcode submit region end(Prohibit modification and deletion)
/*
解题思路：一次遍历
创建临时节点
cur 遍历 cur.Next 和 cur.Next.Next 是否相等，相等的话 for 循环直到 cur.Next 指向下一个非重复值
*/

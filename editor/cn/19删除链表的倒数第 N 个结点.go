// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
//
//
// 示例 2：
//
//
// 输入：head = [1], n = 1
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：head = [1,2], n = 1
// 输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
// Related Topics 链表 双指针 👍 1904 👎 0

package main

import "fmt"

func main() {
	result := removeNthFromEnd(new(ListNode), 1)
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
// 1,2,3,4,5,6
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{0, head}
	r := head
	l := tmp
	for i := 0; i < n; i++ {
		r = r.Next
	}
	for r != nil {
		r = r.Next
		l = l.Next
	}
	l.Next = l.Next.Next
	return tmp.Next
}

// leetcode submit region end(Prohibit modification and deletion)
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	tmp := &ListNode{0, head}
	r := head
	l := tmp
	for i := 0; i < n; i++ {
		r = r.Next
	}
	for l != nil {
		if r == nil {
			l.Next = l.Next.Next
			break
		}
		r = r.Next
		l = l.Next
	}
	return tmp.Next
}

/*
解题思路：双指针
让右指针快走一个。
当右指针到 nil 时，l.Next = l.Next.Next
其它情况 l,r 都右移
*/

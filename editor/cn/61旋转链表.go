// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
//
//
// 示例 1：
//
//
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
//
//
// 示例 2：
//
//
// 输入：head = [0,1,2], k = 4
// 输出：[2,0,1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 10⁹
//
// Related Topics 链表 双指针 👍 767 👎 0

package main

import "fmt"

func main() {
	result := rotateRight(&ListNode{Val: 1, Next: nil}, 0)
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	} // 有多少节点
	mov := n - k%n // 需要移动的位置
	if mov == n {
		return head
	}
	cur.Next = head // 将尾节点和头节点链接
	for mov > 0 {
		cur = cur.Next
		mov--
	} // 遍历到新的尾节点
	head = cur.Next
	cur.Next = nil
	return head
}

// leetcode submit region end(Prohibit modification and deletion)

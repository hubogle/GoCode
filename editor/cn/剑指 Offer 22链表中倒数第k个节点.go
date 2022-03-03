// 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
// 例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
//
//
//
// 示例：
//
//
// 给定一个链表: 1->2->3->4->5, 和 k = 2.
//
// 返回链表 4->5.
// Related Topics 链表 双指针 👍 330 👎 0

package main

import "fmt"

func main() {
	node2 := ListNode{Val: 2}
	node1 := ListNode{Val: 1, Next: &node2}
	node0 := ListNode{Val: 0, Next: &node1}
	result := getKthFromEnd(&node0, 2)
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
func getKthFromEnd(head *ListNode, k int) *ListNode {
	l, r := head, head
	for i := 0; i < k; i++ {
		r = r.Next
	} // 快慢指针间隔 k 个元素
	for r != nil {
		r = r.Next
		l = l.Next
	}
	return l
}

// leetcode submit region end(Prohibit modification and deletion)
/*
解题思路：双指针移动
时间复杂度：O(n)
空间复杂度：O(1)
快指针指向链表 `k+1` 个节点，慢指针与快指针间隔 `k` 个节点。
两个指针同步往后移动，当快指针到达尾部节点，慢指针正好到链表倒数 `k` 个节点。
*/

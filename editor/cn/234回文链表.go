// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
// 输入：head = [1,2,2,1]
// 输出：true
//
//
// 示例 2：
//
//
// 输入：head = [1,2]
// 输出：false
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围[1, 10⁵] 内
// 0 <= Node.val <= 9
//
//
//
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
// Related Topics 栈 递归 链表 双指针 👍 1309 👎 0

package main

import (
	"fmt"
)

func main() {
	result := isPalindrome(new(ListNode))
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
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	low := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}
	end := low
	endStart := reverse(end.Next)
	start := head
	for endStart != nil {
		if start.Val != endStart.Val {
			return false
		}
		start = start.Next
		endStart = endStart.Next
	}
	return true
}

func reverse(node *ListNode) *ListNode {
	var tmp *ListNode
	head := node
	for head != nil {
		val := head.Next
		head.Next = tmp
		tmp = head
		head = val
	}
	return tmp
}

// leetcode submit region end(Prohibit modification and deletion)
/*
解题思路：
快慢指针，确定中间节点，将后半部分链表进行反转。
与 环形链表 的区别在于，快慢指针判断时环链表 for low != fast ，快慢指针写法不能相同
一个指针在头部，一个指针在尾部，同时遍历遇到不同的值直接返回 false
*/

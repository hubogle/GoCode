// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 进阶：
//
//
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
//
// 示例 1：
//
//
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]
//
//
// 示例 3：
//
//
// 输入：head = [1,2,3,4,5], k = 1
// 输出：[1,2,3,4,5]
//
//
// 示例 4：
//
//
// 输入：head = [1], k = 1
// 输出：[1]
//
//
//
//
//
// 提示：
//
//
// 列表中节点的数量在范围 sz 内
// 1 <= sz <= 5000
// 0 <= Node.val <= 1000
// 1 <= k <= sz
//
// Related Topics 递归 链表 👍 1524 👎 0

package main

import "fmt"

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	result := reverseKGroup(node, 2)
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	tmp := &ListNode{Next: head}
	pre := tmp
	for head != nil {
		tail := pre // 寻找反转的尾部节点
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return tmp.Next
			}
		} // 反转位置
		next := tail.Next                // 尾部节点的下一个
		head, tail = reverse(head, tail) // 反转
		pre.Next = head                  // 链接上一个
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return tmp.Next
}

// 反转头尾节点
func reverse(start, end *ListNode) (*ListNode, *ListNode) {
	pre := end.Next
	cur := start
	for pre != end {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return end, start
}

// leetcode submit region end(Prohibit modification and deletion)

func reverseKGroup1(head *ListNode, k int) *ListNode {
	var cur *ListNode
	var pre *ListNode
	var ans int
	temp := &ListNode{Next: head}
	cur = head
	pre = temp
	for cur != nil {
		stack := make([]*ListNode, 0, k)
		for ans != k && cur != nil { // 栈里存储足够反转的节点
			stack = append(stack, cur)
			cur = cur.Next
			ans++
		}
		if len(stack) == k { // 满足数量进行反转
			stack[0].Next = cur
			for i := len(stack) - 1; i >= 0; i-- {
				pre.Next = stack[i]
				pre = pre.Next
			}
			ans = 0
		}
	}
	return temp.Next
}

/*
解题思路：
如何寻找反转节点的开始和结束位置。
以及保存开始节点 pre 上个位置，以及尾部节点的开始搜索位置。

		next := tail.Next                // 尾部节点的下一个
		head, tail = reverse(head, tail) // head 和 tail 位置反转
		pre.Next = head                  // pre 保存开始节点上个位置，连接头部节点
		tail.Next = next                 // 连接尾部节点
		pre = tail
		head = tail.Next

*/

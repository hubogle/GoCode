// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//
//
//
//
//
// 示例 1：
//
//
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
//
//
// 示例 2：
//
//
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]
//
//
// 示例 3：
//
//
// 输入：head = []
// 输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 5 * 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
//
//
//
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
// Related Topics 链表 双指针 分治 排序 归并排序 👍 1592 👎 0

package main

import "fmt"

func main() {
	result := sortList(new(ListNode))
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
func sortList(head *ListNode) *ListNode {
	return MergeSort(head)
}

func MergeSort(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	slow := node
	fast := node
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	node2 := slow.Next
	slow.Next = nil
	node1 := node
	return Merge(MergeSort(node1), MergeSort(node2))
}

func Merge(node1 *ListNode, node2 *ListNode) *ListNode {
	tmp := new(ListNode)
	head := tmp
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			tmp.Next = node1
			node1 = node1.Next
		} else {
			tmp.Next = node2
			node2 = node2.Next
		}
		tmp = tmp.Next
	}
	if node1 == nil && node2 == nil {
		return head.Next
	}
	if node1 == nil {
		tmp.Next = node2
	} else {
		tmp.Next = node1
	}
	return head.Next
}

// leetcode submit region end(Prohibit modification and deletion)
/*
解题思路：归并排序
首先快慢指针寻找到中间节点，从改节点一分为二。
创建 merge 拼接两个链表合并为一个链表
*/

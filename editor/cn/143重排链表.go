// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//
// L0 → L1 → … → Ln - 1 → Ln
//
//
// 请将其重新排列后变为：
//
//
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//
//
// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]
//
// 示例 2：
//
//
//
//
// 输入：head = [1,2,3,4,5]
// 输出：[1,5,2,4,3]
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 5 * 10⁴]
// 1 <= node.val <= 1000
//
// Related Topics 栈 递归 链表 双指针 👍 837 👎 0

package main

import "fmt"

func main() {
	result := reorderList(new(ListNode))
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
func reorderList(head *ListNode) {
	low, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}
	start := low // 开始反转的位置
	right := reverse(start.Next)
	start.Next = nil // 注意这里，链表找到中间位置后
	left := head
	for right != nil {
		leftNext := left.Next
		rightNext := right.Next
		left.Next = right
		right.Next = leftNext
		left = leftNext
		right = rightNext
	}
}
func reverse(node *ListNode) *ListNode {
	var tmp *ListNode
	for node != nil {
		next := node.Next
		node.Next = tmp
		tmp = node
		node = next
	}
	return tmp
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：
快慢指针，找到中间节点，反转后半部节点。
注意左边部分最后一个节点，是右边反转后的最后一个节点，要将左边最后一个节点设为 nil
左右节点同时遍历，右半部分链表节点数量少。
*/

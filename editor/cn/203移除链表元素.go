// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
//
//
// 示例 1：
//
//
// 输入：head = [1,2,6,3,4,5,6], val = 6
// 输出：[1,2,3,4,5]
//
//
// 示例 2：
//
//
// 输入：head = [], val = 1
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：head = [7,7,7,7], val = 7
// 输出：[]
//
//
//
//
// 提示：
//
//
// 列表中的节点数目在范围 [0, 10⁴] 内
// 1 <= Node.val <= 50
// 0 <= val <= 50
//
// Related Topics 递归 链表 👍 812 👎 0

package main

import "fmt"

func main() {
	node3 := ListNode{3, nil}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	node0 := ListNode{0, &node1}
	result := removeElements(&node0, 2)
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
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
func removeElements(head *ListNode, val int) *ListNode {
	ans := new(ListNode) // 新建一个虚拟头结点
	ans.Next = head
	cur := ans
	for cur.Next != nil && cur != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return ans.Next
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：移除链表中的指定元素
新建一个虚拟节点，这样的话处理逻辑可以统一。
如果创建虚拟节点的话，要单独处理 head 节点。
*/

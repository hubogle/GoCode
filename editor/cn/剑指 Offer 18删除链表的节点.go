// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
// 返回删除后的链表的头节点。
//
// 注意：此题对比原题有改动
//
// 示例 1:
//
// 输入: head = [4,5,1,9], val = 5
// 输出: [4,1,9]
// 解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
//
//
// 示例 2:
//
// 输入: head = [4,5,1,9], val = 1
// 输出: [4,5,9]
// 解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
//
//
//
//
// 说明：
//
//
// 题目保证链表中节点的值互不相同
// 若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点
//
// Related Topics 链表 👍 195 👎 0

package main

import "fmt"

func main() {
	node2 := ListNode{Val: 2}
	node1 := ListNode{Val: 1, Next: &node2}
	node0 := ListNode{Val: 0, Next: &node1}
	result := deleteNode(&node0, 1)
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
func deleteNode(head *ListNode, val int) *ListNode {
	// if head.Val == val {
	// 	return head.Next
	// }
	// pre, cur := head, head.Next
	// for cur != nil && cur.Val != val {
	// 	pre, cur = cur, cur.Next
	// }
	// pre.Next = cur.Next
	empty := new(ListNode)
	empty.Next = head
	pre, cur := empty, empty.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			return empty.Next
		} else {
			pre, cur = cur, cur.Next
		}
	}
	return empty.Next
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解题思路：
定义空指针头，指向 head 这样可以统一跳过逻辑
然后定义当前指针，和下一个指针，判断下一个指针指向的值是否为 val ，判断是否跳过该指针。
注意：定义两个指针的指向起始位置，empty 和 empty.Next
*/

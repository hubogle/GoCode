// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
// 示例 1：
//
// 输入：head = [1,3,2]
// 输出：[2,3,1]
//
//
//
// 限制：
//
// 0 <= 链表长度 <= 10000
// Related Topics 栈 递归 链表 双指针 👍 241 👎 0

package main

import "fmt"

func main() {
	node3 := ListNode{3, nil}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	result := reversePrint(&node1)
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

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var node *ListNode // 定义空指针，用作哨兵
	for head != nil {  // 这里不能这样判断 head.Next != nil
		nextNode := head.Next
		head.Next = node // 指向上一个
		node = head
		head = nextNode
	}
	res := []int{}
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)

/*
解析思路：链表反转，用一个哨兵简单化代码
⚠️ 链表循环时终止条件是 `head != nil` 不要进行 `head.Next != nil` 判断

*/

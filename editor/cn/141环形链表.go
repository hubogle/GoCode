// 给你一个链表的头节点 head ，判断链表中是否有环。
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到
// 链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
//
// 如果链表中存在环 ，则返回 true 。 否则，返回 false 。
//
//
//
// 示例 1：
//
//
//
//
// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。
//
//
// 示例 2：
//
//
//
//
// 输入：head = [1,2], pos = 0
// 输出：true
// 解释：链表中有一个环，其尾部连接到第一个节点。
//
//
// 示例 3：
//
//
//
//
// 输入：head = [1], pos = -1
// 输出：false
// 解释：链表中没有环。
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 10⁴]
// -10⁵ <= Node.val <= 10⁵
// pos 为 -1 或者链表中的一个 有效索引 。
//
//
//
//
// 进阶：你能用 O(1)（即，常量）内存解决此问题吗？
// Related Topics 哈希表 链表 双指针 👍 1408 👎 0

package main

import "fmt"

func main() {
	result := hasCycle(new(ListNode))
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

func hasCycle(head *ListNode) bool {
	fast := head
	low := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		low = low.Next
		if fast == low {
			return true
		}
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)

// 用 map 标记是否访问过
func hasCycle1(head *ListNode) bool {
	cache := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := cache[head]; ok {
			return true
		}
		cache[head] = struct{}{}
		head = head.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	low := head
	fast := head.Next
	for low != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		low = low.Next
		fast = fast.Next.Next
	}
}

/*
	快慢指针理解：
	初始化慢指针在 head，快指针在 head.Next
	慢指针每次移动一个，快指针每次移动两个。
	循环条件：low != fast 终止条件 fast.Next != nil && fast.Next.Next != nil

	为什么快慢初始化不在同一个，幻想head 前面有一个临时指针，for 循环下一步的情况就是 head 和 head.Next
*/

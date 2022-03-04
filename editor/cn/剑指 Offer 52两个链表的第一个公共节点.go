// 输入两个链表，找出它们的第一个公共节点。
//
// 如下面的两个链表：
//
//
//
// 在节点 c1 开始相交。
//
//
//
// 示例 1：
//
//
//
// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2,
// skipB = 3
// 输出：Reference of the node with value = 8
// 输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1
// ,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
//
//
//
//
// 示例 2：
//
//
//
// 输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB =
// 1
// 输出：Reference of the node with value = 2
// 输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4
// ]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
//
//
//
//
// 示例 3：
//
//
//
// 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// 输出：null
// 输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而
// skipA 和 skipB 可以是任意值。
// 解释：这两个链表不相交，因此返回 null。
//
//
//
//
// 注意：
//
//
// 如果两个链表没有交点，返回 null.
// 在返回结果后，两个链表仍须保持原有的结构。
// 可假定整个链表结构中没有循环。
// 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
// 本题与主站 160 题相同：https://leetcode-cn.com/problems/intersection-of-two-linked-
// lists/
//
// Related Topics 哈希表 链表 双指针 👍 450 👎 0

package main

func main() {
	// result := getIntersectionNode()
	// fmt.Println(result)
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	l, r := headA, headB
	for l != r {
		if l == nil {
			l = headB
		} else {
			l = l.Next
		}
		if r == nil {
			r = headA
		} else {
			r = r.Next
		}
	}
	return l
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：寻找两个链表的公共节点位置
解题思路：如果两个链表相交，想象一下，将两个链表拼接成一个环，只不过这个环有个尾巴，也就是双指针遍历两个链表的情景。
时间复杂度为 `O(n)` 的情况，只能是遍历两个链表

情况1: 长度相同
	两个指针分别同时遍历两个链表，如果同时指向同一个节点（可能为空），就直接返回。
情况2：长度不同
	两个指针分别同时遍历两个链表，短链表先遍历完，然后遍历长链表。两个指针分别遍历完长短链表，如果相交就返回节点，不相交返回 nil

*/

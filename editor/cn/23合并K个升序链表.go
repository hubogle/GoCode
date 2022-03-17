// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//  1->4->5,
//  1->3->4,
//  2->6
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
// 输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
// 输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 1818 👎 0

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	result := mergeKLists(nil)
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
type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	h := new(minHeap)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}
	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
		pre.Next = tmp
		pre = pre.Next
	}
	return dummyHead.Next
}

// leetcode submit region end(Prohibit modification and deletion)
// 顺序合并
func mergeKLists1(lists []*ListNode) *ListNode {
	var cur *ListNode
	for i := 0; i < len(lists); i++ {
		cur = merge(cur, lists[i])
	}
	return cur
}

// 分治合并 时间复杂度 O(kn * log k)
func mergeKLists2(lists []*ListNode) *ListNode {
	var mergeList func(l, r int) *ListNode
	mergeList = func(l, r int) *ListNode {
		if l == r {
			return lists[l]
		}
		if l > r {
			return nil
		}
		mid := (l + r) >> 1
		return merge(mergeList(l, mid), mergeList(mid+1, r))
	}
	return mergeList(0, len(lists)-1)
}
func merge(node1 *ListNode, node2 *ListNode) *ListNode {
	tmp := new(ListNode)
	head := tmp
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			head.Next = node1
			node1 = node1.Next
		} else {
			head.Next = node2
			node2 = node2.Next
		}
		head = head.Next
	}
	if node1 != nil {
		head.Next = node1
	}
	if node2 != nil {
		head.Next = node2
	}
	return tmp.Next
}

/*
解题思路：
* 顺序合并
依次遍历每个链表，依次将每个链表合并
* 分治合并
* 优先队列 （使用小堆实现）
*/

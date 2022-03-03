// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
//
//
// 示例 1:
//
//
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
//
//
// 示例 2:
//
//
// 输入: nums = [1], k = 1
// 输出: [1]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
//
//
//
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
// Related Topics 数组 哈希表 分治 桶排序 计数 快速选择 排序 堆（优先队列） 👍 1047 👎 0

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	result := topKFrequent(nums, 2)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)

type Heap [][2]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i][1] < h[j][1] } // 构建一个小堆，最后留下的三个就是最大值
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func topKFrequent(nums []int, k int) []int {
	ans := map[int]int{}
	for _, num := range nums {
		ans[num]++
	} // key 是数值， val 是出现的次数
	h := &Heap{}
	heap.Init(h) // 实现 堆
	for key, v := range ans {
		heap.Push(h, [2]int{key, v})
		if h.Len() > k {
			heap.Pop(h)
		} // 弹出顶部的最小值
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
// 。
//
// 返回 滑动窗口中的最大值 。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
// 输入：nums = [1], k = 1
// 输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 1432 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// nums := []int{7, 2, 4}
	result := maxSlidingWindow(nums, 3)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	var queue []int // 维护一个双端队列
	push := func(i int) {
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] { // 新进元素大于等于递增元素
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i) // 坐标是递增
	} // 维护一个坐标递增，坐标对应值递减的队列
	ans := make([]int, 1, len(nums)-k+1)
	for i := 0; i < k; i++ {
		push(i)
	}
	fmt.Println(queue)
	ans[0] = nums[queue[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		for queue[0] < i-k+1 { // 首坐标已超过滑动窗口范围，弹出
			queue = queue[1:]
		}
		ans = append(ans, nums[queue[0]])
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func maxSlidingWindow1(nums []int, k int) []int {
	ans := []int{}
	push := func(i int) {
		for len(ans) > 0 && nums[i] >= nums[ans[len(ans)-1]] {
			ans = ans[:len(ans)-1]
		}
		ans = append(ans, i)
	} // 维护一个单调递减队列
	for i := 0; i < k; i++ {
		push(i)
	}
	n := len(nums)
	res := make([]int, 1, n-k+1)
	res[0] = nums[ans[0]]
	for i := k; i < n; i++ {
		push(i)
		for ans[0] <= i-k {
			res = res[1:]
		}
		res = append(res, nums[ans[0]])
	}
	return ans
}

/*
解体思路：单调队列
时间复杂度：O(n)
空间复杂度：O(k) 使用双向队列，从队首弹出元素，保证数量不超过 k + 1 个元素

1. 需要求滑动窗口最大值，若 i < j 且 num[i] <= num[j]；
2. 当右移时，只要 i 还在窗口，那么 j 一定在窗口，因此可以将 num[i] 移除，它一定不是最大值；
3. 使用队列存储未移除的下标，下标从小到大的顺序存储，对应的 num 值是单调递减的；
4. 窗口右移时，将新元素加入队列，不断的对新的元素与队尾元素比较
5. 若新元素大于队尾元素，则队尾元素可以移除，直到队列为空或新元素小于队尾元素；
6. 队列下标对应的元素是递减的，因此队列第一个下标对应的值是最大值；
7. 此时最大值可能在左边界，随着右移，可能要队首弹出元素，直到新队首元素出现在窗口；

需要可以弹出队首和队尾元素，使用双端队列，并且这种队列叫做 【单调队列】

队列中坐标单调递增，坐标对应的值单调递减。
*/

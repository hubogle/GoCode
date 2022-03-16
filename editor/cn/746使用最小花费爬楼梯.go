// 给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
// 请你计算并返回达到楼梯顶部的最低花费。
//
//
//
// 示例 1：
//
//
// 输入：cost = [10,15,20]
// 输出：15
// 解释：你将从下标为 1 的台阶开始。
// - 支付 15 ，向上爬两个台阶，到达楼梯顶部。
// 总花费为 15 。
//
//
// 示例 2：
//
//
// 输入：cost = [1,100,1,1,1,100,1,1,100,1]
// 输出：6
// 解释：你将从下标为 0 的台阶开始。
// - 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
// - 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
// - 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
// - 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
// - 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
// - 支付 1 ，向上爬一个台阶，到达楼梯顶部。
// 总花费为 6 。
//
//
//
//
// 提示：
//
//
// 2 <= cost.length <= 1000
// 0 <= cost[i] <= 999
//
// Related Topics 数组 动态规划 👍 813 👎 0

package main

import "fmt"

func main() {
	// cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	cost := []int{10, 15, 20}
	result := minCostClimbingStairs(cost)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
// dp[i] = mix(dp[i-1], dp[i-2])
// 前两步耗费体力，后两步不耗费体力
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
		// dp[1], dp[0] = min(dp[0], dp[1])+cost[i], dp[1]
	}
	return min(dp[n-1], dp[n-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 前面两步不耗费体力，后两步耗费体力
func minCostClimbingStairs1(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = 0 // 默认第一步不花体力
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

// leetcode submit region end(Prohibit modification and deletion)

/*
题目：数组每间隔 1、2 进行相加，似的总和最小。
dp[i] = min(dp[i-1], dp[i-2]) 推导出来的
易错点注意取的是 最后一步和最后第二步的最小值

解体思路1: 常规动态规划解题
解体思路2: 因为题目存在，第一步不花费体力\最后一步不花费体力，额外申请一位数组空间存储最后一步和最后二步的最小值。gst
*/

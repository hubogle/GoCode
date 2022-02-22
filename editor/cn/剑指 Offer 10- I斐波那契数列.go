// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
//
//
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
//
// 斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
//
//
// 示例 1：
//
//
// 输入：n = 2
// 输出：1
//
//
// 示例 2：
//
//
// 输入：n = 5
// 输出：5
//
//
//
//
// 提示：
//
//
// 0 <= n <= 100
//
// Related Topics 记忆化搜索 数学 动态规划 👍 300 👎 0

package main

import "fmt"

func main() {
	result := fib(3)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 1, 0
	for i := 2; i <= n; i++ {
		a, b = a+b, a
		a = a % 1000000007
	}
	return a
}

// leetcode submit region end(Prohibit modification and deletion)
/*
注意：起始计算条件 F(n) = F(n-1) + F(n-2) n >= 2
F(2) = 1, F(1) = 1, F(0) = 0
*/

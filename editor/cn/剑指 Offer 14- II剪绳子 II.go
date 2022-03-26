// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1]
// 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘
// 积是18。
//
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
//
//
// 示例 1：
//
// 输入: 2
// 输出: 1
// 解释: 2 = 1 + 1, 1 × 1 = 1
//
// 示例 2:
//
// 输入: 10
// 输出: 36
// 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
//
//
//
// 提示：
//
//
// 2 <= n <= 1000
//
//
// 注意：本题与主站 343 题相同：https://leetcode-cn.com/problems/integer-break/
// Related Topics 数学 动态规划 👍 178 👎 0

package main

import "fmt"

func main() {
	result := cuttingRope(88)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	mod := 1000000007
	ans := 1
	for i := 1; i <= n/3-1; i++ { // 留下一个 3 和余数做最大相乘
		ans = ans * 3 % mod
	}
	switch n % 3 {
	case 0:
		return ans * 3 % mod
	case 1:
		return ans * 4 % mod
	default:
		return ans * 3 * 2 % mod
	}
}

// leetcode submit region end(Prohibit modification and deletion)
/*
解题思路：无法使用动态规划，因为会取余
如 1000000008 则会变为 1 后续的相乘就失效了
n = 4；最大 2 * 2
n = 5；最大 2 * 3

n = 6；最大 3 * 3
n = 7；最大 3 * 2 * 2
n = 8；最大 3 * 3 * 2
n = 9；最大 3 * 3 * 3
都转换为 3，与最后的余数做最大相乘。
*/

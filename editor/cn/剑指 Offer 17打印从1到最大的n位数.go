// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
//
// 示例 1:
//
// 输入: n = 1
// 输出: [1,2,3,4,5,6,7,8,9]
//
//
//
//
// 说明：
//
//
// 用返回一个整数列表来代替打印
// n 为正整数
//
// Related Topics 数组 数学 👍 189 👎 0

package main

import "fmt"

func main() {
	n := 1
	result := printNumbers(n)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func printNumbers(n int) []int {
	nums := 2
	for i := 1; i <= n; i++ {
		nums *= 10
	}
	ans := make([]int, nums-1, nums-1)
	for i := 1; i < nums; i++ {
		ans[i-1] = i
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

/*
打印数字范围，注意审题：看答案样例。
*/

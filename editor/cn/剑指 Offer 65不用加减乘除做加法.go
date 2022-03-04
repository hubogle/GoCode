// 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
//
//
//
// 示例:
//
// 输入: a = 1, b = 1
// 输出: 2
//
//
//
// 提示：
//
//
// a, b 均可能是负数或 0
// 结果不会溢出 32 位整数
//
// Related Topics 位运算 数学 👍 267 👎 0

package main

import "fmt"

func main() {
	// result := add(11, 21)
	fmt.Println(1 & 1 << 1)
	// fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func add(a int, b int) int {
	for b != 0 {
		c := a & b << 1 // 进位 +
		a = a ^ b       // 无进位
		b = c
	}
	return a
}

// leetcode submit region end(Prohibit modification and deletion)

/*
无进位求和：^(亦或) 想象10进制下的模拟情况：（如:19+1=20；无进位求和就是10，而非20；因为它不管进位情况）
进位求和：&(与) 想象10进制下模拟情况：（9+1=10，如果是用&的思路来处理，则9+1得到的进位数为1，而不是10，所以要用<<1向左再移动一位，这样就变为10了）；

这样公式就是：（a^b) ^ ((a&b)<<1) 即：每次无进位求 + 每次得到的进位数--------我们需要不断重复这个过程，直到进位数为0为止；

无进位加法 a ^ b，101 + 010 最好的理解方式就是盯住每一位
如果有进位的话，a & b << 1
没有进位表示一步到位了
有进位则再做一次加法。如100 100 产生进位 做完一次无进位加法后 000, 再做一次加法
*/

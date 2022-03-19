// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
//
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
//
//
//
// 示例 1：
//
//
// 输入：num1 = "11", num2 = "123"
// 输出："134"
//
//
// 示例 2：
//
//
// 输入：num1 = "456", num2 = "77"
// 输出："533"
//
//
// 示例 3：
//
//
// 输入：num1 = "0", num2 = "0"
// 输出："0"
//
//
//
//
//
//
// 提示：
//
//
// 1 <= num1.length, num2.length <= 10⁴
// num1 和num2 都只包含数字 0-9
// num1 和num2 都不包含任何前导零
//
// Related Topics 数学 字符串 模拟 👍 525 👎 0

package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := addStrings("456", "777")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	n1 := len(num1) - 1
	n2 := len(num2) - 1
	ans := ""
	tmp := 0
	for ; n1 >= 0 || n2 >= 0 || tmp != 0; n1, n2 = n1-1, n2-1 {
		var x, y int
		if n1 >= 0 {
			x = int(num1[n1] - '0')
		}
		if n2 >= 0 {
			y = int(num2[n2] - '0')
		}
		res := x + y + tmp
		ans = strconv.Itoa(res%10) + ans
		tmp = res / 10
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
/*
解题思路：申请字符串，和临时进位数
从字符串尾部开始累加，若大于 10 则累计到进位，一直遍历累加
*/

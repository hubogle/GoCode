// 根据 逆波兰表示法，求表达式的值。
//
// 有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
//
// 注意 两个整数之间的除法只保留整数部分。
//
// 可以保证给定的逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
//
//
//
// 示例 1：
//
//
// 输入：tokens = ["2","1","+","3","*"]
// 输出：9
// 解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
//
//
// 示例 2：
//
//
// 输入：tokens = ["4","13","5","/","+"]
// 输出：6
// 解释：该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
//
//
// 示例 3：
//
//
// 输入：tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
// 输出：22
// 解释：该算式转化为常见的中缀算术表达式为：
//  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22
//
//
//
// 提示：
//
//
// 1 <= tokens.length <= 10⁴
// tokens[i] 是一个算符（"+"、"-"、"*" 或 "/"），或是在范围 [-200, 200] 内的一个整数
//
//
//
//
// 逆波兰表达式：
//
// 逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。
//
//
// 平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
// 该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。
//
//
// 逆波兰表达式主要有以下两个优点：
//
//
// 去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
// 适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中
//
// Related Topics 栈 数组 数学 👍 469 👎 0

package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	result := evalRPN(tokens)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func evalRPN(tokens []string) int {
	// +、-、*、/
	var stack []int
	var ans int
	for _, v := range tokens {
		val, err := strconv.Atoi(v)
		if err == nil { // 数字
			stack = append(stack, val)
		} else {
			l, r := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				ans = l + r
			case "-":
				ans = l - r
			case "*":
				ans = l * r
			case "/":
				ans = l / r
			}
			stack = append(stack, ans)
		}
	}
	return stack[0]
}

// leetcode submit region end(Prohibit modification and deletion)

func evalRPN1(tokens []string) int {
	stack := make([]int, (len(tokens)+1)/2)
	index := -1
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			index++
			stack[index] = val // 存储当前操作数
		} else {
			index-- // 取出左右两个数进行计算
			switch token {
			case "+":
				stack[index] += stack[index+1]
			case "-":
				stack[index] -= stack[index+1]
			case "*":
				stack[index] *= stack[index+1]
			default:
				stack[index] /= stack[index+1]
			}
		}
	}
	return stack[0]
}

/*
思路1：利用栈的特性，从左到右计算，用栈存储操作数
1. 遇到操作数，将数据入栈
2. 遇到运算符，从栈中取出前两个数，利用运算符计算数值后将数据入栈

思路2：数组模拟栈
一个有效的计算公式，长度 `n` 一定是奇数，操作数比运算符多一个。
操作数有 `(n+1)/2` 运算符有 `(n-1)/2`
数组最坏情况占用空间 `(n+1)/2`

数组模拟栈 stack ，数组下标 0 对应栈底，定义 index 表示元素下标位置，初始化 index = -1
如果遇到操作数，则将 index 的值加 1，然后将操作数赋给 stack[index]；
如果遇到运算符，则将 index 的值减 1，此时 stack[index] 和 stack[index+1] 的元素分别是左操作数和右操作数；

时间复杂度：O(n) 遍历数组所有元素
空间复杂度：O(n) 为操作数申请的空间大小
*/

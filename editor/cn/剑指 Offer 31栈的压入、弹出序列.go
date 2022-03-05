// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈
// 的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
//
//
//
// 示例 1：
//
// 输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
// 输出：true
// 解释：我们可以按以下顺序执行：
// push(1), push(2), push(3), push(4), pop() -> 4,
// push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//
//
// 示例 2：
//
// 输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
// 输出：false
// 解释：1 不能在 2 之前弹出。
//
//
//
//
// 提示：
//
//
// 0 <= pushed.length == popped.length <= 1000
// 0 <= pushed[i], popped[i] < 1000
// pushed 是 popped 的排列。
//
//
// 注意：本题与主站 946 题相同：https://leetcode-cn.com/problems/validate-stack-sequences/
// Related Topics 栈 数组 模拟 👍 288 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	result := validateStackSequences(pushed, popped)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func validateStackSequences(pushed []int, popped []int) bool {
	var stack = make([]int, 0, len(pushed))
	index := 0
	for _, v := range pushed {
		stack = append(stack, v)
		n := len(stack)
		for n > 0 && popped[index] == stack[n-1] {
			stack = stack[:n-1]
			index++
			n = len(stack)
		}
	}
	return len(stack) == 0
}

// leetcode submit region end(Prohibit modification and deletion)

func validateStackSequences1(pushed []int, popped []int) bool {
	stack := list.New() // 辅助栈
	index := 0
	for _, v := range pushed {
		stack.PushBack(v)
		for stack.Len() > 0 && popped[index] == stack.Back().Value.(int) {
			stack.Remove(stack.Back())
			index++
		}
	}
	return stack.Len() == 0
}

/*
题目：栈的入栈顺序，是否为预期的弹出顺序
解题思路：构造一个模拟栈
1. 模拟栈正常入栈。
2. 如果模拟栈最后一个，与弹出队列第一位相同。
3. 模拟栈弹出最后一个，弹出队列删除第一位值。
4. 重复 2.3 步骤
5. 最后判断模拟栈是否为空
*/

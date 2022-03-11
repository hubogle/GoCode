// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
//
//
// 只使用数字1到9
// 每个数字 最多使用一次
//
//
// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
//
//
//
// 示例 1:
//
//
// 输入: k = 3, n = 7
// 输出: [[1,2,4]]
// 解释:
// 1 + 2 + 4 = 7
// 没有其他符合的组合了。
//
// 示例 2:
//
//
// 输入: k = 3, n = 9
// 输出: [[1,2,6], [1,3,5], [2,3,4]]
// 解释:
// 1 + 2 + 6 = 9
// 1 + 3 + 5 = 9
// 2 + 3 + 4 = 9
// 没有其他符合的组合了。
//
// 示例 3:
//
//
// 输入: k = 4, n = 1
// 输出: []
// 解释: 不存在有效的组合。
// 在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
//
//
//
//
// 提示:
//
//
// 2 <= k <= 9
// 1 <= n <= 60
//
// Related Topics 数组 回溯 👍 432 👎 0

package main

import "fmt"

func main() {
	result := combinationSum3(3, 7)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var path []int
	var backtracking func(index, sum int)
	backtracking = func(index, sum int) {
		// if sum > n {
		// 	return
		// }
		// if sum == n {
		// 	if len(path) == k {
		// 		ans = append(ans, append([]int{}, path...))
		// 	}
		// 	return
		// } // 注意这里减枝操作的细节，判断总和和数量范围
		if sum == n && len(path) == k {
			ans = append(ans, append([]int{}, path...))
		}
		for i := index; i <= 9-(k-len(path))+1 && sum < n && len(path) < k; i++ { // 如果减枝操作在这里效率更高
			path = append(path, i)
			backtracking(i+1, sum+i)
			path = path[:len(path)-1]
		}
	}
	backtracking(1, 0)
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

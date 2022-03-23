// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
//
// 示例 1:
//
//
// 输入：s = "abaccdeff"
// 输出：'b'
//
//
// 示例 2:
//
//
// 输入：s = ""
// 输出：' '
//
//
//
//
// 限制：
//
// 0 <= s 的长度 <= 50000
// Related Topics 队列 哈希表 字符串 计数 👍 194 👎 0

package main

import "fmt"

func main() {
	result := firstUniqChar("abaccdeff")
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) byte {
	cache := [26]int{}
	queue := make([]byte, 0)
	n := len(s)
	for i := range cache {
		cache[i] = n
	}
	for k, v := range s {
		if cache[v-'a'] == n {
			cache[v-'a'] = k
			queue = append(queue, byte(v))
		} else {
			cache[v-'a'] = -1
			for len(queue) > 0 && cache[queue[0]-'a'] == -1 {
				queue = queue[1:]
			}
		}
	}
	if len(queue) > 0 {
		return s[cache[queue[0]-'a']]
	}
	return ' '
}

// leetcode submit region end(Prohibit modification and deletion)
func firstUniqChar1(s string) byte {
	cache := make([]int, 'z'-'a'+1)
	for _, v := range s {
		cache[v-'a']++
	}
	for _, v := range s {
		if cache[v-'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}

/*
解题思路：26位数组当作 map
第一次遍历字符串放到 map 里累加值，第二次遍历那个字符串数量为 1

队列，借助 cache 缓存判断是否有重复的
*/

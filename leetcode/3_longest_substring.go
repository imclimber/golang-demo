package main

import (
	"log"
)

func main() {
	input := `pwwkew`
	res := lengthOfLongestSubstring(input)

	log.Println("res: ", res)
}

// func lengthOfLongestSubstring(s string) int {
// 	max := 0
// 	dataMap := make(map[byte]byte)
// 	index := 0
// 	leng := len(s)
// 	for i := 0; i < leng; i++ {
// 		if i > 0 {
// 			delete(dataMap, s[i-1])
// 		}
// 		for j := index; j < leng; j++ {
// 			if _, ok := dataMap[s[index]]; !ok {
// 				dataMap[s[index]]++
// 				index++
// 			} else {
// 				break
// 			}
// 		}
// 		if max < len(dataMap) {
// 			max = len(dataMap)
// 		}
// 		if index == leng-1 {
// 			break
// 		}
// 	}

// 	return max
// }

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		if _, ok := m[s[rk+1]]; !ok {
			log.Println("m[s[rk+1]]: ", m[s[rk+1]])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			log.Println("in:", m[s[rk+1]])

			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

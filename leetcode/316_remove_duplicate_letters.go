package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "cbacdcbc"
	// s := "bcabc"
	// s := "abacb"
	res := removeDuplicateLetters_V2(s)
	fmt.Println("res:", res)
}

func removeDuplicateLetters_V2(s string) string {
	// 获取每个元素的数量，看过的元素计数不能出现在其中
	existsCount := make([]int, 26)
	for i, _ := range s {
		fmt.Printf("s[i]: %c\n", s[i])

		existsCount[s[i]-'a']++
	}
	// fmt.Println("existsCount: ", existsCount)

	// 缓存栈中所有元素，用于判定重复元素
	visited := make([]bool, 26)

	stack := make([]byte, 0)

	for i, _ := range s {
		ch := s[i]

		// 元素已经存在于缓存中，无需处理
		if visited[ch-'a'] {
			existsCount[ch-'a']-- // 次数需要减一，保证看过的元素不出现在统计信息中
			continue
		}

		// 元素小于栈顶元素字典序，且后续元素中存在相同栈顶元素，出栈，删除缓存中元素
		for len(stack) > 0 && ch < stack[len(stack)-1] && existsCount[stack[len(stack)-1]-'a'] > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			visited[tmp-'a'] = false
		}

		// 元素不存在于缓存中，加入缓存，同时加入栈
		visited[ch-'a'] = true
		stack = append(stack, ch)
		existsCount[ch-'a']-- // 次数需要减一，保证看过的元素不出现在统计信息中

		fmt.Println("visited: ", visited)
		fmt.Println("stack: ", stack)
		fmt.Println("existsCount: ", existsCount)
	}

	return string(stack)
}

func removeDuplicateLetters(s string) string {
	srune := []rune(s)

	strs := make([]string, 0)
	for i, _ := range srune {
		strs = append(strs, string(srune[i]))
	}
	slen := len(strs)

	var stack ArrayStack
	visitedMap := make(map[string]int)

	for i, _ := range strs {
		fmt.Println(strs[i], stack.peek(), isExists(strs, i+1, slen-1, stack.peek()), strs[i+1:slen])

		// 栈中已经存在，跳过
		if _, ok := visitedMap[strs[i]]; ok {
			continue
		}

		// 之前元素存在字典序大的，判定之后元素中存在则需要弹出
		for !stack.isEmpty() && strs[i] < stack.peek() && isExists(strs, i+1, slen-1, stack.peek()) {
			delete(visitedMap, stack.peek())
			stack.pop()
		}

		// 之前不存在则推入
		if _, ok := visitedMap[strs[i]]; !ok {
			stack.push(strs[i])
			visitedMap[strs[i]] = 1
		}

		fmt.Println("stack data: ", stack.datas)
	}

	var resBld strings.Builder
	resMiddle := make([]string, 0)
	for !stack.isEmpty() {
		resMiddle = append(resMiddle, stack.pop())
	}
	resLen := len(resMiddle)
	for i := resLen - 1; i >= 0; i-- {
		resBld.Write([]byte(resMiddle[i]))
	}

	return resBld.String()
}

func isExists(nums []string, start, end int, des string) bool {
	slen := len(nums)
	if start > slen-1 || end < 0 {
		return false
	}

	for i := start; i <= end; i++ {
		if nums[i] == des {
			return true
		}
	}

	return false
}

type ArrayStack struct {
	datas []string
	size  int
}

func (a *ArrayStack) isEmpty() bool {
	if a.size == 0 {
		return true
	}

	return false
}

func (a *ArrayStack) push(data string) {
	a.datas = append(a.datas, data)
	a.size++
}

func (a *ArrayStack) peek() string {
	if a.size > 0 {
		return a.datas[a.size-1]
	}

	return ""
}

func (a *ArrayStack) pop() string {
	if a.isEmpty() {
		return ""
	}

	tmp := a.datas[a.size-1]
	a.datas = a.datas[:a.size-1]
	a.size--

	return tmp
}

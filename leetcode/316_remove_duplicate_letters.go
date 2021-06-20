package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "cbacdcbc"
	// s := "bcabc"
	// s := "abacb"
	res := removeDuplicateLetters(s)
	fmt.Println("res:", res)
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

		// 之前元素存在字典序大的，判定之后元素中存在则需要弹出
		_, ok := visitedMap[strs[i]]
		for !ok && strs[i] < stack.peek() && isExists(strs, i+1, slen-1, stack.peek()) {
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

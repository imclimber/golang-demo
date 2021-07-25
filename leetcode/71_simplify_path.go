package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make([]int, 0)
	a = append(a, 1)
	b := a[:0]
	fmt.Println("b: ", b)

	path := "/a/./b/../../c/x/y/z/s"
	// path := "/../"
	// path := "/home//foo/"
	resPath := simplifyPath_V2(path)
	fmt.Println("resPath:", resPath)
}

func simplifyPath_V2(path string) string {
	dirs := strings.Split(path, "/")
	fmt.Println("dirs: ", len(dirs), dirs)

	var res []string
	for i, _ := range dirs {
		if dirs[i] == "." || dirs[i] == "" {
			continue
		}
		if dirs[i] == ".." {
			slen := len(res)
			if slen > 0 {
				res = res[:slen-1]
			}
			continue
		}

		res = append(res, dirs[i])
	}

	var resPath strings.Builder
	if len(res) == 0 {
		resPath.Write([]byte("/"))
		return resPath.String()
	}

	for i, _ := range res {
		resPath.Write([]byte("/"))
		resPath.Write([]byte(res[i]))
	}
	return resPath.String()
}

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	var stack ArrayStack

	fmt.Println("dirs: ", len(dirs), dirs)

	for i, _ := range dirs {
		if dirs[i] == "." || dirs[i] == "" {
			continue
		}

		if dirs[i] == ".." {
			if !stack.isEmpty() {
				stack.pop()
				fmt.Println("stack: ", stack.size, stack.datas)
			}

			continue
		}

		stack.push(dirs[i])
		fmt.Println("stack: ", stack.size, stack.datas)
	}

	fmt.Println("stack: ", stack.size, stack.datas)

	resMiddle := make([]string, 0)
	for !stack.isEmpty() {
		resMiddle = append(resMiddle, stack.pop())
	}
	fmt.Println("resMiddle: ", resMiddle)

	// 无需翻转，从尾往前遍历不就好了么
	res := reverse(resMiddle)
	var resPath strings.Builder

	if len(res) == 0 {
		return "/"
	}

	for i, _ := range res {
		resPath.Write([]byte("/"))
		resPath.Write([]byte(res[i]))
	}

	return resPath.String()
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

func (a *ArrayStack) pop() string {
	if a.isEmpty() {
		return ""
	}

	tmp := a.datas[a.size-1]
	a.datas = a.datas[:a.size-1]
	a.size--

	return tmp
}

func reverse(input []string) []string {
	slen := len(input)

	for i, j := 0, slen-1; i < j; {
		input[i], input[j] = input[j], input[i]
		i++
		j--
	}

	return input
}

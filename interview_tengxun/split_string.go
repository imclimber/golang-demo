// 1、有一个串 abc=1&d=2  ，实现字符串分割，导出abc=>1，b=>2的KV结构，请充分考虑对异常情况的兼容￼
package main

import (
	"fmt"
)

func main() {
	input := "abc=1&d=2&e*2"
	res := getResult(input)

	fmt.Println(res)
}

func getResult(input string) map[string]string {
	mapRes := make(map[string]string)
	// leng := len(input)
	left := -1
	strs := make([]string, 0)
	indexArray := getTargetLocationFirst(input)
	fmt.Println(indexArray)
	for i, _ := range indexArray {
		var str string
		if left == -1 {
			str = getKVString(input, 0, indexArray[i]-1)
			left = indexArray[i] + 1
		} else {
			str = getKVString(input, left, indexArray[i]-1)
			left = indexArray[i] + 1
		}

		strs = append(strs, str)
	}

	leng1 := len(strs)
	for i := 0; i < leng1; i++ {
		fmt.Println("-----------------------")
		str := strs[i]
		fmt.Println("str: ", str)
		kvIndex, hasGet := getTargetLocationSecond(str)
		if !hasGet {
			continue
		}
		fmt.Println("kvIndex: ", kvIndex)

		fmt.Println("*****************")
		leftPart := getKVString(str, 0, kvIndex-1)
		rightPart := getKVString(str, kvIndex+1, len(str)-1)
		fmt.Println("leftPart, rightPart: ", leftPart, rightPart)
		mapRes[leftPart] = rightPart
	}

	return mapRes
}

func getTargetLocationFirst(input string) []int {
	res := make([]int, 0)
	for i, _ := range input {
		if input[i] == '&' {
			res = append(res, i)
		}
	}
	res = append(res, len(input))

	return res
}

func getTargetLocationSecond(input string) (int, bool) {
	for i, _ := range input {
		if input[i] == '=' {
			return i, true
		}
	}

	return 0, false
}

func getKVString(input string, left, right int) string {
	fmt.Println("getKVString, left, right:", left, right)

	leng := len(input)
	var res string

	for i, _ := range input {
		if i >= left && i <= right && right < leng {
			res += string(input[i])
		}
	}

	// 禁止使用库函数
	// var buf bytes.Buffer
	// for i, _ := range input {
	// 	if i>=left && i<= right && right < leng {
	// 		buf.WriteByte(input[i])
	// 	}
	// }

	// res = buf.String()
	fmt.Println("getKVString: ", res)
	return res
}
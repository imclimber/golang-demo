package main

import (
	"fmt"
	"strings"
)

func main() {
	count := 0
	fmt.Scan(&count)
	var builder strings.Builder

	input := ""
	for i := 0; i < count; i++ {
		fmt.Scan(&input)

		// res := getRightResult(input)
		res := getRightResult_V2(input)
		builder.WriteString(res)
		if i != count-1 {
			builder.WriteString("\n")
		}
	}

	fmt.Println(builder.String())
}

func getRightResult(input string) string {
	leng := len(input)
	var res strings.Builder
	mid := make([]byte, 0)

	for i := 0; i < leng; i++ {
		if i < 2 || input[i] != mid[len(mid)-1] || input[i] != mid[len(mid)-2] {
			if i < 3 || input[i] != mid[len(mid)-1] || mid[len(mid)-2] != mid[len(mid)-3] {
				res.WriteByte(input[i])
				mid = append(mid, input[i])
			}
		}
	}

	return res.String()
}

func getRightResult_V2(input string) string {
	leng := len(input)
	var res strings.Builder
	mid := make([]byte, 0)

	for i := 0; i < leng; i++ {
		if i >= 2 && input[i] == mid[len(mid)-1] && input[i] == mid[len(mid)-2] {
			continue
		}

		if i >= 3 && input[i] == mid[len(mid)-1] && mid[len(mid)-2] == mid[len(mid)-3] {
			continue
		}

		mid = append(mid, input[i])
		res.WriteByte(input[i])
	}

	return res.String()
}

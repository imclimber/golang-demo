package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 5, 7}
	res := getResult(a, 8)
	fmt.Println(res)
}

func getResult(inputs []int, target int) [][]int {
	leng := len(inputs)

	middle := make([]int, 0)
	results := make([][]int, 0)

	resultMap := make(map[int]int)
	for i := 0; i < leng; i++ {
		if _, ok := resultMap[target-inputs[i]]; !ok {
			resultMap[target-inputs[i]] = inputs[i]
		}

		if _, ok := resultMap[inputs[i]]; ok {
			middle = append(middle, target-inputs[i])
		}
	}

	for i := len(middle) - 1; i >= 0; i-- {
		results = append(results, []int{middle[i], target - middle[i]})
	}

	return results
}

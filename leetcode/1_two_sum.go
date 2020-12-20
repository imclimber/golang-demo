package main

import (
	"log"
)

func main() {
	nums := []int{11, 5, 56, 2, 7}
	// result1 := twoSumByDoubleFor(nums, 9)
	// log.Println(result1)

	result2 := twoSumByMap(nums, 9)
	log.Println(result2)
}

func twoSumByDoubleFor(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSumByMap(nums []int, target int) []int {
	dataMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := dataMap[target-nums[i]]; ok {
			return []int{dataMap[target-nums[i]], i}
		}

		dataMap[nums[i]] = i
	}

	return []int{}
}

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("before sort, nums: ", nums)
	nums = insert_sort(nums, 0)
	fmt.Println("after sort, nums: ", nums)
}

func insert_sort(nums []int, target int) []int {
	len1 := len(nums)
	nums = append(nums, 100)

	fmt.Println("inner sort before, nums: ", nums)
	for i := len1 - 1; i >= 0; i-- {
		if nums[i] > target {
			nums[i+1] = nums[i]
			continue
		}
		nums[i] = target
	}
	fmt.Println("inner sort after, nums: ", nums)

	return nums
}

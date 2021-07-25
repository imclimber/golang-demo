package main

import "fmt"

func main() {
	nums := []int{8, 6, 7, 2, 9, 4, 3, 1, 0, 5}

	fmt.Println("before sort, nums: ", nums)
	nums = insertSort(nums)
	fmt.Println("after sort, nums: ", nums)
}

func insertSort(nums []int) []int {
	length := len(nums)
	temp := 0

	for i := 1; i < length; i++ {
		temp = nums[i]
		j := i

		for ; j > 0 && nums[j-1] > temp; j-- {
			nums[j] = nums[j-1]
		}

		nums[j] = temp
	}

	return nums
}

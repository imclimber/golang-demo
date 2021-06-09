package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("before, nums1: ", nums1)
	rotate1(nums1, 3)
	fmt.Println("after, nums1: ", nums1)
}

func rotate(nums []int, k int) {
	le := len(nums)

	k = k % le

	tmp := make([]int, k)
	for i := 0; i < k; i++ {
		tmp[i] = nums[le-k+i]
	}

	for j := le - 1; j >= k; j-- {
		nums[j] = nums[j-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
}

func rotate1(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}

	copy(nums, newNums)
}

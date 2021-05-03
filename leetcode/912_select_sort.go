package main

import "fmt"

func main() {
	nums := []int{8, 6, 7, 2, 9, 4, 3, 1, 0, 5}

	fmt.Println("before sort, nums: ", nums)
	select_sort(nums)
	fmt.Println("after sort, nums: ", nums)
}

func select_sort(nums []int) {
	len := len(nums)

	for i := 0; i < len-1; i++ {
		// 假设每一次循环第一个元素为本轮最小值
		minIndex := i
		for j := i + 1; j < len; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j // 只需要记录最小值，内部循环无需交换
			}
		}

		// 本轮最小值已经找到，跟设定位置交换
		swap(nums, i, minIndex)

		// fmt.Println("nums: ", nums)
	}
}

func swap(nums []int, i int, minIndex int) {
	nums[minIndex], nums[i] = nums[i], nums[minIndex]
}

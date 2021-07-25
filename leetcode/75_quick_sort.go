package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	depthForDebug := 0

	// nums := []int{2, 1, 3, 4, 7, 5, 6}
	// quickSort_V1(nums, 0, len(nums)-1, depthForDebug)
	// fmt.Println("after sort, nums: ", nums)

	// nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// quickSort_V2(nums, 0, len(nums)-1, depthForDebug)
	// fmt.Println("after sort, nums: ", nums)

	nums := []int{2, 3, 2, 1, 2, 2, 2, 3}
	quickSort_V3(nums, 0, len(nums)-1, depthForDebug)
	fmt.Println("after sort, nums: ", nums)
}

// 循环不变量：[left, lt]  [lt+1, right]
func quickSort_V1(nums []int, left, right int, depthForDebug int) {
	// 递归深度打印
	for i := 0; i < depthForDebug; i++ {
		fmt.Printf("%+v", "-")
	}
	fmt.Println("divide (", left, ", ", right, ")")

	// 递归终止条件
	if left >= right {
		return
	}

	// handle
	lt := left
	for i := left + 1; i <= right; i++ {
		if nums[i] < nums[left] {
			lt++
			swap(nums, lt, i)
		}
	}
	swap(nums, lt, left) // 待排序元素归位

	// 递归
	quickSort_V1(nums, left, lt-1, depthForDebug+1)
	quickSort_V1(nums, lt+1, right, depthForDebug+1)
}

// 随机选择 pivot 解决有序数组问题。
// 循环不变量：[left, lt]  [lt+1, right]
func quickSort_V2(nums []int, left, right int, depthForDebug int) {
	// 递归深度打印
	for i := 0; i < depthForDebug; i++ {
		fmt.Printf("%+v", "-")
	}
	fmt.Println("divide (", left, ", ", right, ")")

	// 递归终止条件
	if left >= right {
		return
	}

	// handle
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(right-left+1) + left
	swap(nums, left, randIndex)
	lt := left
	for i := left + 1; i <= right; i++ {
		if nums[i] < nums[left] {
			lt++
			swap(nums, lt, i)
		}
	}
	swap(nums, lt, left) // 待排序元素归位

	// 递归
	quickSort_V2(nums, left, lt-1, depthForDebug+1)
	quickSort_V2(nums, lt+1, right, depthForDebug+1)
}

// 三分区间解决重复元素问题。
// 随机选择 pivot 解决有序数组问题。
// 循环不变量：[left, lt]  [lt+1, i)  [gt, right]
func quickSort_V3(nums []int, left, right int, depthForDebug int) {
	// 递归深度打印
	for i := 0; i < depthForDebug; i++ {
		fmt.Printf("%+v", "-")
	}
	fmt.Println("divide (", left, ", ", right, ")")

	// 递归终止条件
	if left >= right {
		return
	}

	// handle
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(right-left+1) + left
	swap(nums, left, randIndex)
	fmt.Println("before:", nums)
	lt, gt := left, right+1

	for i := left + 1; i < gt; {
		if nums[i] < nums[left] {
			lt++
			swap(nums, lt, i)
			i++
		} else if nums[i] > nums[left] {
			gt--
			swap(nums, gt, i)
		} else {
			i++
		}
	}
	swap(nums, lt, left) // 待排序元素归位
	fmt.Println("after:", nums)

	// 递归
	quickSort_V3(nums, left, lt-1, depthForDebug+1)
	quickSort_V3(nums, gt, right, depthForDebug+1)
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

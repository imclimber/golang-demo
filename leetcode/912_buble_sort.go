package main

import "fmt"

func main() {
	nums := []int{8, 6, 7, 2, 9, 4, 3, 1, 0, 5}

	fmt.Println("before sort, nums: ", nums)
	bubbleSort(nums)
	fmt.Println("after sort, nums: ", nums)
}

func bubbleSort(nums []int) []int {
	len := len(nums)

	// 选择排序（远程可以交换）
	// for i:= 0; i< len-1; i++ {
	//     minIndex := i
	//     for j := i + 1; j<len;j++{
	//         if nums[j] < nums[minIndex]{
	//             minIndex = j
	//         }
	//     }

	//     nums[i], nums[minIndex] = nums[minIndex], nums[i]
	// }

	// 冒泡排序(邻居才可以交换)
	for i := 0; i < len-1; i++ {
		for j := len - 1; j > i; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}

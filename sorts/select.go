package sorts

import (
	"testgo/utils"
)

// SelectSort ...
func SelectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {

		exIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[exIndex] > nums[j] {
				exIndex = j
			}
		}

		if exIndex != i {
			utils.SwapDirectly(&nums[i], &nums[exIndex])
		}
	}
}

// func main() {

// 	nums := []int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9}

// 	log.Println(nums)

// 	SelectSort(nums)

// 	log.Println(nums)
// }

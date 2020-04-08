package main

import (
	"log"
	"testgo/utils"
)

// BubleSort ...
func BubleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				utils.SwapDirectly(&nums[j], &nums[j+1])
			}
		}
	}
}

func main() {

	nums := []int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9}

	log.Println(nums)

	BubleSort(nums)

	log.Println(nums)
}

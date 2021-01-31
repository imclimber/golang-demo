package main

import (
	"log"
	"time"
)

func main() {
	a := fun([]int{-1, 0, 3, 5, 9, 12}, 2)
	log.Println("a: ", a)
}

func fun(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0

	for left <= right {
		mid = left + (right-left)/2

		log.Println(left, mid, right)
		time.Sleep(1 * time.Second)

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

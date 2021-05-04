package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	res := maxArea(nums)
	fmt.Println("after, res: ", res)
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1

	max := 0.0
	for l < r {
		max = math.Max(math.Min(float64(height[l]), float64(height[r]))*float64((r-l)), max)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return int(max)
}

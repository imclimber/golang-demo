package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	res := trap(nums)
	fmt.Println("after, res: ", res)
}

func getLeftMax(height []int) int {
	hlen := len(height)
	leftMax := make([]int, hlen)
	for i := 1; i < hlen; i++ {
		leftMax[i] = getMax(leftMax[i-1], height[i-1])
	}

	fmt.Println("leftMax: ", leftMax)
	return 0
}

func trap(height []int) int {
	hlen := len(height)
	if hlen < 3 {
		return 0
	}

	sumall := 0
	for i := 0; i < hlen-1; {
		if height[i] == 0 {
			i++
			continue
		}

		rindex := 0
		max := 0

		for j := i + 1; j < hlen; j++ {
			if height[j] >= height[i] {
				rindex = j
				max = height[j]
				fmt.Println("first: i, j, rindex, max: ", i, j, rindex, max)
				break
			}

			if height[j] >= max {
				max = height[j]
				rindex = j
				fmt.Println("second: i, j, rindex, max: ", i, j, rindex, max)
			}
		}

		sumall += (rindex-i-1)*getMin(max, height[i]) - sumarray(height, i+1, rindex-1)
		fmt.Println("i, sumall:", i, sumall)

		i = rindex
	}

	return sumall
}

func sumarray(nums []int, left, right int) int {
	if left > right {
		return 0
	}

	sum := 0
	for i := left; i <= right; i++ {
		sum += nums[i]
	}

	return sum
}

func getMin(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func getMax(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}

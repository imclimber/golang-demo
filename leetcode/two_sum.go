package leetcode

func twoSumByDoubleFor(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSumByMap(nums []int, target int) []int {
	dataMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := dataMap[target-nums[i]]; ok {
			return []int{i, dataMap[target-nums[i]]}
		}

		dataMap[nums[i]] = i
	}

	return []int{}
}

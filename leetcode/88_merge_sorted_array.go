package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	nums2 := []int{0, 6, 7, 8, 9}

	fmt.Println("before sort, nums1: ", nums1)
	merge(nums1, 5, nums2, 5)
	fmt.Println("after sort, nums1: ", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 归并排序, source1, nums2 => nums1
	// source1 := nums1[:m]
	source1 := make([]int, m)
	for x := 0; x < m; x++ {
		source1[x] = nums1[x]
	}
	fmt.Println("source1: ", source1)

	i, j, k := 0, 0, 0
	for ; k < m+n; k++ {
		// if i < m {
		// 	fmt.Printf("source1[%+v] = %+v\n", i, source1[i])
		// }
		// if j < n {
		// 	fmt.Printf("nums2[%+v]= %+v\n", j, nums2[j])
		// }

		// 边界
		if i == m {
			nums1[k] = nums2[j]
			j++
		} else if j == n {
			fmt.Println("i: ", i)
			nums1[k] = source1[i]
			i++
		} else if source1[i] <= nums2[j] {
			nums1[k] = source1[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
	}
}

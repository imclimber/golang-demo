package leetcode

import (
	"log"
	"testing"
)

func Test_twoSumByDoubleFor(t *testing.T) {
	nums := []int{11, 5, 56, 2, 7}
	result := twoSumByDoubleFor(nums, 9)
	log.Println(result)
}

func Test_twoSumByMap(t *testing.T) {
	nums := []int{11, 5, 56, 2, 7}
	result := twoSumByMap(nums, 9)
	log.Println(result)
}

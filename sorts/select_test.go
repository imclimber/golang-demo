package sorts

import (
	"testing"
)

func TestSelectSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "A",
			nums: []int{
				1, 10, 2, 9, 3, 8, 4, 7, 5, 6,
			},
		},
		{
			name: "B",
			nums: []int{
				10, 100, 20, 90, 30, 80, 40, 70, 50, 60,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("nums:", tt.nums)
			SelectSort(tt.nums)
			t.Log("after sort, nums:", tt.nums)
		})
	}
}

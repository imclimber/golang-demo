package main

import (
	"fmt"
)

func main() {
	var n, d, e int
	fmt.Scanf("%d %d", &n, &d)

	inputs := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &e)
		inputs = append(inputs, e)
	}

	//     fmt.Println(n, d, inputs)
	count, err := getResult(n, d, inputs)
	if err != nil {
		fmt.Println("err: ", err)
	}
	// fmt.Println(count, res)
	fmt.Println(count)
}

// 超时
func getResult(n, d int, locs []int) (count int64, err error) {
	leng := len(locs)
	for i, j := 0, 0; i+2 < leng; i++ {
		j = i + 2
		for j < leng && locs[j]-locs[i] <= d {
			j++
		}
		j--
		count += int64((j - i) * (j - i - 1) / 2)

		// fmt.Println(i, j, count)
	}
	count %= 99997867
	return count, nil
}

func getResult_V2(n, d int, locs []int) (count int64, err error) {
	leng := len(locs)
	for i, j := 0, 0; i < leng; i++ {
		for i >= 2 && locs[i]-locs[j] > d {
			j++
		}
		count += int64((i - j) * (i - j - 1) / 2)
	}

	count %= 99997867

	return count, nil
}

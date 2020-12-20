package main

import (
	"log"
)

func main() {
	// grid := [][]byte{
	// 	{'1', '1', '0', '0', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '1', '0', '0'},
	// 	{'0', '0', '0', '1', '1'},
	// }

	grid := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}

	log.Printf("grid: %+v", grid)

	num := numIslands(grid)
	log.Println("num: ", num)

	log.Printf("grid: %+v", grid)
}

func numIslands(grid [][]byte) int {
	rowNum := len(grid)
	columnNum := len(grid[0])

	num := 0
	for i := 0; i < rowNum; i++ {
		for j := 0; j < columnNum; j++ {
			if grid[i][j] == '1' {
				num++
				findNumAndSetZero(grid, i, j, rowNum, columnNum)
				// log.Printf("grid: %+v", grid)
			}
		}
	}

	return num
}

func findNumAndSetZero(grid [][]byte, a, b, rowNum, columnNum int) {

	if a >= rowNum || b >= columnNum || grid[a][b] == '0' {
		return
	}

	grid[a][b] = '0'

	// 向右
	if b+1 < columnNum && grid[a][b+1] == '1' {
		findNumAndSetZero(grid, a, b+1, rowNum, columnNum)
	}

	// 向左
	if b-1 >= 0 && grid[a][b-1] == '1' {
		findNumAndSetZero(grid, a, b-1, rowNum, columnNum)
	}

	// 向下
	if a+1 < rowNum && grid[a+1][b] == '1' {
		findNumAndSetZero(grid, a+1, b, rowNum, columnNum)
	}

	// 向上
	if a-1 >= 0 && grid[a-1][b] == '1' {
		findNumAndSetZero(grid, a-1, b, rowNum, columnNum)
	}
}

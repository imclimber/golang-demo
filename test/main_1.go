package main

import(
	"fmt"
	"time"
)

func main(){
	nums := []int{1,2,3,4,5,6,7,8,9,10}

	strs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	printCrossResult(nums, strs)
}



func printCrossResult(nums []int, strs []string){
	chan1 := make(chan int)

	go func(input []int){
		leng := len(input)
		for i:= 0; i< leng; i++{		
			chan1 <- input[i]
		}
	}(nums)

	go func(input []string){
		leng1 := len(strs)
		for j:= 0; j<leng1; j++{
			tmp := <- chan1
			fmt.Print(tmp)
			fmt.Print(input[j])
		}
	}(strs)

	time.Sleep(time.Second * 10)
}
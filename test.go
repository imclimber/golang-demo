package main

import "fmt"

var block = "package"

func main() {
	block := "function"
	{
		block := 1
		fmt.Printf("The block is %+v \n", block)
	}
	fmt.Printf("The block is %+v  \n", block)

	var a interface{}
	a = 1
	if a != nil {
		a := "2"
		fmt.Println(a)
	}
	fmt.Println(a)
}

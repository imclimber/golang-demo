package main

import (
	"fmt"
)

func foo(a []int) {
	fmt.Printf("foo point of a: %+p\n", &a)
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("after append foo point of a: %+p\n", &a)
	a[0] = 200
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Printf("before point of a: %+p\n", &a)
	foo(a)
	fmt.Printf("after point of a: %+p\n", &a)
	fmt.Println(a)

	b := a[1:3]
	fmt.Printf("b: %+p\n", &b)
	b = a[1:cap(a)]
	fmt.Printf("after b: %+p\n", &b)

	c := make([]int, 1)
	c = append(c, 1)
	fmt.Println("c: ", c)
}

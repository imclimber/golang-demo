package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	//Store
	m.Store(1, "a")
	m.Store(2, "b")

	//LoadOrStore
	//若key不存在，则存入key和value，返回false和输入的value
	v, ok := m.LoadOrStore("1", "aaa")
	fmt.Println(ok, v) //false aaa

	//若key已存在，则返回true和key对应的value，不会修改原来的value
	v, ok = m.LoadOrStore(1, "aaa")
	fmt.Println(ok, v) //true a

	//Load
	v, ok = m.Load(1)
	if ok {
		fmt.Println("it's an existing key,value is ", v)
	} else {
		fmt.Println("it's an unknown key")
	}

	//Range
	//遍历sync.Map, 要求输入一个func作为参数
	f := func(k, v interface{}) bool {
		//这个函数的入参、出参的类型都已经固定，不能修改
		//可以在函数体内编写自己的代码，调用map中的k,v

		fmt.Println(k, v)
		return true
	}
	m.Range(f)

	//Delete
	fmt.Println(m.Load(1))
	m.Delete(1)
	fmt.Println(m.Load(1))

}

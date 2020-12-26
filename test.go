package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// res := make([]Person, 0)

	// res = append(res, Person{
	// 	Name: "a",
	// 	Age:  10,
	// })
	// res = append(res, Person{
	// 	Name: "b",
	// 	Age:  20,
	// })
	// function(res)
	// log.Printf("%+v", res)

	// dataMap := make(map[int]string)
	// dataMap[1] = "1"
	// dataMap[2] = "2"

	// log.Println("dataMap[3]: ", dataMap[3])

	t := Teacher{}
	t.ShowA()
}

type Pers struct{}
type Teacher struct {
	Pers
}

func (p *Pers) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *Pers) ShowB() {
	fmt.Println("showB")
}

func (t *Teacher) ShowB() {
	fmt.Println("Teacher showB")
}

func function(input []Person) {
	input[0].Name = "aaaaaa"
	input[0].Age = 100
	input[1].Name = "bbbbb"
	input[1].Age = 200
}

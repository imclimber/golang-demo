package main

import (
	// "fmt"
    // "crypto/md5"
    "log"
)

type Person struct {
    Name string
    Age int
}
 
func main() {
    // md5Res := md5.New()
    // md5Res.Write([]byte("100"))
    // res := md5Res.Sum(nil)
    // log.Println("res:", fmt.Sprintf("%x", res))

    // res := md5.Sum([]byte("100"))
    // log.Println("res:", fmt.Sprintf("%x", res))
    person := Person{
        Name: "zhangshan",
        Age: 20,
    }

    personTwo := Person{
        Name: "lisi",
        Age: 40,
    }


    // 引用类型作为值传递
    slice := make([]Person, 0)
    slice = append(slice, person, personTwo)
    changePerson(slice)
    log.Printf("not pointer slice: person.Name[%+v], person.Age[%+v]", slice[0].Name, slice[0].Age)
    log.Printf("not pointer slice: person.Name[%+v], person.Age[%+v]", slice[1].Name, slice[1].Age)

    slicePointer := make([]*Person, 0)
    slicePointer = append(slicePointer, &person, &personTwo)
    changePersonPointer(slicePointer)
    log.Printf("pointer slice: person.Name[%+v], person.Age[%+v]", slicePointer[0].Name, slicePointer[0].Age)
    log.Printf("pointer slice: person.Name[%+v], person.Age[%+v]", slicePointer[1].Name, slicePointer[1].Age)

    personThree := Person{
        Name: "zhangshan",
        Age: 20,
    }
    personMap := make(map[string][]Person)
    personMap["p"] = []Person{personThree}
    changePersonMap(personMap)
    log.Printf("map: person.Name[%+v], person.Age[%+v]", personMap["p"][0].Name, personMap["p"][0].Age)
}

func changePerson(slice []Person) error {
    sliceRes := make([]*Person, 0)
    
    for i, person := range slice {
        personInner := person
        sliceRes = append(sliceRes, &personInner)
        if i == 0{
            person.Name = "1111111"
            person.Age = 100
        }else {
            person.Name = "22222"
            person.Age = 200
        }
    }
    
    log.Printf("sliceRes: %+v", *sliceRes[0])
    log.Printf("sliceRes: %+v", *sliceRes[1])

    return nil
}

func changePersonPointer(slice []*Person) error {
    sliceRes := make([]*Person, 0)
    
    for i, person := range slice {
        personInner := person
        sliceRes = append(sliceRes, personInner)
        if i == 0{
            person.Name = "1111111"
            person.Age = 100
        }else {
            person.Name = "22222"
            person.Age = 200
        }
    }
    
    log.Printf("sliceRes: %+v", *sliceRes[0])
    log.Printf("sliceRes: %+v", *sliceRes[1])

    return nil
}

func changePersonMap(personMap map[string][]Person) error {
    for _, person := range personMap {
        log.Printf("map.Name: %+v", person[0].Name)
        log.Printf("map.Age: %+v", person[0].Age)

        person[0].Name = "1111111"
        person[0].Age = 100
    }

    return nil
}
package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "James", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Moh", age: 40})

	s := person{name: "Alfred", age: 40}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
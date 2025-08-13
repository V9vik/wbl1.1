package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) Print() {
	fmt.Printf("Human{%s,%d}\n", h.name, h.age)
}

func (h Human) Hello() {
	fmt.Printf("Hello,%s!\n", h.name)
}

type Action struct {
	Human
}

func main() {
	human := Human{
		name: "Vladimir",
		age:  20,
	}
	human.Hello()
	human.Print()

	action := Action{Human: human}

	action.Hello()
	action.Print()
}

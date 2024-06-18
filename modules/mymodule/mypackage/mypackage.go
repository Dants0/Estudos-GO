package mypackage

import "fmt"

type Person struct {
	name string
	age  int
}

func CreatePerson(name string, age int) Person{
	var person Person
	person.name = name
	person.age = age
	fmt.Println(person)
	return person
}

func PrintHello() {
	fmt.Println("Hello from my package!")
}
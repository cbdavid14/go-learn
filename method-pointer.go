package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary float64
}

func (e employee) details() {
	fmt.Printf("Name: %s\n", e.name)
}

func (e *employee) setNewName(newName string) {
	e.name = newName
}

func main() {
	emp := employee{name: "david", age: 31, salary: 14500}
	emp.details()

	//(*employee).setNewName(&emp, "Jhonatan")

	emp.setNewName("Jhonatan")
	fmt.Printf("name %+v\n", emp)
}

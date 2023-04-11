package main

import "fmt"

func main() {
	persona := Person{
		ID:          1,
		Name:        "Jose",
		DateOfBirth: "15/08/1993",
	}

	empleado := Employee{
		Person:   persona,
		ID:       111,
		Position: "Manager",
	}

	empleado.PrintEmployee()
}

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Person   Person
	ID       int
	Position string
}

func (employee Employee) PrintEmployee() {
	fmt.Println("ID Personal:", employee.Person.ID)
	fmt.Println("Nombre:", employee.Person.Name)
	fmt.Println("Fecha de Nacimiento:", employee.Person.DateOfBirth)

	fmt.Println("ID Empleado:", employee.ID)
	fmt.Println("Cargo:", employee.Position)
}

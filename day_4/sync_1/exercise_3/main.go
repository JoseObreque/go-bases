package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int

	// Entrada de salario por consola
	salary, err := ReadInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Comprobación de monto del salario
	if salary <= 10000 {
		fmt.Println(errors.New("Error: el salario ingresado no alcanza el mínimo imponible."))
	} else {
		fmt.Println("Debe pagar impuesto.")
	}
}

/*
Método que permite leer un valor entero por consola y retorna asigna dicho valor a la variable
entera entregada como parámetro.
*/
func ReadInput() (int, error) {
	var number int

	fmt.Printf("Ingresar salario: ")
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		return 0, errors.New("Entrada inválida")
	}

	return number, nil
}

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
	if salary < 150000 {
		err := fmt.Errorf("%w", &myError{msg: "El salario ingresado no alcanza el mínimo imponible"})
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto.")
	}
}

// Estructura de error
type myError struct {
	msg string
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

/*
Método que retorna un mensaje de error (string) en caso de que el salario del empleado
es menor a 150.000.
*/
func (error *myError) Error() string {
	return error.msg
}

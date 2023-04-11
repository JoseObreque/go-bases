package main

import "fmt"

func main() {
	// Datos del trabajador
	edad := 29
	esEmpleado := true
	antiguedadLaboral := 3
	salario := 250000

	// Condicional 1: Verifica si no se cumplen condiciones para el préstamo
	// Condicional 2: Verifica si se otorga préstamo con intereses o no

	if edad <= 22 || !esEmpleado || antiguedadLaboral <= 1 {
		fmt.Println("No se otorgará préstamo")
	} else if salario > 100000 {
		fmt.Println("Se otorgará préstamo con intereses")
	} else {
		fmt.Println("Se otorgará préstamo sin intereses")
	}
}

package main

import "fmt"

func main() {
	var employees = map[string]int{
		"Benjamín": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30,
	}

	// Obtención de la edad de Benjamín
	fmt.Println("Edad de Benjamín:", employees["Benjamín"])

	// Obtención de la cantidad de empleados mayores de 21 años
	adultsOver21 := 0
	for _, age := range employees {
		if age > 21 {
			adultsOver21++
		}
	}
	fmt.Println("Cantidad de empleados mayores de 21 años:", adultsOver21)

	// Adición de un nuevo empleado (Federico - 25 años)
	employees["Federico"] = 25

	// Eliminación de un empleado
	delete(employees, "Pedro")

	// Estado actual de los datos
	fmt.Println("Datos:", employees)
}

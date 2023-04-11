package main

import (
	"errors"
	"fmt"
)

// Slice que contiene clientes
var customers = []customer{
	{
		legajo:    "1234",
		nombre:    "Carlos",
		dni:       "12345678-9",
		telefono:  "+56911111111",
		direccion: "Calle Falsa 123",
	},
	{
		legajo:    "4321",
		nombre:    "Charles",
		dni:       "98765432-1",
		telefono:  "+56922222222",
		direccion: "Calle Verdadera 321",
	},
}

func main() {
	// Manejo de panic: customerExists
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)

			// Mensajes requeridos
			fmt.Println("\nFin de la ejecución")
			fmt.Println("Se detectaron múltiples errores en tiempo de ejecución")
		}
	}()

	// Búsqueda de cliente por DNI (puede arrojar un panic)
	_, err := customerExists("12345678-0")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("El cliente no existe.")
	}

	// Checkeo de datos de cliente a registrar
	possibleCustomer := customer{
		legajo:    "0011",
		nombre:    "Pedro",
		dni:       "12345678-1",
		telefono:  "+56988776655",
		direccion: "Calle Neutral 1111",
	}

	_, err2 := checkCustomerData(possibleCustomer)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Todos los datos son correctos")
	}

	fmt.Println("Fin de la ejecución")
}

// Estructura que representa a un cliente.
type customer struct {
	legajo    string
	nombre    string
	dni       string
	telefono  string
	direccion string
}

/*
Función que busca un cliente en el registro mediante su DNI.

Si se encuentra un resultado, retorna el valor true junto a un error. En caso
contrario retorna el valor false (sin errores asociados).
*/
func customerExists(dni string) (bool, error) {
	for _, registeredCustomer := range customers {
		if registeredCustomer.dni == dni {
			return true, errors.New("Error: el cliente ya existe.")
		}
	}
	return false, nil
}

/*
Función que verifica si los datos de un cliente a registrar son válidos, es decir, no están
en blanco (con su valor por defecto).

Retorna un valor true si todos los datos son correctos. En caso de existir al menos un campo
en blanco, retorna un valor false junto a un error.
*/
func checkCustomerData(customerData customer) (bool, error) {
	// Array que contiene los datos del cliente de forma ordenada (iterable)
	fieldValues := [5]interface{}{
		customerData.legajo,
		customerData.nombre,
		customerData.dni,
		customerData.direccion,
		customerData.telefono,
	}

	// Búsqueda de posibles campos en blanco
	for _, value := range fieldValues {
		if value == "" {
			return false, errors.New("No pueden existir campos vacíos.")
		}
	}

	return true, nil
}

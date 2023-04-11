package exercise_5

import (
	"errors"
)

/*
Función que recibe una especie de animal (string) y retorna una función que
permite calcular la cantidad de alimento necesaria para un número determinado
de este tipo de animales.
*/
func animal(especie string) (func(int) float64, error) {
	switch especie {
	case "perro":
		return alimentoPerros, nil
	case "gato":
		return alimentoGatos, nil
	case "hamster":
		return alimentoHamsters, nil
	case "tarántula":
		return alimentoTarantulas, nil
	default:
		return nil, errors.New("El animal indicado no es válido")
	}
}

/*
Función que permite calcular la cantidad de alimento (float64) necesario para dar de
comer a una cantidad de perros dada (int).
*/
func alimentoPerros(cantidadAnimales int) float64 {
	const alimentoPorAnimal float64 = 10
	return alimentoPorAnimal * float64(cantidadAnimales)
}

/*
Función que permite calcular la cantidad de alimento (float64) necesario para dar de
comer a una cantidad de gatos dada (int).
*/
func alimentoGatos(cantidadAnimales int) float64 {
	const alimentoPorAnimal float64 = 5
	return alimentoPorAnimal * float64(cantidadAnimales)
}

/*
Función que permite calcular la cantidad de alimento (float64) necesario para dar de
comer a una cantidad de hamsters dada (int).
*/
func alimentoHamsters(cantidadAnimales int) float64 {
	const alimentoPorAnimal float64 = 0.250
	return alimentoPorAnimal * float64(cantidadAnimales)
}

/*
Función que permite calcular la cantidad de alimento (float64) necesario para dar de
comer a una cantidad de tarántulas dada (int).
*/
func alimentoTarantulas(cantidadAnimales int) float64 {
	const alimentoPorAnimal float64 = 0.150
	return alimentoPorAnimal * float64(cantidadAnimales)
}

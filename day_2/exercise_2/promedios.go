package exercise_2

import "errors"

/*
Función que toma un número variable de notas (enteros) como parámetros y
retorna su promedio (float64).
*/
func calcularPromedio(notas ...int) (float64, error) {
	suma := 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("No se admiten notas negativas")
		}
		suma += nota
	}
	return float64(suma) / float64(len(notas)), nil
}

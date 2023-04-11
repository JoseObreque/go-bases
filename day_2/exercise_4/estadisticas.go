package exercise_4

/*
Función que recibe una operación deseada (string) y retorna una función apropiada
para calcular dicha operación estadística.
*/
func estadistica(operacion string) func(notas ...int) float64 {
	switch operacion {
	case "minimo":
		return buscarMinimo
	case "maximo":
		return buscarMaximo
	case "promedio":
		return promedio
	default:
		return nil
	}
}

/*
Función que recibe una cantidad variable de notas (int) como parámetros de entrada y
retorna la nota mínima (float64).
*/
func buscarMinimo(notas ...int) float64 {
	min := notas[0]

	for i := 1; i < len(notas); i++ {
		if notas[i] < min {
			min = notas[i]
		}
	}

	return float64(min)
}

/*
Función que recibe una cantidad variable de notas (int) como parámetros de entrada y
retorna la nota máxima (float64).
*/
func buscarMaximo(notas ...int) float64 {
	max := notas[0]

	for i := 1; i < len(notas); i++ {
		if notas[i] > max {
			max = notas[i]
		}
	}

	return float64(max)
}

/*
Función que recibe una cantidad variable de notas (int) como parámetros de entrada y
retorna el promedio de las calificaciones (float64).
*/
func promedio(notas ...int) float64 {
	sumatoria := 0

	for _, nota := range notas {
		sumatoria += nota
	}

	return float64(sumatoria) / float64(len(notas))
}

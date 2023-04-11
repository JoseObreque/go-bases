package exercise_3

import "errors"

/*
Función que toma la cantidad de minutos trabajados (int) por un colaborador y
su categoría (string) y retorna su salario mensual (float64).
*/
func calcularSalario(minutosTrabajados int, categoria string) (float64, error) {
	// Cálculo de horas trabajadas
	horasTrabajadas := float64(minutosTrabajados) / 60

	// Cálculo de exercise_3
	switch categoria {
	case "A":
		return horasTrabajadas * 1000, nil
	case "B":
		// Bono: 20% adicional sobre salario mensual
		return (horasTrabajadas * 1500) * 1.20, nil
	case "C":
		// Bono: 50% adicional sobre salario mensual
		return (horasTrabajadas * 3000) * 1.50, nil
	default:
		return 0, errors.New("Categoría no válida")
	}
}

package exercise_1

import "errors"

/*
Función que calcula los exercise_1 (float64) sobre el salario (int) de un trabajador.

En caso de ingresar un salario negativo, la función retornará un error.
*/
func calcularImpuestos(salario int) (float64, error) {
	impuestoBase := 0.17

	switch {
	case salario > 150000:
		return float64(salario) * (0.10 + impuestoBase), nil
	case salario > 50000:
		return float64(salario) * impuestoBase, nil
	case salario > 0:
		return 0, nil
	default:
		return 0, errors.New("Salario no válido")
	}
}

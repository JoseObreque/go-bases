package main

import (
	"errors"
	"fmt"
)

func main() {
	salario, err := calcularSalario(2000, 1000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Salario:", salario)
	}
}

func calcularSalario(horasTrabajadas float64, valorHora float64) (float64, error) {
	if horasTrabajadas < 80 {
		return 0, errors.New("Error: el trabajador no puede haber trabajado menos de 80 hrs mensuales")
	}

	salarioMensual := horasTrabajadas * valorHora

	if salarioMensual >= 150000 {
		return salarioMensual * 0.90, nil
	}

	return salarioMensual, nil
}

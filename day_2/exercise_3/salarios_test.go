package exercise_3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuestos(t *testing.T) {
	t.Run("Salario Categoría A", func(t *testing.T) {
		var (
			minutosTrabajados int    = 9600
			categoria         string = "A"
			salarioEsperado   float64
		)

		salarioEsperado = (float64(minutosTrabajados) / 60) * 1000

		salario, err := calcularSalario(minutosTrabajados, categoria)

		assert.Equal(t, salarioEsperado, salario)
		assert.Equal(t, nil, err)

	})

	t.Run("Salario Categoría B", func(t *testing.T) {
		var (
			minutosTrabajados int    = 9600
			categoria         string = "B"
			salarioEsperado   float64
		)

		salarioEsperado = ((float64(minutosTrabajados) / 60) * 1500) * 1.20

		salario, err := calcularSalario(minutosTrabajados, categoria)

		assert.Equal(t, salarioEsperado, salario)
		assert.Equal(t, nil, err)

	})

	t.Run("Salario Categoría C", func(t *testing.T) {
		var (
			minutosTrabajados int    = 9600
			categoria         string = "C"
			salarioEsperado   float64
		)

		salarioEsperado = ((float64(minutosTrabajados) / 60) * 3000) * 1.50

		salario, err := calcularSalario(minutosTrabajados, categoria)

		assert.Equal(t, salarioEsperado, salario)
		assert.Equal(t, nil, err)

	})

	t.Run("Salario Categoría B", func(t *testing.T) {
		var (
			minutosTrabajados int    = 9600
			categoria         string = "B"
			salarioEsperado   float64
		)

		salarioEsperado = ((float64(minutosTrabajados) / 60) * 1500) * 1.20

		salario, err := calcularSalario(minutosTrabajados, categoria)

		assert.Equal(t, salarioEsperado, salario)
		assert.Equal(t, nil, err)

	})

	t.Run("Salario Categoría Inválida", func(t *testing.T) {
		var (
			minutosTrabajados int     = 9600
			categoria         string  = "D"
			salarioEsperado   float64 = 0
		)

		salario, err := calcularSalario(minutosTrabajados, categoria)

		assert.Equal(t, salarioEsperado, salario)
		assert.Error(t, err)

	})
}

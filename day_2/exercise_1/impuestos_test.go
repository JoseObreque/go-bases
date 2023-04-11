package exercise_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuestos(t *testing.T) {
	t.Run("Salario sobre 100K", func(t *testing.T) {
		var (
			salario          int     = 200000
			impuestoEsperado float64 = float64(salario) * 0.27
		)
		impuesto, err := calcularImpuestos(salario)
		assert.Equal(t, impuestoEsperado, impuesto)
		assert.Equal(t, nil, err)
	})

	t.Run("Salario sobre 50K", func(t *testing.T) {
		var (
			salario          int     = 75000
			impuestoEsperado float64 = float64(salario) * 0.17
		)
		impuesto, err := calcularImpuestos(salario)
		assert.Equal(t, impuestoEsperado, impuesto)
		assert.Equal(t, nil, err)
	})

	t.Run("Salario bajo 50K", func(t *testing.T) {
		var (
			salario          int     = 20000
			impuestoEsperado float64 = 0
		)
		impuesto, err := calcularImpuestos(salario)
		assert.Equal(t, impuestoEsperado, impuesto)
		assert.Equal(t, nil, err)
	})

	t.Run("Salario Negativo", func(t *testing.T) {
		var (
			salario          int     = -10000
			impuestoEsperado float64 = 0
		)
		impuesto, err := calcularImpuestos(salario)
		assert.Equal(t, impuestoEsperado, impuesto)
		assert.Error(t, err)
	})
}

package exercise_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularPromedio(t *testing.T) {
	t.Run("Promedio notas válidas", func(t *testing.T) {
		var promedioEsperado float64 = 4
		promedio, err := calcularPromedio(1, 2, 3, 4, 10)
		assert.Equal(t, promedioEsperado, promedio)
		assert.Equal(t, nil, err)
	})
	t.Run("Promedio notas no válidas", func(t *testing.T) {
		var promedioEsperado float64 = 0
		promedio, err := calcularPromedio(1, -2, 3, 4, 10)
		assert.Equal(t, promedioEsperado, promedio)
		assert.Error(t, err)
	})
}

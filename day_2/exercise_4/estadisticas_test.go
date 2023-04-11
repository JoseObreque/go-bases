package exercise_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var notas = []int{10, 5, 7, 1, 8}

func TestBuscarMinimo(t *testing.T) {
	var minimoEsperado float64 = 1

	minimo := buscarMinimo(notas...)

	assert.Equal(t, minimoEsperado, minimo)
}

func TestBuscarMaximo(t *testing.T) {
	var maximoEsperado float64 = 10
	notas := []int{10, 5, 7, 1, 8}

	maximo := buscarMaximo(notas...)

	assert.Equal(t, maximoEsperado, maximo)
}

func TestPromedio(t *testing.T) {
	var promedioEsperado float64 = 6.2

	promedio := promedio(notas...)

	assert.Equal(t, promedioEsperado, promedio)
}

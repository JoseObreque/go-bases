package exercise_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var notas = []int{10, 5, 7, 1, 8}

func TestAlimentoPerro(t *testing.T) {
	var cantidadAlimentoEsperado float64 = 100
	cantidadPerros := 10

	cantidadAlimento := alimentoPerros(cantidadPerros)

	assert.Equal(t, cantidadAlimentoEsperado, cantidadAlimento)
}

func TestAlimentoGato(t *testing.T) {
	var cantidadAlimentoEsperado float64 = 50
	cantidadGatos := 10

	cantidadAlimento := alimentoGatos(cantidadGatos)

	assert.Equal(t, cantidadAlimentoEsperado, cantidadAlimento)
}

func TestAlimentoHamsters(t *testing.T) {
	var cantidadAlimentoEsperado float64 = 2.50
	cantidadHamsters := 10

	cantidadAlimento := alimentoHamsters(cantidadHamsters)

	assert.Equal(t, cantidadAlimentoEsperado, cantidadAlimento)
}

func TestAlimentoTarantulas(t *testing.T) {
	var cantidadAlimentoEsperado float64 = 1.50
	cantidadTarantulas := 10

	cantidadAlimento := alimentoTarantulas(cantidadTarantulas)

	assert.Equal(t, cantidadAlimentoEsperado, cantidadAlimento)
}

func TestAnimal(t *testing.T) {
	t.Run("Especie: perro", func(t *testing.T) {
		especie := "perro"
		funcionEsperada := alimentoPerros

		funcion, err := animal(especie)

		assert.IsType(t, funcionEsperada, funcion)
		assert.Nil(t, err)
	})

	t.Run("Especie: gato", func(t *testing.T) {
		especie := "gato"
		funcionEsperada := alimentoGatos

		funcion, err := animal(especie)

		assert.IsType(t, funcionEsperada, funcion)
		assert.Nil(t, err)
	})

	t.Run("Especie: hamster", func(t *testing.T) {
		especie := "hamster"
		funcionEsperada := alimentoHamsters

		funcion, err := animal(especie)

		assert.IsType(t, funcionEsperada, funcion)
		assert.Nil(t, err)
	})

	t.Run("Especie: tarantula", func(t *testing.T) {
		especie := "tarántula"
		funcionEsperada := alimentoTarantulas

		funcion, err := animal(especie)

		assert.IsType(t, funcionEsperada, funcion)
		assert.Nil(t, err)
	})

	t.Run("Especie: gato", func(t *testing.T) {
		especie := "gato"
		funcionEsperada := alimentoGatos

		funcion, err := animal(especie)

		assert.IsType(t, funcionEsperada, funcion)
		assert.Nil(t, err)
	})

	t.Run("Especie Inexistente", func(t *testing.T) {
		especie := "elefante"

		funcion, err := animal(especie)

		assert.Nil(t, funcion)
		assert.Error(t, err)
	})
}

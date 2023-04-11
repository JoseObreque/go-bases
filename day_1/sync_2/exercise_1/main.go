package main

import "fmt"

func main() {
	// Palabra a analizar
	palabra := "árbol"
	fmt.Println("Palabra:", palabra)

	// Obtención de longitud de palabra e impresión de letras por consola
	fmt.Println("Letras:")
	largo := 0
	for _, letra := range palabra {
		largo++
		fmt.Println(string(letra))
	}

	// Impresión de longitud de palabra por consola
	fmt.Println("Largo de la palabra:", largo)
}

package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Ejecución finalizada")

	filename := "./day_4/sync_2/exercise_1/customers.txt"
	data := readFile(filename)
	fmt.Println(data)
}

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado.")
	}
	return string(data)
}

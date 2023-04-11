package main

import "fmt"

func main() {
	// Número de mes a buscar
	numeroMes := 10

	// Map con los meses del año y su número de mes correspondiente
	meses := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	// Extracción del mes solicitado
	mes, existe := meses[numeroMes]
	if !existe {
		fmt.Println("Número de mes no válido")
	} else {
		fmt.Println(mes)
	}
}

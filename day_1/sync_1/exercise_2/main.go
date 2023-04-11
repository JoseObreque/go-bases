package main

import "fmt"

func main() {
	var (
		temperatura int = 13   // Temperatura (°C)
		humedad     int = 68   // Humedad relativa del aire (%)
		presion     int = 1017 // Presión atmosférica (hPa)
	)
	fmt.Println("Temperatura:", temperatura, "°C")
	fmt.Println("Humedad relativa:", humedad, "%")
	fmt.Println("Presión Atmosférica:", presion, "hPa")
}
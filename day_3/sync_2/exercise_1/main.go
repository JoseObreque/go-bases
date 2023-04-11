package main

import "fmt"

func main() {
	alumno := Alumno{} // Nuevo alumno sin info registrada

	alumno.detalle() // Se muestran los detalles del alumno por defecto

	alumno.modificarNombre("John")
	alumno.modificarApellido("Doe")
	alumno.modificarDNI("99999999-9")
	alumno.modificarFechaIngreso("28/03/2023")

	alumno.detalle() // Se muestran los detalles del alumno (con modificaciones)
}

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

// Función que permite modificar el nombre de un alumno.
func (alumno *Alumno) modificarNombre(nuevoNombre string) {
	alumno.Nombre = nuevoNombre
}

// Función que permite modificar el apellido de un alumno.
func (alumno *Alumno) modificarApellido(nuevoApellido string) {
	alumno.Apellido = nuevoApellido
}

// Función que permite modificar el DNI registrado de un alumno.
func (alumno *Alumno) modificarDNI(nuevoDNI string) {
	alumno.DNI = nuevoDNI
}

// Función que permite modificar la fecha de ingreso de un alumno.
func (alumno *Alumno) modificarFechaIngreso(nuevaFechaIngreso string) {
	alumno.Fecha = nuevaFechaIngreso
}

// Función que muestra el detalle de la información general de un alumno.
func (alumno Alumno) detalle() {
	fmt.Println("\nInformación del Alumno")
	fmt.Println("Nombre:", alumno.Nombre)
	fmt.Println("Apellido:", alumno.Apellido)
	fmt.Println("DNI:", alumno.DNI)
	fmt.Println("Fecha de Ingreso:", alumno.Fecha)
}

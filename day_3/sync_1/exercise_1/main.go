package main

import (
	"errors"
	"fmt"
)

var Products = []Product{
	{
		ID:          1,
		Name:        "Macbook Pro",
		Price:       2000000,
		Description: "Macbook Pro (14 pulgadas) con chip M1",
		Category:    "Tecnología",
	},
}

func main() {
	// Definición de nuevo producto
	nuevoProducto := Product{
		ID:          2,
		Name:        "Iphone 14",
		Price:       1100000,
		Description: "Iphone 14 (nuevo)",
		Category:    "Tecnología",
	}

	// Almacenado del producto en slice global
	nuevoProducto.Save()

	// Presentación de los contenidos del slice global por consola
	nuevoProducto.GetAll()

	// Búsqueda de producto por ID en el slice global (existente)
	producto, err := GetById(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(producto)
	}
}

// Declaración de estructura
type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
	Category    string
}

/*
Método que permite almacenar el producto actual.
*/
func (p Product) Save() {
	Products = append(Products, p)
}

/*
Método que permite presentar los productos almacenados por consola.
*/
func (p Product) GetAll() {
	for i, product := range Products {
		fmt.Println("Producto", i)
		fmt.Println(product)
	}
}

/*
Función que permite buscar un determinado producto en base a su ID.
*/
func GetById(id int) (Product, error) {
	for _, product := range Products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, errors.New("Producto no encontrado")
}
